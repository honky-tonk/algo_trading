package assets

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Stocks_Data struct {
	Name          string    `yaml:"Name"`
	Start_Time    time.Time `yaml:"Start_Time"`
	End_Time      time.Time `yaml:"End_Time"`
	Time_Interval string    `yaml:"Time_Interval"`
}

type Stocks struct {
	Datas []Stocks_Data `yaml:"Stock_Datas"`
}

func ParseConfig() *Stocks {
	//read config from local
	file, err := os.ReadFile("config/config.yaml")
	// out, err := exec.Command("pwd").Output()
	// fmt.Println("current work dir is: ", string(out))
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
