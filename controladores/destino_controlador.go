package controladores

import (
	"net/http"

	"de-mala-pronta-api/dtos"
	"de-mala-pronta-api/modelos"
	"de-mala-pronta-api/servicos"
)

type DestinoControlador struct {
	servico *servicos.DestinoServico
}

func NovoDestinoControlador(servico *servicos.DestinoServico) *DestinoControlador {
	return &DestinoControlador{servico: servico}
}

func destinoParaDTO(d modelos.Destino) dtos.DestinoDTO {
	return dtos.DestinoDTO{ID: d.ID, ViagemID: d.ViagemID, Cidade: d.Cidade, Pais: d.Pais,
		DataChegada: d.DataChegada, DataSaida: d.DataSaida}
}

func dtoParaDestino(d dtos.DestinoDTO) *modelos.Destino {
	return &modelos.Destino{ID: d.ID, ViagemID: d.ViagemID, Cidade: d.Cidade, Pais: d.Pais,
		DataChegada: d.DataChegada, DataSaida: d.DataSaida}
}

func (c *DestinoControlador) Cadastrar(w http.ResponseWriter, r *http.Request) {
	var dto dtos.DestinoDTO
	if !lerJSON(w, r, &dto) {
		return
	}
	d := dtoParaDestino(dto)
	if err := c.servico.Cadastrar(d); err != nil {
		tratarErro(w, err)
		return
	}
	responderJSON(w, http.StatusCreated, destinoParaDTO(*d))
}

func (c *DestinoControlador) Listar(w http.ResponseWriter, r *http.Request) {
	viagemID, ok := viagemIDDaQuery(w, r)
	if !ok {
		return
	}
	destinos, err := c.servico.Listar(viagemID)
	if err != nil {
		tratarErro(w, err)
		return
	}
	resposta := []dtos.DestinoDTO{}
	for _, d := range destinos {
		resposta = append(resposta, destinoParaDTO(d))
	}
	responderJSON(w, http.StatusOK, resposta)
}

func (c *DestinoControlador) BuscarPorID(w http.ResponseWriter, r *http.Request) {
	id, ok := idDaURL(w, r)
	if !ok {
		return
	}
	d, err := c.servico.BuscarPorID(id)
	if err != nil {
		tratarErro(w, err)
		return
	}
	responderJSON(w, http.StatusOK, destinoParaDTO(*d))
}

func (c *DestinoControlador) Atualizar(w http.ResponseWriter, r *http.Request) {
	id, ok := idDaURL(w, r)
	if !ok {
		return
	}
	var dto dtos.DestinoDTO
	if !lerJSON(w, r, &dto) {
		return
	}
	dto.ID = id
	d := dtoParaDestino(dto)
	if err := c.servico.Atualizar(d); err != nil {
		tratarErro(w, err)
		return
	}
	responderJSON(w, http.StatusOK, destinoParaDTO(*d))
}

func (c *DestinoControlador) Excluir(w http.ResponseWriter, r *http.Request) {
	id, ok := idDaURL(w, r)
	if !ok {
		return
	}
	if err := c.servico.Excluir(id); err != nil {
		tratarErro(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
