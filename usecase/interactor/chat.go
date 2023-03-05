package interactor

import (
	"context"
	"fmt"
	"log"

	"github.com/hidenari-yuda/ai-market-go/pb"
	"github.com/hidenari-yuda/ai-market-go/usecase"
	"golang.org/x/sync/errgroup"
)

type ChatInteractor interface {
	// Gest API
	// Create
	Create(Chat *pb.Chat) (*pb.Chat, error)

	// Update
	Update(Chat *pb.Chat) (bool, error)

	// Delete
	Delete(param *pb.ChatIdRequest) (bool, error)

	// Get
	GetById(param *pb.ChatIdRequest) (*pb.Chat, error)

	// get list by user id
	GetListByGroup(param *pb.ChatGroupIdRequest) ([]*pb.Chat, error)

	// stream
	GetStream(ctx context.Context, stream chan<- pb.Chat) error
}

type ChatInteractorImpl struct {
	firebase usecase.Firebase
	// chatRepository      usecase.ChatRepository
	chatRepository usecase.ChatRepository
	// chatUserRepository  usecase.ChatRepository
}

func NewChatInteractorImpl(
	fb usecase.Firebase,
	// cR usecase.ChatRepository,
	cgR usecase.ChatRepository,
	// cuR usecase.ChatUserRepository,
) ChatInteractor {
	return &ChatInteractorImpl{
		firebase: fb,
		// chatRepository:      cR,
		chatRepository: cgR,
		// chatUserRepository:  cuR,
	}
}

func (i *ChatInteractorImpl) Create(chat *pb.Chat) (*pb.Chat, error) {

	// ユーザー登録
	err := i.chatRepository.Create(chat)
	if err != nil {
		return chat, err
	}

	return chat, nil
}

func (i *ChatInteractorImpl) Update(chat *pb.Chat) (bool, error) {

	// ユーザー登録
	err := i.chatRepository.Update(chat)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (i *ChatInteractorImpl) Delete(param *pb.ChatIdRequest) (bool, error) {

	// ユーザー登録
	err := i.chatRepository.Delete(param.Id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (i *ChatInteractorImpl) GetById(param *pb.ChatIdRequest) (*pb.Chat, error) {
	var (
		chat *pb.Chat
		err  error
	)

	// ユーザー登録
	chat, err = i.chatRepository.GetById(param.Id)
	if err != nil {
		log.Println("error is:", err)
		return chat, err
	}

	return chat, nil
}

func (i *ChatInteractorImpl) GetListByGroup(param *pb.ChatGroupIdRequest) ([]*pb.Chat, error) {
	var (
		chats []*pb.Chat
		err   error
	)

	// ユーザー登録
	chats, err = i.chatRepository.GetListByGroup(param.GroupId)
	if err != nil {
		log.Println("error is:", err)
		return chats, err
	}

	return chats, nil
}

// get stream
func (i *ChatInteractorImpl) GetStream(ctx context.Context, stream chan<- pb.Chat) error {
	defer close(stream)
	eg, _ := errgroup.WithContext(ctx)
	eg.Go(func() error {
		err := i.firebase.GetChatStream(ctx, stream)
		if err != nil {
			return err
		}
		return nil
	})
	if err := eg.Wait(); err != nil {
		return fmt.Errorf("failed to GetMessageStreamService.Handle: %s", err)
	}
	return nil
}
