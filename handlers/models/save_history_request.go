package models

type HistorySaveRequest struct {
	Name  string `json:"name"`
	Path  string `json:"path"`
	Score int    `json:"score"`
}
