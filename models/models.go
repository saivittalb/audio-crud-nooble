package models

// Short schema of the short table
type Short struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Category     string `json:"category"`
	AudioFile    string `json:"audiofile"`
	CreatorName  string `json:"creatorname"`
	CreatorEmail string `json:"creatoremail"`
	Date         string `json:"date"`
}
