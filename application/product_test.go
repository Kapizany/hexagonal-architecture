package application_test

import (
	"testing"

	"github.com/Kapizany/hexagonal-architecture/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Test Product"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "THE PRICE MUST BE GREATER THAN ZERO TO ENABLE THE PRODUCT", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Test Product"
	product.Status = application.ENABLED
	product.Price = 10

	err := product.Disable()
	require.Equal(t, "THE PRICE MUST BE ZERO IN ORDER TO HAVE THE PRODUCT DISABLED", err.Error())

	product.Price = 0
	err = product.Disable()
	require.Nil(t, err)
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Test Product"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "STATUS MUST BE ENABLED OR DISABLED", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "THE PRICE MUST BE GREATER OR EQUAL TO ZERO", err.Error())

	product.Price = 0
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Name = ""
	_, err = product.IsValid()
	require.Error(t, err)

	product.Name = ""
	_, err = product.IsValid()
	require.Error(t, err)

	product.Name = "Product"
	_, err = product.IsValid()
	require.Nil(t, err)
}
