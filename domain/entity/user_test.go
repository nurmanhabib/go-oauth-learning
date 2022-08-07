package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/nurmanhabib/go-oauth2/domain/entity"
)

func TestUser_GetAuthIdentifier(t *testing.T) {
	var nilUser *entity.User
	var emptyUser entity.User

	sampleUser := entity.User{
		ID: 123,
	}

	t.Run("if nil user given", func(t *testing.T) {
		identifier := nilUser.GetAuthIdentifier()

		assert.Equal(t, "0", identifier)
	})

	t.Run("if empty user given", func(t *testing.T) {
		identifier := emptyUser.GetAuthIdentifier()

		assert.Equal(t, "0", identifier)
	})

	t.Run("if sampe user given", func(t *testing.T) {
		identifier := sampleUser.GetAuthIdentifier()

		assert.NotEmpty(t, identifier)
		assert.Equal(t, "123", identifier)
	})
}

func TestUser_GetAuthPassword(t *testing.T) {
	var nilUser *entity.User
	var emptyUser entity.User

	sampleUser := entity.User{
		Password: "secret",
	}

	t.Run("if nil user given", func(t *testing.T) {
		password := nilUser.GetAuthPassword()

		assert.Empty(t, password)
		assert.Equal(t, "", password)
	})

	t.Run("if empty user given", func(t *testing.T) {
		password := emptyUser.GetAuthPassword()

		assert.Empty(t, password)
		assert.Equal(t, "", password)
	})

	t.Run("if sampe user given", func(t *testing.T) {
		password := sampleUser.GetAuthPassword()

		assert.NotEmpty(t, password)
		assert.Equal(t, "secret", password)
	})
}
