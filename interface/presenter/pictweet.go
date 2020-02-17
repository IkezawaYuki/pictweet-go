package presenter

import "github.com/IkezawaYuki/pictweet-go/interface/port"

type pictweetPresenter struct {
}

func NewPictweetPresenter() port.OutputPort {
	return &pictweetPresenter{}
}

func (p *pictweetPresenter) Index() {

}
