package service

import (
	"github.com/ha-t2/di-sample/model"
)

type ProductService struct {
	Repo ProductRepoInterface
}

func NewProductService(repo ProductRepoInterface) ProductService {
	return ProductService{Repo: repo}
}

// repo(Repository)を呼び出して、存在チェックをする。
func (s *ProductService) Exist(id int) bool {
	product := s.Repo.Get(id)
	return product.Id == id
}

type ProductRepoInterface interface {
	Get(id int) model.Product
}
