package main

import (
	"fmt"
	"net/http"
)

func main() {

	for {
		res, err := iniciarMonitoramento()


		if res.StatusCode == 200 {
			fmt.Println("Response: ", res.Status)
		} else {
			fmt.Println("Request Error: ", err, res.Status)
		}

	}
}

func iniciarMonitoramento() (*http.Response, error) {
	fmt.Println("Monitorando...")
	siteUrl := "https://httpstat.us/random/200,201,400-404,500-504"
	// siteUrl = "https://www.aludsra.com.br/"
	res, err := http.Get(siteUrl)

	return res, err
}
