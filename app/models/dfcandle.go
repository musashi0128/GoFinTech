package models

import (
	"time"
)

// チャートのキャンドルスティックの構造体
// 配列のような挙動をする参照型 -> スライスは動的である
// ここでは送られてくるCandleを随時保管する
type DataFrameCandle struct {
	ProductCode string        `json:"product_code"`
	Duration    time.Duration `json:"duration"`
	Candles     []Candle      `json:"candles"`
	Smas        []Sma         `json:"smas,omitempty"`
	Emas        []Ema         `json:"emas,omitempty"`
}

// チャートの単純移動平均線の構造体
type Sma struct {
	Period int       `json:"period,omitempty"`
	Values []float64 `json:"values,omitempty"`
}

// チャートの指数平滑移動平均線の構造体
type Ema struct {
	Period int       `json:"period,omitempty"`
	Values []float64 `json:"values,omitempty"`
}

func (df *DataFrameCandle) Times() []time.Time {
	s := make([]time.Time, len(df.Candles))
	for i, candle := range df.Candles {
		s[i] = candle.Time
	}
	return s
}

func (df *DataFrameCandle) Opens() []float64 {
	s := make([]float64, len(df.Candles))
	for i, candle := range df.Candles {
		s[i] = candle.Open
	}
	return s
}

func (df *DataFrameCandle) Closes() []float64 {
	s := make([]float64, len(df.Candles))
	for i, candle := range df.Candles {
		s[i] = candle.Close
	}
	return s
}

func (df *DataFrameCandle) Highs() []float64 {
	s := make([]float64, len(df.Candles))
	for i, candle := range df.Candles {
		s[i] = candle.High
	}
	return s
}

func (df *DataFrameCandle) Low() []float64 {
	s := make([]float64, len(df.Candles))
	for i, candle := range df.Candles {
		s[i] = candle.Low
	}
	return s
}

func (df *DataFrameCandle) Volume() []float64 {
	s := make([]float64, len(df.Candles))
	for i, candle := range df.Candles {
		s[i] = candle.Volume
	}
	return s
}

// DataFrameCandleにデータを追加する
// smaのデータをtalib.Smaで算出して追加していく
// func (df *DataFrameCandle) AddSma(period int) bool {
// 	if len(df.Candles) > period {
// 		df.Smas = append(df.Smas, Sma{
// 			Period: period,
// 			Values: talib.Sma(df.Closes(), period),
// 		})
// 		return true
// 	}
// 	return false
// }

// // smaのデータをtalib.Emaで算出して追加していく
// func (df *DataFrameCandle) AddEma(period int) bool {
// 	if len(df.Candles) > period {
// 		df.Emas = append(df.Emas, Ema{
// 			Period: period,
// 			Values: talib.Ema(df.Closes(), period),
// 		})
// 		return true
// 	}
// 	return false
// }
