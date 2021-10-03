# GoLang

## Diretórios base em um projeto golang

Depois de instalar o sdk do golang no pc e adicionar o path do go ao PATH do sistema é necessário criar o diretório `go` na home do seu usuário.
O diretório `~/go` deve possuir 3 sub diretórios: `src, pkg, bin`.
O diretório `~/go/src` será onde colocaremos os códigos fonte das aplicações.


```sh
mkdir -p ~/go/{pkg,src,bin}
mkdir ~/go/src/hello
touch ~/go/src/hello/hello.go
code ~/go/src
```

## Hello World em Go

```go
// hello.go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
```

## Anotações básicas

- Go não precisa de ponto e virgula
- O arquivo de entrada da aplicação deve declarar o `package main` e a função de entrada deve ser a função `main`.
- O comando `go build main.go` faz a compilação do arquivo main.go e gera um executável com mesmo nome porem sem a extensão `go`.
- O comando `go run main.go` compila e executa o arquivo, porém não gera um arquivo executável como saída.
- Variáveis declaradas porém não utilizadas geram erro de compilação e devem ser removidas.
- A declaração de variáveis e tipo pode ser explicita ou implícita.
- A palavra chave `var` é usada para declarar uma variável, após ela informamos o nome da variável e após o nome informamos o tipo. Exemplo `var name string = "John"`.
- Existe uma forma curta de declarar e inicializar variáveis com o operador `:=`. O que estiver a esquerda do operador define o nome da variável e o que estiver à direita define o valor. Exemplo: `name := "John"`.
