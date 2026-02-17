package pacoteExemplo // 1. Declaração do pacote (pode ser diferente de 'main')
import (
	"estudosGo/pacoteExemplo/internal/pacoteInterno"
	"fmt"
)

var VariavelExemploA = "Esta é uma variável do pacote exemplo A" // 2. Declaração de uma variável pública (iniciada com letra maiúscula)
var VariavelExemploC = variavelExemploC                          // 3. Variável pública que referencia a variável privada (variávelExemploC) do mesmo pacote

func PrintVariavelPacoteInterno() {
	fmt.Println(pacoteInterno.VariavelInterna) // 4. Função pública que imprime o valor da variável do pacote interno
}