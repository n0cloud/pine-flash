/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/N0Cloud/pine-flash/blclient"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pine-flash",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		portName, err := cmd.Flags().GetString("port")
		if err != nil {
			return err
		}

		var debug bool
		if d, err := cmd.Flags().GetBool("debug"); err != nil {
			debug = d
		}
		l := newLogger(debug)

		client, err := blclient.Dial(portName, blclient.WithLogger(l))
		if err != nil {
			return err
		}
		defer client.Close()

		l.Info("handshake success")

		_, err = client.GetBootInfo()
		if err != nil {
			l.Error(err)
		}

		_, err = client.ReadJedecId()
		if err != nil {
			l.Error(err)
		}

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	flags := rootCmd.Flags()

	flags.String("port", "", "serial port to use")
	rootCmd.MarkFlagRequired("port")
	flags.Bool("debug", false, "enable debug mode")
}
