# {{app_name}}

> {{app_description}}

---

## Configuracoes do Projeto

---

## Como Rodar

- First run

Antes de tudo, rode este comando para evitar problemas durante a instalação das dependências da aplicação:
```
git config --global url."git@github.com:argosapitech".insteadOf "https://github.com/argosapitech"
go env -w GOPRIVATE="github.com/argosapitech/*"
```

Então, você pode executar as dependências rodando:

```
go mod tidy && make deps/local/up
```

- Run API

```
make run/api
```

- Run Consumer

```
make run/consumer
```
- Unit Test
```
make test/local/unit
```
- Integration Test
```
make test/local/live-integration
```