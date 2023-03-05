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

	// get list by type
	GetListByType(userType *pb.User_Type) ([]*pb.User, error)

	// get top list
	GetTopList(param *pb.UserTopRequest) ([]*pb.User, error)

	// get list by search
	GetListBySearch(param *pb.UserSearchRequest) ([]*pb.User, error)

	// admin
	GetAll() ([]*pb.User, error)

	// auth
	// signin
	GetByFirebaseId(firebaseId string) (*pb.User, error)
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

	// get list by user id
	GetListByUser(userId int64) ([]*pb.Billing, error)
}

/******************* item *******************/
// item
type ItemRepository interface {
	// Gest API
	// Create
	Create(param *pb.Item) error

	// Update
	Update(param *pb.Item) error

	// Delete
	Delete(id int64) error

	// Get
	GetById(id int64) (*pb.Item, error)

	// get list by type
	GetListByUser(userId int64) ([]*pb.Item, error)

	// get list by category
	GetListByFreeWord(freeWord string) ([]*pb.Item, error)

	// get by latest id=user_id
	GetLatestList(userId int64) ([]*pb.Item, error)

	// get by trend id=user_id
	GetTrendList(userId int64) ([]*pb.Item, error)

	// get recommended list
	GetRecommendedListByUser(userId int64) ([]*pb.Item, error)

	// get by sold id=user_id
	GetSoldListByUser(userId int64) ([]*pb.Item, error)

	// get by purchased id=user_id
	GetPurchasedListByUser(userId int64) ([]*pb.Item, error)

	// get by liked id=user_id
	GetLikedListByUser(userId int64) ([]*pb.Item, error)

	// get list by id list
  GetListByIdList(idList []int64) ([]*pb.Item, error)

	// admin
	GetAll() ([]*pb.Item, error)
}

// itemContent
type ItemContentRepository interface {
	// Gest API
	// Create
	Create(param *pb.ItemContent) error

	// Update
	Delete(id int64) error

	// Get
	GetById(id int64) (*pb.ItemContent, error)

	// get list by item id
	GetListByItem(itemId int64) ([]*pb.ItemContent, error)

	// get list by id list
	GetListByIdList(idList []int64) ([]*pb.ItemContent, error)
}

// itemTool
type ItemToolRepository interface {
	// Gest API
	// Create
	Create(param *pb.ItemTool) error

	// Update
	Delete(id int64) error

	// Get
	GetById(id int64) (*pb.ItemTool, error)

	// get list by item id
	GetListByItem(itemId int64) ([]*pb.ItemTool, error)

	// get list by id list
	GetListByIdList(idList []int64) ([]*pb.ItemTool, error)
}

// itemCategory
type ItemCategoryRepository interface {
	// Gest API
	// Create
	Create(param *pb.ItemCategory) error

	// Update
	Delete(id int64) error

	// Get
	GetById(id int64) (*pb.ItemCategory, error)

	// get list by item id
	GetListByItem(itemId int64) ([]*pb.ItemCategory, error)

	// get list by id list
	GetListByIdList(idList []int64) ([]*pb.ItemCategory, error)
}

// itemSubCategory
type ItemSubCategoryRepository interface {
	// Gest API
	// Create
	Create(param *pb.ItemSubCategory) error

	// Update
	Delete(id int64) error

	// Get
	GetById(id int64) (*pb.ItemSubCategory, error)

	// get list by item id
	GetListByItem(itemId int64) ([]*pb.ItemSubCategory, error)

	// get list by id list
	GetListByIdList(idList []int64) ([]*pb.ItemSubCategory, error)
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

	// get list by item id
	GetListByItem(itemId int64) ([]*pb.Like, error)

	// get list by id list
	GetListByIdList(idList []int64) ([]*pb.Like, error)

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

	// get list by item id
	GetListByItem(itemId int64) ([]*pb.View, error)

	// get list by id list
	GetListByIdList(idList []int64) ([]*pb.View, error)

	// admin
	GetAll() ([]*pb.View, error)
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

	// get list by user id
	GetListByUser(userId int64) ([]*pb.Review, error)

	// get list by item id
	GetListByItem(itemId int64) ([]*pb.Review, error)

	// get list by id list
	GetListByIdList(idList []int64) ([]*pb.Review, error)

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

	// get list by user id
	GetListByUser(userId int64) ([]*pb.Asp, error)

	// get list by item id
	GetListByItem(itemId int64) ([]*pb.Asp, error)

	// get list by service
	GetListByService(service int64) ([]*pb.Asp, error)

	// get list by id list
	GetListByIdList(idList []int64) ([]*pb.Asp, error)

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

	// get list by user id
	GetFollowingListByUser(userId int64) ([]*pb.Following, error)

	// get list by user id
	GetFollowedListByUser(userId int64) ([]*pb.Following, error)

	// get list by id list
	GetListByIdList(idList []int64) ([]*pb.Following, error)

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

	// get sold list by user id
	GetSoldListByUser(userId int64) ([]*pb.Order, error)

	// get purchased list by user id
	GetPurchasedListByUser(userId int64) ([]*pb.Order, error)

	// get list by id list
	GetListByIdList(idList []int64) ([]*pb.Order, error)

	// admin
	GetAll() ([]*pb.Order, error)
}

/******************* cashback *******************/
// cashback
type CashbackRepository interface {
	// Gest API
	// Create
	Create(param *pb.Cashback) error

	// Update
	Update(param *pb.Cashback) error

	// Delete
	Delete(id int64) error

	// Get
	GetById(id int64) (*pb.Cashback, error)

	// get sold list by user id
	GetListByUser(userId int64) ([]*pb.Cashback, error)

	// get purchased list by user id
	GetPurchasedListByUser(userId int64) ([]*pb.Cashback, error)

	// get list by id list
	GetListByIdList(idList []int64) ([]*pb.Cashback, error)

	// admin
	GetAll() ([]*pb.Cashback, error)
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

	// Delete
	Delete(id int64) error

	// Get
	GetById(id int64) (*pb.ChatUser, error)

	// get list by group id
	GetListByGroup(groupId int64) ([]*pb.ChatUser, error)
}