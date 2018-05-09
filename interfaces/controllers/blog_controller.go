package controllers

import (
	"github.com/morix1500/clean-architecture-sample/domain"
	"github.com/morix1500/clean-architecture-sample/interfaces/database"
	"github.com/morix1500/clean-architecture-sample/usecase"
	pb "github.com/morix1500/clean-architecture-sample/proto"
)

type BlogController struct {
	Interactor usecase.BlogInteractor
}

func NewBlogController(sqlHandler database.SqlHandler) *BlogController {
	return &BlogController{
		Interactor: usecase.BlogInteractor{
			BlogRepository: &database.BlogRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (ctr *BlogController) Insert(req *pb.InsertRequest) (*pb.InsertResponse, error) {
	b := domain.Blog{}
	b.Id = req.Id
	b.Title = req.Title
	b.Content = req.Content
	err := ctr.Interactor.Add(b)
	if err != nil {
		return &pb.InsertResponse{}, err
	}
	return &pb.InsertResponse{}, nil
}

func (ctr *BlogController) Select(req *pb.SelectRequest) (*pb.SelectResponse, error) {
	id := req.Id
	b, err := ctr.Interactor.BlogById(id)
	if err != nil {
		panic(err)
	}
	return &pb.SelectResponse{
		Id: b.Id,
		Title: b.Title,
		Content: b.Content,
	}, nil
}
