package factories

import (
	"github.com/matheus-dr/go-test/src/controllers"
	"github.com/matheus-dr/go-test/src/repositories"
	"github.com/matheus-dr/go-test/src/services"
)

func CategoryFactory() controllers.CategoryController {
	repository := repositories.CategoryRepository{}

	service := services.CategoryService{Repository: repository}

	return controllers.CategoryController{Service: service}
}
