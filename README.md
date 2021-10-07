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
- Os tipos básicos de variaveis no go são int, float (32 e 64 bits), bool, string.
- Para saber o tipo de uma variável usa-se a função `reflect.TypeOf(variable)`.
- Para sair do programa com um código de erro no terminal usa-se a função `os.Exit(errorCode)`.
- Para fazer requisições http no go temos o pacote `net/http` onde podemos usar funções como a `http.Get(utl)`. 
- Variaveis float no go podem ser de 2 tipos: float32 e float64.
- As funções `fmt.Scanf("%s", $variable)` e `fmt.Scan(&variable)` são usadas para fazer leitura do teclado armazenando valores em ponteiros de variáveis.
- Instruções como if e for não precisam de parênteses.
- A instrução switch case no go não precisa do comando break.
- Em go não existe a instrução while, usa-se apenas o for no lugar.
- Funções no go podem retornar multiplos valores e todos eles devem ser atribuidos a variáveis. Caso não queira usar um dos valores, pode-se nomeá-lo como _.
- Há uma diferença entre arrays e slices no go. A diferença é que o array tem um número definido de posições já o slice tem posições variáveis.
- Arrays e slices possuem um tamanho que pode ser obtido com a função `len(arrayOrSlice)` e também tem uma quantidade de elementos possiveis obtidos com a função `cap(arrayOrSlice)` (capacidade).
- Para adicionar um novo elemento em um slice, usa-se a função `append(slice, novoItem)` que retorna um novo slice sem modificar o anterior.
