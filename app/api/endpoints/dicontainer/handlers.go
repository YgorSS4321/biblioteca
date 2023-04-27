package dicontainer

import "treinamento/app/api/endpoints/handlers"

func GetBookHandlers() *handlers.BookHandlers {
	return handlers.NewBookHandlers(GetBookServices())
}
