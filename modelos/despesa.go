package modelos

type Despesa struct {
	ID        int
	ViagemID  int
	Descricao string
	Valor     float64
	Data      string
	Categoria string
}
