package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/subhasishgoswami/superform/common"
	"github.com/subhasishgoswami/superform/server"
	"github.com/subhasishgoswami/superform/smartcontract"
	"github.com/subhasishgoswami/superform/subgraph"
)

var (
	apiListenAddr        string
	subgraphURL          string
	subgraphAPIKey       string
	rpcURL               string
	smartContractAddress string
)

func init() {
	rootCmd.AddCommand(apiCmd)

	apiCmd.Flags().StringVar(&apiListenAddr, "listen-addr", "localhost:9000", "listen address for webserver")

	apiCmd.Flags().StringVar(&subgraphURL, "subgraph-url", "https://api.thegraph.com/subgraphs/name/subhasishgoswami/superform", "Subgraph URL")
	apiCmd.Flags().StringVar(&subgraphAPIKey, "subgraph-api-key", "", "Subgraph API Key")

	apiCmd.Flags().StringVar(&rpcURL, "rpc-url", "https://endpoints.omniatech.io/v1/bsc/mainnet/public", "RPC URL")
	apiCmd.Flags().StringVar(&smartContractAddress, "smart-contract-address", "0x889447dFab0611F714a12306cB66392EC38d2211", "Smart Contract Addres")
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start the API server",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		superformSubgraph := subgraph.NewSuperFormInterface(subgraphURL, subgraphAPIKey)
		log := common.LogSetup(false, "debug").WithFields(logrus.Fields{
			"service": "server/api",
			"version": "0.1",
		})
		smartContract := smartcontract.NewGethInterface(rpcURL, smartContractAddress)
		log.Info("Starting Server")
		api := server.NewServer(apiListenAddr, *superformSubgraph, smartContract, *log)

		err = api.StartServer()
		if err != nil {
			log.WithError(err).Fatal("server error")
		}
	},
}
