package interactor

import (
	"log"

	"github.com/hidenari-yuda/ai-market-go/pb"
	"github.com/hidenari-yuda/ai-market-go/usecase"
)

type ItemInteractor interface {
	// Gest API
	// Create
	Create(Item *pb.Item) (*pb.Item, error)

	// Update
	Update(Item *pb.Item) (bool, error)

	// Delete
	Delete(param *pb.ItemIdRequest) (bool, error)

	// Get
	GetById(param *pb.ItemIdRequest) (*pb.Item, error)

	// get list by user id
	GetListByUser(param *pb.ItemUserIdRequest) ([]*pb.Item, error)

	// get by type
	GetListBySearch(param *pb.ItemSearchRequest) ([]*pb.Item, error)

	// get latest id=user_id
	GetLatestList(param *pb.ItemUserIdRequest) ([]*pb.Item, error)

	// get trend list by user id
	GetTrendList(param *pb.ItemUserIdRequest) ([]*pb.Item, error)

	// get recommended list by user id
	GetRecommendedListByUser(param *pb.ItemUserIdRequest) ([]*pb.Item, error)

	// get sold list by user id
	GetSoldListByUser(param *pb.ItemUserIdRequest) ([]*pb.Item, error)

	// get purchased list by user id
	GetPurchasedListByUser(param *pb.ItemUserIdRequest) ([]*pb.Item, error)

	// get liked list by user id
	GetLikedListByUser(param *pb.ItemUserIdRequest) ([]*pb.Item, error)

	// get list by id list
	GetListByIdList(param *pb.ItemIdListRequest) ([]*pb.Item, error)
}

type ItemInteractorImpl struct {
	firebase       usecase.Firebase
	itemRepository usecase.ItemRepository
}

func NewItemInteractorImpl(
	fb usecase.Firebase,
	iR usecase.ItemRepository,
) ItemInteractor {
	return &ItemInteractorImpl{
		firebase:       fb,
		itemRepository: iR,
	}
}

func (i *ItemInteractorImpl) Create(item *pb.Item) (*pb.Item, error) {

	err := i.itemRepository.Create(item)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (i *ItemInteractorImpl) Update(item *pb.Item) (bool, error) {

	err := i.itemRepository.Update(item)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (i *ItemInteractorImpl) Delete(param *pb.ItemIdRequest) (bool, error) {

	err := i.itemRepository.Delete(param.Id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (i *ItemInteractorImpl) GetById(param *pb.ItemIdRequest) (*pb.Item, error) {
	var (
		item *pb.Item
		err  error
	)

	item, err = i.itemRepository.GetById(param.Id)
	if err != nil {
		log.Println("error is:", err)
		return item, err
	}

	return item, nil
}

func (i *ItemInteractorImpl) GetListByUser(param *pb.ItemUserIdRequest) ([]*pb.Item, error) {
	var (
		items []*pb.Item
		err   error
	)

	items, err = i.itemRepository.GetListByUser(param.UserId)
	if err != nil {
		log.Println("error is:", err)
		return items, err
	}

	return items, nil
}

// get by type
// rpc GetListByCategory (ItemCategoryRequest) returns (ItemList) {}
func (i *ItemInteractorImpl) GetListBySearch(param *pb.ItemSearchRequest) ([]*pb.Item, error) {
	var (
		items []*pb.Item
		err   error
	)

	// categoryStrList := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(input.IdList)), ", "), "[]")

	items, err = i.itemRepository.GetListByFreeWord(param.FreeWord)
	if err != nil {
		log.Println("error is:", err)
		return items, err
	}

	return items, nil
}

// // get by latest id=user_id
// rpc GetLatestList(ItemIdRequest) returns (ItemList) {}
func (i *ItemInteractorImpl) GetLatestList(param *pb.ItemUserIdRequest) ([]*pb.Item, error) {
	var (
		items []*pb.Item
		err   error
	)

	items, err = i.itemRepository.GetLatestList(param.UserId)
	if err != nil {
		log.Println("error is:", err)
		return items, err
	}

	return items, nil
}

// // get by trend id=user_id
// rpc GetTrendList(ItemIdRequest) returns (ItemList) {}
func (i *ItemInteractorImpl) GetTrendList(param *pb.ItemUserIdRequest) ([]*pb.Item, error) {
	var (
		items []*pb.Item
		err   error
	)

	items, err = i.itemRepository.GetTrendList(param.UserId)
	if err != nil {
		log.Println("error is:", err)
		return items, err
	}

	return items, nil
}

// // get by recommend id=user_id
// rpc GetRecommendedListByUser(ItemIdRequest) returns (ItemList) {}
func (i *ItemInteractorImpl) GetRecommendedListByUser(param *pb.ItemUserIdRequest) ([]*pb.Item, error) {
	var (
		items []*pb.Item
		err   error
	)

	items, err = i.itemRepository.GetRecommendedListByUser(param.UserId)
	if err != nil {
		log.Println("error is:", err)
		return items, err
	}

	return items, nil
}

// // get by sold id=user_id
// rpc GetSoldListByUser(ItemIdRequest) returns (ItemList) {}
func (i *ItemInteractorImpl) GetSoldListByUser(param *pb.ItemUserIdRequest) ([]*pb.Item, error) {
	var (
		items []*pb.Item
		err   error
	)

	items, err = i.itemRepository.GetSoldListByUser(param.UserId)
	if err != nil {
		log.Println("error is:", err)
		return items, err
	}

	return items, nil
}

// rpc GetPurchasedListByUser(ItemIdRequest) returns (ItemList) {}
func (i *ItemInteractorImpl) GetPurchasedListByUser(param *pb.ItemUserIdRequest) ([]*pb.Item, error) {
	var (
		items []*pb.Item
		err   error
	)

	items, err = i.itemRepository.GetPurchasedListByUser(param.UserId)
	if err != nil {
		log.Println("error is:", err)
		return items, err
	}

	return items, nil
}

// // get by liked id=user_id
// rpc GetLikedListByUser(ItemIdRequest) returns (ItemList) {}
func (i *ItemInteractorImpl) GetLikedListByUser(param *pb.ItemUserIdRequest) ([]*pb.Item, error) {
	var (
		items []*pb.Item
		err   error
	)

	items, err = i.itemRepository.GetLikedListByUser(param.UserId)
	if err != nil {
		log.Println("error is:", err)
		return items, err
	}

	return items, nil
}

// // get list by id list
// rpc GetListByIdList (ItemIdListRequest) returns (ItemList) {}
func (i *ItemInteractorImpl) GetListByIdList(param *pb.ItemIdListRequest) ([]*pb.Item, error) {
	var (
		items []*pb.Item
		err   error
	)

	items, err = i.itemRepository.GetListByIdList(param.Id)
	if err != nil {
		log.Println("error is:", err)
		return items, err
	}

	return items, nil
}

// admin
// get all
func (i *ItemInteractorImpl) GetAll() ([]*pb.Item, error) {
	var (
		items []*pb.Item
		err   error
	)

	items, err = i.itemRepository.GetAll()
	if err != nil {
		log.Println("error is:", err)
		return items, err
	}

	return items, nil
}
