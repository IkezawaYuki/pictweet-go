package repository

type TweetRepository interface {
	FindByID(int)
	FindAll()
	Create()
	Update()
	Delete()
}
