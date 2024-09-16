package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/yahn1ukov/personal-blog/internal/app"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "configs/config.yaml", "Config for api")
	flag.Parse()
}

func main() {
	signals := []os.Signal{
		os.Kill,
		os.Interrupt,
	}

	ch := make(chan os.Signal, 1)

	signal.Notify(ch, signals...)

	application := app.New(configPath)

	if err := application.Start(context.Background()); err != nil {
		log.Fatalln(err)
	}

	<-ch

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	if err := application.Stop(ctx); err != nil {
		log.Fatalln(err)
	}
}
