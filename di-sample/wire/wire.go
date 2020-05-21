package wire

import (
	"github.com/google/wire"
	"github.com/ha-t2/di-sample/repo"
	"github.com/ha-t2/di-sample/service"
)

func InitializeProductService() service.ProductService {
	wire.Build(service.NewProductService, repo.NewProductRepo)
	return service.ProductService{}
}
