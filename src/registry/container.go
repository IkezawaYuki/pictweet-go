package registry

import (
	"github.com/IkezawaYuki/pictweet-go/src/infrastructure/datastore"
	"github.com/IkezawaYuki/pictweet-go/src/usecase"
	"github.com/sarulabs/di"
)

type Container struct {
	ctn di.Container
}

func NewContainer() (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	if err := builder.Add([]di.Def{
		{
			Name:  "tweet-usecase",
			Build: buildTweetUsecase,
		},
	}...); err != nil {
		return nil, err
	}
	return &Container{
		ctn: builder.Build(),
	}, nil
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func (c *Container) Clean() error {
	return c.ctn.Clean()
}

func buildTweetUsecase(ctn di.Container) (interface{}, error) {
	repo := datastore.NewSqlHandler()
	return usecase.NewPictweetUsecase(repo), nil
}
