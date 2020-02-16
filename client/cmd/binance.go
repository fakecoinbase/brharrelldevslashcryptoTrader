package cmd

import (
	"context"
	"fmt"
	api_binance "github.com/brharrelldev/cryptoTrader/api/binance"
	"github.com/brharrelldev/cryptoTrader/ct_config"
	"github.com/brharrelldev/cryptoTrader/exchanges/binance_api/general"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

const (
	baseURL = "https://api.binance.com"
)

//BinanceCmd is for the binance sub-command
var BinanceCmd cli.Command

func init() {

	BinanceCmd.Name = "binance"
	BinanceCmd.Usage = "binance <action>"
	BinanceCmd.Subcommands = []cli.Command{
		{
			Name:    "server-time",
			Aliases: []string{"st"},
			Action:  serverTimeAction,
		},
		{
			Name:   "ping",
			Action: pingAction,
		},
		{
			Name:    "market-depth",
			Aliases: []string{"md"},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "symbol",
					Usage:    "market-depth --symbol=<trading pairs>",
					Required: true,
				},
				cli.IntFlag{
					Name:  "limit",
					Usage: "display depth limit",
				},
			},
			Action: getMarketDepth,
		},
	}
}

func pingAction(c *cli.Context) error {
	conf := &ct_config.Config{
		BinanceConfig: &ct_config.BinanceConfig{
			BaseURL: baseURL,
		},
	}
	ping, err := general.NewGeneralAPI(conf)
	if err != nil {
		return fmt.Errorf("could not instantiate general api")
	}

	pingResp, err := ping.GetPing()
	if err != nil {
		return fmt.Errorf("error occurred when checking server time %v", err)
	}

	pingJson, err := pingResp.ToJson()
	if err != nil {
		return fmt.Errorf("error serializing json object %v", err)
	}

	fmt.Println(pingJson)

	return nil
}

func serverTimeAction(c *cli.Context) error {
	conf := &ct_config.Config{
		BinanceConfig: &ct_config.BinanceConfig{
			BaseURL: baseURL,
		},
	}

	st, err := general.NewGeneralAPI(conf)
	if err != nil {
		return fmt.Errorf("could not instantiate general api")
	}

	stResp, err := st.CheckServiceTime()
	if err != nil {
		return fmt.Errorf("error occurred when checking server time %v", err)
	}

	stJson, err := stResp.ToJson()
	if err != nil {
		return fmt.Errorf("error serializing json object %v", err)
	}

	fmt.Println(stJson)

	return nil
}

func getMarketDepth(c *cli.Context) error {

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("could not dial grpc server %v", err)
	}

	defer conn.Close()

	client := api_binance.NewBinanceMarketDataClient(conn)

	getMarketDepth, err := client.GetBinanceMarketDepth(context.Background(), &api_binance.GetBinanceMarketDepthRequest{
		Symbol: c.String("symbol"),
		Limit:  int32(c.Int("limit")),
	})

	fmt.Println(getMarketDepth)

	return nil

}
