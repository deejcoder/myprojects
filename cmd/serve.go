package cmd

import (
	"context"
	"os"
	"os/signal"

	"github.com/deejcoder/myprojects/api"
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

		// cancel when api exits e.g unexpectedly
		defer cancel()
		api.Start(ctx)

		return nil
	},
}
