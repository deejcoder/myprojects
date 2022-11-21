package cmd

import (
	"os"

	"github.com/deejcoder/myprojects/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "myprojects",
	Short: "myprojects provides a RESTful API for the frontend (made in React), to showcase our projects",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

/*
Execute is a wrapper for the command executor; executing commands
*/
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

// init : before a command is executed
func init() {
	config.InitConfig()
}
