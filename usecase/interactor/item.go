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
	Delete(id uint) (bool, error)

	// Get
	GetById(id uint) (*pb.Item, error)

	// get list by user id
	GetListByUser(userId uint) ([]*pb.Item, error)
}

type ItemInteractorImpl struct {
	firebase            usecase.Firebase
	chatGroupRepository usecase.ItemRepository
	chatUserRepository  usecase.ChatUserRepository
}

func NewItemInteractorImpl(
	fb usecase.Firebase,
	cgR usecase.ItemRepository,
	cuR usecase.ChatUserRepository,
) ItemInteractor {
	return &ItemInteractorImpl{
		firebase:            fb,
		chatGroupRepository: cgR,
		chatUserRepository:  cuR,
	}
}

func (i *ItemInteractorImpl) Create(chatGroup *pb.Item) (*pb.Item, error) {

	// ユーザー登録
	err := i.chatGroupRepository.Create(chatGroup)
	if err != nil {
		return chatGroup, err
	}

	return chatGroup, nil
}

func (i *ItemInteractorImpl) Update(chatGroup *pb.Item) (bool, error) {

	// ユーザー登録
	err := i.chatGroupRepository.Update(chatGroup)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (i *ItemInteractorImpl) Delete(id uint) (bool, error) {

	// ユーザー登録
	err := i.chatGroupRepository.Delete(id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (i *ItemInteractorImpl) GetById(id uint) (*pb.Item, error) {
	var (
		chatGroup *pb.Item
		err       error
	)

	// ユーザー登録
	chatGroup, err = i.chatGroupRepository.GetById(id)
	if err != nil {
		log.Println("error is:", err)
		return chatGroup, err
	}

	return chatGroup, nil
}

func (i *ItemInteractorImpl) GetListByUser(userId uint) ([]*pb.Item, error) {
	var (
		chatGroups []*pb.Item
		err        error
	)

	// ユーザー登録
	chatGroups, err = i.chatGroupRepository.GetListByUser(userId)
	if err != nil {
		log.Println("error is:", err)
		return chatGroups, err
	}

	return chatGroups, nil
}
