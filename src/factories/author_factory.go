package factories

import (
	"github.com/matheus-dr/go-test/src/controllers"
	"github.com/matheus-dr/go-test/src/repositories"
	"github.com/matheus-dr/go-test/src/services"
)

func AuthorFactory() controllers.AuthorController {
	repository := repositories.AuthorRepository{}

	service := services.AuthorService{Repository: repository}

	return controllers.AuthorController{Service: service}
}
