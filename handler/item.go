package handler

import (
	"context"

	"github.com/hidenari-yuda/ai-market-go/infra/database"
	"github.com/hidenari-yuda/ai-market-go/infra/driver"
	"github.com/hidenari-yuda/ai-market-go/pb"
	"github.com/hidenari-yuda/ai-market-go/usecase"
	"github.com/hidenari-yuda/ai-market-go/usecase/interactor"
)

type ItemServiceServer struct {
	pb.UnimplementedItemServiceServer
	ItemInteractor interactor.ItemInteractor
	Db                *database.Db
	Firebase          usecase.Firebase
}

func NewItemSercviceServer(ItemInteractor interactor.ItemInteractor) *ItemServiceServer {
	return &ItemServiceServer{
		ItemInteractor: ItemInteractor,
		Db:                database.NewDb(),
		Firebase:          driver.NewFirebaseImpl(),
	}
}

// create chat group
func (s *ItemServiceServer) Create(ctx context.Context, req *pb.Item) (*pb.Item, error) {

	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}

	res, err := s.ItemInteractor.Create(req)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return res, nil
}

// update chat group
func (s *ItemServiceServer) Update(ctx context.Context, req *pb.Item) (*pb.ChatBoolResponse, error) {

	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}

	res, err := s.ItemInteractor.Update(req)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &pb.ChatBoolResponse{Error: res}, nil
}

// delete chat group
func (s *ItemServiceServer) Delete(ctx context.Context, req *pb.ChatIdRequest) (*pb.ChatBoolResponse, error) {

	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}

	res, err := s.ItemInteractor.Delete(req.Id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &pb.ChatBoolResponse{Error: res}, nil
}

// get chat group by id
func (s *ItemServiceServer) GetById(ctx context.Context, req *pb.ChatIdRequest) (*pb.Item, error) {

	res, err := s.ItemInteractor.GetById(req.Id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// get chat group by user id
func (s *ItemServiceServer) GetListByUser(ctx context.Context, req *pb.ChatIdRequest) (*pb.ItemList, error) {

	res, err := s.ItemInteractor.GetListByUser(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.ItemList{ItemList: res}, nil
}
