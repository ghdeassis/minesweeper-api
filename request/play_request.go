package request

type PlayRequest struct {
	GameID *int `json:"gameId" binding:"required"`
	Row    *int `json:"row"    binding:"required"`
	Column *int `json:"column" binding:"required"`
}
