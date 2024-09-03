# pos-go-expert

## Principais comandos

### Executando um arquivo GO

```
go run main.go 
```

### Gerando o binário de arquivo especifico

```
go build main.go
```

### Gerando o binário de arquivo especifico para Windows

```
GOOS=windows go build main.go
```

### Gerando o binário de arquivo especifico para Linux

```
GOOS=linux go build main.go
```

### Gerando o binário de arquivo especifico para Mac

```
GOOS=darwin go build main.go
```

### Gerando o binário de arquivo especifico para Mac e especificando a arquitetura

```
GOOS=darwin GOARC=amd64 go build main.go
```

### Gerando o binário do projeto

```
go build
```
Irá gerar um binário do projeto todo com o nome do modulo definido no *go.mod*.

### Gerando o binário do projeto especificando o nome do binário

```
go build -o novo_nome
```

### Mostrando as arquiteturas disponíveis

```
go tool dist list
```

## Links Utéis
[## Aplicações para diferentes SO e Arquiteturas](https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures)
