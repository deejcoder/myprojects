package cmd

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"github.com/Dilicor/myprojects/api"
	"github.com/Dilicor/myprojects/storage"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCommand)
}

var serveCommand = &cobra.Command{
	Use:   "serve",
	Short: "serves the REST API",
	RunE: func(cmd *cobra.Command, args []string) error {

		// create a context which can be cancelled
		ctx, cancel := context.WithCancel(context.Background())

		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, os.Interrupt)
			<-ch
			log.Info("signal caught. shutting down...")
			// cancel context when user sends interupt signal (CTRL + C); waits until signal received
			cancel()
		}()

		var waitGroup sync.WaitGroup
		waitGroup.Add(2)

		go func() {
			defer waitGroup.Done()
			defer cancel()

			api.Serve(ctx)
		}()

		go func() {
			defer waitGroup.Done()
			defer cancel()

			storage.Connect(ctx)
		}()

		waitGroup.Wait()

		return nil
	},
}
