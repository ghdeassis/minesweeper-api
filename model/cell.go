package model

type Cell struct {
	HasMine bool `json:"hasMine"`
	Flagged bool `json:"flagged"`
}
