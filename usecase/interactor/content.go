package interactor

import (
	"log"

	"github.com/hidenari-yuda/ai-market-go/pb"
	"github.com/hidenari-yuda/ai-market-go/usecase"
)

type ContentInteractor interface {
	// Gest API
	// Create
	Create(Content *pb.Content) (*pb.Content, error)

	// Update
	Update(Content *pb.Content) (bool, error)

	// Delete
	Delete(param *pb.ContentIdRequest) (bool, error)

	// Get
	GetById(param *pb.ContentIdRequest) (*pb.Content, error)

	// get list by user id
	GetListByUser(param *pb.ContentUserIdRequest) ([]*pb.Content, error)

	// get by type
	GetListBySearch(param *pb.ContentSearchRequest) ([]*pb.Content, error)

	// get latest id=user_id
	GetLatestList(param *pb.ContentUserIdRequest) ([]*pb.Content, error)

	// get trend list by user id
	GetTrendList(param *pb.ContentUserIdRequest) ([]*pb.Content, error)

	// get recommended list by user id
	GetRecommendedListByUser(param *pb.ContentUserIdRequest) ([]*pb.Content, error)

	// get sold list by user id
	GetSoldListByUser(param *pb.ContentUserIdRequest) ([]*pb.Content, error)

	// get purchased list by user id
	GetPurchasedListByUser(param *pb.ContentUserIdRequest) ([]*pb.Content, error)

	// get liked list by user id
	GetLikedListByUser(param *pb.ContentUserIdRequest) ([]*pb.Content, error)

	// get list by id list
	GetListByIdList(param *pb.ContentIdListRequest) ([]*pb.Content, error)
}

type ContentInteractorImpl struct {
	firebase                  usecase.Firebase
	itemRepository            usecase.ContentRepository
	itemContentRepository     usecase.ContentDetailRepository
	itemToolRepository        usecase.ContentToolRepository
	itemCategoryRepository    usecase.ContentCategoryRepository
	itemSubCategoryRepository usecase.ContentSubCategoryRepository
}

func NewContentInteractorImpl(
	fb usecase.Firebase,
	cR usecase.ContentRepository,
	cdR usecase.ContentDetailRepository,
	ctR usecase.ContentToolRepository,
	ccR usecase.ContentCategoryRepository,
	cscR usecase.ContentSubCategoryRepository,
) ContentInteractor {
	return &ContentInteractorImpl{
		firebase:                  fb,
		itemRepository:            cR,
		itemContentRepository:     cdR,
		itemToolRepository:        ctR,
		itemCategoryRepository:    ccR,
		itemSubCategoryRepository: cscR,
	}
}

func (i *ContentInteractorImpl) Create(item *pb.Content) (*pb.Content, error) {

	err := i.itemRepository.Create(item)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (i *ContentInteractorImpl) Update(item *pb.Content) (bool, error) {

	err := i.itemRepository.Update(item)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (i *ContentInteractorImpl) Delete(param *pb.ContentIdRequest) (bool, error) {

	err := i.itemRepository.Delete(param.Id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (i *ContentInteractorImpl) GetById(param *pb.ContentIdRequest) (*pb.Content, error) {
	var (
		item *pb.Content
		err  error
	)

	item, err = i.itemRepository.GetById(param.Id)
	if err != nil {
		log.Println("error is:", err)
		return item, err
	}

	return item, nil
}

func (i *ContentInteractorImpl) GetListByUser(param *pb.ContentUserIdRequest) ([]*pb.Content, error) {
	var (
		items []*pb.Content
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
// rpc GetListByCategory (ContentCategoryRequest) returns (ContentList) {}
func (i *ContentInteractorImpl) GetListBySearch(param *pb.ContentSearchRequest) ([]*pb.Content, error) {
	var (
		items []*pb.Content
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
// rpc GetLatestList(ContentIdRequest) returns (ContentList) {}
func (i *ContentInteractorImpl) GetLatestList(param *pb.ContentUserIdRequest) ([]*pb.Content, error) {
	var (
		items []*pb.Content
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
// rpc GetTrendList(ContentIdRequest) returns (ContentList) {}
func (i *ContentInteractorImpl) GetTrendList(param *pb.ContentUserIdRequest) ([]*pb.Content, error) {
	var (
		items []*pb.Content
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
// rpc GetRecommendedListByUser(ContentIdRequest) returns (ContentList) {}
func (i *ContentInteractorImpl) GetRecommendedListByUser(param *pb.ContentUserIdRequest) ([]*pb.Content, error) {
	var (
		items []*pb.Content
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
// rpc GetSoldListByUser(ContentIdRequest) returns (ContentList) {}
func (i *ContentInteractorImpl) GetSoldListByUser(param *pb.ContentUserIdRequest) ([]*pb.Content, error) {
	var (
		items []*pb.Content
		err   error
	)

	items, err = i.itemRepository.GetSoldListByUser(param.UserId)
	if err != nil {
		log.Println("error is:", err)
		return items, err
	}

	return items, nil
}

// rpc GetPurchasedListByUser(ContentIdRequest) returns (ContentList) {}
func (i *ContentInteractorImpl) GetPurchasedListByUser(param *pb.ContentUserIdRequest) ([]*pb.Content, error) {
	var (
		items []*pb.Content
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
// rpc GetLikedListByUser(ContentIdRequest) returns (ContentList) {}
func (i *ContentInteractorImpl) GetLikedListByUser(param *pb.ContentUserIdRequest) ([]*pb.Content, error) {
	var (
		items []*pb.Content
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
// rpc GetListByIdList (ContentIdListRequest) returns (ContentList) {}
func (i *ContentInteractorImpl) GetListByIdList(param *pb.ContentIdListRequest) ([]*pb.Content, error) {
	var (
		items []*pb.Content
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
func (i *ContentInteractorImpl) GetAll() ([]*pb.Content, error) {
	var (
		items []*pb.Content
		err   error
	)

	items, err = i.itemRepository.GetAll()
	if err != nil {
		log.Println("error is:", err)
		return items, err
	}

	return items, nil
}
