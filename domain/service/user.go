package service

import "github.com/IkezawaYuki/pictweet-go/domain/repository"

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepo,
	}
}

func (service *UserService) Exist(email string) bool {
	user, _ := service.userRepository.FindByEmail(email)
	if user != nil {
		return true
	}
	return false
}
