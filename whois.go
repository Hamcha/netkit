package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strings"
	"flag"
)

const whoisService string = "http://whomsy.com/api/"

type APIResponse struct {
	Domain  string
	Message string
}

func main() {
	flag.Parse()
	host := flag.Args()

	if len(host) < 1 {
		fmt.Println("USAGE: whois DOMAIN")
		return
	}

	resp, err := http.Get(whoisService+host[0])

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var apiresp APIResponse
	err = json.Unmarshal(body, &apiresp)

	if strings.Index(apiresp.Message,"\n") >= 0 {
		lines := strings.Split(apiresp.Message,"\n")

		for _,x := range lines {
			if (strings.Index(x,":") < len(x)-1) && (strings.Index(x,":") > 0) {
				fmt.Println(x);
			}
		}
	} else {
		fmt.Println(apiresp.Message);
	}

}
