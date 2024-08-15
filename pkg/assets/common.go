package assets

import (
	"algo_trading/net"
	"encoding/csv"
	"fmt"
	"strconv"
	"time"
)

func getFromYahoo(s *Stock_Data) error {
	//var prices []yahoo_price
	var price Prices

	yahoourl := fmt.Sprintf("https://query1.finance.yahoo.com/v7/finance/download/%s?period1=%s&period2=%s&interval=1d&events=history&includeAdjustedClose=true", s.Name, fmt.Sprintf("%d", s.Start_Time.Unix()), fmt.Sprintf("%d", s.End_Time.Unix()))
	resp, err := net.Get(yahoourl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	//fill price
	for i, record := range records {
		if i == 0 {
			continue //跳过表格第一行
		}
		//[Date Open High Low Close Adj Close Volume]
		//...
		price.Date, err = time.Parse("2006-01-02", record[0])
		if err != nil {
			return err
		}

		price.Open, err = strconv.ParseFloat(record[1], 64)
		if err != nil {
			return err
		}

		price.High, err = strconv.ParseFloat(record[2], 64)
		if err != nil {
			return err
		}

		price.Low, err = strconv.ParseFloat(record[3], 64)
		if err != nil {
			return err
		}

		price.Close, err = strconv.ParseFloat(record[4], 64)
		if err != nil {
			return err
		}

		price.Volume, err = strconv.ParseInt(record[6], 10, 64)
		if err != nil {
			return err
		}
		s.Stock_Prices = append(s.Stock_Prices, price)
	}
	// for _, v := range s.Stock_Prices {
	// 	fmt.Println(v)
	// }
	return nil
}

func getFromBloomberg(s *Stock_Data) error {
	return nil
}
