package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const token = "bfU1dADSYq1a7NHfgBrS2E"

type APIResponse struct {
	Results []struct {
		Price float64 `json:"regularMarketPrice"`
	} `json:"results"`
}

type StockResult struct {
	Ticker string
	Price  float64
	Error  error
}

func fetchStock(ticker string) StockResult {
	var data APIResponse

	url := fmt.Sprintf("https://brapi.dev/api/quote/%s?token=%s", ticker, token)
	resp, err := http.Get(url)
	if err != nil {
		return StockResult{Ticker: ticker, Error: fmt.Errorf("error fetching data for %s: %s", ticker, err)}
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return StockResult{Ticker: ticker, Error: fmt.Errorf("error decoding data for %s: %s", ticker, err)}
	}

	if len(data.Results) == 0 {
		return StockResult{Ticker: ticker, Error: fmt.Errorf("no data found for %s", ticker)}
	}

	price := data.Results[0].Price
	return StockResult{Ticker: ticker, Price: price, Error: nil}
}

func main() {
	startNow := time.Now()
	defer func() { fmt.Printf("This operation took: %s\n", time.Since(startNow)) }()

	tickers := []string{"MGLU3", "PETR4", "KLBN4", "VALE3", "ITUB4", "BBDC4", "BBAS3", "ABEV3", "B3SA3", "PETR3", "VVAR3", "GNDI3", "LREN3", "NTCO3", "RENT3", "WEGE3", "IRBR3", "CSNA3", "SUZB3", "JBSS3", "BRFS3", "CIEL3", "ELET3", "ELET6", "CMIG4", "ENBR3", "CPFE3", "SBSP3", "BBSE3", "ITSA4", "ITUB3", "BBDC3", "BBAS3", "SANB11", "BPAC11", "BIDI11", "BRML3", "BRPR3", "BRKM5", "BRAP4", "BRDT3", "BRML3", "BRPR3", "BRKM5", "BRAP4", "BRDT3", "BRFS3", "BRKM5", "BRML3", "BRPR3", "BRSR6", "B3SA3", "BTOW3", "CCRO3", "CIEL3", "CMIG4", "COGN3", "CPFE3", "CPLE6", "CRFB3", "CSAN3", "CSMG3", "CSNA3", "CVCB3", "CYRE3", "DIRR3", "DTEX3", "ECOR3", "EGIE3", "ELET3", "ELET6", "EMBR3", "ENBR3", "ENEV3", "ENGI11", "EQTL3", "EZTC3", "FLRY3", "GGBR4", "GNDI3", "GOAU4", "GOLL4", "GRND3", "GUAR3", "HAPV3", "HGTX3", "HYPE3", "IGTA3", "IRBR3", "ITSA4", "ITUB4", "JBSS3", "JHSF3", "KLBN4"}

	for _, ticker := range tickers {
		result := fetchStock(ticker)
		if result.Error != nil {
			fmt.Printf("Error for ticker %s: %s\n", result.Ticker, result.Error)
		} else {
			fmt.Printf("Price of stock %s: R$ %.2f\n", result.Ticker, result.Price)
		}
	}
}
