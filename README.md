## Ao longo desse readme teremos algumas anotações feitas durante o estudo da linguagem

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

### Estruturas de repetição
- Em go não temos o while, quando queremos um loop infinito, utilizamos o for sem nenhum argumento
```go
  for {
    // instrução aqui
  }
```
- Também podemos utilizar a forma padrão de um for
```go
  for i := 0; i < len(arrayX); i++ {
      fmt.Println(arrayX[i])
  }
```
- Outra forma é utilizar o `range`
```go
   for i, element := range arrayX {
        fmt.Println("Estou passando na posição", i,
            "o valor do elemento é", element)
    }
```
### Array
- Array tem um tamanho fixo, não podemos modificar o tamanho
- Para tamanhos dinâmicos temos outra estrutura chamada de slice
```go
  var names [4]string
  names[0] = "Gabriel"
  names[1] = "João"
  names[2] = "Toy"
```

### Slices
- Similar é um array, mas com tamanho dinâmico
- Faz alocação de memória conforme a necessidade
- Podemos usar a func `append` para adicionar elementos ao slice
- Podemos usar a func `len` para verificar o tamanho
- Podemos usar a func `cap` para verificar a capacidade

```go
  names := []string {"Gabriel", "João", "Toy"}
  append(names, "Aparecida")

  fmt.Println("O meu slice tem", len(names), "itens")
  fmt.Println("O meu slice tem capacidade para", cap(names), "itens")

```

## Módulos
- No Go lidamos muito com pacotes, para começar a criar pacotes, devemos utilizar o comando
  ```bash
    go mod init nomeModulo
  ```
- Por padrão, funções dentro de um módulo para serem expostas para outros arquivos, devem iniciar com letra maiúscula, caso seja minuscula, ela irá se comportar como uma função privada daquele arquivo
- Basta declarar no início do arquivo o nome do pacote
  Ex:  
  ```go
  package testModules

  import "fmt"

  func TestPrint() {
    fmt.Println("Func test")
  }
  ```
- Após isso só usar em outro arquivo
  ```go
  package main

  import testModules "testModules/module-test"

  func main() {
    testModules.TestPrint()
  }
  ```

  ## Ponteiros
  - Ponteiro é um tema muito grande e pode ser um pouco complexo, mas tentarei resumir de forma simples
    - Basicamente um ponteiro aponta para um endereço de memória, onde podemos controlar o que está em determinado endereço de memória
    - Há duas formas de passar uma variável para uma função, passar por cópia ou referência (usando ponteiros), veremos um exemplo a seguir
    ```go
    package main

    import "fmt"

    func main() {
      var value int32 = 9

      fmt.Println("address memory is", &value, "and value:", value)
      printAsCopy(value)
      fmt.Println("After function address memory is", &value, "and value:", value)

      fmt.Println("")
      fmt.Println("")

      fmt.Println("address memory is", &value, "and value:", value)
      printAsReference(&value)
      fmt.Println("After function address memory is", &value, "and value:", value)
    }

    func printAsReference(value *int32) {
      fmt.Println("printAsReference -> address memory is", value, "and value:", *value)
      *value = 16 // Alterando o valor na memória, teremos efeito diretamente na função main
    }

    func printAsCopy(value int32) {
      fmt.Println("printAsCopy -> address memory is", &value, "and value:", value)
      value = 16 // como estamos mudando o valor somente aqui na variável de cópia, não teremos
      // o efeito na função main
    }
    ```

## Maps
- Utiliza o formato key:value.
- E.g. nome e telefone
- Performance excelente para lookups.
- map[key]value{ key: value }
- Acesso: m[key]
- Key sem value retorna zero. Isso pode trazer problemas.
- Para verificar: comma ok idiom.
- v, ok := m[key]
- ok é um boolean, true/false
- Na prática: if v, ok := m[key]; ok { }
- Para adicionar um item: m[v] = value
- Maps não tem ordem.
- Ex: 
  ```go
  package main

  import (
    "fmt"
  )

  func main() {

    amigos := map[string]int{
      "alfredo": 5551234,
      "joana":   9996674,
    }

    fmt.Println(amigos)
    fmt.Println(amigos["joana"])

    amigos["gopher"] = 444444

    fmt.Println(amigos)
    fmt.Println(amigos["gopher"], "\n\n")


    // comma ok idiom
    if será, ok := amigos["fantasma"]; !ok {
      fmt.Println("não tem!")
    } else {
      fmt.Println(será)
    }

  }
  ```
- Maps: range & deletando
- Range: for k, v := range map { }
- Reiterando: maps não tem ordem e um range usará uma ordem aleatória.
- Ex:
  ```go
    package main

    import (
      "fmt"
    )

    func main() {

      qualquercoisa := map[int]string{
        123: "muito legal",
        98:  "menos legal um pouquinho",
        983: "esse é massa",
        19:  "idade de ir pra festa",
      }

      fmt.Println(qualquercoisa)

      total := 0

      for key, _ := range qualquercoisa {
        total += key
      }
      
      fmt.Println(total)

    }

  ```

