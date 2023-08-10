package core_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/martoc/semver/core"
)

func TestClass_Get(t *testing.T) {
	t.Parallel()
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// When
	mockPlaceHolder := core.NewMockPlaceHolder(ctrl) // Create a mock implementation
	mockPlaceHolder.EXPECT().Get().Return()          // Expect the Get() method to be called
	class := &core.Class{
		PlaceHolder: mockPlaceHolder, // Set the mock implementation on the Class struct
	}
	// Then
	class.Get()
}
