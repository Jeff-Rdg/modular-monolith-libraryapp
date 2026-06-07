package author

import (
	"modular-monolith-libraryApp/modules/author/port"
	"modular-monolith-libraryApp/modules/author/repository"
	"modular-monolith-libraryApp/modules/author/usecase"

	"go.uber.org/fx"
)

var Module = fx.Module("author",
	fx.Provide(
		fx.Annotate(
			repository.NewAuthorRepository,
			fx.As(new(port.AuthorRepository)),
		),
		usecase.NewCreateAuthor,
		//handler.NewHandler,
	),
)
