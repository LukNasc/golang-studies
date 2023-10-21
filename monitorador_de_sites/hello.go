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

const monitoramentos = 5
const delay = 5

func main() {
	exibeIntroducao()
	for {
		exibeMenu()

		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			exibirLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Nenhuma opção válida")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Lucas"
	versao := 1.1

	fmt.Println("Olá ", nome)
	fmt.Println("Este programa está na versão ", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := lerSitesDoArquivo()

	if len(sites) > 0 {
		for i := 0; i < monitoramentos; i++ {
			for _, site := range sites {
				testaSite(site)
			}
			time.Sleep(delay * time.Second)
			fmt.Println("")
		}

		fmt.Println("")
	}

}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro ao carregar o site", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
	} else {
		fmt.Println("Site:", site, "está com problemas com status code:", resp.StatusCode)
	}

	registraLog(site, resp.StatusCode == 200)
}

func lerSitesDoArquivo() []string {
	resultado := []string{}
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	read := bufio.NewReader(arquivo)
	for {
		linha, err := read.ReadString('\n')
		linha = strings.TrimSpace(linha)
		resultado = append(resultado, linha)
		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return resultado
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()

}

func exibirLogs() {
	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorre um erro ao abrir o arquivo", err)
	}

	fmt.Println(string(arquivo))
}
