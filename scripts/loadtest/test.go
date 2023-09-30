package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func main() {
	rate := vegeta.Rate{Freq: 10000, Per: time.Second}
	duration := 30 * time.Second

	data := url.Values{}
	data.Set("uid", "b563feb7b2b84b6test")

	header := http.Header{}
	header.Add("Accept", "application/x-www-form-urlencoded")
	header.Add("Content-Type", "application/x-www-form-urlencoded")

	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    "http://localhost:8186/order",
		Body:   []byte(data.Encode()),
		Header: header,
	})
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()

	fmt.Println("Время мин / сред / макс (µs - микросек) :", metrics.Latencies.Min,
		metrics.Latencies.Mean, metrics.Latencies.Max)
	fmt.Println("Status Codes:", metrics.StatusCodes)
	fmt.Println("Кол-во запросов:", metrics.Requests)
}
