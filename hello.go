package main //Pacote principal

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
)

func main() {
	criarArray()
	criarSlice()
	criaVariavel()
	for {

		comando := leComando()

		exibeMenu(comando)
	}

	// nome, idade := devolveNomeIdade()

	// fmt.Println("Nome:", nome+"Idade", idade)

	///Se não quero utilizar todos os retornos da function
	// texto, _ := devolveNomeIdade()
	// fmt.Println(texto)

	// _, num := devolveNomeIdade()
	// fmt.Println(num)

	//Function

}

func criaVariavel() {
	//Criando variavel
	var nome string = "Igor"
	var versao32 float32 = 1.1
	var versao64 float64 = 10.14
	var texto = "Variável texto"
	versao3 := "Versão 45.45.1"

	fmt.Println("Olá,", nome)
	fmt.Println("Este programa está ne versao,", versao32)
	fmt.Println("Este programa está ne versao,", versao64)

	fmt.Println(texto)
	fmt.Println("O tipo da varivavel versao32 é", reflect.TypeOf(versao32))
	fmt.Println(versao3)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func exibeMenu(comando int) {

	// //IF
	// if comando == 1 {
	// 	fmt.Println("Monitorando")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo logs")
	// } else if comando == 0 {
	// 	fmt.Println("saindo fo programa")
	// } else {
	// 	fmt.Println("comando desconhecido")
	// }

	//Switch
	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo logs")
	case 0:
		fmt.Println("saindo fo programa")
		os.Exit(0)
	default:
		fmt.Println("Comando desconhecido")
		os.Exit(-1)
	}
}

func leComando() int {
	var comandoLido int
	//fmt.Scanf("%d", &comando) //o & está indicando que é o endereço (ponteiro) da memoria
	fmt.Scan(&comandoLido) // não precisa fornecer o indicafor da variavel %f

	fmt.Println("O comando escolhido foi", comandoLido)

	fmt.Println("O endereço da minha variável é", &comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando")
	site := "https://www.alura.com.br"
	resposta, _ := http.Get(site)
	//fmt.Println(resposta)

	if resposta.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
	} else {
		fmt.Println("Site", site, "está com problemas. Status code:", resposta.StatusCode)
	}
}

// / Função com 2 retornos
func devolveNomeIdade() (string, int) {
	nome := "Douglas"
	idade := 30

	return nome, idade
}

// Array com go
func criarArray() {
	var sites [2]string

	sites[0] = "primeiro"
	sites[1] = "segundo"

	fmt.Println("array", sites)

}

func criarSlice() {
	nomes := []string{"primeiro nome", "segundo nome", "terceiro nome"}
	fmt.Println("slice", nomes)
	fmt.Println("quantidade", len(nomes))
	fmt.Println("capacidade do slice", cap(nomes))
}

// For
func criarFor() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	sites := [2]string{"", ""}

	for j, site := range sites {
		fmt.Println("posicao - j", j)
		fmt.Println("For com range - site", site)
	}
}
