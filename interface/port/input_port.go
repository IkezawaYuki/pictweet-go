package port

import "github.com/IkezawaYuki/pictweet-go/domain/model"

type InputPort interface {
	Index() (*model.Tweets, error)
}
