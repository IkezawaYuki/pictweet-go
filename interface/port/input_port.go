package port

import (
	"github.com/IkezawaYuki/pictweet-go/domain/entity"
)

type InputPort interface {
	Index() (*entity.Tweets, error)
}
