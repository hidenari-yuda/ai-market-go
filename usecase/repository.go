package usecase

import "github.com/hidenari-yuda/ai-market-go/pb"

type UserRepository interface {
	// Gest API
	// Create
	Create(param *pb.User) error

	// Update
	Update(param *pb.User) error

	// Delete
	Delete(id int64) error

	// Get
	GetById(id int64) (*pb.User, error)

	// get by uuid
	GetByUuid(uuid string) (*pb.User, error)

	// get latest list
	GetLatestList() ([]*pb.User, error)

	// get trend list
	GetTrendList() ([]*pb.User, error)

	// get list by search
	GetListBySearch(freeWord string) ([]*pb.User, error)

	// admin
	GetAll() ([]*pb.User, error)

	// auth
	// signin
	GetByFirebaseId(firebaseId string) (*pb.User, error)
}

type UserMediaRepository interface {
	// Gest API
	// Create
	Create(param *pb.UserMedia) error

	// Update
	Update(param *pb.UserMedia) error

	// Delete
	Delete(id int64) error

	// delete by user id
	DeleteByUser(userId int64) error

	// Get
	GetById(id int64) (*pb.UserMedia, error)

	// get by uuid
	GetByUuid(uuid string) (*pb.UserMedia, error)

	// get list by user id
	GetListByUser(userId int64) ([]*pb.UserMedia, error)

	// get list by user id list
	GetListByUserList(idList string) ([]*pb.UserMedia, error)
}

/******************* billing *******************/
// billing
type BillingRepository interface {
	// Gest API
	// Create
	Create(param *pb.Billing) error

	// Update
	Update(param *pb.Billing) error

	// Delete
	Delete(id int64) error

	// Get
	GetById(id int64) (*pb.Billing, error)

	// get by uuid
	GetByUuid(uuid string) (*pb.Billing, error)

	// get list by user id
	GetListByUser(userId int64) ([]*pb.Billing, error)
}

/******************* content *******************/
// content
type ContentRepository interface {
	// Gest API
	// Create
	Create(param *pb.Content) error

	// Update
	Update(param *pb.Content) error

	// update impression
	UpdateImpressionByIdList(idList string) error

	// update view
	UpdateView(id int64) error

	// update click
	UpdateClick(id int64) error

	// update like
	UpdateLike(id int64) error

	// Delete
	Delete(id int64) error

	// Get
	GetById(id int64) (*pb.Content, error)

	// get by uuid
	GetByUuid(uuid string) (*pb.Content, error)

	// get list by type
	GetListByUser(userId int64) ([]*pb.Content, error)

	// get list by category
	GetListByFreeWord(freeWord string) ([]*pb.Content, error)

	// get by latest id=user_id
	GetLatestList() ([]*pb.Content, error)

	// get by trend id=user_id
	GetTrendList() ([]*pb.Content, error)

	// get recommended list
	GetRecommendedListByUser(userId int64) ([]*pb.Content, error)

	// get recommended list by content
	GetRecommendedListByContent() ([]*pb.Content, error)

	// get by sold id=user_id
	GetSoldListByUser(userId int64) ([]*pb.Content, error)

	// get by purchased id=user_id
	GetPurchasedListByUser(userId int64) ([]*pb.Content, error)

	// get by liked id=user_id
	GetLikedListByUser(userId int64) ([]*pb.Content, error)

	// get list by id list
	GetListByIdList(idList string) ([]*pb.Content, error)

	// admin
	GetAll() ([]*pb.Content, error)
}

// contentContent
type ContentDetailRepository interface {
	// Gest API
	// Create
	Create(param *pb.ContentDetail) error

	Update(param *pb.ContentDetail) error

	// Update
	Delete(id int64) error

	DeleteByContent(contentId int64) error

	// Get
	GetById(id int64) (*pb.ContentDetail, error)

	// get list by content id
	GetListByContent(contentId int64) ([]*pb.ContentDetail, error)

	// get list by user
	GetListByUser(userId int64) ([]*pb.ContentDetail, error)

	// get list by id list
	GetListByIdList(idList string) ([]*pb.ContentDetail, error)

	// get list by content id list
	GetListByContentIdList(contentIdList string) ([]*pb.ContentDetail, error)
}

// contentTool
type ContentToolRepository interface {
	// Gest API
	// Create
	Create(param *pb.ContentTool) error

	Update(param *pb.ContentTool) error

	// Update
	Delete(id int64) error

	DeleteByContent(contentId int64) error

	// Get
	GetById(id int64) (*pb.ContentTool, error)

	// get list by content id
	GetListByContent(contentId int64) ([]*pb.ContentTool, error)

	// get list by user
	GetListByUser(userId int64) ([]*pb.ContentTool, error)

	// get list by id list
	GetListByIdList(idList string) ([]*pb.ContentTool, error)

	// get list by content id list
	GetListByContentIdList(contentIdList string) ([]*pb.ContentTool, error)
}

