package alphavantage_client

import (
	go_logger "github.com/pefish/go-logger"
	go_test_ "github.com/pefish/go-test"
	"testing"
)

const apiKey = "UBKBWL7XMB2Y6XQ4"

var logger = go_logger.Logger.CloneWithLevel("debug")

func TestClientType_TreasuryYield(t *testing.T) {
	client := NewClient(logger, apiKey)
	results, err := client.TreasuryYield(&TreasuryYieldParams{
		Maturity: TreasuryYieldMaturityType_5year,
	})
	go_test_.Equal(t, nil, err)
	go_test_.Equal(t, true, len(results) > 0)
}
