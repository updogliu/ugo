package uws

import (
	gw "github.com/gorilla/websocket"
)

// Message types copied from (and should be synced with) gorilla/websocket.
//
// The message types are defined in RFC 6455, section 11.8.
const (
	// TextMessage denotes a text data message. The text message payload is
	// interpreted as UTF-8 encoded text data.
	TextMessage = gw.TextMessage

	// BinaryMessage denotes a binary data message.
	BinaryMessage = gw.BinaryMessage

	// CloseMessage denotes a close control message. The optional message
	// payload contains a numeric code and text. Use the FormatCloseMessage
	// function to format a close message payload.
	CloseMessage = gw.CloseMessage

	// PingMessage denotes a ping control message. The optional message payload
	// is UTF-8 encoded text.
	PingMessage = gw.PingMessage

	// PongMessage denotes a pong control message. The optional message payload
	// is UTF-8 encoded text.
	PongMessage = gw.PongMessage
)
