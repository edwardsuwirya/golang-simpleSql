package appStatus

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestAppStatus(t *testing.T) {
	t.Log("It should success return correct message")
	statusMessageTest := StatusText(0)
	assert.Equal(t, statusMessageTest, "Sukses")
}
