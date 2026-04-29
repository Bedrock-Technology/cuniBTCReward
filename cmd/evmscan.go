/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cuniBTCReward/pkg/slack"
	"cuniBTCReward/service/crons"
	"cuniBTCReward/service/evmscan"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

// evmscanCmd represents the evmscan command
var evmscanCmd = &cobra.Command{
	Use:   "evmscan",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("evmscan called")
		var c ServiceConfig
		conf.MustLoad(cfgFile, &c)
		c.EvmScanConf.MustSetUp()
		// fmt.Printf("evmscanconfi:%v", c.EvmScanConf)

		//log
		if c.LogSlack != "" {
			logx.AddWriter(logx.NewWriter(slack.NewSlackWriter(c.LogSlack)))
			logx.AddGlobalFields(logx.Field("server", c.EvmScanConf.Name))
			defer logx.Close()
		}

		//cron
		crontab := crons.New()
		evmScan := evmscan.NewScanner(&c.EvmScanConf)
		if c.EvmScanConf.LogsScanSpec != "" {
			_, err := crontab.AddFunc(c.EvmScanConf.LogsScanSpec, evmScan.LogScan)
			logx.Must(err)
			logx.Infof("add cron evmlogs scan spec: %v", c.EvmScanConf.LogsScanSpec)
		}

		//servers
		serverGroup := service.NewServiceGroup()
		serverGroup.Add(crontab)
		serverGroup.Start()
		serverGroup.Stop()
	},
}

func init() {
	rootCmd.AddCommand(evmscanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// evmscanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// evmscanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
