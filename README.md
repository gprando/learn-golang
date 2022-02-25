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

## Orientação a Objetos em GO

### Classes e encapsulamento 
- Golang não tem classes, mas tem algo equivalente
- O que são classes?
  - Basicamente coleção de campos de dados e funções que compartilham uma responsabilidade bem definida
- Em go temos receptores para um tipo onde podemos simular um classe, geralmente usamos uma struct para ser receptor, mas não é obrigatório

ex usando int:
```go
  package main

  import "fmt"

  type MyClass int

  func (value MyClass) Pow() int {
    return int(value * value)
  }

  func main() {
    num := MyClass(10)

    fmt.Println("num: ", num)
    fmt.Println("pow of num: ", num.Pow())
  }

``` 

ex usando struct:

```go
  package main

  import "fmt"

  type rect struct {
      width, height int
  }

  func (r *rect) area() int {
      return r.width * r.height
  }

  func (r rect) perim() int {
      return 2*r.width + 2*r.height
  }

  func main() {
      r := rect{width: 10, height: 5}

      fmt.Println("area: ", r.area())
      fmt.Println("perim:", r.perim())

      rp := &r
      fmt.Println("area: ", rp.area())
      fmt.Println("perim:", rp.perim())
  }
```

### Polimorfismo
- Polimorfismo são diferentes formas de fazer uma mesma coisa, a alto nível as coisas são iguais, mas a baixo nível são diferente, por exemplo calculo área de figuras geométricas, a essência é a mesma, mas a fórmula é diferente
- Em linguagens tradicionar podemos fazer isso através de herança, mas golang não tem classes nem herança

### Interfaces
- Em go interfaces são um agrupamento de assinatura de métodos
- Não há implementação
- Usado para expressar similaridade conceitual entre tipos 
ex: 
```go
package main

import (
    "fmt"
    "math"
)

type geometry interface {
    area() float64
    perim() float64
}

type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)
    measure(c)
}
```

## Concorrência vs Paralelismo
- Execução paralela necessita de hardware diferente, núcleos diferentes
- Execução concorrente pode executar no mesmo hardware


### Paralelismo
- Execução paralela é quando 2 coisas executam exatamente no mesmo tempo
- No processador cada núcleo é feita e capaz de executar uma coisa de cada vez
- Para conseguirmos ter execução paralela temos que ter um processador com pelo menos 2 núcleos
- Por que fazer execução paralela?
  - Terminar tarefas mais rapidamente
  - Só podemos fazer coisas em paralelo quando as mesmas não tem dependências entre si
  
### Concorrência 
- Pro que usar concorrência?
  - Falamos que tarefas são concorrentes quando a hora de início e a hora de término das tarefas se sobrepóesm 
  - Não significa que estão executando ao mesmo tempo


### Processos
- Basicamente uma instância de um programa em execução
- Cada processo tem coisas exclusivas
  - Como um pedaço de memória
  - Seu próprio código, stack, heap

### Schedule
- Task principal de um processador
- Da uma ilusão de paralelismo em execução
- Alterna execução entre os processos de forma concorrente
- Temos diversos algoritmos de scheduling, com diferentes formas de priorização

### Context Switch
- Quando queremos trocar de processo por exemplo do processo A para o processo B, temos que salvar o estado (Contexto) do processo anterior (A) para usar posteriormente quando retornarmos a ele
  - Contexto do processo, engloba tudo sobre ele, endereços de memória, stack, heap, etc...
- Para definir quando o processador irá alternar entre cada processo, temos auxílio de um timer, onde o processador define um tempo, e a cada tempo ele alterna o contexto para outro processo

## Tópicos e goroutines

#### Threads Vs Processos
- Um processo pode ter várias threads dentro dele
- Thread compartilham contexto

#### Goroutines 
- São threads em Go
- Muitas goroutines executam dentro de uma mesma thread do sistema operacional
- Podemos ter várias goroutines dentro de uma thread, onde estamos alternando entre cada execução, do ponto de vista do sistema operacional, tudo é um processo só, mas internamente no go, ele pode alternar em uma thread vários goroutines. Para isso acontecer, temos auxílio do *go runtime scheduler*

#### Go runtime scheduler
- Usa um processador lógico
  - É mapeado para uma thread
- Faz agendamento dentro de uma thread do SO
- O sistema operacional agenda qual thread irá executar/alternar, uma vez que o SO agendou e executou uma thread, o GO pode escolher diferentes goroutines para executar em momentos diferentes dentro da exução dessa thread

### Concorrência e intercalação
- Intercalação são as instruções em duas tarefas diferentes

