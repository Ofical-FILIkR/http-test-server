package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ApiResponseIP struct {
	Country    string `json:"country"`
	RegionName string `json:"regionName"`
	City       string `json:"city"`
	Query      string `json:"query"`
	Mobile     bool   `json:"mobile"`
	Proxy      bool   `json:"proxy"`
	Hosting    bool   `json:"hosting"`
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Пожалуйста, передайте хотя бы один аргумент.")
		return
	}

	arg := os.Args[1]

	url := "http://ip-api.com/json/" + arg

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Ошибка при выполнении запроса: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка при чтении ответа: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Неправильный код ответа: %d", resp.StatusCode)
	}

	var data ApiResponseIP
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatalf("Ошибка при парсинге JSON: %v", err)
	}

	fmt.Printf("Query: %s, Country: %s, RegionName %s, City: %s, Mobile: %t, Proxy: %t, Hosting: %t \n",
		data.Query, data.Country, data.RegionName, data.City, data.Mobile, data.Proxy, data.Hosting)
}
