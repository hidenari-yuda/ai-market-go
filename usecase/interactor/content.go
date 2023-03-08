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

	// get by uuid
	GetByUuid(param *pb.ContentUuidRequest) (*pb.Content, error)

	// get list by user id
	GetListByUser(param *pb.ContentUserIdRequest) ([]*pb.Content, error)

	// get by type
	GetListBySearch(param *pb.ContentSearchRequest) ([]*pb.Content, error)

	// get latest id=user_id
	GetLatestList() ([]*pb.Content, error)

	// get trend list by user id
	GetTrendList() ([]*pb.Content, error)

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
	firebase                     usecase.Firebase
	contentRepository            usecase.ContentRepository
	contentContentRepository     usecase.ContentDetailRepository
	contentToolRepository        usecase.ContentToolRepository
	contentCategoryRepository    usecase.ContentCategoryRepository
	contentSubCategoryRepository usecase.ContentSubCategoryRepository
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
		firebase:                     fb,
		contentRepository:            cR,
		contentContentRepository:     cdR,
		contentToolRepository:        ctR,
		contentCategoryRepository:    ccR,
		contentSubCategoryRepository: cscR,
	}
}

func (i *ContentInteractorImpl) Create(content *pb.Content) (*pb.Content, error) {

	if err := i.contentRepository.Create(content); err != nil {
		return content, err
	}

	for _, detail := range content.Details {
		detail.ContentId = content.Id
		if err := i.contentContentRepository.Create(detail);err != nil {
			return content, err
		}
	}

	for _, tool := range content.Tools {
		tool.ContentId = content.Id
		if err := i.contentToolRepository.Create(tool);err != nil {
			return content, err
		}
	}

	for _, category := range content.Categories {
		category.ContentId = content.Id
		if err := i.contentCategoryRepository.Create(category);err != nil {
			return content, err
		}
	}

	for _, subCategory := range content.SubCategories {
		subCategory.ContentId = content.Id
		if err := i.contentSubCategoryRepository.Create(subCategory); err != nil {
			return content, err
		}
	}

	return content, nil
}

func (i *ContentInteractorImpl) Update(content *pb.Content) (bool, error) {

	if err := i.contentRepository.Update(content); err != nil {
		return false, err
	}

	if err := i.contentContentRepository.DeleteByContent(content.Id); err != nil {
		return false, err
	}

	for _, detail := range content.Details {
		detail.ContentId = content.Id
		if err := i.contentContentRepository.Create(detail);err != nil {
			return false, err
		}
	}

	if err := i.contentToolRepository.DeleteByContent(content.Id); err != nil {
		return false, err
	}

	for _, tool := range content.Tools {
		tool.ContentId = content.Id
		if err := i.contentToolRepository.Create(tool);err != nil {
			return false, err
		}
	}

	if err := i.contentCategoryRepository.DeleteByContent(content.Id); err != nil {
		return false, err
	}

	for _, category := range content.Categories {
		category.ContentId = content.Id
		if err := i.contentCategoryRepository.Create(category);err != nil {
			return false, err
		}
	}

	if err := i.contentSubCategoryRepository.DeleteByContent(content.Id); err != nil {
		return false, err
	}

	for _, subCategory := range content.SubCategories {
		subCategory.ContentId = content.Id
		if err := i.contentSubCategoryRepository.Create(subCategory); err != nil {
			return false, err
		}
	}

	return true, nil
}

