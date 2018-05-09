package database

import (
	"github.com/morix1500/clean-architecture-sample/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"reflect"
)

var testData = domain.Blog{
	Id:      1,
	Title:   "Test",
	Content: "Hoge",
}

type dummySqlHandler struct {}

func (dummySqlHandler) Execute(q string, i ...interface{}) (Result, error) {
	return dummyResult{},nil
}

func (dummySqlHandler) Query(q string, i ...interface{}) (Row, error) {
	return dummyRow{}, nil
}

type dummyResult struct{}
func (dummyResult) LastInsertId() (int64, error) {
	return 0, nil
}

func (dummyResult) RowsAffected() (int64, error) {
	return 0, nil
}

type dummyRow struct{}

func (dummyRow) Scan(dest ...interface{}) error {
	for i, _ := range dest {
		dpv := reflect.ValueOf(dest[i])
		dv  := reflect.Indirect(dpv)

		switch i {
		case 0:
			var num int32 = 1
			sv := reflect.ValueOf(num)
			dv.Set(sv)
		case 1:
			sv := reflect.ValueOf("Test")
			dv.Set(sv)
		case 2:
			sv := reflect.ValueOf("Hoge")
			dv.Set(sv)
		}
	}

	return nil
}

func (dummyRow) Next() bool {
	return false
}

func (dummyRow) Close() error {
	return nil
}

func TestInsert(t *testing.T) {
	blogRepository := BlogRepository{SqlHandler: dummySqlHandler{}}
	err := blogRepository.Insert(testData)
	assert.Equal(t, err, nil)
}

func TestSelect(t *testing.T) {
	blogRepository := BlogRepository{SqlHandler: dummySqlHandler{}}
	var id int32 = 1
	b, err := blogRepository.Select(id)
	expect := testData

	assert.Equal(t, b, expect)
	assert.Equal(t, err, nil)
}
