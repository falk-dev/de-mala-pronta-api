package dtos

type DespesaDTO struct {
	ID        int     `json:"id"`
	ViagemID  int     `json:"viagem_id"`
	Descricao string  `json:"descricao"`
	Valor     float64 `json:"valor"`
	Data      string  `json:"data"`
	Categoria string  `json:"categoria"`
}
