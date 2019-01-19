package dbrepository

import "github.com/priteshgudge/mongorestaurantsample/domain"

//Reader read from db
type Reader interface {
	Get(id domain.ID) (*domain.Restaurant, error)
	GetAll() ([]*domain.Restaurant, error)
	//Regex Substring Match on the name field
	FindByName(name string) ([]*domain.Restaurant, error)
}

//Writer  write to db
type Writer interface {
	//Create Or update
	Store(b *domain.Restaurant) (domain.ID, error)
	Delete(id domain.ID) error
}

type Filter interface {
	FindByTypeOfFood(foodType string) ([]*domain.Restaurant, error)
	FindByTypeOfPostCode(postCode string) ([]*domain.Restaurant, error)
	//Search --> across all string fields regex match with case insensitive
	//substring match accross all string fields
	Search(query string) ([]*domain.Restaurant, error)
}

//Repository db interface
type Repository interface {
	Reader
	Writer
	//Filter
}
