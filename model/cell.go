package model

type Cell struct {
	HasMine    bool `json:"hasMine"`
	IsFlagged  bool `json:"isFlagged"`
	IsRevealed bool `json:"isRevealed"`
}
