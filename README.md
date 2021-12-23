## Ao longo desse readme teremos algumas anotações feitas durante o estuda da linguagem

## Como executar um programa em Go

Para executar, basta compilar o programa e depois executar, algo semelhando ao C, quando usamos o GCC para compilar.

```go
go build nome-arquivo.go
./nome-arquivo
```

Para tornar isso menos verboso, podemos utilizar outro recurso do go, que compila e executa tudo no mesmo comando, que é o `go run`

```go
go run nome-arquivo.go
```

### Convenções/Padrões
- Por padrão/convenção, todo nome de função de pacote importado em nossos programas `iniciam com letra maiúscula`, isso também se da ao fato de funções com letra minúscula serem funções privadas do arquivo de origem
- Ponto e vírgula `;` é opcional, por padrão não é usado
- Variáveis declaradas e não utilizadas causam erro ao compilar

## Go Workspace
- A pasta bin para os arquivos binários, compilados
- A pasta src para o código fonte
- A pasta pkg para os pacotes compartilhados entre as aplicações

## Tipos de variáveis e declarações

#### Tipos primitivos
- int
  - int
  - int8  
  - int16  
  - int32  
  - int64
  - uint
  - uint8
  - uint16
  - uint32
  - uint64
  - uintptr 
- byte // alias for uint8
- float
  - float32 
  - float64
- complex
  - complex64
  - complex128
- string
- bool

#### Declaração

```go
  var name string = "Gabriel"

```

- Podemos declarar com `var` no iníciom seguido do `nome-variavel`, `tipo`(opcional), `=`, `valor`
- há várias formas de declaração
- Const é algo que não poderá ter seu valor alterado
- Temos inferências de tipo, ou seja, temos tipagem automática
```go
  var foo int
  foo = 32

  // OR:
  var foo int = 32

  // OR:
  var foo = 32

  // OR:
  foo := 32

  // OR:
  const foo = 32
```

#### Operador de atribuição `=`
Uso:
- Reatribuição de variáveis

#### Operador curto de declaração `:=`
Uso:
- Tipagem automática
- Só pode repetir se houverem variáveis novas
!= do assignment operator (operador de atribuição)
- Só funciona dentro de codeblocks
- := utilizado pra criar novas variáveis, dentro de code blocks
- = para atribuir valores a variáveis já existentes
- Não pode usar `:=` fora de funcs. Porque, fora de qualquer função, uma instrução deve começar com uma palavra-chave/reservada. 
- Você pode usar a declaração curta para declarar uma variável em um escopo mais recente, mesmo que a variável já tenha sido declarada com o mesmo nome antes:

```go
  var foo int = 34

  func some() {
    // because foo here is scoped to some func
    foo := 42  // <-- legal
    foo = 314  // <-- legal
  }
```
> Referências de alguns itens a cima https://stackoverflow.com/questions/17891226/difference-between-and-operators-in-go/45654233#45654233 

### Controle de fluxo com if
- Não precisa de parenteses na expresão lógica

```go
   if expressao == 1 {

    } else if expressao == 2 {

    } else if expressao == 0 {

    } else {

    }
```
### Controle de fluxo com switch
- Não precisa de parenteses na expresão lógica
- Não precisa break nas expressões de cada case

```go
   switch expressao {
    case 1:
        fmt.Println("1...")
    case 2:
        fmt.Println("2...")
    case 0:
        fmt.Println("0...")
    default:
        fmt.Println("default ...")
  }
```
### Funções
- Tem a estrutura bem parecida com outras linguagens
```go
   func nameFunc(arg1 int, arg2 float32) int {
   }
```

- Podem retornar mais do que um valor
```go
   func nameFuncReturn2Values() (int, string) {
     name := "gabriel"
     years := 21

     return years, name
   }

   years, name := nameFuncReturn2Values()
  
  // também podemos descartar algum dos retornos utilizando o "_"  
   _, name := nameFuncReturn2Values()
```

#### Estruturas de repetição
- Em go não temos o while, quando queremos um loop infinito, utilizamos o for sem nenhum argumento
```go
  for {
    // instrução aqui
  }
```