package service

import (
	"github.com/ha-t2/di-sample/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

// serviceのテストにmysqlが不要になった
func TestProductService_Exist(t *testing.T) {
	var cases = []struct {
		name     string
		input    int
		expected bool
	}{
		{
			name:     "exist",
			input:    1,
			expected: true,
		},
		{
			name:     "not exist",
			input:    2,
			expected: false,
		},
	}

	mockRepo := &MockConstRepo{model.Product{
		Id:    1,
		Price: 100,
	}}
	s := ProductService{Repo: mockRepo}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := s.Exist(c.input)
			assert.Equal(t, got, c.expected)
		})
	}
}

// 固定で一つの値を返すrepository
type MockConstRepo struct {
	product model.Product
}

func (m *MockConstRepo) Get(i int) model.Product {
	return m.product
}
