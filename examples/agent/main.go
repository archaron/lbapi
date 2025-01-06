package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/im-kulikov/gonfig"

	"github.com/archaron/lbapi"
	"github.com/archaron/lbapi/events"
)

var mskLocation, _ = time.LoadLocation("Europe/Moscow")

func main() {
	var (
		lb  *lbapi.Client
		err error
	)

	// Set location for stringers
	time.Local = mskLocation

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// Use LB_API_ADDRESS, LB_API_USERNAME,LB_API_PASSWORD environment variables, or specify credentials here:
	config := lbapi.ClientConfig{}

	if err := gonfig.New(gonfig.Config{}).Load(&config); err != nil {
		panic(err)
	}

	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   false,
		Level:       slog.LevelDebug,
		ReplaceAttr: nil,
	}))

	lb = lbapi.NewClient(config, log)

	err = lb.ConnectAndLogin(ctx)

	if err != nil {
		panic(err)
	}

	defer func(lb *lbapi.Client) {
		if err := lb.Close(); err != nil {
			panic(err)
		}
	}(lb)

	options, err := lb.GetSyncOptions(context.Background(), lbapi.GetSyncOptionsRequest{
		AgentID: 1,
	})

	if err != nil {
		panic(err)
	}

	spew.Dump(options)

	ok, err := lb.Subscribe(ctx, events.AllLBEvents)
	if err != nil {
		panic(err)
	}

	if !ok {
		panic("Unable to subscribe")
	}

	ch := lb.GetEventsChannel()

	for {
		select {
		case <-ctx.Done():
			return
		case e := <-ch:
			spew.Dump(e)
		}
	}
}
