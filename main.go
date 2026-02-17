package main // 1. Declaração obrigatória do pacote

import (
	"estudosGo/funcao"
	"estudosGo/pacoteExemplo"
	"fmt"
) // 2. Importação de pacotes (format)

// 3. Função principal (Entry Point)
func main() {
    fmt.Println("Hello World")
		fmt.Println(pacoteExemplo.VariavelExemploA)
		fmt.Println(pacoteExemplo.VariavelExemploB)
		// fmt.Println(pacoteExemplo.variavelExemploC) // 4. Acessando variáveis do pacote exemplo (variável privada não pode ser acessada)
		fmt.Println(pacoteExemplo.VariavelExemploC)
		fmt.Println("Acessando variável do pacote interno através da função pública do pacote exemplo:")
		pacoteExemplo.PrintVariavelPacoteInterno() // 5. Acessando a variável do pacote interno através da função pública do pacote exemplo

		fmt.Println("=====")

		fmt.Println(funcao.Dividir(10,2))
		fmt.Println(funcao.Dividir2(10,2))

		resultado, resto := funcao.Dividir2(10,2)

		fmt.Println("RESULTADO:", resultado, "RESTO:", resto)

		fmt.Println(funcao.SomarVarios(1,2,3,4,5))

		// UsarHighOrder1 não recebe argumentos; chame-a assim:
		fmt.Println("UsarHighOrder1():", funcao.UsarHighOrder1())

		// Existe também UsarHighOrder2 que faz a mesma operação direto:
		fmt.Println("UsarHighOrder2():", funcao.UsarHighOrder2())

		// fmt.Println(pacoteInterno.VariavelInterna) // 6. Tentando acessar a variável do pacote interno diretamente (não é possível, pois o pacote interno não é acessível fora do pacote exemplo)
}