package request

type MessageFormRequest struct {
	Type       int      `json:"type"`
	UserID     []uint64 `json:"userId"`
	FormUserID uint64   `json:"formUserId"`
	Icon       string   `json:"icon"`
	Image      string   `json:"image"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
}
