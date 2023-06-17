package spec

import (
	"github.com/stretchr/testify/assert"
	"quebrada_api/internal/domain/spec"
	"testing"
)

func TestGetUsersActivedSpec(t *testing.T) {
	userSpec := spec.GetUsersActivedSpec()
	query := userSpec.GetQuery()
	value := userSpec.GetValues()
	assert.Equal(t, "active = ?", query)
	assert.Equal(t, []interface{}{true}, value)
}

func TestGetUserWithEmailSpec(t *testing.T) {
	userSpec := spec.GetUserWithEmailSpec("marcos@gmail.com")
	query := userSpec.GetQuery()
	value := userSpec.GetValues()
	assert.Equal(t, "email = ?", query)
	assert.Equal(t, []interface{}{"marcos@gmail.com"}, value)
}
