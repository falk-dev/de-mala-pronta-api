package dtos

type ViagemDTO struct {
	ID         int     `json:"id"`
	Titulo     string  `json:"titulo"`
	DataInicio string  `json:"data_inicio"`
	DataFim    string  `json:"data_fim"`
	Orcamento  float64 `json:"orcamento"`
	Status     string  `json:"status"`
}

type TotalDespesasDTO struct {
	ViagemID int     `json:"viagem_id"`
	Total    float64 `json:"total"`
}

type ErroDTO struct {
	Erro string `json:"erro"`
}
