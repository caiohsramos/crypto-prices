package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/mitchellh/mapstructure"
)

type CMCService struct {
	client *resty.Client
}

type Quote struct {
	ID                int         `mapstructure:"id"`
	Name              string      `mapstructure:"name"`
	Symbol            string      `mapstructure:"symbol"`
	Slug              string      `mapstructure:"slug"`
	IsActive          int         `mapstructure:"is_active"`
	IsFiat            int         `mapstructure:"is_fiat"`
	CirculatingSupply int         `mapstructure:"circulating_supply"`
	TotalSupply       int         `mapstructure:"total_supply"`
	MaxSupply         int         `mapstructure:"max_supply"`
	DateAdded         string      `mapstructure:"date_added"` // TODO: Treat as time.Time
	NumMarketPairs    int         `mapstructure:"num_market_pairs"`
	CmcRank           int         `mapstructure:"cmc_rank"`
	LastUpdated       string      `mapstructure:"last_updated"` // TODO: Treat as time.Time
	Tags              []string    `mapstructure:"tags"`
	Platform          interface{} `mapstructure:"platform"`
	Quote             struct {
		Usd struct {
			Price            float64 `mapstructure:"price"`
			Volume24H        float64 `mapstructure:"volume_24h"`
			PercentChange1H  float64 `mapstructure:"percent_change_1h"`
			PercentChange24H float64 `mapstructure:"percent_change_24h"`
			PercentChange7D  float64 `mapstructure:"percent_change_7d"`
			PercentChange30D float64 `mapstructure:"percent_change_30d"`
			MarketCap        float64 `mapstructure:"market_cap"`
			LastUpdated      string  `mapstructure:"last_updated"` // TODO: Treat as time.Time
		} `mapstructure:"USD"`
	} `mapstructure:"quote"`
}

func NewCMCService(apiKey string) *CMCService {
	client := resty.New()
	client.SetHostURL("https://pro-api.coinmarketcap.com")
	client.SetHeader("X-CMC_PRO_API_KEY", apiKey)

	return &CMCService{client: client}
}

func (r *CMCService) GetQuotes(symbols []string) ([]*Quote, error) {
	symbolsList := strings.Join(symbols, ",")

	resp, err := r.client.R().SetQueryParam("symbol", symbolsList).Get("/v1/cryptocurrency/quotes/latest")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Status code: %d\nMessage: %s", resp.StatusCode(), resp.String()))
	}
	var jsonResponse map[string]interface{}
	err = json.Unmarshal(resp.Body(), &jsonResponse)
	if err != nil {
		return nil, err
	}

	var quotes []*Quote
	for _, item := range jsonResponse["data"].(map[string]interface{}) {
		quote := &Quote{}
		err = mapstructure.Decode(item, quote)
		if err != nil {
			return nil, err
		}

		quotes = append(quotes, quote)
	}

	return quotes, nil
}
