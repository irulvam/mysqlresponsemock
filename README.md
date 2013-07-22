## REALLY simple tool to mock MySQL responses

I needed a simple response mock for the [github.com/ziutek/mymysql/mysql](github.com/ziutek/mymysql/mysql) responses. Using [gomock](http://godoc.org/code.google.com/p/gomock/gomock) seamed like overkill.

## Example usage

```go

import (
	"github.com/rzajac/assert/assert"
)

type testStruct struct {
	Id    string
	Field string
}

func Test_initFromDb(t *testing.T) {
	row, res := mysqlresultmock.NewResponseMock("Id", 4, "Field", 5)

	tStruct := new(testStruct)
	err := tStruct.initFromDB(row, res)
	assert.NotError(t, err)

	assert.Equal(t, "4", tStruct.Id)
	assert.Equal(t, "5", tStruct.Field)
}
```

## License

Released under the MIT License.

Mysqlresponsemock (c) Rafal Zajac <rzajac@gmail.com>