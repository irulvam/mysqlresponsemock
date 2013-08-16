## REALLY BASIC tool to mock MySQL responses

I needed a simple mock for the [github.com/ziutek/mymysql/mysql](github.com/ziutek/mymysql/mysql) responses. Using [gomock](http://godoc.org/code.google.com/p/gomock/gomock) seamed like overkill.

## Example usage

```go

import (
	"github.com/rzajac/assert/assert"
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
```

## Run tests

    go test github.com/rzajac/mysqlresponsemock

## Installation

    go get github.com/rzajac/mysqlresponsemock

## License

Released under the MIT License.

Mysqlresponsemock (c) Rafal Zajac <rzajac@gmail.com>