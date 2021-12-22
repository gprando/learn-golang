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

### Convenções
- Por padrão/convenção, todo nome de função de pacote importado em nossos programas `iniciam com letra maiúscula`, isso também se da ao fato de funções com letra minúscula serem funções privadas do arquivo de origem