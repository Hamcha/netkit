package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"flag"
)

const geoipService string = "http://ip-api.com/json"

type APIResponse struct {
	Status string
	Country string
	CountryCode string
	Region string
	RegionName string
	City string
	Zip string
	Lat string
	Lon string
	Timezone string
	Isp string
	Org string
	As string
	Query string
	Message string
}

func main() {
	flag.Parse()
	host := flag.Args()

	var resp *http.Response
	var err error

	if len(host) < 1 {
		resp, err = http.Get(geoipService)
	} else {
		resp, err = http.Get(geoipService+"/"+host[0])
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var apiresp APIResponse
	err = json.Unmarshal(body, &apiresp)

	if apiresp.Status == "fail" {
		fmt.Println("FAIL: %s",apiresp.Message)
	} else {
		fmt.Printf("IP Address: %s\n",apiresp.Query)
		fmt.Printf("Country: %s (%s)\n",apiresp.Country,apiresp.CountryCode)
		fmt.Printf("Region: %s (%s)\n",apiresp.RegionName, apiresp.Region)
		fmt.Printf("City: %s\n",apiresp.City)
		fmt.Printf("Zip: %s\n",apiresp.Zip)
		fmt.Printf("Lat: %s\n",apiresp.Lat)
		fmt.Printf("Lon: %s\n",apiresp.Lon)
		fmt.Printf("Timezone: %s\n",apiresp.Timezone)
		fmt.Printf("ISP: %s (%s)\n",apiresp.Isp, apiresp.Org)
		fmt.Printf("As: %s\n",apiresp.As)
	}

}
