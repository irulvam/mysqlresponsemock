package mysqlresponsemock

import (
	"github.com/rzajac/assert/assert"
	"testing"
)

func Test_NewResponseMock(t *testing.T) {
	row, resp := NewResponseMock("Id", 1, "Name", "Tom", "Balance", 123.12)

	assert.Equal(t, 0, resp.Map("Id"))
	assert.Equal(t, 1, resp.Map("Name"))
	assert.Equal(t, 2, resp.Map("Balance"))

	assert.Equal(t, 1, row.Int(0))
	assert.Equal(t, "Tom", row.Str(1))
	assert.Equal(t, 123.12, row.Float(2))
}
