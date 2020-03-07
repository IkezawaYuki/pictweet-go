package dto

type ToggleFavoriteVideoResponse struct {
	TweetID    string `json:"video_id"`
	IsFavorite bool   `json:"is_favorite"`
}
