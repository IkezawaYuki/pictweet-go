package port

import "github.com/IkezawaYuki/pictweet-go/domain/model"

type OutputPort interface {
	Index()
}

type OutputData struct {
	Tweets   *model.Tweets   `json:tweets,omitempty`
	Tweet    *model.Tweet    `json:"tweet,omitempty"`
	Comments *model.Comments `json:"comments,omitempty"`
	Comment  *model.Comment  `json:"comment,omitempty"`
	Users    *model.Users    `json:"users,omitempty"`
	User     *model.User     `json:"user,omitempty"`
}
