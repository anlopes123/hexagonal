package application_test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/anlopes123/hexagonal/application"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED 
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)
}
