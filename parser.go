package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Candle struct{
		Time time.Time   //Время начала свечи	
		Open float64 	 // — цена открытия	
		High float64 	 //— максимальная цена	
		Low float64 	 //— минимальная цена	
		Close float64	 //— цена закрытия
		Volume float64   //— объём	
		Turnover float64 //— оборот USDT
	}

func Parser()(candle Candle, err error){
	
	url := "https://api.bybit.com/v5/market/kline?category=linear&symbol=BTCUSDT&interval=1&limit=1"

	resp, err:= http.Get(url)

	kline, err := io.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil{}

	os.WriteFile("BtcKline.json", kline, 0777)

	type List struct{
		List [][]string `json:"list"`
	}
	type Result struct{
		Result List `json:"result"`
	}

	var res Result

	err = json.Unmarshal(kline, &res)

	candleStr := res.Result.List[0]

	var cand []float64

	for _, num := range candleStr{
		
		c, err := strconv.ParseFloat(num, 64)
		if err != nil{
			return
		}
		cand = append(cand, c)
	}
	candTime := NormTime(cand[0])

	candle = Candle{
		Time: candTime,
		Open: cand[1],
		High: cand[2],
		Low: cand[3],
		Close: cand[4],
		Volume: cand[5],
		Turnover: cand[6],
	}
	fmt.Print(candle)
	return
}

func NormTime(timeMl float64)(data time.Time){

	times := int64(timeMl)
data = time.UnixMilli(times) 

return
}