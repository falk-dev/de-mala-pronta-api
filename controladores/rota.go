package controladores

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"de-mala-pronta-api/dtos"
	"de-mala-pronta-api/servicos"
)

func Rotas(viagem *ViagemControlador, destino *DestinoControlador, despesa *DespesaControlador) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/viagens", viagem.Cadastrar)
	mux.HandleFunc("GET /api/viagens", viagem.Listar)
	mux.HandleFunc("GET /api/viagens/{id}", viagem.BuscarPorID)
	mux.HandleFunc("PUT /api/viagens/{id}", viagem.Atualizar)
	mux.HandleFunc("DELETE /api/viagens/{id}", viagem.Excluir)
	mux.HandleFunc("GET /api/viagens/{id}/total-despesas", despesa.TotalPorViagem)

	mux.HandleFunc("POST /api/destinos", destino.Cadastrar)
	mux.HandleFunc("GET /api/destinos", destino.Listar)
	mux.HandleFunc("GET /api/destinos/{id}", destino.BuscarPorID)
	mux.HandleFunc("PUT /api/destinos/{id}", destino.Atualizar)
	mux.HandleFunc("DELETE /api/destinos/{id}", destino.Excluir)

	mux.HandleFunc("POST /api/despesas", despesa.Cadastrar)
	mux.HandleFunc("GET /api/despesas", despesa.Listar)
	mux.HandleFunc("GET /api/despesas/{id}", despesa.BuscarPorID)
	mux.HandleFunc("PUT /api/despesas/{id}", despesa.Atualizar)
	mux.HandleFunc("DELETE /api/despesas/{id}", despesa.Excluir)

	mux.Handle("GET /scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("scripts"))))
	mux.Handle("GET /css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	mux.Handle("GET /", http.FileServer(http.Dir("telas")))

	return mux
}

func responderJSON(w http.ResponseWriter, status int, dado any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(dado)
}

func responderErro(w http.ResponseWriter, status int, msg string) {
	responderJSON(w, status, dtos.ErroDTO{Erro: msg})
}

func tratarErro(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, servicos.ErrInvalido):
		responderErro(w, http.StatusBadRequest, err.Error())
	case errors.Is(err, servicos.ErrNaoEncontrado):
		responderErro(w, http.StatusNotFound, err.Error())
	case errors.Is(err, servicos.ErrConflito):
		responderErro(w, http.StatusConflict, err.Error())
	default:
		responderErro(w, http.StatusInternalServerError, "erro interno")
	}
}

func idDaURL(w http.ResponseWriter, r *http.Request) (int, bool) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		responderErro(w, http.StatusBadRequest, "id inválido")
		return 0, false
	}
	return id, true
}

func lerJSON(w http.ResponseWriter, r *http.Request, destino any) bool {
	if err := json.NewDecoder(r.Body).Decode(destino); err != nil {
		responderErro(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return false
	}
	return true
}

func viagemIDDaQuery(w http.ResponseWriter, r *http.Request) (int, bool) {
	texto := r.URL.Query().Get("viagem_id")
	if texto == "" {
		return 0, true
	}
	id, err := strconv.Atoi(texto)
	if err != nil || id <= 0 {
		responderErro(w, http.StatusBadRequest, "viagem_id inválido")
		return 0, false
	}
	return id, true
}
