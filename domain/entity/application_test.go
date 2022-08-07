package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/nurmanhabib/go-oauth2/domain/entity"
)

func TestApplication(t *testing.T) {
	var applicationNil *entity.Application
	var applicationStruct entity.Application

	applicationV1 := entity.Application{Version: entity.ApplicationV1}
	applicationV2 := entity.Application{Version: entity.ApplicationV2}

	t.Run("if empty application given", func(t *testing.T) {
		assert.NotNil(t, applicationStruct)
		assert.False(t, applicationStruct.IsV1())
		assert.False(t, applicationStruct.IsV2())
	})

	t.Run("if nil application given", func(t *testing.T) {
		assert.Nil(t, applicationNil)
		assert.False(t, applicationNil.IsV1())
		assert.False(t, applicationNil.IsV2())
	})

	t.Run("if application v1 given", func(t *testing.T) {
		assert.NotNil(t, applicationV1)
		assert.True(t, applicationV1.IsV1())
		assert.False(t, applicationV1.IsV2())
	})

	t.Run("if application v2 given", func(t *testing.T) {
		assert.NotNil(t, applicationV2)
		assert.False(t, applicationV2.IsV1())
		assert.True(t, applicationV2.IsV2())
	})
}
