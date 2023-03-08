package cli_test
import(
	"testing"
	"github.com/golang/mock/gomock"
	mock_application "github.com/anlopes123/hexagonal/application/mocks"
	"github.com/stretchr/testify/require"
	"fmt"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName:= "Product Test"
	productPrice:= 25.99
	productStatus:= "enabled"
	productId:= "abc"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gpmock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gpmock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s with the name %s has been create with the price %f and status %s", 
	      productId, productName, productPrice, productStatus)

	result, err:= cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(resultExpected, result)	  

}