// contentCategory
type ContentCategoryRepository interface {
	// Gest API
	// Create
	Create(param *pb.ContentCategory) error

	Update(param *pb.ContentCategory) error

	// Update
	Delete(id int64) error

	DeleteByContent(contentId int64) error

	// Get
	GetById(id int64) (*pb.ContentCategory, error)

	// get list by content id
	GetListByContent(contentId int64) ([]*pb.ContentCategory, error)

	// get list by user
	GetListByUser(userId int64) ([]*pb.ContentCategory, error)

	// get list by id list
	GetListByIdList(idList string) ([]*pb.ContentCategory, error)

	// get list by content id list
	GetListByContentIdList(contentIdList string) ([]*pb.ContentCategory, error)
}

// contentSubCategory
type ContentSubCategoryRepository interface {
	// Gest API
	// Create
	Create(param *pb.ContentSubCategory) error

	Update(param *pb.ContentSubCategory) error

	// Update
	Delete(id int64) error

	DeleteByContent(contentId int64) error

	// Get
	GetById(id int64) (*pb.ContentSubCategory, error)

	// get list by content id
	GetListByContent(contentId int64) ([]*pb.ContentSubCategory, error)

	// get list by user
	GetListByUser(userId int64) ([]*pb.ContentSubCategory, error)

	// get list by id list
	GetListByIdList(idList string) ([]*pb.ContentSubCategory, error)

	// get list by content id list
	GetListByContentIdList(contentIdList string) ([]*pb.ContentSubCategory, error)
}

/******************* like *******************/
// like
type LikeRepository interface {
	// Gest API
	// Create
	Create(param *pb.Like) error

	// Update
	Update(param *pb.Like) error

	// Delete
	Delete(id int64) error

	// Get
	GetById(id int64) (*pb.Like, error)

	// get list by user id
	GetListByUser(userId int64) ([]*pb.Like, error)

	// get list by content id
	GetListByContent(contentId int64) ([]*pb.Like, error)

	// get list by content id
	GetBoolByUserAndContent(userId, contentId int64) (bool, error)

	// get list by id list
	GetListByIdList(idList string) ([]*pb.Like, error)

	// admin
	GetAll() ([]*pb.Like, error)
}

/******************* view *******************/
// view
type ViewRepository interface {
	// Gest API
	// Create
	Create(param *pb.View) error

	// Update
	Update(param *pb.View) error

	// Delete
	Delete(id int64) error

	// Get
	GetById(id int64) (*pb.View, error)

	// get list by user id
	GetListByUser(userId int64) ([]*pb.View, error)

	// get list by content id
	GetListByContent(contentId int64) ([]*pb.View, error)

	// get list by id list
	GetListByIdList(idList string) ([]*pb.View, error)

	// admin
	GetAll() ([]*pb.View, error)
}

/******************* click *******************/
// click
type ClickRepository interface {
	// Gest API
	// Create
	Create(param *pb.Click) error

	// Update
	Update(param *pb.Click) error

	// Delete
	Delete(id int64) error

	// Get
	GetById(id int64) (*pb.Click, error)

	// get list by user id
	GetListByUser(userId int64) ([]*pb.Click, error)

	// get list by content id
	GetListByContent(contentId int64) ([]*pb.Click, error)

	// get list by id list
	GetListByIdList(idList string) ([]*pb.Click, error)

	// admin
	GetAll() ([]*pb.Click, error)
}

/******************* review *******************/
// review
type ReviewRepository interface {
	// Gest API
	// Create
	Create(param *pb.Review) error

	// Update
	Update(param *pb.Review) error

	// Delete
	Delete(id int64) error

	// Get
	GetById(id int64) (*pb.Review, error)

	// get by uuid
	GetByUuid(uuid string) (*pb.Review, error)

	// get list by user id
	GetListByUser(userId int64) ([]*pb.Review, error)

	// get list by content id
	GetListByContent(contentId int64) ([]*pb.Review, error)

	// get list by id list
	GetListByIdList(idList string) ([]*pb.Review, error)

	// admin
	GetAll() ([]*pb.Review, error)
}