## Goroutines
- Uma goroutine é criada automaticamente a partir da função *main()* quando o programa executa
- Outras goroutines podem poder ser criadas utilizando a palavra reservada `go`
- Uma goroutine termina quando o código nela completar
- Quando a goroutine main parar a execução *todas as outras goroutines irão parar*


### Sincronização básica
- Sincroização é quando várias threads concordam com um tempo de um evento
- Uma thread nunca sabe sobre o tempo de uma outra goroutine
- Podemos sincronizar threads através de eventos globais, onde temos a emissão de um evento que é visto por todas as threads

### Wait groups
- Sincronização mais comum no go
- Temos um pacote para isso `sync`
  - Vários métodos para sincronização, sendo o mais comum `sync.WaitGroup`, que faz o programa esperar todas as threads terminem a execução para continuar 
    - Funciona com um contador interno, para cada nova goroutine que adicionamos ao grupo ele incrementa 1, ao completar uma goroutine ele decrementa, quando esperamos todas as goroutines, queremos que o contador esteja < 0 

ex: 
```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int) {
    fmt.Printf("Worker %d starting\n", id)

    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {

    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)

        i := i

        go func() {
            defer wg.Done()
            worker(i)
        }()
    }

    wg.Wait()

}
```
```bash
$ go run waitgroups.go
Worker 5 starting
Worker 3 starting
Worker 4 starting
Worker 1 starting
Worker 2 starting
Worker 4 done
Worker 1 done
Worker 2 done
Worker 5 done
Worker 3 done
```

### Tópicos/Comunicação em threads
- Por mais que goroutines rodem de forma independente, geralmente teremos dentro de um programa rodando coisas do mesmo contexto, pois se tivessemos coisas totalmente independentes, tanto em execução e contexto, faria muito mais sentido estarem em programas separados
  - Geralmente são partes menores de uma tarefa maior

#### Channels
- Transfere dados entre goroutines 
- É tipado
- Podemos especificar a direção do canal
ex: 
```go
package main

import "fmt"

func main() {

    messages := make(chan string)

    go func() { messages <- "ping" }()

    msg := <-messages
    fmt.Println(msg)
}
```

##### Blocking in channels (bloqueio em canais)
- Por padrão um canal é chamado de unbuffered
  - Quando você cria o canal ele é sem buffer
    - Os canais sem buffer não podem armazenar dados em transito


ex: 
```go
package main

import "fmt"

func main() {

    // com buffer igual a 2, o default é sem buffer é = 0 
    messages := make(chan string, 2)

    messages <- "buffered"
    messages <- "channel"

    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
```

### Select em channels
- A seleção do Go permite que você aguarde operações de vários canais. Combinar goroutines e canais com select é um recurso poderoso do Go.
```go
package main

import (
    "fmt"
    "time"
)

func main() {

    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}
```
### Mutex

- Vimos como os canais são ótimos para comunicação entre goroutines.

- Mas e se não precisarmos de comunicação? E se quisermos apenas garantir que apenas uma goroutine possa acessar uma variável por vez para evitar conflitos?

- Esse conceito é chamado de exclusão mútua, e o nome convencional para a estrutura de dados que o fornece é mutex.

- A biblioteca padrão do Go fornece exclusão mútua com sync.Mutex e seus dois métodos:
  - Lock
  - Unlock
  - 
- Podemos definir um bloco de código a ser executado em exclusão mútua envolvendo-o com uma chamada para Lock and Unlock conforme mostrado no método Inc.

- Também podemos usar defer para garantir que o mutex seja desbloqueado como no método Value.
ex:
```go
package main

import (
    "fmt"
    "sync"
)

type Container struct {
    mu       sync.Mutex
    counters map[string]int
}

func (c *Container) inc(name string) {

    c.mu.Lock()
    defer c.mu.Unlock()
    c.counters[name]++
}

func main() {
    c := Container{

        counters: map[string]int{"a": 0, "b": 0},
    }

    var wg sync.WaitGroup

    doIncrement := func(name string, n int) {
        for i := 0; i < n; i++ {
            c.inc(name)
        }
        wg.Done()
    }

    wg.Add(3)
    go doIncrement("a", 10000)
    go doIncrement("a", 10000)
    go doIncrement("b", 10000)

    wg.Wait()
    fmt.Println(c.counters)
}
```
- O contêiner contém um mapa de contadores; como queremos atualizá-lo simultaneamente de várias goroutines, adicionamos um Mutex para sincronizar o acesso. Observe que os mutexes não devem ser copiados, portanto, se esse struct for passado, deve ser feito por ponteiro.
- Bloqueie o mutex antes de acessar os contadores; desbloqueie-o no final da função usando uma instrução defer.