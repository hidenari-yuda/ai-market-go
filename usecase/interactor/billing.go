package interactor

import (
	"log"

	"github.com/hidenari-yuda/ai-market-go/pb"
	"github.com/hidenari-yuda/ai-market-go/usecase"
)

type BillingInteractor interface {
	// Gest API
	// Create
	Create(Billing *pb.Billing) (*pb.Billing, error)

	// Update
	Update(Billing *pb.Billing) (bool, error)

	// Delete
	Delete(param *pb.BillingIdRequest) (bool, error)

	// Get
	GetById(param *pb.BillingIdRequest) (*pb.Billing, error)

	// get by uuid
	GetByUuid(param *pb.BillingUuidRequest) (*pb.Billing, error)

	// get list by user

	// get list by id list
	// GetListByIdList(param *pb.BillingIdListRequest) ([]*pb.Billing, error)
}

type BillingInteractorImpl struct {
	firebase                  usecase.Firebase
	orderRepository						usecase.BillingRepository
}

func NewBillingInteractorImpl(
	fb usecase.Firebase,
	oR usecase.BillingRepository,
) BillingInteractor {
	return &BillingInteractorImpl{
		firebase:                  fb,
		orderRepository:						oR,
	}
}

func (i *BillingInteractorImpl) Create(order *pb.Billing) (*pb.Billing, error) {

	err := i.orderRepository.Create(order)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (i *BillingInteractorImpl) Update(order *pb.Billing) (bool, error) {

	err := i.orderRepository.Update(order)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (i *BillingInteractorImpl) Delete(param *pb.BillingIdRequest) (bool, error) {

	err := i.orderRepository.Delete(param.Id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (i *BillingInteractorImpl) GetById(param *pb.BillingIdRequest) (*pb.Billing, error) {
	var (
		order *pb.Billing
		err  error
	)

	order, err = i.orderRepository.GetById(param.Id)
	if err != nil {
		log.Println("error is:", err)
		return order, err
	}

	return order, nil
}

func (i *BillingInteractorImpl) GetByUuid(param *pb.BillingUuidRequest) (*pb.Billing, error) {
	var (
		order *pb.Billing
		err  error
	)

	order, err = i.orderRepository.GetByUuid(param.Uuid)
	if err != nil {
			log.Println("error is:", err)
			return order, err
	}

	return order, nil
}

// get list by user
func (i *BillingInteractorImpl) GetListByUser(param *pb.BillingUserIdRequest) ([]*pb.Billing, error) {
	var (
		orders []*pb.Billing
		err   error
	)

	orders, err = i.orderRepository.GetListByUser(param.UserId)
	if err != nil {
		log.Println("error is:", err)
		return orders, err
	}

	return orders, nil
}

// // get list by id list
// rpc GetListByIdList (BillingIdListRequest) returns (BillingList) {}
// func (i *BillingInteractorImpl) GetListByIdList(param *pb.BillingIdListRequest) ([]*pb.Billing, error) {
// 	var (
// 		orders []*pb.Billing
// 		err   error
// 	)

// 	orders, err = i.orderRepository.GetListByIdList(param.Id)
// 	if err != nil {
// 		log.Println("error is:", err)
// 		return orders, err
// 	}

// 	return orders, nil
// }

// // admin
// // get all
// func (i *BillingInteractorImpl) GetAll() ([]*pb.Billing, error) {
// 	var (
// 		orders []*pb.Billing
// 		err   error
// 	)

// 	orders, err = i.orderRepository.GetAll()
// 	if err != nil {
// 		log.Println("error is:", err)
// 		return orders, err
// 	}

// 	return orders, nil
// }
