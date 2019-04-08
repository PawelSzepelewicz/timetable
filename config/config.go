package config

import (
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	TelegramToken string `json:"telegram_token"`
	DatabasePath string `json:"database_path"`
	TodayTemplate string `json:"today_template"`
	NextdayTemplate string `json:"nextday_template"`
	TodayWeekdayTemplate string `json:"today_weekday_template"`
	NextdayWeekdayTemplate string `json:"nextday_weekday_template"`
	OracleTemplate string `json:"oracle_template"`
}

func New(path string) (*Config, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := &Config{}

	err = json.Unmarshal(file, c)
	if err != nil {
		return nil, err
	}
	
	return c, nil
}
