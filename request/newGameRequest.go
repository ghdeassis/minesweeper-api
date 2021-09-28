package request

type NewGameRequest struct {
	Rows    *int `json:"rows"    binding:"required"`
	Columns *int `json:"columns" binding:"required"`
	Mines   *int `json:"mines"   binding:"required"`
}
