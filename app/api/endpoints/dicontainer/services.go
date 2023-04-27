package dicontainer

import (
	"treinamento/core/interfaces/primary"
	"treinamento/core/services"
)

func GetBookServices() primary.BookManager {
	return services.NewBookServices(GetBookRepository())
}
