//+build wireinject

package wire

import (
	"github.com/google/wire"
	"my-clean-rchitecture/api/handler"
	"my-clean-rchitecture/app"
	"my-clean-rchitecture/domain/repo"
	"my-clean-rchitecture/domain/service"
)

func I() app.Server {
	wire.Build(
		app.NewDB,
		repo.NewRepos,
		service.NewServices,
		handler.NewHandlers,
		app.NewServer,
	)

	return app.Server{}
}
