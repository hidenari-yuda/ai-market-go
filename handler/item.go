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
	Db             *database.Db
	Firebase       usecase.Firebase
}

func NewItemSercviceServer(ItemInteractor interactor.ItemInteractor) *ItemServiceServer {
	return &ItemServiceServer{
		ItemInteractor: ItemInteractor,
		Db:             database.NewDb(),
		Firebase:       driver.NewFirebaseImpl(),
	}
}

// create Item group
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

// update Item group
func (s *ItemServiceServer) Update(ctx context.Context, req *pb.Item) (*pb.ItemBoolResponse, error) {

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

	return &pb.ItemBoolResponse{Error: res}, nil
}

// delete Item group
func (s *ItemServiceServer) Delete(ctx context.Context, req *pb.ItemIdRequest) (*pb.ItemBoolResponse, error) {

	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}

	res, err := s.ItemInteractor.Delete(req)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &pb.ItemBoolResponse{Error: res}, nil
}

// get Item group by id
func (s *ItemServiceServer) GetById(ctx context.Context, req *pb.ItemIdRequest) (*pb.Item, error) {

	res, err := s.ItemInteractor.GetById(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// get Item group by user id
func (s *ItemServiceServer) GetListByUser(ctx context.Context, req *pb.ItemUserIdRequest) (*pb.ItemList, error) {

	res, err := s.ItemInteractor.GetListByUser(req)
	if err != nil {
		return nil, err
	}

	return &pb.ItemList{Item: res}, nil
}

// get list by search
func (s *ItemServiceServer) GetListBySearch(ctx context.Context, req *pb.ItemSearchRequest) (*pb.ItemList, error) {

	res, err := s.ItemInteractor.GetListBySearch(req)
	if err != nil {
		return nil, err
	}

	return &pb.ItemList{Item: res}, nil
}

// get latest id=user_id
func (s *ItemServiceServer) GetLatestList(ctx context.Context, req *pb.ItemUserIdRequest) (*pb.ItemList, error) {

	res, err := s.ItemInteractor.GetLatestList(req)
	if err != nil {
		return nil, err
	}

	return &pb.ItemList{Item: res}, nil
}

// get trend list by user id
func (s *ItemServiceServer) GetTrendList(ctx context.Context, req *pb.ItemUserIdRequest) (*pb.ItemList, error) {

	res, err := s.ItemInteractor.GetTrendList(req)
	if err != nil {
		return nil, err
	}

	return &pb.ItemList{Item: res}, nil
}

// get recommended list by user id
func (s *ItemServiceServer) GetRecommendedListByUser(ctx context.Context, req *pb.ItemUserIdRequest) (*pb.ItemList, error) {

	res, err := s.ItemInteractor.GetRecommendedListByUser(req)
	if err != nil {
		return nil, err
	}

	return &pb.ItemList{Item: res}, nil
}

// get sold list by user id
func (s *ItemServiceServer) GetSoldListByUser(ctx context.Context, req *pb.ItemUserIdRequest) (*pb.ItemList, error) {

	res, err := s.ItemInteractor.GetSoldListByUser(req)
	if err != nil {
		return nil, err
	}

	return &pb.ItemList{Item: res}, nil
}

// get purchased list by user id
func (s *ItemServiceServer) GetPurchasedListByUser(ctx context.Context, req *pb.ItemUserIdRequest) (*pb.ItemList, error) {

	res, err := s.ItemInteractor.GetPurchasedListByUser(req)
	if err != nil {
		return nil, err
	}

	return &pb.ItemList{Item: res}, nil
}

// get liked list by user id
func (s *ItemServiceServer) GetLikedListByUser(ctx context.Context, req *pb.ItemUserIdRequest) (*pb.ItemList, error) {

	res, err := s.ItemInteractor.GetLikedListByUser(req)
	if err != nil {
		return nil, err
	}

	return &pb.ItemList{Item: res}, nil
}

// get list by id list
func (s *ItemServiceServer) GetListByIdList(ctx context.Context, req *pb.ItemIdListRequest) (*pb.ItemList, error) {

	res, err := s.ItemInteractor.GetListByIdList(req)
	if err != nil {
		return nil, err
	}

	return &pb.ItemList{Item: res}, nil
}
