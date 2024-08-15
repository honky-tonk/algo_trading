package assets

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Prices struct {
	Date   time.Time
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume int64
}

type Stock_Data struct {
	Name          string    `yaml:"Name"`
	DataFrom      string    `yaml:"DataFrom"`
	Start_Time    time.Time `yaml:"Start_Time"`
	End_Time      time.Time `yaml:"End_Time"`
	Time_Interval string    `yaml:"Time_Interval"`

	Stock_Prices []Prices
}

type Stocks struct {
	Datas []Stock_Data `yaml:"Stock_Datas"`
}

func ParseConfig() *Stocks {
	//read config from local
	file, err := os.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("Can't read config.yaml file: %v", err.Error())
	}

	//parse yaml to struct
	var stocks Stocks
	if err = yaml.Unmarshal(file, &stocks); err != nil {
		log.Fatalf("Can't unmarshal yaml file to struct: %v", err.Error())
	}

	return &stocks
}

// for backtest
func (s *Stocks) GetPrice() error {
	//loop all stocks
	for _, v := range s.Datas {
		fmt.Println(v.Name, " datafrom is :", v.DataFrom)
		switch {
		case v.DataFrom == "Yahoo":
			if err := getFromYahoo(&v); err != nil {
				return err
			}
			continue
		case v.DataFrom == "Bloomberg":
			if err := getFromBloomberg(&v); err != nil {
				return err
			}
			continue
		default:
			return errors.New("Could Not find Data Source for Stock")

		}
	}

	return nil
}
