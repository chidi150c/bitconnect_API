package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	//t := twilio.NewTwilioAPI()
	var resource Bitconnect

	bitResp, err := http.Get("https://bitconnect.co/api/info/BTC_BCC")
	if err != nil {
		log.Fatalf("could not connecct to bitconnet", err)
	}
	bitdecoder := json.NewDecoder(bitResp.Body)
	err = bitdecoder.Decode(&resource)
	if err != nil {
		log.Fatalf("could no json decode from bitconnect bitResp.Body", err)
	}

	Bid := resource.Markets[0].Bid
	Lastprice := resource.Markets[0].Lastprice
	Volume24h := resource.Markets[0].Volume24h
	Currency := resource.Markets[0].Currency
	Marketname := resource.Markets[0].Marketname
	Ask := resource.Markets[0].Ask
	Low24h := resource.Markets[0].Low24h
	Change24h := resource.Markets[0].Change24h
	High24h := resource.Markets[0].High24h
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

	//if Lastprice <= Low24h

	//t.TwilioTextMessage("how for")
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
