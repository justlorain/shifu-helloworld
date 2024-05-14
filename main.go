package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	BaseURL         = "http://deviceshifu-plate-reader.deviceshifu.svc.cluster.local"
	PullingInterval = time.Second * 5
)

func main() {
	ticker := time.NewTicker(PullingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			nums, err := GetMeasurement()
			if err != nil {
				fmt.Printf("Error get measurement: %v\n", err.Error())
				continue
			}
			avg := CalculateAverage(nums)
			fmt.Printf("Microplate reader measurement average: %v\n", avg)
		}
	}
}

func GetMeasurement() ([]float64, error) {
	url := BaseURL + "/get_measurement"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fields := strings.Fields(string(body))
	// as spec in doc, res is 8*12
	res := make([]float64, len(fields))
	for i, field := range fields {
		v, err := strconv.ParseFloat(field, 64)
		if err != nil {
			return nil, err
		}
		res[i] = v
	}
	return res, nil
}

func CalculateAverage(nums []float64) float64 {
	var sum float64
	for _, num := range nums {
		sum += num
	}
	return sum / float64(len(nums))
}
