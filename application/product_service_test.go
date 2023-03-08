package application_test

import (
	"github.com/anlopes123/hexagonal/application"
	mock_application "github.com/anlopes123/hexagonal/application/mocks"
	"testing"
	"github.com/stretchr/testify/require"	
	"github.com/golang/mock/gomock"	
)
func TestProductService_Get(t *testing.T) {
	ctrl:= gomock.NewController(t)
	defer ctrl.Finish()

	product:= mock_application.NewMockProductInterface(ctrl)
	persistense:= mock_application.NewMockProductPersistenceInterface(ctrl)
	persistense.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service:= application.ProductService {
		Persistence: persistense,
	}

	result, err:= service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)

}

func TestProductService_Create(t *testing.T) {
	ctrl:= gomock.NewController(t)
	defer ctrl.Finish()

	product:= mock_application.NewMockProductInterface(ctrl)
	persistense:= mock_application.NewMockProductPersistenceInterface(ctrl)
	persistense.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service:= application.ProductService {
		Persistence: persistense,
	}
	result, err:= service.Create("Product 1", 10)
	require.Nil(t, err)
	require.Equal(t, product, result)

}

func TestProductService_Enable(t *testing.T) {
	ctrl:= gomock.NewController(t)
	defer ctrl.Finish()

	product:= mock_application.NewMockProductInterface(ctrl)
	product.EXPECT().Enable().Return(nil)
	persistense:= mock_application.NewMockProductPersistenceInterface(ctrl)
	persistense.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service:= application.ProductService {
		Persistence: persistense,
	}

	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}
	
func TestProductService_Disable(t *testing.T) {
	ctrl:= gomock.NewController(t)
	defer ctrl.Finish()

	product:= mock_application.NewMockProductInterface(ctrl)
	product.EXPECT().Disable().Return(nil)
	persistense:= mock_application.NewMockProductPersistenceInterface(ctrl)
	persistense.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service:= application.ProductService {
		Persistence: persistense,
	}

	result, err := service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}
	