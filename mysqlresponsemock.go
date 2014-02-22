// REALLY Basic primitives to mock github.com/ziutek/mymysql responses
//
// Mysqlresponsemock (c) Rafal Zajac <rzajac@gmail.com>
// https://github.com/rzajac/mysqlresponsemock
//
// Licensed under the MIT license

// REALLY Basic primitives to mock github.com/ziutek/mymysql responses
package mysqlresponsemock

import (
	"github.com/ziutek/mymysql/mysql"
)

// ResultMock is coppied from the original code.
type ResultMock struct {
	Field_count int
	Fieldsf     []*mysql.Field // Fields table
	Fc_map      map[string]int // Maps field name to column number
}

// NewResponseMock mocks MySQL response.
//
// Example:
//
//     row, resp := NewResponseMock("Id", 1, "Name", "Tom")
//
func NewResponseMock(v ...interface{}) (mysql.Row, *ResultMock) {

	if len(v)%2 != 0 {
		panic("Number of arguments must be even")
	}

	var colValues []interface{}
	var colName string
	var ok bool
	colNames := []string{}

	for i := 0; i < len(v); i += 2 {
		colName, ok = v[i].(string)
		if !ok {
			panic("Columns must be strings")
		}
		colNames = append(colNames, colName)
		colValues = AddColValue(colValues, v[i+1])
	}

	return colValues, NewResultMock(colNames)
}

// AddColValue adds column value to the end of the mocked mysql.Row.
func AddColValue(colValues []interface{}, v interface{}) []interface{} {
	switch v.(type) {
	case int:
		colValues = append(colValues, int32(v.(int)))
	case uint:
		colValues = append(colValues, uint32(v.(uint)))

	default:
		colValues = append(colValues, v)
	}

	return colValues
}

// NewResultMock mocks mysql.Result
// NOTE: Right now it only sets mapping between column indexes and column names.
func NewResultMock(columns []string) *ResultMock {
	rm := new(ResultMock)
	rm.Field_count = len(columns)
	rm.Fc_map = make(map[string]int)

	for idx, colName := range columns {
		rm.Fc_map[colName] = idx
	}
	return rm
}

// AddColumn adds one more column to the mocked mysql.Result.
func (res *ResultMock) AddColumn(colName string) {
	res.Fc_map[colName] = res.Field_count
	res.Field_count++
}

// Returns index for given name or -1 if field of that name doesn't exist
func (res *ResultMock) Map(field_name string) int {
	if fi, ok := res.Fc_map[field_name]; ok {
		return fi
	}
	return -1
}

// Returns true if this is status result that includes no result set
func (res *ResultMock) StatusOnly() bool {
	panic("Not implemented.")
}

// Returns a table containing descriptions of the columns
func (res *ResultMock) Fields() []*mysql.Field {
	panic("Not implemented.")
}

func (res *ResultMock) Message() string {
	panic("Not implemented.")
}

func (res *ResultMock) AffectedRows() uint64 {
	panic("Not implemented.")
}

func (res *ResultMock) InsertId() uint64 {
	panic("Not implemented.")
}

func (res *ResultMock) WarnCount() int {
	panic("Not implemented.")
}

func (res *ResultMock) MakeRow() mysql.Row {
	panic("Not implemented.")
}

func (r *ResultMock) End() error {
	panic("Not implemented.")
}

func (r *ResultMock) GetFirstRow() (mysql.Row, error) {
	panic("Not implemented.")
}

func (r *ResultMock) GetLastRow() (mysql.Row, error) {
	panic("Not implemented.")
}

func (r *ResultMock) GetRow() (mysql.Row, error) {
	panic("Not implemented.")
}

func (r *ResultMock) GetRows() ([]mysql.Row, error) {
	panic("Not implemented.")
}

func (r *ResultMock) MoreResults() bool {
	return false
}

func (r *ResultMock) NextResult() (mysql.Result, error) {
	panic("Not implemented.")
}

func (r *ResultMock) ScanRow(row mysql.Row) error {
	panic("Not implemented.")
}
