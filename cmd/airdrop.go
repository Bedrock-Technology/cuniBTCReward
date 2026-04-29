/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cuniBTCReward/pkg/slack"
	"cuniBTCReward/service/airdrop"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

// airdropCmd represents the airdrop command
var airdropCmd = &cobra.Command{
	Use:   "airdrop",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Run: func(cmd *cobra.Command, args []string) {},
}

var block uint64
var chainId uint
var checkBalance bool
var contracts []string
var token string

var file string
var epoch uint64
var contract string
var dryRun bool

func init() {
	rootCmd.AddCommand(airdropCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// airdropCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// airdropCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	airdropCmd.AddCommand(sharesCmd)
	sharesCmd.Flags().Uint64VarP(&block, "block", "b", 0, "block number to check balance at, default is latest")
	sharesCmd.Flags().UintVarP(&chainId, "chainid", "c", 1, "chain ID to check balance on, default is 1 for Ethereum mainnet")
	sharesCmd.Flags().BoolVarP(&checkBalance, "checkbalance", "k", true, "whether to check token balance")
	sharesCmd.Flags().StringSliceVarP(&contracts, "contracts", "r", []string{}, "list of contracts to check balance for")
	sharesCmd.Flags().StringVarP(&token, "token", "t", "", "token address to check balance for")

	airdropCmd.AddCommand(airdropCreateCmd)
	airdropCreateCmd.Flags().UintVarP(&chainId, "chainid", "c", 1, "chain ID to check balance on, default is 1 for Ethereum mainnet")
	airdropCreateCmd.Flags().StringVarP(&contract, "contract", "r", "", "contract address of airdrop")
	_ = airdropCreateCmd.MarkFlagRequired("contract")
	airdropCreateCmd.Flags().StringVarP(&file, "file", "f", "", "path to the file containing the airdrop data")
	_ = airdropCreateCmd.MarkFlagRequired("file")
	airdropCreateCmd.Flags().Uint64VarP(&epoch, "epoch", "e", 0, "epoch number for the airdrop")
	_ = airdropCreateCmd.MarkFlagRequired("epoch")
	airdropCreateCmd.Flags().BoolVarP(&dryRun, "dryrun", "d", false, "whether to run in dry run mode")

}

// sharesCmd represents the shares command
var sharesCmd = &cobra.Command{
	Use:   "shares",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
	
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("airdrop shares called")
		var c ServiceConfig
		conf.MustLoad(cfgFile, &c)
		c.AirdropConf.MustSetUp()
		// fmt.Printf("evmscanconfi:%v", c.EvmScanConf)

		//log
		if c.LogSlack != "" {
			logx.AddWriter(logx.NewWriter(slack.NewSlackWriter(c.LogSlack)))
			logx.AddGlobalFields(logx.Field("server", c.AirdropConf.Name))
			defer logx.Close()
		}

		airdropIns := airdrop.NewAirdrop(&c.AirdropConf)
		getBlock, err := airdropIns.GetCursor(chainId)
		if err != nil {
			slack.SendTo(c.AirdropConf.NotifySlack, fmt.Sprintf("[%s] %v", c.AirdropConf.Name, err))
			fmt.Println(err)
			return
		}
		if block == 0 {
			block = getBlock.BlockNumber
		}
		if block > getBlock.BlockNumber {
			fmt.Printf("block number [%d] is greater than current cursor block number [%d], no need to check\n", block, getBlock.BlockNumber)
			return
		}
		fmt.Printf("Block: %d, Chain ID: %d, Check Balance: %t\n", block, chainId, checkBalance)
		result, err := airdropIns.GetAllAddressAtBlock(block, chainId, checkBalance, contracts, token)
		if err != nil {
			slack.SendTo(c.AirdropConf.NotifySlack, fmt.Sprintf("[%s] %v", c.AirdropConf.Name, err))
			fmt.Println(err)
			return
		}
		resultJson, err := json.Marshal(&result)
		if err != nil {
			slack.SendTo(c.AirdropConf.NotifySlack, fmt.Sprintf("[%s] %v", c.AirdropConf.Name, err))
			fmt.Println(err)
			return
		}
		fmt.Println(string(resultJson))
	},
}

type AirdropInfo struct {
	Address string
	Shares  string
	Amount  string
}

// airdropCreateCmd represents the create command
var airdropCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
	
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("airdrop create called")
		var c ServiceConfig
		conf.MustLoad(cfgFile, &c)
		c.AirdropConf.MustSetUp()
		// fmt.Printf("evmscanconfi:%v", c.EvmScanConf)

		//log
		if c.LogSlack != "" {
			logx.AddWriter(logx.NewWriter(slack.NewSlackWriter(c.LogSlack)))
			logx.AddGlobalFields(logx.Field("server", c.AirdropConf.Name))
			defer logx.Close()
		}

		fileHandle, err := os.Open(file)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer fileHandle.Close()
		info := make([]AirdropInfo, 0)
		reader := csv.NewReader(fileHandle)
		for {
			record, err := reader.Read()
			if err != nil {
				if err != io.EOF {
					fmt.Println("Read err:", err)
					return
				}
				break
			}
			// fmt.Println(record) // Each record is a []string
			if len(record) < 3 {
				fmt.Println("Invalid record:", record)
				return
			}
			if record[0] == "" || record[1] == "" || record[2] == "" {
				fmt.Println("Empty address or amount in record:", record)
				return
			}
			if shares, err := decimal.NewFromString(record[1]); err != nil || shares.IsNegative() || shares.IsZero() {
				fmt.Println("Invalid shares in record:", record)
				return
			}
			if amount, err := decimal.NewFromString(record[2]); err != nil || amount.IsNegative() || amount.IsZero() {
				fmt.Println("Invalid amount in record:", record)
				return
			}
			info = append(info, AirdropInfo{
				Address: record[0],
				Shares:  record[1],
				Amount:  record[2],
			})
		}
		if len(info) == 0 {
			fmt.Println("No valid records found in the file")
			return
		}
		leaves := lo.Map(info, func(item AirdropInfo, _ int) airdrop.TreeLeaf {
			return airdrop.TreeLeaf{
				Address: item.Address,
				Amount:  item.Amount,
			}
		})
		shares := lo.Map(info, func(item AirdropInfo, _ int) string {
			return item.Shares
		})

		airdropIns := airdrop.NewAirdrop(&c.AirdropConf)
		root, err := airdropIns.CreateAirdropEpoch(chainId, contract, epoch, leaves, shares, dryRun)
		if err != nil {
			slack.SendTo(c.AirdropConf.NotifySlack, fmt.Sprintf("[%s] %v", c.AirdropConf.Name, err))
			fmt.Println(err)
			return
		}
		fmt.Printf("Airdrop epoch created successfully, length of leaves: %d\n", len(info))
		fmt.Printf("Epoch: %d, Contract: %s, Root: %s\n", epoch, contract, hexutil.Encode(root))
	},
}
