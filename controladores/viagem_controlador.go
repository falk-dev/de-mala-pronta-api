package controladores

import (
	"net/http"

	"de-mala-pronta-api/dtos"
	"de-mala-pronta-api/modelos"
	"de-mala-pronta-api/servicos"
)

type ViagemControlador struct {
	servico *servicos.ViagemServico
}

func NovoViagemControlador(servico *servicos.ViagemServico) *ViagemControlador {
	return &ViagemControlador{servico: servico}
}

func viagemParaDTO(v modelos.Viagem) dtos.ViagemDTO {
	return dtos.ViagemDTO{ID: v.ID, Titulo: v.Titulo, DataInicio: v.DataInicio,
		DataFim: v.DataFim, Orcamento: v.Orcamento, Status: v.Status}
}

func dtoParaViagem(d dtos.ViagemDTO) *modelos.Viagem {
	return &modelos.Viagem{ID: d.ID, Titulo: d.Titulo, DataInicio: d.DataInicio,
		DataFim: d.DataFim, Orcamento: d.Orcamento, Status: d.Status}
}

func (c *ViagemControlador) Cadastrar(w http.ResponseWriter, r *http.Request) {
	var dto dtos.ViagemDTO
	if !lerJSON(w, r, &dto) {
		return
	}
	v := dtoParaViagem(dto)
	if err := c.servico.Cadastrar(v); err != nil {
		tratarErro(w, err)
		return
	}
	responderJSON(w, http.StatusCreated, viagemParaDTO(*v))
}

func (c *ViagemControlador) Listar(w http.ResponseWriter, r *http.Request) {
	viagens, err := c.servico.Listar()
	if err != nil {
		tratarErro(w, err)
		return
	}
	resposta := []dtos.ViagemDTO{}
	for _, v := range viagens {
		resposta = append(resposta, viagemParaDTO(v))
	}
	responderJSON(w, http.StatusOK, resposta)
}

func (c *ViagemControlador) BuscarPorID(w http.ResponseWriter, r *http.Request) {
	id, ok := idDaURL(w, r)
	if !ok {
		return
	}
	v, err := c.servico.BuscarPorID(id)
	if err != nil {
		tratarErro(w, err)
		return
	}
	responderJSON(w, http.StatusOK, viagemParaDTO(*v))
}

func (c *ViagemControlador) Atualizar(w http.ResponseWriter, r *http.Request) {
	id, ok := idDaURL(w, r)
	if !ok {
		return
	}
	var dto dtos.ViagemDTO
	if !lerJSON(w, r, &dto) {
		return
	}
	dto.ID = id
	v := dtoParaViagem(dto)
	if err := c.servico.Atualizar(v); err != nil {
		tratarErro(w, err)
		return
	}
	responderJSON(w, http.StatusOK, viagemParaDTO(*v))
}

func (c *ViagemControlador) Excluir(w http.ResponseWriter, r *http.Request) {
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
