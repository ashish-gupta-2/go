package main

import (
	"embed"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	config "ashish.com/m/cmd/config"
	"ashish.com/m/internal/assets"
	app "ashish.com/m/pkg/app"
	log "github.com/sirupsen/logrus"
)

//go:embed data/*
var embededAssets embed.FS

func main() {
	// print application banner
	footerText := fmt.Sprintf("api-%s", app.Version().Version)
	fmt.Println(app.CreateBanner(footerText))

	// initialize log config and create root context
	app.InitLog(app.LogLevelDebug, os.Stdout, true)
	ctx := app.NewRootContext()

	// initialize all embedded assets
	mappedAssets := readAssets("data", embededAssets)
	err := assets.Init(ctx, mappedAssets)
	if err != nil {
		log.WithContext(ctx).Fatalf("data initialization error: %v", err)
	}

	// create and initialize the application
	config := config.New(config.NewConfig())
	config.Init(ctx)

	// start the application
	go func() {
		config.Run(ctx)
	}()

	// configure linux intrupt signal
	chQuit := make(chan os.Signal, 1)
	signal.Notify(chQuit, syscall.SIGINT, syscall.SIGTERM)

	// wait for the intrupt signal, and shutdown the application
	<-chQuit
	log.WithContext(ctx).Info("caught os signal, gracefully stopping the app")
	config.Shutdown(ctx)
}

func readAssets(parent string, fs embed.FS) map[string][]byte {
	tmpMap := make(map[string][]byte)
	dirs, _ := fs.ReadDir(parent)
	for _, dir := range dirs {
		path := fmt.Sprintf("%s/%s", parent, dir.Name())
		if dir.IsDir() {
			for p, content := range readAssets(path, fs) {
				tmpMap[p] = content
			}
		} else {
			content, _ := fs.ReadFile(path)
			tmpMap[path] = content
		}
	}
	return tmpMap
}
