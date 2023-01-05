package jwt_authentication

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJwtAuth(t *testing.T) {
	userId := "1234"
	token := GenerateToken(userId)
	t.Log(token)
	actualUserId, err := VerifyToken(token)
	assert.Nil(t, err)
	assert.Equal(t, userId, actualUserId)
}
