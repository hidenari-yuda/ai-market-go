package interactor

import (
	"log"

	"github.com/hidenari-yuda/ai-market-go/pb"
	"github.com/hidenari-yuda/ai-market-go/usecase"
)

type PointInteractor interface {
	// Gest API
	// Create
	Create(param *pb.Point) (*pb.Point, error)

	// Update
	Update(param *pb.Point) (bool, error)

	// Delete
	Delete(param *pb.PointIdRequest) (bool, error)

	// Get
	GetById(param *pb.PointIdRequest) (*pb.Point, error)

	// get by uuid
	GetByUuid(param *pb.PointUuidRequest) (*pb.Point, error)

	// get list by id list
	GetListByIdList(param *pb.PointIdListRequest) ([]*pb.Point, error)
}

type PointInteractorImpl struct {
	firebase                  usecase.Firebase
	orderRepository						usecase.PointRepository
}

func NewPointInteractorImpl(
	fb usecase.Firebase,
	oR usecase.PointRepository,
) PointInteractor {
	return &PointInteractorImpl{
		firebase:                  fb,
		orderRepository:						oR,
	}
}

func (i *PointInteractorImpl) Create(param *pb.Point) (*pb.Point, error) {

	err := i.orderRepository.Create(param)
	if err != nil {
		return param, err
	}

	return param, nil
}

func (i *PointInteractorImpl) Update(param *pb.Point) (bool, error) {

	err := i.orderRepository.Update(param)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (i *PointInteractorImpl) Delete(param *pb.PointIdRequest) (bool, error) {

	err := i.orderRepository.Delete(param.Id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (i *PointInteractorImpl) GetById(param *pb.PointIdRequest) (*pb.Point, error) {
	var (
		order *pb.Point
		err  error
	)

	order, err = i.orderRepository.GetById(param.Id)
	if err != nil {
		log.Println("error is:", err)
		return order, err
	}

	return order, nil
}

func (i *PointInteractorImpl) GetByUuid(param *pb.PointUuidRequest) (*pb.Point, error) {
	var (
		order *pb.Point
		err  error
	)

	order, err = i.orderRepository.GetByUuid(param.Uuid)
	if err != nil {
			log.Println("error is:", err)
			return order, err
	}

	return order, nil
}

// // get list by id list
// rpc GetListByIdList (paramIdListRequest) returns (paramList) {}
func (i *PointInteractorImpl) GetListByIdList(param *pb.PointIdListRequest) ([]*pb.Point, error) {
	var (
		orders []*pb.Point
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
func (i *PointInteractorImpl) GetAll() ([]*pb.Point, error) {
	var (
		orders []*pb.Point
		err   error
	)

	orders, err = i.orderRepository.GetAll()
	if err != nil {
		log.Println("error is:", err)
		return orders, err
	}

	return orders, nil
}
