package database_sqlite

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUser(t *testing.T) {
	err := OpenSqlite("./go-in-action.db")
	assert.Nil(t, err)
	expectedser := &User{
		Name:     "gopher",
		Password: "654321",
	}
	err = CreateUser(expectedser)
	assert.Nil(t, err)
	assert.Greater(t, expectedser.Id, int64(0))
	actualUser, err := QueryUser(expectedser.Id)
	assert.Nil(t, err)
	assert.Equal(t, expectedser, actualUser)
	err = CloseSqlite()
	assert.Nil(t, err)
}

func TestDeleteUser(t *testing.T) {
	err := OpenSqlite("./go-in-action.db")
	assert.Nil(t, err)

	expectedser := &User{
		Name:     "gopher",
		Password: "654321",
	}
	err = CreateUser(expectedser)
	assert.Nil(t, err)
	assert.Greater(t, expectedser.Id, int64(0))
	err = DeleteUser(expectedser.Id)
	assert.Nil(t, err)

	err = CloseSqlite()
	assert.Nil(t, err)
}
