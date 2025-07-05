package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	iniciarMonitoramento()

}

const monitoramento = 3
const delay = 5

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := lerSitesDoArquivo()

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Estou passando na posição", i, "do meu slice e essa posição tem o site:", site)
			testaSite(site)
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("_______________Fim do monitoramento_______________")
	imprimeLogs()
}

func testaSite(url string) {
	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Request Error: ", err)
	}


	if res.StatusCode == 200 {
		registraLog(url, true)
	} else {
		registraLog(url, false)
	}

	fmt.Println("Response: ", res.Status)

}

func lerSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("alura/sites.txt") //retorna o ponteiro de um onde o arquivo está na memória
	// arquivo, err := os.ReadFile("alura/sites.txt") //retorna um array de bytes

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Ocorreu um erro: ", err)
		}
	}
	
	arquivo.Close()

	fmt.Println(sites)

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("alura/logs.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666);

	if err != nil {
		fmt.Println("Ocorreu um erro ao tentar registrar log", err)
	} else {
		fmt.Println(arquivo)
	} 

	arquivo.WriteString(time.Now().Format("02/01/2006 - 15:04:05 ") + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
} 

func imprimeLogs() {
	arquivo, err := os.ReadFile("./alura/logs.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro ao imprimir os logs: ", err)
	} 

	fmt.Println(string(arquivo))
}
