package dtos

type DestinoDTO struct {
	ID          int    `json:"id"`
	ViagemID    int    `json:"viagem_id"`
	Cidade      string `json:"cidade"`
	Pais        string `json:"pais"`
	DataChegada string `json:"data_chegada"`
	DataSaida   string `json:"data_saida"`
}
