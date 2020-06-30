package utils

import (
	"github.com/edwardsuwirya/simpleSql/utils/appStatus"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestAppStatus(t *testing.T) {
	t.Log("It should success return correct message")
	statusMessageTest := appStatus.StatusText(0)
	assert.Equal(t, statusMessageTest, "Sukses")
}
