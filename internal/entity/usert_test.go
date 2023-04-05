package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("MizunoTeste", "mizuno@gmial.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "MizunoTeste", user.Name)
	assert.Equal(t, "mizuno@gmial.com", user.Email)
}

func TestUser(t *testing.T) {
	user, err := NewUser("MizunoTeste", "mizuno@gmial.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.NotEqual(t, "123456", user.Password)
}
