package application_test

import (
	mock_application "github.com/anlopes123/hexagonal/application/mocks/application"
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
		Persistense: persistense,
	}

	result, err:= service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)

}