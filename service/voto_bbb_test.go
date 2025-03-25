package service

import (
	"testing"

	"go.uber.org/mock/gomock"
)

func Test_CreateVoto_WhenReturn_Sucess(t *testing.T) {
ctrl := gomock.NewController(t)

m := NewMockService(ctrl)

m.EXPECT()
}