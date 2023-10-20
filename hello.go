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

const monitoramento = 5
const delay = 3

func main() {

	for {
		exibeIntroducao()
		iniciarMonitoramento()
		comando := escolherComando()

		switch comando {
		case 1:
			monitorar()
		case 2:
			fmt.Println("Exibindo Logs:")
			imprimeLogs()
		case 3:
			fmt.Printf("Saindo do programa...")
			os.Exit(1)
		default:
			fmt.Println("Comando não localizado!")
		}
	}

}

func exibeIntroducao() {
	nome := "Felipe"
	versao := 1.2
	fmt.Println("Olã, Sr.", nome)
	fmt.Println("Este programa esta na versão", versao)
}

func iniciarMonitoramento() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Apresentar logs")
	fmt.Println("3 - Sair do programa")
}

func escolherComando() int {
	var comando int
	fmt.Scan(&comando)

	fmt.Println("O comando escolhido foi", comando)
	return comando
}

func monitorar() {
	fmt.Println("Monitorando...")

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		fmt.Println("Recarregando testes!")
		time.Sleep(delay * time.Second)

	}
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro ao monitorar o site, erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O Site", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("O Site", site, "está com problema", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var listaSite []string

	arquivo, err := os.Open("sites.txt")
	//arquivo, err := os.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo, erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err1 := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		listaSite = append(listaSite, linha)

		if err1 == io.EOF {
			break
		}
	}
	arquivo.Close()
	return listaSite
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro na leitura:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " Online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo: ", err)
	}

	fmt.Println(string(arquivo))
}
