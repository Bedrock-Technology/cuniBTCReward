/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	export "cuniBTCReward/api/export"
	exportconfig "cuniBTCReward/api/export/config"
	"cuniBTCReward/pkg/slack"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("api called")
		var c ServiceConfig
		conf.MustLoad(cfgFile, &c)

		//log
		if c.LogSlack != "" {
			logx.AddWriter(logx.NewWriter(slack.NewSlackWriter(c.LogSlack)))
			logx.AddGlobalFields(logx.Field("server", c.ApiConf.Name))
			defer logx.Close()
		}

		server := rest.MustNewServer(c.ApiConf.RestConf,
			rest.WithFileServer("/docs", http.Dir("./api/docs")),
			rest.WithCors("*"))

		defer server.Stop()

		setupConfig := exportconfig.Config{
			Config: c.ApiConf.Config,
		}
		export.Setup(server, setupConfig)

		fmt.Printf("Starting server at %s:%d...\n", c.ApiConf.Host, c.ApiConf.Port)
		server.Start()
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
