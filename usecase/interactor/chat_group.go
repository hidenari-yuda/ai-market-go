package interactor

import (
	"github.com/hidenari-yuda/ai-market-go/pb"
	"github.com/hidenari-yuda/ai-market-go/usecase"
)

type ChatGroupInteractor interface {
	// Gest API
	// Create
	Create(ChatGroup *pb.ChatGroup) (*pb.ChatGroup, error)

	// Update
	Update(ChatGroup *pb.ChatGroup) (bool, error)

	// Delete
	Delete(param *pb.ChatIdRequest) (bool, error)

	// Get
	GetById(param *pb.ChatIdRequest) (*pb.ChatGroup, error)

	// get by uuid
	GetByUuid(param *pb.ChatUuidRequest) (*pb.ChatGroup, error)

	// get list by user id
	GetListByUser(param *pb.ChatUserIdRequest) ([]*pb.ChatGroup, error)
}

type ChatGroupInteractorImpl struct {
	firebase usecase.Firebase
	// chatGroupRepository usecase.ChatGroupRepository
	// chatUserRepository  usecase.ChatUserRepository
}

func NewChatGroupInteractorImpl(
	fb usecase.Firebase,
	// cgR usecase.ChatGroupRepository,
	// cuR usecase.ChatUserRepository,
) ChatGroupInteractor {
	return &ChatGroupInteractorImpl{
		firebase: fb,
		// chatGroupRepository: cgR,
		// chatUserRepository:  cuR,
	}
}

func (i *ChatGroupInteractorImpl) Create(chatGroup *pb.ChatGroup) (*pb.ChatGroup, error) {

	// ユーザー登録
	// err := i.chatGroupRepository.Create(chatGroup)
	// if err != nil {
	// 	return chatGroup, err
	// }

	return chatGroup, nil
}

func (i *ChatGroupInteractorImpl) Update(chatGroup *pb.ChatGroup) (bool, error) {

	// ユーザー登録
	// err := i.chatGroupRepository.Update(chatGroup)
	// if err != nil {
	// 	return false, err
	// }

	return true, nil
}

func (i *ChatGroupInteractorImpl) Delete(param *pb.ChatIdRequest) (bool, error) {

	// ユーザー登録
	// err := i.chatGroupRepository.Delete(param.Id)
	// if err != nil {
	// 	return false, err
	// }

	return true, nil
}

func (i *ChatGroupInteractorImpl) GetById(param *pb.ChatIdRequest) (*pb.ChatGroup, error) {
	var (
		chatGroup *pb.ChatGroup
		// err       error
	)

	// ユーザー登録
	// chatGroup, err = i.chatGroupRepository.GetById(param.Id)
	// if err != nil {
	// 	log.Println("error is:", err)
	// 	return chatGroup, err
	// }

	return chatGroup, nil
}

func (i *ChatGroupInteractorImpl) GetByUuid(param *pb.ChatUuidRequest) (*pb.ChatGroup, error) {
	var (
		chatGroup *pb.ChatGroup
		// err       error
	)

	// ユーザー登録
	// chatGroup, err = i.chatGroupRepository.GetByUuid(param.Uuid)
	// if err != nil {
	// 	log.Println("error is:", err)
	// 	return chatGroup, err
	// }

	return chatGroup, nil
}

func (i *ChatGroupInteractorImpl) GetListByUser(param *pb.ChatUserIdRequest) ([]*pb.ChatGroup, error) {
	var (
		chatGroups []*pb.ChatGroup
		// err        error
	)

	// ユーザー登録
	// chatGroups, err = i.chatGroupRepository.GetListByUser(param.UserId)
	// if err != nil {
	// 	log.Println("error is:", err)
	// 	return chatGroups, err
	// }

	return chatGroups, nil
}
