package usql

import (
	"github.com/siddontang/go-mysql/canal"
	"github.com/updogliu/ugo/ulog"
)

type TableListener struct {
	// User can listen for row events at this channel.
	C chan *canal.RowsEvent

	source *canal.Canal

	// Embed placeholder methods to satisfy interface `DummyEventHandler`
	canal.DummyEventHandler
}

// Example:
//     listener := NewTableListener(config, []string{"table_foo", "table_bar"})
//     for rowEvent := range listener.C {}
//
func NewTableListener(addr, user, password, dbname string, tables []string) *TableListener {
	if len(tables) == 0 {
		ulog.Panic("Should not create a listener of no table")
	}

	cfg := canal.NewDefaultConfig()
	cfg.Addr = addr
	cfg.User = user
	cfg.Password = password
	cfg.Dump.ExecutionPath = ""
	cfg.ParseTime = true

	for _, table := range tables {
		cfg.IncludeTableRegex = append(cfg.IncludeTableRegex, dbname+"\\."+table)
	}

	source, err := canal.NewCanal(cfg)
	if err != nil {
		ulog.Panic("Failed to NewTableListener: ", err)
	}

	tl := &TableListener{
		source: source,
		C:      make(chan *canal.RowsEvent, 4096),
	}

	ulog.Info("New TableListener run on tables: ", tables)
	go tl.run()

	return tl
}

func (tl *TableListener) run() {
	tl.source.SetEventHandler(tl)

	pos, err := tl.source.GetMasterPos()
	if err != nil {
		ulog.Panic("Failed to GetMasterPos: ", err)
	}

	err = tl.source.RunFrom(pos)
	if err != nil {
		ulog.Panic("Got an error in running TableListener: ", err)
	}
}

func (tl *TableListener) OnRow(e *canal.RowsEvent) error {
	tl.C <- e
	return nil
}
