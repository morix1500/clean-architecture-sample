package controllers

import (
	"github.com/morix1500/clean-architecture-sample/domain"
	"github.com/morix1500/clean-architecture-sample/interfaces/database"
	pb "github.com/morix1500/clean-architecture-sample/proto"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var testData = domain.Blog{
	Id:      1,
	Title:   "Test",
	Content: "Hoge",
}

type dummySqlHandler struct{}

func (dummySqlHandler) Execute(q string, i ...interface{}) (database.Result, error) {
	return dummyResult{}, nil
}

func (dummySqlHandler) Query(q string, i ...interface{}) (database.Row, error) {
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
		dv := reflect.Indirect(dpv)

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
	blogController := NewBlogController(dummySqlHandler{})

	req := pb.InsertRequest{
		Id:      testData.Id,
		Title:   testData.Title,
		Content: testData.Content,
	}
	_, err := blogController.Insert(&req)
	assert.Equal(t, err, nil)
}

func TestSelect(t *testing.T) {
	blogController := NewBlogController(dummySqlHandler{})

	req := pb.SelectRequest{
		Id: 1,
	}
	b, err := blogController.Select(&req)
	expect := &pb.SelectResponse{
		Id:      testData.Id,
		Title:   testData.Title,
		Content: testData.Content,
	}
	assert.Equal(t, b, expect)
	assert.Equal(t, err, nil)
}
