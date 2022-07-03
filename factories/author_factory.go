package factories

import (
	"github.com/matheus-dr/go-test/controllers"
	"github.com/matheus-dr/go-test/repositories"
	"github.com/matheus-dr/go-test/services"
)

func AuthorFactory() controllers.AuthorController {
	repository := repositories.AuthorRepository{}

	service := services.AuthorService{Repository: repository}

	return controllers.AuthorController{Service: service}
}
