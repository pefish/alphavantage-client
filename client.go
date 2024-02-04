package alphavantage_client

import (
	go_format "github.com/pefish/go-format"
	go_http "github.com/pefish/go-http"
	go_logger "github.com/pefish/go-logger"
	"time"
)

type ClientType struct {
	logger  go_logger.InterfaceLogger
	apiKey  string
	timeout time.Duration
}

type DataType string

const (
	DataType_Json DataType = "json"
	DataType_Csv  DataType = "csv"
)

type TreasuryYieldMaturityType string

const (
	TreasuryYieldMaturityType_3month TreasuryYieldMaturityType = "3month"
	TreasuryYieldMaturityType_2year  TreasuryYieldMaturityType = "2year"
	TreasuryYieldMaturityType_5year  TreasuryYieldMaturityType = "5year"
	TreasuryYieldMaturityType_7year  TreasuryYieldMaturityType = "7year"
	TreasuryYieldMaturityType_10year TreasuryYieldMaturityType = "10year"
	TreasuryYieldMaturityType_30year TreasuryYieldMaturityType = "30year"
)

type IntervalType string

const (
	IntervalType_Daily   IntervalType = "daily"
	IntervalType_Weekly  IntervalType = "weekly"
	IntervalType_Monthly IntervalType = "monthly"
)

const BASE_URL = "https://www.alphavantage.co/query"

func NewClient(
	logger go_logger.InterfaceLogger,
	apiKey string,
) *ClientType {
	return &ClientType{
		logger:  logger,
		apiKey:  apiKey,
		timeout: 10 * time.Second,
	}
}

func (c *ClientType) SetTimeout(timeout time.Duration) *ClientType {
	c.timeout = timeout
	return c
}

type TreasuryYieldParams struct {
	Interval IntervalType              `json:"interval,omitempty" default:"monthly"`
	Maturity TreasuryYieldMaturityType `json:"maturity,omitempty" default:"10year"` // 默认 10 年期国债收益率
	DataType DataType                  `json:"datatype,omitempty" default:"json"`
}

type TreasuryYieldResult struct {
	Date  string `json:"date"`
	Value string `json:"value"`
}

func (c *ClientType) TreasuryYield(treasuryYieldParams *TreasuryYieldParams) ([]TreasuryYieldResult, error) {
	var result struct {
		Data []TreasuryYieldResult `json:"data"`
	}
	err := c.RequestForStruct(
		"TREASURY_YIELD",
		treasuryYieldParams,
		&result,
	)
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}

func (c *ClientType) RequestForStruct(
	function string,
	params interface{},
	obj interface{},
) error {
	paramsMap := go_format.FormatInstance.StructToMap(params)
	paramsMap["function"] = function
	paramsMap["apikey"] = c.apiKey

	_, _, err := go_http.NewHttpRequester(
		go_http.WithTimeout(c.timeout),
		go_http.WithLogger(c.logger),
	).GetForStruct(
		go_http.RequestParam{
			Url:    BASE_URL,
			Params: paramsMap,
		},
		obj,
	)
	if err != nil {
		return err
	}
	return nil
}