/******************* asp *******************/
// asp
type AspRepository interface {
	// Gest API
	// Create
	Create(param *pb.Asp) error

	// Update
	Update(param *pb.Asp) error

	// Delete
	Delete(id int64) error

	// Get
	GetById(id int64) (*pb.Asp, error)

	// get by uuid
	GetByUuid(uuid string) (*pb.Asp, error)

	// get list by user id
	GetListByUser(userId int64) ([]*pb.Asp, error)

	// get list by content id
	GetListByContent(contentId int64) ([]*pb.Asp, error)

	// get list by service
	GetListByService(service int64) ([]*pb.Asp, error)

	// get list by id list
	GetListByIdList(idList string) ([]*pb.Asp, error)

	// admin
	GetAll() ([]*pb.Asp, error)
}

/******************* following *******************/
// following
type FollowingRepository interface {
	// Gest API
	// Create
	Create(param *pb.Following) error

	// Update
	Update(param *pb.Following) error

	// Delete
	Delete(id int64) error

	// Get
	GetById(id int64) (*pb.Following, error)

	// get by uuid
	GetByUuid(uuid string) (*pb.Following, error)

	// get list by user id
	GetFollowingListByUser(userId int64) ([]*pb.Following, error)

	// get list by user id
	GetFollowedListByUser(userId int64) ([]*pb.Following, error)

	// get list by id list
	GetListByIdList(idList string) ([]*pb.Following, error)

	// admin
	GetAll() ([]*pb.Following, error)
}

/******************* order *******************/
// order
type OrderRepository interface {
	// Gest API
	// Create
	Create(param *pb.Order) error

	// Update
	Update(param *pb.Order) error

	// Delete
	Delete(id int64) error

	// Get
	GetById(id int64) (*pb.Order, error)

	// get by uuid
	GetByUuid(uuid string) (*pb.Order, error)

	// get sold list by user id
	GetSoldListByUser(userId int64) ([]*pb.Order, error)

	// get purchased list by user id
	GetPurchasedListByUser(userId int64) ([]*pb.Order, error)

	// get list by id list
	GetListByIdList(idList string) ([]*pb.Order, error)

	// admin
	GetAll() ([]*pb.Order, error)
}

/******************* point *******************/
// point
type PointRepository interface {
	// Gest API
	// Create
	Create(param *pb.Point) error

	// Update
	Update(param *pb.Point) error

	// Delete
	Delete(id int64) error

	// Get
	GetById(id int64) (*pb.Point, error)

	// get by uuid
	GetByUuid(uuid string) (*pb.Point, error)

	// get sold list by user id
	GetListByUser(userId int64) ([]*pb.Point, error)

	// get list by id list
	GetListByIdList(idList string) ([]*pb.Point, error)

	// admin
	GetAll() ([]*pb.Point, error)
}

/******************* pointHistory *******************/
// pointHistory
type PointHistoryRepository interface {
	// Gest API
	// Create
	Create(param *pb.PointHistory) error

	// Update
	Update(param *pb.PointHistory) error

	// Delete
	Delete(id int64) error

	// Get
	GetById(id int64) (*pb.PointHistory, error)

	// get by uuid
	GetByUuid(uuid string) (*pb.PointHistory, error)

	// get sold list by user id
	GetListByUser(userId int64) ([]*pb.PointHistory, error)

	// get purchased list by user id
	GetPurchasedListByUser(userId int64) ([]*pb.PointHistory, error)

	// get list by id list
	GetListByIdList(idList string) ([]*pb.PointHistory, error)

	// admin
	GetAll() ([]*pb.PointHistory, error)
}

type ChatGroupRepository interface {
	// Gest API
	// Create
	Create(param *pb.ChatGroup) error

	// Update
	Update(param *pb.ChatGroup) error

	// Delete
	Delete(id int64) error

	// Get
	GetById(id int64) (*pb.ChatGroup, error)

	// get by uuid
	GetByUuid(uuid string) (*pb.ChatGroup, error)

	// get list by user id
	GetListByUser(userId int64) ([]*pb.ChatGroup, error)

	// admin
	GetAll() ([]*pb.ChatGroup, error)
}

type ChatRepository interface {
	// Gest API
	// Create
	Create(param *pb.Chat) error

	// Update
	Update(param *pb.Chat) error

	// Delete
	Delete(id int64) error

	// Get
	GetById(id int64) (*pb.Chat, error)

	// get list by user id
	GetListByGroup(groupId int64) ([]*pb.Chat, error)

	// get stream by group
	GetStreamByGroup(groupId int64) ([]*pb.Chat, error)

	// get latest chat for notification
	GetLatestChatByUser(userId int64) (*pb.Chat, error)
}

type ChatUserRepository interface {
	// Gest API
	// Create
	Create(param *pb.ChatUser) error

	Update(param *pb.ChatUser) error

	// Delete
	Delete(id int64) error

	// Get
	GetById(id int64) (*pb.ChatUser, error)

	// get list by group id
	GetListByGroup(groupId int64) ([]*pb.ChatUser, error)
}
