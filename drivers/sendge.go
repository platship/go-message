package drivers

import (
	"errors"
)

type Sendge struct {
	APIkey string `json:"apiKey"`
	APIUrl string `json:"apiUrl"`
}

func (s *Sendge) Init() (err error) {
	if s.APIkey == "" {
		return errors.New("API Key is undefined")
	}
	if s.APIUrl == "" {
		return errors.New("API Url is undefined")
	}
	return
}

func (s *Sendge) Send() error {
	// 发送信息
	return nil
}
