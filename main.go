package main

import (
	"log"
	"net/http"

	"de-mala-pronta-api/conexao"
	"de-mala-pronta-api/controladores"
	"de-mala-pronta-api/daos"
	"de-mala-pronta-api/servicos"
)

func main() {
	db, err := conexao.Abrir("de-mala-pronta.db")
	if err != nil {
		log.Fatal("erro ao abrir o banco: ", err)
	}
	defer db.Close()

	viagemDAO := daos.NovoViagemDAO(db)
	destinoDAO := daos.NovoDestinoDAO(db)
	despesaDAO := daos.NovoDespesaDAO(db)

	viagemControlador := controladores.NovoViagemControlador(servicos.NovoViagemServico(viagemDAO))
	destinoControlador := controladores.NovoDestinoControlador(servicos.NovoDestinoServico(destinoDAO, viagemDAO))
	despesaControlador := controladores.NovoDespesaControlador(servicos.NovoDespesaServico(despesaDAO, viagemDAO))

	rotas := controladores.Rotas(viagemControlador, destinoControlador, despesaControlador)

	log.Println("Servidor em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", rotas))
}
