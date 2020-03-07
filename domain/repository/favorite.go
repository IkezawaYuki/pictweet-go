package repository

type FavoriteRepository interface {
	Toggle(uint, uint) (bool, error)
}
