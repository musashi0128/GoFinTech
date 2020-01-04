package main

import (
	"GoFinTech/app/controllers"
	"GoFinTech/config"
	"GoFinTech/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	controllers.StartWebServer()
}

/*
func main() {
	utils.LoggingSettings(config.Config.LogFile)
	fmt.Println(models.DbConnection)
}
*/

/*
func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	tickerChannel := make(chan bitflyer.Ticker)
	go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannel) // 1.ProductCode, 2.tickerChannel
	for ticker := range tickerChannel {
		fmt.Println(ticker)
		fmt.Println(ticker.GetMidPrice())
		fmt.Println(ticker.DateTime())
		fmt.Println(ticker.TruncateDateTime(time.Hour))
	}
}
*/

/* 売買 or 予約用で実行させる
func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)

	// 1BTCが$7000になったら0.01BTC購入
	order := &bitflyer.Order{
		ProductCode:     config.Config.ProductCode,
		ChildOrderType:  "LIMIT",
		Side:            "BUY",
		Price:           7000,
		Size:            0.01,
		MinuteToExpires: 1,
		TimeInForce:     "GTC",
	}
	// 0.01BTCの購入
	order := &bitflyer.Order{
		ProductCode:     config.Config.ProductCode,
		ChildOrderType:  "MARKET",
		Side:            "BUY",
		Size:            0.01,
		MinuteToExpires: 1,
		TimeInForce:     "GTC",
	}

	// 0.01BTCの売却
	order := &bitflyer.Order{
		ProductCode:     config.Config.ProductCode,
		ChildOrderType:  "MARKET",
		Side:            "SELL",
		Size:            0.01,
		MinuteToExpires: 1,
		TimeInForce:     "GTC",
	}

	res, _ := apiClient.SendOrder(order)
	fmt.Println(res.ChildOrderAcceptanceID)
}
*/
