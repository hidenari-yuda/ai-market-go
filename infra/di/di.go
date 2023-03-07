package di

import (
	"context"

	"github.com/hidenari-yuda/ai-market-go/handler"
	"github.com/hidenari-yuda/ai-market-go/infra/database"
	"github.com/hidenari-yuda/ai-market-go/pb"
	"github.com/hidenari-yuda/ai-market-go/repository"
	"github.com/hidenari-yuda/ai-market-go/usecase"
	"github.com/hidenari-yuda/ai-market-go/usecase/interactor"
	"google.golang.org/grpc"
)

// RegisterServiceServer is a function to register service server
func RegisterServiceServer(ctx context.Context, s *grpc.Server, db *database.Db, firebase usecase.Firebase) {
	regsiterUserServiceServer(ctx, s, db, firebase)
	registerContentServiceServer(ctx, s, db, firebase)
	registerChatServiceServer(ctx, s, db, firebase)
	registerChatGroupServiceServer(ctx, s, db, firebase)
}

// regsiterUserServiceServer is a function to register user service server
func regsiterUserServiceServer(ctx context.Context, s *grpc.Server, db *database.Db, firebase usecase.Firebase) {
	userRepository := repository.NewUserRepositoryImpl(db)
	pb.RegisterUserServiceServer(s, handler.NewUserSercviceServer(interactor.NewUserInteractorImpl(firebase, userRepository)))
}

// content
func registerContentServiceServer(ctx context.Context, s *grpc.Server, db *database.Db, firebase usecase.Firebase) {
	contentRepository := repository.NewContentRepositoryImpl(db)
	contentContentRepository := repository.NewContentDetailRepositoryImpl(db)
	contentToolRepository := repository.NewContentToolRepositoryImpl(db)
	contentCategoryRepository := repository.NewContentCategoryRepositoryImpl(db)
	contentSubCategoryRepository := repository.NewContentSubCategoryRepositoryImpl(db)
	pb.RegisterContentServiceServer(s, handler.NewContentSercviceServer(interactor.NewContentInteractorImpl(firebase, contentRepository, contentContentRepository, contentToolRepository, contentCategoryRepository, contentSubCategoryRepository)))
}

// chatGroup
func registerChatGroupServiceServer(ctx context.Context, s *grpc.Server, db *database.Db, firebase usecase.Firebase) {
	// chatGroupRepository := repository.NewChatGroupRepositoryImpl(db)
	// chatUserRepository := repository.NewChatUserRepositoryImpl(db)
	pb.RegisterChatGroupServiceServer(s, handler.NewChatGroupSercviceServer(interactor.NewChatGroupInteractorImpl(firebase)))
}

// chat
func registerChatServiceServer(ctx context.Context, s *grpc.Server, db *database.Db, firebase usecase.Firebase) {
	// chatRepository := repository.NewChatRepositoryImpl(db)
	pb.RegisterChatServiceServer(s, handler.NewChatSercviceServer(interactor.NewChatInteractorImpl(firebase)))
}