- delete(map, key)
- Deletar uma key não-existente não retorna erros!
Ex:
  ```go
    package main

    import (
      "fmt"
    )

    func main() {

      qualquercoisa := map[int]string{
        123: "muito legal",
        98:  "menos legal um pouquinho",
        983: "esse é massa",
        19:  "idade de ir pra festa",
      }

      fmt.Println(qualquercoisa)

      delete(qualquercoisa, 123)

      fmt.Println(qualquercoisa)

    }

  ```

## Strucs
- É um tipo de dados composto, agrupa tipo de dados arbitrários em um objeto
Ex: 
```go
  type struct Person {
    name  string 
    addr  string 
    phone string 
  }

  // declaração sem iniciar
  var p1 Person
  
  // ou iniciar um objeto com todos os valores "zero"
  p1 := new(Person)
  
  // ou declarar e iniciar usando struct literal
  p1 := Person(name: "gabriel", addr: "rua ...", phone: "1234")
```

## Funções
- É apenas um conjunto de instruções com um nome
- Ótimo para operações comuns, utilizadas em mais de um lugar
- Ótimo para separar responsabilidade por contexto

```go
//chave func, nome da função (args) retorno { }
func main() {
  // conteúdo da função aqui ....
}
```
#### Argumentos em funções
- Vem após o nome dentro de parâmetros ex: ```func hello(parametro int)....```
- Cada parâmetro é separado por virgula
- Podemos encurar a tipagem dos parâmetros, quando tempos uma sequência de parêmetros do mesmo tipo
    ex: 
```go
  func hello(firstName, lastName string) {
    fmt.Println(firstName + lastName)
  }
```

#### Retornos
- Toda função pode retornar algo ou não (void)
- Podemos ter mais do que um retorno para uma função
- Sempre que a função tiver retorno, em sua chamada, devemos atribuir o valor de retorno a uma variável

ex: 
```go
//  Sem retorno
func foo(x int) {
  fmt.Println(x+1);
}

foo()

//  um retorno
func foo1(x int) int {
  return x+1;
}

a := foo1()

//  mais de um retorno
func foo1(x int) (int, int) {
  return x, x+1;
}

a, b := foo2()
```

#### Passagem de parâmetros por cópia vs referência
- Temos essas duas formas de passar parêmetros pras funções
  - Cópia, forma padrão, uma cópia do parâmetro é passado para dentro da funçõa, ao alterar o valor dentro da função, não irá refletir no valor originar passado
  - Referência, passamos o endereço de memória da variável, ao alterar o valor, alteramos o valor no endereço de memória e consequentemente na variável original
  ex: 
```go
package main

import "fmt"

func main() {
	var value int32 = 9

	fmt.Println("address memory is", &value, "and value:", value)
	printAsCopy(value)
	fmt.Println("After function address memory is", &value, "and value:", value)

	fmt.Println("")
	fmt.Println("")

	fmt.Println("address memory is", &value, "and value:", value)
	printAsReference(&value)
	fmt.Println("After function address memory is", &value, "and value:", value)
}

func printAsReference(value *int32) {
	fmt.Println("printAsReference -> address memory is", value, "and value:", *value)
	*value = 16 // Alterando o valor na memória, teremos efeito diretamente na função main
}

func printAsCopy(value int32) {
	fmt.Println("printAsCopy -> address memory is", &value, "and value:", value)
	value = 16 // como estamos mudando o valor somente aqui na variável de cópia, não teremos
	// o efeito na função main
}
```

#### Function as Arguments
  - Podemos criar uma função que recebe outra função como parâmetro
  ex: 
```go
// quando chamada, irá executar uma função que criamos, aplicando o valor
func applyIt(outraFunc func (int) int, 
  valor int) int {
    return outraFunc(valor)
}

func minhaFuncSoma(value int) {
  return value + 1
}
func minhaFuncSubtrai(value *int32) {
  return value - 1
}

func main() {
  fmt.Println(applyIt(minhaFuncSoma, 10)) // 11
  fmt.Println(applyIt(minhaFuncSubtrai, 10)) // 9
}
```

#### Funções anônimas
- Funções sem nome
  ex: 
```go
// quando chamada, irá executar uma função que criamos, aplicando o valor
func applyIt(outraFunc func (int) int, 
  valor int) int {
    return outraFunc(valor)
}

func main() {
  fmt.Println(applyIt(func (x int) int {
    return x+1
  }, 10)) // 11
}
```
#### Funções variádicas (com número de parâmetros variados) 
- Uma função pode ter número de parâmetros variados
- Os parâmetros se tornam um slice dentro da função
ex: 
```go
  func getMax(values ...int) int {
    max := -1 
    for _, v := range values {
      if v > max {
        vax = v
      }
    }

    return max
  }


  maxValue := getMax(2,34,5,1,2,3,44,5,9)

  // também posso passar um slice
  vslice := [] int{1,2,3,4,53,2,33,45}
  maxValue := getMax(vslice...)
``` 

#### Defer Function call (Adiar a chamada de uma função)  
- Podemos chamar uma função com a declaração de `defer`, que irá adiar a chamada da função
- Os argumentos não são adiados, serão utilizados na ordem correta, caso a função recebesse um `X` como parâmetro e o valor dele fosse 2, caso mudasse para 4, o valor da execução da função no defer será 2
ex: 
```go
  func main() {
    defer fmt.Println("Bye!")
    
    
    fmt.Println("Hello")
  }
//  output -> 
// Hello
// Bye!
``` 