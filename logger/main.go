package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	for {
		request, err := http.NewRequest("GET", "http://server:8080", strings.NewReader(``))
		if err != nil {
			panic(err)
		}

		client := &http.Client{}
		resp, err := client.Do(request)

		if err == nil {
			fmt.Println(resp.StatusCode)
			body, _ := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			fmt.Println(string(body))
		} else {
			fmt.Println("Error when curling:" + err.Error())
		}
		time.Sleep(1 * time.Second)
	}
}