func (i *ContentInteractorImpl) Delete(param *pb.ContentIdRequest) (bool, error) {

	err := i.contentRepository.Delete(param.Id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (i *ContentInteractorImpl) GetById(param *pb.ContentIdRequest) (*pb.Content, error) {
	var (
		content *pb.Content
		err     error
	)

	content, err = i.contentRepository.GetById(param.Id)
	if err != nil {
		log.Println("error is:", err)
		return content, err
	}

	details, err := i.contentContentRepository.GetListByContent(content.Id)
	if err != nil {
		log.Println("error is:", err)
		return content, err
	}

	tools, err := i.contentToolRepository.GetListByContent(content.Id)
	if err != nil {
		log.Println("error is:", err)
		return content, err
	}

	categories, err := i.contentCategoryRepository.GetListByContent(content.Id)
	if err != nil {
		log.Println("error is:", err)
		return content, err
	}

	subCategories, err := i.contentSubCategoryRepository.GetListByContent(content.Id)
	if err != nil {
		log.Println("error is:", err)
		return content, err
	}

	content.Details = details
	content.Tools = tools
	content.Categories = categories
	content.SubCategories = subCategories

	return content, nil
}

func (i *ContentInteractorImpl) GetByUuid(param *pb.ContentUuidRequest) (*pb.Content, error) {
	var (
		content *pb.Content
		err     error
	)

	content, err = i.contentRepository.GetByUuid(param.Uuid)
	if err != nil {
		log.Println("error is:", err)
		return content, err
	}

	details, err := i.contentContentRepository.GetListByContent(content.Id)
	if err != nil {
		log.Println("error is:", err)
		return content, err
	}

	tools, err := i.contentToolRepository.GetListByContent(content.Id)
	if err != nil {
		log.Println("error is:", err)
		return content, err
	}

	categories, err := i.contentCategoryRepository.GetListByContent(content.Id)
	if err != nil {
		log.Println("error is:", err)
		return content, err
	}

	subCategories, err := i.contentSubCategoryRepository.GetListByContent(content.Id)
	if err != nil {
		log.Println("error is:", err)
		return content, err
	}

	content.Details = details
	content.Tools = tools
	content.Categories = categories
	content.SubCategories = subCategories

	return content, nil
}

func (i *ContentInteractorImpl) GetListByUser(param *pb.ContentUserIdRequest) ([]*pb.Content, error) {
	var (
		contents []*pb.Content
		err      error
	)

	contents, err = i.contentRepository.GetListByUser(param.UserId)
	if err != nil {
		log.Println("error is:", err)
		return contents, err
	}

		details, err := i.contentContentRepository.GetListByUser(param.UserId)
	if err != nil {
		log.Println("error is:", err)
		return contents, err
	}

	tools, err := i.contentToolRepository.GetListByUser(param.UserId)
	if err != nil {
		log.Println("error is:", err)
		return contents, err
	}

	categories, err := i.contentCategoryRepository.GetListByUser(param.UserId)
	if err != nil {
		log.Println("error is:", err)
		return contents, err
	}

	subCategories, err := i.contentSubCategoryRepository.GetListByUser(param.UserId)
	if err != nil {
		log.Println("error is:", err)
		return contents, err
	}

	for _, content := range contents {
		content.Details = details
		content.Tools = tools
		content.Categories = categories
		content.SubCategories = subCategories
	}

	return contents, nil
}

// get by type
// rpc GetListByCategory (ContentCategoryRequest) returns (ContentList) {}
func (i *ContentInteractorImpl) GetListBySearch(param *pb.ContentSearchRequest) ([]*pb.Content, error) {
	var (
		contents []*pb.Content
		err      error
		contentIdList  []int64
	)

	// categoryStrList := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(input.IdList)), ", "), "[]")

	contents, err = i.contentRepository.GetListByFreeWord(param.FreeWord)
	if err != nil {
		log.Println("error is:", err)
		return contents, err
	}

	for _, content := range contents {
		contentIdList = append(contentIdList, content.Id)
	}

	details, err := i.contentContentRepository.GetListByIdList(contentIdList)
	if err != nil {
		log.Println("error is:", err)
		return contents, err
	}

	tools, err := i.contentToolRepository.GetListByIdList(contentIdList)
	if err != nil {
		log.Println("error is:", err)
		return contents, err
	}

	categories, err := i.contentCategoryRepository.GetListByIdList(contentIdList)
	if err != nil {
		log.Println("error is:", err)
		return contents, err
	}

	subCategories, err := i.contentSubCategoryRepository.GetListByIdList(contentIdList)
	if err != nil {
		log.Println("error is:", err)
		return contents, err
	}

	for _, content := range contents {
		content.Details = details
		content.Tools = tools
		content.Categories = categories
		content.SubCategories = subCategories
	}

	return contents, nil
}

// // get by latest id=user_id
// rpc GetLatestList(ContentIdRequest) returns (ContentList) {}
func (i *ContentInteractorImpl) GetLatestList() ([]*pb.Content, error) {
	var (
		contents []*pb.Content
		err      error
	)

	contents, err = i.contentRepository.GetLatestList()
	if err != nil {
		log.Println("error is:", err)
		return contents, err
	}

	return contents, nil
}

// // get by trend id=user_id
// rpc GetTrendList(ContentIdRequest) returns (ContentList) {}
func (i *ContentInteractorImpl) GetTrendList() ([]*pb.Content, error) {
	var (
		contents []*pb.Content
		err      error
	)

	contents, err = i.contentRepository.GetTrendList()
	if err != nil {
		log.Println("error is:", err)
		return contents, err
	}

	return contents, nil
}

// // get by recommend id=user_id
// rpc GetRecommendedListByUser(ContentIdRequest) returns (ContentList) {}
func (i *ContentInteractorImpl) GetRecommendedListByUser(param *pb.ContentUserIdRequest) ([]*pb.Content, error) {
	var (
		contents []*pb.Content
		err      error
	)

	contents, err = i.contentRepository.GetRecommendedListByUser(param.UserId)
	if err != nil {
		log.Println("error is:", err)
		return contents, err
	}

	return contents, nil
}

// // get by sold id=user_id
// rpc GetSoldListByUser(ContentIdRequest) returns (ContentList) {}
func (i *ContentInteractorImpl) GetSoldListByUser(param *pb.ContentUserIdRequest) ([]*pb.Content, error) {
	var (
		contents []*pb.Content
		err      error
	)

	contents, err = i.contentRepository.GetSoldListByUser(param.UserId)
	if err != nil {
		log.Println("error is:", err)
		return contents, err
	}

	return contents, nil
}

// rpc GetPurchasedListByUser(ContentIdRequest) returns (ContentList) {}
func (i *ContentInteractorImpl) GetPurchasedListByUser(param *pb.ContentUserIdRequest) ([]*pb.Content, error) {
	var (
		contents []*pb.Content
		err      error
	)

	contents, err = i.contentRepository.GetPurchasedListByUser(param.UserId)
	if err != nil {
		log.Println("error is:", err)
		return contents, err
	}

	return contents, nil
}

// // get by liked id=user_id
// rpc GetLikedListByUser(ContentIdRequest) returns (ContentList) {}
func (i *ContentInteractorImpl) GetLikedListByUser(param *pb.ContentUserIdRequest) ([]*pb.Content, error) {
	var (
		contents []*pb.Content
		err      error
	)

	contents, err = i.contentRepository.GetLikedListByUser(param.UserId)
	if err != nil {
		log.Println("error is:", err)
		return contents, err
	}

	return contents, nil
}

// // get list by id list
// rpc GetListByIdList (IdListRequest) returns (ContentList) {}
func (i *ContentInteractorImpl) GetListByIdList(param *pb.ContentIdListRequest) ([]*pb.Content, error) {
	var (
		contents []*pb.Content
		err      error
	)

	contents, err = i.contentRepository.GetListByIdList(param.Id)
	if err != nil {
		log.Println("error is:", err)
		return contents, err
	}

	return contents, nil
}

// admin
// get all
func (i *ContentInteractorImpl) GetAll() ([]*pb.Content, error) {
	var (
		contents []*pb.Content
		err      error
	)

	contents, err = i.contentRepository.GetAll()
	if err != nil {
		log.Println("error is:", err)
		return contents, err
	}

	return contents, nil
}
