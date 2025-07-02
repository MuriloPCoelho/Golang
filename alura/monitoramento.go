package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	iniciarMonitoramento()
	lerSitesDoArquivo()
}

const monitoramento = 3
const delay = 5

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{
		"https://ufcspa.edu.br/",
		"https://www.alura.com.br/",
		"https://qi.edu.br/faqi/",
	}

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Estou passando na posição", i, "do meu slice e essa posição tem o site:", site)
			testaSite(site)
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testaSite(url string) {
	res, err := http.Get(url)

	if res.StatusCode == 200 {
		fmt.Println("Response: ", res.Status)
	} else {
		fmt.Println("Request Error: ", err)
	}
}

func lerSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("alura/sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	fmt.Println(arquivo)
	return sites
}
