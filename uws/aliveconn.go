package uws

import (
	"fmt"
	"net/http"
	"time"

	gw "github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/updogliu/ugo/utime"
)

// A managed websocket connection that will auto-reconnect on too many consecutive errors.
// NOT safe for concurrent usage.
type AliveConn struct {
	url string
	cfg AliveConnConfig

	numReadErrs      int
	numReconnectErrs int // clear this only after successfully read a msg on the new conn
	conn             *gw.Conn
}

// Default values will be used for unset fields (fields with zero-values).
type AliveConnConfig struct {
	// Reconnect after failed to read message for more than this many times in succession.
	MaxReadErrs int

	// Panic after failed to reconnect for more than this many times in succession.
	MaxReconnectErrs int

	ReadErrPauseMs      int64
	ReconnectErrPauseMs int64

	// Text message to send on every new underlying connection. Empty string means no message to send.
	InitMsgToSend string
}

func NewAliveConn(url string, cfg AliveConnConfig) *AliveConn {
	// Fill unset fields of `cfg` with default values.
	if cfg.MaxReadErrs == 0 {
		cfg.MaxReadErrs = 20
	}
	if cfg.MaxReconnectErrs == 0 {
		cfg.MaxReconnectErrs = 20
	}
	if cfg.ReadErrPauseMs == 0 {
		cfg.ReadErrPauseMs = 100
	}
	if cfg.ReconnectErrPauseMs == 0 {
		cfg.ReconnectErrPauseMs = 3000
	}

	return &AliveConn{url: url, cfg: cfg}
}

func (ac *AliveConn) ReadMessage() ( /*msgType*/ int /*msg*/, []byte, error) {
	if ac.conn == nil {
		if err := ac.reconnect(); err != nil {
			return 0, nil, err
		}
	}

	msgType, msg, err := ac.conn.ReadMessage()
	if err != nil {
		ac.numReadErrs++
		if ac.numReadErrs > ac.cfg.MaxReadErrs {
			ac.dropConn() // stop retrying on this conn
		} else {
			utime.SleepMs(ac.cfg.ReadErrPauseMs)
		}
		return 0, nil, err
	}

	ac.numReadErrs = 0
	ac.numReconnectErrs = 0
	return msgType, msg, nil
}

// Drop the current underlying connection, if any.
func (ac *AliveConn) dropConn() {
	if ac.conn != nil {
		if err := ac.conn.Close(); err != nil {
			// Log at info level because this is on a best-efforts basis.
			log.Info("Got an error when close the dropped conn: ", err)
		}
		ac.conn = nil
	}
}

func (ac *AliveConn) reconnect() (retErr error) {
	var dialer = &gw.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 45 * time.Second,
		// According to the doc, we should limit the buffer sizes to the maximum expected message size.
		// Buffers larger than the largest message do not provide any benefit.
		ReadBufferSize:  4 * 1024 * 1024,
		WriteBufferSize: 64 * 1024,
	}

	ac.dropConn()

	// Create a new underlying connection.
	conn, _, err := dialer.Dial(ac.url, nil)
	if err != nil {
		retErr = fmt.Errorf("Failed to dial %v: %v", ac.url, err)
		goto reconnectFailed
	}
	ac.conn = conn

	// Send the InitMsgToSend, if any.
	if ac.cfg.InitMsgToSend != "" {
		err := ac.conn.WriteMessage(TextMessage, []byte(ac.cfg.InitMsgToSend))
		if err != nil {
			retErr = fmt.Errorf("Failed to send InitMsgToSend: %v", err)
			goto reconnectFailed
		}
	}

	// Clear error counters.
	ac.numReadErrs = 0
	return nil

reconnectFailed:
	ac.numReconnectErrs++
	if ac.numReconnectErrs > ac.cfg.MaxReconnectErrs {
		log.Panicf("Failed to reconnect after %v retries. Last err: %v", ac.numReconnectErrs, retErr)
	}
	utime.SleepMs(ac.cfg.ReconnectErrPauseMs)
	return retErr
}
