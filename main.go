package main

import (
	"bitconnect_API/twilio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {

	t := twilio.NewTwilioAPI()
	var resource Bitconnect
	lcount := 0
	hcount := 0
	for {
		bitResp, err := http.Get("https://bitconnect.co/api/info/BTC_BCC")
		if err != nil {
			log.Fatalf("could not connecct to bitconnet", err)
		}
		bitdecoder := json.NewDecoder(bitResp.Body)
		err = bitdecoder.Decode(&resource)
		if err != nil {
			log.Fatalf("could no json decode from bitconnect bitResp.Body", err)
		}
		//ParseFloat(s string, bitSize int) (float64, error)
		Bid, _ := strconv.ParseFloat(resource.Markets[0].Bid, 64)
		Lastprice, _ := strconv.ParseFloat(resource.Markets[0].Lastprice, 64)
		Volume24h, _ := strconv.ParseFloat(resource.Markets[0].Volume24h, 64)
		Currency := resource.Markets[0].Currency
		Marketname := resource.Markets[0].Marketname
		Ask, _ := strconv.ParseFloat(resource.Markets[0].Ask, 64)
		Low24h, _ := strconv.ParseFloat(resource.Markets[0].Low24h, 64)
		High24h, _ := strconv.ParseFloat(resource.Markets[0].High24h, 64)
		Change24h := High24h - Low24h
		Basecurrency := resource.Markets[0].Basecurrency

		fmt.Println("Bid : ", Bid)
		fmt.Println("Lastprice : ", Lastprice)
		fmt.Println("Volume24h : ", Volume24h)
		fmt.Println("Currency : ", Currency)
		fmt.Println("Marketname : ", Marketname)
		fmt.Println("Ask : ", Ask)
		fmt.Println("Low24h : ", Low24h)
		fmt.Println("Change24h : ", Change24h)
		fmt.Println("High24h : ", High24h)
		fmt.Println("Basecurrency : ", Basecurrency)

		if (Lastprice <= Low24h) && (Change24h >= 0.00047) {
			message := fmt.Sprintf("The price of bitconnet coin currently %f has fallen low by %f in the last 24hours", Lastprice, Change24h)
			fmt.Println(message)
			//_ = t
			t.TwilioTextMessage(message)
			lcount++
		}

		if (Lastprice >= High24h) && (Change24h >= 0.00047) {
			message := fmt.Sprintf("The price of bitconnet coin currently %f has risen high by %f in the last 24hours", Lastprice, Change24h)
			fmt.Println(message)
			//_ = t
			t.TwilioTextMessage(message)
			hcount++
		}

		time.Sleep(time.Minute * 10)
		if (lcount == 2) || (hcount == 2) {
			time.Sleep(time.Minute * 60 * 12)
			lcount = 0
			hcount = 0
		}
	}
}

type Bitconnect struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Markets []Market `json:"markets"`
}
type Market struct {
	Bid          string `json:"bid"`
	Lastprice    string `json:"last_price"`
	Volume24h    string `json:"volume24h"`
	Currency     string `json:"currency"`
	Marketname   string `json:"marketname"`
	Ask          string `json:"ask"`
	Low24h       string `json:"low24h"`
	Change24h    string `json:"change24h"`
	High24h      string `json:"high24h"`
	Basecurrency string `json:"basecurrency"`
}
