package resources

import (
	"github.com/stretchr/testify/assert"
	"quebrada_api/resources"
	"testing"
)

func TestLoadResource(t *testing.T) {
	err := resources.LoadTemplates()
	assert.Nil(t, err)
	temp := resources.VerificationEmailTemplate.GetTemplate()
	assert.NotNil(t, temp.Name())
}
