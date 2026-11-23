# De Mala Pronta

Aplicação web para organização de viagens, cadastro de destinos e controle de despesas.
Projeto acadêmico em Go com programação em camadas, CRUD, banco relacional e API REST.

## Tecnologias

- Go (net/http, database/sql)
- SQLite com o driver `modernc.org/sqlite` (sem CGO)
- HTML e JavaScript puro (Fetch API)

## Estrutura de pastas

```
de-mala-pronta-api/
├── main.go                 # inicia conexão, serviços, controladores e servidor
├── conexao/                # conexão e criação das tabelas SQLite
├── controladores/          # rotas e tratamento HTTP/JSON
├── daos/                   # comandos SQL
├── dtos/                   # structs de entrada/saída (JSON)
├── modelos/                # entidades
├── servicos/               # regras de negócio e validações
├── telas/                  # páginas HTML
├── scripts/                # JavaScript que consome a API
├── css/                    # estilos CSS
├── go.mod
└── README.md
```

## Como instalar as dependências

Requer Go 1.22 ou superior.

```
go mod tidy
```

## Como executar

```
go run main.go
```

Acesse http://localhost:8080. O banco `de-mala-pronta.db` é criado automaticamente na primeira execução.

## Endpoints

| Método | Rota | Descrição |
|--------|------|-----------|
| POST | /api/viagens | Cadastrar viagem |
| GET | /api/viagens | Listar viagens |
| GET | /api/viagens/{id} | Buscar viagem |
| PUT | /api/viagens/{id} | Atualizar viagem |
| DELETE | /api/viagens/{id} | Excluir viagem |
| GET | /api/viagens/{id}/total-despesas | Soma das despesas da viagem |
| POST | /api/destinos | Cadastrar destino |
| GET | /api/destinos | Listar destinos (filtro `?viagem_id=`) |
| GET | /api/destinos/{id} | Buscar destino |
| PUT | /api/destinos/{id} | Atualizar destino |
| DELETE | /api/destinos/{id} | Excluir destino |
| POST | /api/despesas | Cadastrar despesa |
| GET | /api/despesas | Listar despesas (filtro `?viagem_id=`) |
| GET | /api/despesas/{id} | Buscar despesa |
| PUT | /api/despesas/{id} | Atualizar despesa |
| DELETE | /api/despesas/{id} | Excluir despesa |

Datas usam o formato `YYYY-MM-DD`.

### Códigos de resposta

200 sucesso · 201 criado · 204 excluído · 400 dados inválidos · 404 não encontrado · 409 conflito de relacionamento (viagem inexistente ou viagem com destinos/despesas ao excluir) · 500 erro interno.

Erros retornam `{"erro": "mensagem"}`.

### Regras de negócio

- Todos os campos são obrigatórios.
- `data_fim` não pode ser anterior a `data_inicio`; `orcamento` não pode ser negativo.
- Destino: `data_chegada` não pode ser posterior a `data_saida`; a viagem deve existir.
- Despesa: `valor` > 0; a viagem deve existir; `data` dentro do período da viagem.

## Exemplos de requisições

Criar viagem:

```
curl -X POST http://localhost:8080/api/viagens \
  -H "Content-Type: application/json" \
  -d '{"titulo":"Rio de Janeiro","data_inicio":"2026-01-10","data_fim":"2026-01-20","orcamento":3000,"status":"planejada"}'
```

Criar destino:

```
curl -X POST http://localhost:8080/api/destinos \
  -H "Content-Type: application/json" \
  -d '{"viagem_id":1,"cidade":"Rio de Janeiro","pais":"Brasil","data_chegada":"2026-01-10","data_saida":"2026-01-20"}'
```

Criar despesa:

```
curl -X POST http://localhost:8080/api/despesas \
  -H "Content-Type: application/json" \
  -d '{"viagem_id":1,"descricao":"Hotel","valor":500.50,"data":"2026-01-11","categoria":"hospedagem"}'
```

Listar despesas de uma viagem e consultar o total:

```
curl "http://localhost:8080/api/despesas?viagem_id=1"
curl http://localhost:8080/api/viagens/1/total-despesas
```

Resposta do total:

```
{"viagem_id":1,"total":500.5}
```
