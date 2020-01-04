package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/ini.v1"
)

// ここでAPIkey/ApiSecret/LogFileを読み込む -> 構造体
type ConfigList struct {
	ApiKey      string
	ApiSecret   string
	LogFile     string
	ProductCode string

	TradeDuration time.Duration
	Durations     map[string]time.Duration
	DbName        string
	SQLDriver     string
	Port          int
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("./config.ini") // configの読み込み
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1) // logfileが読み込めない場合はエラーコード1で抜ける
	}

	// 時間によって設定を使い分けるための変数
	durations := map[string]time.Duration{
		"1s": time.Second,
		"1m": time.Minute,
		"1h": time.Hour,
	}

	// configより各種keyの取得
	Config = ConfigList{
		ApiKey:        cfg.Section("bitflyer").Key("api_key").String(),
		ApiSecret:     cfg.Section("bitflyer").Key("api_secret").String(),
		LogFile:       cfg.Section("gofintech").Key("log_file").String(), //トレードの履歴や情報をログとして残しておく
		ProductCode:   cfg.Section("gofintech").Key("product_code").String(),
		Durations:     durations,
		TradeDuration: durations[cfg.Section("gofintech").Key("trade_duration").String()],
		DbName:        cfg.Section("db").Key("name").String(),
		SQLDriver:     cfg.Section("db").Key("driver").String(),
		Port:          cfg.Section("web").Key("port").MustInt(),
	}
}
