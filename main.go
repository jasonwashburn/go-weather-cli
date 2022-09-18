package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type owmResponse struct {
	Main map[string]interface{} `json:"main"`
}

func main() {

    zipCode := os.Args[1]
    if zipCode == "" {
        log.Fatal("You must enter a zip code")
    }
	countryCode := "us"
	var appID string = os.Getenv("OWM_API_KEY")
	var address string = "https://api.openweathermap.org/data/2.5/weather?zip=" + zipCode + "," + countryCode + "&appid=" + appID
	fmt.Println(address)
	resp, respErr := http.Get(address)
	if respErr != nil {
		log.Fatal(respErr)
	}
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	respObj := owmResponse{}
	jsonErr := json.Unmarshal(body, &respObj)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	for key, value := range respObj.Main {
		fmt.Println("Key: ", key, "=> Value: ", value)
	}
}
