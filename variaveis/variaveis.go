package variaveis

// Variáveis em Go são declaradas usando a palavra-chave var, seguida pelo nome da variável e seu tipo.
var Nome string // Declara uma variável do tipo string chamada nome.
var Idade int // Declara uma variável do tipo int chamada idade.
var A, B string // Declara duas variáveis do tipo string, a e b.

// Variáveis também podem ser declaradas e inicializadas ao mesmo tempo.
var Cidade string = "São Paulo" // Declara e inicializa a variável cidade com o valor "São Paulo".
var Altura, Peso float64 = 1.75, 70.5 // Declara e inicializa as variáveis altura e peso com os valores 1.75 e 70.5, respectivamente.

// O Go também suporta a inferência de tipo, onde o tipo da variável é determinado pelo valor atribuído a ela.
var Sobrenome = "Silva" // O tipo da variável sobrenome é inferido como string.
var Idade2 = 30 // O tipo da variável idade2 é inferido como int.

// Bloco de declaração de variáveis, onde podemos declarar várias variáveis em um único bloco usando parênteses.
var (
	PrimeiroNome string = "João" // Declara e inicializa a variável primeiroNome com o valor "João".
	UltimoNome  string = "Pereira" // Declara e inicializa a variável ultimoNome com o valor "Pereira".
	Profissao   string = "Engenheiro" // Declara e inicializa a variável profissao com o valor "Engenheiro".
)

// Para declarar e inicializar variáveis dentro de uma função, podemos usar a sintaxe curta :=, que é uma forma mais concisa de declarar e atribuir valores.
func DeclararVariaveis() {
	nome := "Maria"	// Declara e inicializa a variável nome com o valor "Maria". O tipo é inferido como string.
	idade := 25 // Declara e inicializa a variável idade com o valor 25. O tipo é inferido como int.
	cidade := "Rio de Janeiro" // Declara e inicializa a variável cidade com o valor "Rio de Janeiro". O tipo é inferido como string.

	// Podemos usar as variáveis declaradas para realizar operações ou imprimir seus valores.
	println("Nome:", nome)
	println("Idade:", idade)
	println("Cidade:", cidade)
}


