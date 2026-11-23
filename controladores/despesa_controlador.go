package controladores

import (
	"net/http"

	"de-mala-pronta-api/dtos"
	"de-mala-pronta-api/modelos"
	"de-mala-pronta-api/servicos"
)

type DespesaControlador struct {
	servico *servicos.DespesaServico
}

func NovoDespesaControlador(servico *servicos.DespesaServico) *DespesaControlador {
	return &DespesaControlador{servico: servico}
}

func despesaParaDTO(d modelos.Despesa) dtos.DespesaDTO {
	return dtos.DespesaDTO{ID: d.ID, ViagemID: d.ViagemID, Descricao: d.Descricao,
		Valor: d.Valor, Data: d.Data, Categoria: d.Categoria}
}

func dtoParaDespesa(d dtos.DespesaDTO) *modelos.Despesa {
	return &modelos.Despesa{ID: d.ID, ViagemID: d.ViagemID, Descricao: d.Descricao,
		Valor: d.Valor, Data: d.Data, Categoria: d.Categoria}
}

func (c *DespesaControlador) Cadastrar(w http.ResponseWriter, r *http.Request) {
	var dto dtos.DespesaDTO
	if !lerJSON(w, r, &dto) {
		return
	}
	d := dtoParaDespesa(dto)
	if err := c.servico.Cadastrar(d); err != nil {
		tratarErro(w, err)
		return
	}
	responderJSON(w, http.StatusCreated, despesaParaDTO(*d))
}

func (c *DespesaControlador) Listar(w http.ResponseWriter, r *http.Request) {
	viagemID, ok := viagemIDDaQuery(w, r)
	if !ok {
		return
	}
	despesas, err := c.servico.Listar(viagemID)
	if err != nil {
		tratarErro(w, err)
		return
	}
	resposta := []dtos.DespesaDTO{}
	for _, d := range despesas {
		resposta = append(resposta, despesaParaDTO(d))
	}
	responderJSON(w, http.StatusOK, resposta)
}

func (c *DespesaControlador) BuscarPorID(w http.ResponseWriter, r *http.Request) {
	id, ok := idDaURL(w, r)
	if !ok {
		return
	}
	d, err := c.servico.BuscarPorID(id)
	if err != nil {
		tratarErro(w, err)
		return
	}
	responderJSON(w, http.StatusOK, despesaParaDTO(*d))
}

func (c *DespesaControlador) Atualizar(w http.ResponseWriter, r *http.Request) {
	id, ok := idDaURL(w, r)
	if !ok {
		return
	}
	var dto dtos.DespesaDTO
	if !lerJSON(w, r, &dto) {
		return
	}
	dto.ID = id
	d := dtoParaDespesa(dto)
	if err := c.servico.Atualizar(d); err != nil {
		tratarErro(w, err)
		return
	}
	responderJSON(w, http.StatusOK, despesaParaDTO(*d))
}

func (c *DespesaControlador) Excluir(w http.ResponseWriter, r *http.Request) {
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

func (c *DespesaControlador) TotalPorViagem(w http.ResponseWriter, r *http.Request) {
	id, ok := idDaURL(w, r)
	if !ok {
		return
	}
	total, err := c.servico.TotalPorViagem(id)
	if err != nil {
		tratarErro(w, err)
		return
	}
	responderJSON(w, http.StatusOK, dtos.TotalDespesasDTO{ViagemID: id, Total: total})
}
