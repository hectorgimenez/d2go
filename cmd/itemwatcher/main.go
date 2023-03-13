package main

import (
	"context"
	itemwatcher "github.com/hectorgimenez/d2go/cmd/itemwatcher/internal"
	"github.com/hectorgimenez/d2go/pkg/memory"
	"github.com/hectorgimenez/d2go/pkg/nip"
	"log"
	"os"
	"os/signal"
)

func main() {
	process, err := memory.NewProcess()
	if err != nil {
		log.Fatalf("error starting process: %s", err.Error())
	}

	gr := memory.NewGameReader(process)

	rules, err := nip.ReadDir("config/itemfilter/")
	if err != nil {
		log.Fatalf("error reading NIP files: %s", err.Error())
	}

	watcher := itemwatcher.NewWatcher(gr, rules)

	ctx := contextWithSigterm(context.Background())
	err = watcher.Start(ctx)
	if err != nil {
		log.Fatalf("error during process: %s", err.Error())
	}
}

func contextWithSigterm(ctx context.Context) context.Context {
	ctxWithCancel, cancel := context.WithCancel(ctx)
	go func() {
		defer cancel()

		signalCh := make(chan os.Signal, 1)
		signal.Notify(signalCh, os.Interrupt)

		select {
		case <-signalCh:
		case <-ctx.Done():
		}
	}()

	return ctxWithCancel
}
