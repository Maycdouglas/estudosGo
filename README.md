# Configuração de Ambiente e Primeiro Programa

## 🛠️ Configuração do VS Code
Para desenvolver em Go no VS Code, é necessário instalar a extensão oficial e as ferramentas auxiliares.

1. **Extensão:** Instale a extensão **Go** (desenvolvida pelo "Go Team").
2. **Ferramentas (Go Tools):**
   - Após instalar a extensão, abra a paleta de comandos (`CTRL + SHIFT + P`).
   - Digite e selecione: `Go: Install/Update Tools`.
   - Selecione todas as opções e clique em **OK**.

### Principais Ferramentas Instaladas
- **`gopls` (Go Please):** É o *Language Server* oficial. Responsável pelo *autocomplete*, *syntax highlight*, formatação e importações automáticas.
- **`dlv` (Delve):** O debugger do Go.
- **`staticcheck`:** Linter para identificar erros comuns e sugerir melhorias no código.
- **`gotests`:** Gera *boilerplate* (esqueleto) de testes automaticamente.
- **`gomodifytags`:** Auxilia na criação de tags em structs (JSON, XML, DB).
- **`impl`:** Gera a implementação de interfaces automaticamente.

---

## 🚀 Inicializando um Projeto
Todo projeto em Go precisa de um módulo (recomendado para versões 1.11+).

1. Crie a pasta do projeto.
2. Abra o terminal na pasta e execute:
   ```bash
   go mod init <nome-do-modulo>
   # Exemplo: go mod init my-first-project
   ```
   *Isso cria o arquivo `go.mod`, que gerencia a versão do Go e as dependências.*

---

## 📝 Estrutura do "Hello World"
Exemplo de arquivo `main.go`:

```go
package main // 1. Declaração obrigatória do pacote

import "fmt" // 2. Importação de pacotes (format)

// 3. Função principal (Entry Point)
func main() {
    fmt.Println("Hello World")
}
```

- **`package main`**: Indica ao compilador que este pacote deve se tornar um **executável** (não uma biblioteca).
- **`func main()`**: É a porta de entrada da execução do programa.

---

## 💻 Comandos de Execução e Compilação

| Comando | Descrição |
| :--- | :--- |
| `go run main.go` | Compila, executa imediatamente e **descarta** o executável ao final. Ideal para desenvolvimento rápido. |
| `go build main.go` | Compila e **gera um executável** binário (ex: `.exe` no Windows) na pasta. |
| `go build -o nome_arq main.go` | Compila definindo um nome específico para o arquivo de saída. |

---

## 🌍 Cross-Compilation (Compilação Cruzada)
O Go permite compilar um binário para um sistema operacional (OS) ou arquitetura diferente da sua máquina atual.

**Variáveis de Ambiente:**
- `GOOS`: Sistema Operacional (ex: `linux`, `windows`, `darwin` para Mac).
- `GOARCH`: Arquitetura do processador (ex: `amd64`, `arm64`).

### Exemplo (No Windows via Terminal):
Para compilar um executável para **Linux** estando no Windows:

1. **Setar as variáveis:**
   ```bash
   go env -w GOOS=linux
   go env -w GOARCH=amd64
   ```

2. **Rodar o build:**
   ```bash
   go build main.go
   ```

3. **IMPORTANTE:** Restaurar para o padrão original após o build:
   ```bash
   go env -w GOOS=windows
   ```

> **Nota:** Em sistemas Linux/Mac, é possível passar as variáveis *inline* sem alterar a configuração global: `GOOS=linux go build main.go`.
