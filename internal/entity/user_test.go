package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Pedro Yoshimura", "ropehapi@gmail.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Pedro Yoshimura", user.Name)
	assert.Equal(t, "ropehapi@gmail.com", user.Email)
}

func TestValidatePassword(t *testing.T){
	user, err := NewUser("Pedro Yoshimura", "ropehapi@gmail.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("123"))
	assert.NotEqual(t, "123456", user.Password)
}