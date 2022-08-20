package main

import (
	"encoding/json"
	"io"
	"os"
)

var carList []car

type car struct {
	Name         string `json:"name"`
	Speed        int    `json:"speed"`
	RacingNumber int    `json:"racing_number"`
	Country      string `json:"country"`
	RacingType   string `json:"racing_type"`
}

func getCars() error {
	file, err := os.Open("cars.json")
	defer file.Close()
	if err != nil {
		return err
	}

	byteVal, _ := io.ReadAll(file)
	json.Unmarshal(byteVal, &carList)
	return nil
}
