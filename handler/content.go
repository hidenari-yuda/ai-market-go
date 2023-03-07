package handler

import (
	"context"

	"github.com/hidenari-yuda/ai-market-go/infra/database"
	"github.com/hidenari-yuda/ai-market-go/infra/driver"
	"github.com/hidenari-yuda/ai-market-go/pb"
	"github.com/hidenari-yuda/ai-market-go/usecase"
	"github.com/hidenari-yuda/ai-market-go/usecase/interactor"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ContentServiceServer struct {
	pb.UnimplementedContentServiceServer
	ContentInteractor interactor.ContentInteractor
	Db                *database.Db
	Firebase          usecase.Firebase
}

func NewContentSercviceServer(ContentInteractor interactor.ContentInteractor) *ContentServiceServer {
	return &ContentServiceServer{
		ContentInteractor: ContentInteractor,
		Db:                database.NewDb(),
		Firebase:          driver.NewFirebaseImpl(),
	}
}

// create Content group
func (s *ContentServiceServer) Create(ctx context.Context, req *pb.Content) (*pb.Content, error) {

	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}

	res, err := s.ContentInteractor.Create(req)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return res, nil
}

// update Content group
func (s *ContentServiceServer) Update(ctx context.Context, req *pb.Content) (*pb.ContentBoolResponse, error) {

	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}

	res, err := s.ContentInteractor.Update(req)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &pb.ContentBoolResponse{Error: res}, nil
}

// delete Content group
func (s *ContentServiceServer) Delete(ctx context.Context, req *pb.ContentIdRequest) (*pb.ContentBoolResponse, error) {

	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}

	res, err := s.ContentInteractor.Delete(req)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &pb.ContentBoolResponse{Error: res}, nil
}

// get Content group by id
func (s *ContentServiceServer) GetById(ctx context.Context, req *pb.ContentIdRequest) (*pb.Content, error) {

	res, err := s.ContentInteractor.GetById(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// get Content group by user id
func (s *ContentServiceServer) GetListByUser(ctx context.Context, req *pb.ContentUserIdRequest) (*pb.ContentList, error) {

	res, err := s.ContentInteractor.GetListByUser(req)
	if err != nil {
		return nil, err
	}

	return &pb.ContentList{Content: res}, nil
}

// get list by search
func (s *ContentServiceServer) GetListBySearch(ctx context.Context, req *pb.ContentSearchRequest) (*pb.ContentList, error) {

	res, err := s.ContentInteractor.GetListBySearch(req)
	if err != nil {
		return nil, err
	}

	return &pb.ContentList{Content: res}, nil
}

// get latest id=user_id
func (s *ContentServiceServer) GetLatestList(ctx context.Context, req *emptypb.Empty) (*pb.ContentList, error) {

	res, err := s.ContentInteractor.GetLatestList()
	if err != nil {
		return nil, err
	}

	return &pb.ContentList{Content: res}, nil
}

// get trend list by user id
func (s *ContentServiceServer) GetTrendList(ctx context.Context, req *emptypb.Empty) (*pb.ContentList, error) {

	res, err := s.ContentInteractor.GetTrendList()
	if err != nil {
		return nil, err
	}

	return &pb.ContentList{Content: res}, nil
}

// get recommended list by user id
func (s *ContentServiceServer) GetRecommendedListByUser(ctx context.Context, req *pb.ContentUserIdRequest) (*pb.ContentList, error) {

	res, err := s.ContentInteractor.GetRecommendedListByUser(req)
	if err != nil {
		return nil, err
	}

	return &pb.ContentList{Content: res}, nil
}

// get sold list by user id
func (s *ContentServiceServer) GetSoldListByUser(ctx context.Context, req *pb.ContentUserIdRequest) (*pb.ContentList, error) {

	res, err := s.ContentInteractor.GetSoldListByUser(req)
	if err != nil {
		return nil, err
	}

	return &pb.ContentList{Content: res}, nil
}

// get purchased list by user id
func (s *ContentServiceServer) GetPurchasedListByUser(ctx context.Context, req *pb.ContentUserIdRequest) (*pb.ContentList, error) {

	res, err := s.ContentInteractor.GetPurchasedListByUser(req)
	if err != nil {
		return nil, err
	}

	return &pb.ContentList{Content: res}, nil
}

// get liked list by user id
func (s *ContentServiceServer) GetLikedListByUser(ctx context.Context, req *pb.ContentUserIdRequest) (*pb.ContentList, error) {

	res, err := s.ContentInteractor.GetLikedListByUser(req)
	if err != nil {
		return nil, err
	}

	return &pb.ContentList{Content: res}, nil
}

// get list by id list
func (s *ContentServiceServer) GetListByIdList(ctx context.Context, req *pb.ContentIdListRequest) (*pb.ContentList, error) {

	res, err := s.ContentInteractor.GetListByIdList(req)
	if err != nil {
		return nil, err
	}

	return &pb.ContentList{Content: res}, nil
}
