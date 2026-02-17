package main // 1. Declaração obrigatória do pacote

import (
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

		// fmt.Println(pacoteInterno.VariavelInterna) // 6. Tentando acessar a variável do pacote interno diretamente (não é possível, pois o pacote interno não é acessível fora do pacote exemplo)
}