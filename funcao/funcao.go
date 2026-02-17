package funcao

// O Go permite retornar múltiplos valores, então podemos retornar tanto o resultado da divisão quanto o resto.
func Dividir(a, b int) (resultado int, resto int) { 
	resultado = a / b
	resto = a % b

	// return // Retorna os valores de resultado e resto. O Go entende que estamos retornando as variáveis nomeadas.
	// Não é recomendado!

	return resultado, resto // Retorna os valores de resultado e resto. Esta é a forma mais clara e recomendada de retornar múltiplos valores.
}

// Outra forma de declarar a função sem nomear os retornos.
func Dividir2(a, b int) (int, int) { 
	resultado := a / b // Usamos := para declarar e atribuir o valor à variável resultado.
	resto := a % b

	return resultado, resto // Retorna os valores de resultado e resto.
}

// Função de alta ordem que retorna outra função.
func funcaoHighOrderSomar(a int) func(int) int {
	return func(b int) int { // Retorna uma função anônima que recebe um inteiro b e retorna um inteiro.
		return a + b // A função anônima tem acesso à variável a da função externa, permitindo somar a e b.
	}
}

func UsarHighOrder1() int {
	somarCom5 := funcaoHighOrderSomar(5) // Cria uma função que soma 5 a qualquer número que receber.
	return somarCom5(10) // Chama a função retornada com o argumento 10, resultando em 15.
}

func UsarHighOrder2() int {
	return funcaoHighOrderSomar(5)(10) // Chama a função de alta ordem diretamente, passando 5 para criar a função e 10 para chamar a função retornada, resultando em 15.
}

// Função que aceita um número variável de argumentos (variadic function).
func SomarVarios(numeros ...int) int { 
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}