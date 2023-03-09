package interactor

import (
	"log"

	"github.com/hidenari-yuda/ai-market-go/pb"
	"github.com/hidenari-yuda/ai-market-go/usecase"
)

type PointHistoryInteractor interface {
	// Gest API
	// Create
	Create(param *pb.PointHistory) (*pb.PointHistory, error)

	// Update
	Update(param *pb.PointHistory) (bool, error)

	// Delete
	Delete(param *pb.PointHistoryIdRequest) (bool, error)

	// Get
	GetById(param *pb.PointHistoryIdRequest) (*pb.PointHistory, error)

	// get by uuid
	GetByUuid(param *pb.PointHistoryUuidRequest) (*pb.PointHistory, error)

	// get purchased list by user id
	GetPurchasedListByUser(param *pb.PointHistoryUserIdRequest) ([]*pb.PointHistory, error)

	// get list by id list
	GetListByIdList(param *pb.PointHistoryIdListRequest) ([]*pb.PointHistory, error)
}

type PointHistoryInteractorImpl struct {
	firebase                  usecase.Firebase
	orderRepository						usecase.PointHistoryRepository
}

func NewPointHistoryInteractorImpl(
	fb usecase.Firebase,
	oR usecase.PointHistoryRepository,
) PointHistoryInteractor {
	return &PointHistoryInteractorImpl{
		firebase:                  fb,
		orderRepository:						oR,
	}
}

func (i *PointHistoryInteractorImpl) Create(param*pb.PointHistory) (*pb.PointHistory, error) {

	err := i.orderRepository.Create(param)
	if err != nil {
		return param, err
	}

	return param, nil
}

func (i *PointHistoryInteractorImpl) Update(param*pb.PointHistory) (bool, error) {

	err := i.orderRepository.Update(param)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (i *PointHistoryInteractorImpl) Delete(param *pb.PointHistoryIdRequest) (bool, error) {

	err := i.orderRepository.Delete(param.Id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (i *PointHistoryInteractorImpl) GetById(param *pb.PointHistoryIdRequest) (*pb.PointHistory, error) {
	var (
		order *pb.PointHistory
		err  error
	)

	order, err = i.orderRepository.GetById(param.Id)
	if err != nil {
		log.Println("error is:", err)
		return order, err
	}

	return order, nil
}

func (i *PointHistoryInteractorImpl) GetByUuid(param *pb.PointHistoryUuidRequest) (*pb.PointHistory, error) {
	var (
		order *pb.PointHistory
		err  error
	)

	order, err = i.orderRepository.GetByUuid(param.Uuid)
	if err != nil {
			log.Println("error is:", err)
			return order, err
	}

	return order, nil
}

// rpc GetPurchasedListByUser(paramIdRequest) returns (paramList) {}
func (i *PointHistoryInteractorImpl) GetPurchasedListByUser(param *pb.PointHistoryUserIdRequest) ([]*pb.PointHistory, error) {
	var (
		orders []*pb.PointHistory
		err   error
	)

	orders, err = i.orderRepository.GetPurchasedListByUser(param.UserId)
	if err != nil {
		log.Println("error is:", err)
		return orders, err
	}

	return orders, nil
}

// // get list by id list
// rpc GetListByIdList (paramIdListRequest) returns (paramList) {}
func (i *PointHistoryInteractorImpl) GetListByIdList(param *pb.PointHistoryIdListRequest) ([]*pb.PointHistory, error) {
	var (
		orders []*pb.PointHistory
		err   error
	)

	orders, err = i.orderRepository.GetListByIdList(param.Id)
	if err != nil {
		log.Println("error is:", err)
		return orders, err
	}

	return orders, nil
}

// admin
// get all
func (i *PointHistoryInteractorImpl) GetAll() ([]*pb.PointHistory, error) {
	var (
		orders []*pb.PointHistory
		err   error
	)

	orders, err = i.orderRepository.GetAll()
	if err != nil {
		log.Println("error is:", err)
		return orders, err
	}

	return orders, nil
}
