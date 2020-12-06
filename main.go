package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	"time"

	"github.com/toyo/go-latlong"
)

type Card struct {
	Id string `json:"id"`
	PropertyId string `json:"property_id"` 
	Price int `json:"price"`
	Date time.Time  `json:"date"`
}

type CardList struct {
	Cards []Card  `json:"cards"`
}

func main() {
	resp, err := http.Get("https://gateway.homes.co.nz/properties/nearby?lat=-41.2804198667&long=174.762992733&limit=48")
	if err != nil {
		log.Fatal("Error getting response. ", err)
	}
	defer resp.Body.Close()
 
	// Read body from response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}

	cardList := &CardList{}
	json.Unmarshal([]byte(body), cardList)
	

	fmt.Printf("%s\n", cardList)
}