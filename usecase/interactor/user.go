package interactor

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log"

	"github.com/hidenari-yuda/ai-market-go/domain/config"
	"github.com/hidenari-yuda/ai-market-go/pb"
	"github.com/hidenari-yuda/ai-market-go/usecase"
	"golang.org/x/xerrors"
)

type UserInteractor interface {
	// Gest API
	// Create
	Create(user *pb.User) (*pb.User, error)

	// Update
	Update(user *pb.User) (bool, error)

	// Delete
	Delete(param *pb.UserIdRequest) (bool, error)

	// Get
	GetById(param *pb.UserIdRequest) (*pb.User, error)

	// get by uuid
	GetByUuid(param *pb.UserUuidRequest) (*pb.User, error)

	// get latest list
	GetLatestList() ([]*pb.User, error)

	// get trend list
	GetTrendList() ([]*pb.User, error)

	// get list by search
	GetListBySearch(param *pb.UserSearchRequest) ([]*pb.User, error)

	// admin API
	GetAll() ([]*pb.User, error)

	// auth
	SignIn(param *pb.SignInRequest) (*pb.User, error)
}

type UserInteractorImpl struct {
	firebase       usecase.Firebase
	userRepository usecase.UserRepository
}

func NewUserInteractorImpl(
	fb usecase.Firebase,
	uR usecase.UserRepository,
) UserInteractor {
	return &UserInteractorImpl{
		firebase:       fb,
		userRepository: uR,
	}
}

// Create
func (i *UserInteractorImpl) Create(user *pb.User) (*pb.User, error) {
	var (
		err error
	)

	// Firebase Authentication Create
	if user.FirebaseId == "" {
		firebaseID, err := i.firebase.CreateUser(
			user.Email,
			user.Password,
		)
		if err != nil {
			return nil, err
		}
		user.FirebaseId = firebaseID
	} else if user.Password == "" {
		mac := hmac.New(sha256.New, []byte(config.App.BasicSecret))
		_, err = mac.Write([]byte(user.FirebaseId))
		if err != nil {
			return nil, xerrors.Errorf("failed to write: %w", err)
		}
		user.Password = hex.EncodeToString(mac.Sum(nil))
	} else {
		err = xerrors.New("firebase id and password are not empty")
		return nil, err
	}

	// ユーザー登録
	if err := i.userRepository.Create(user); err != nil {
		return user, err
	}

	return user, nil
}

// Update
func (i *UserInteractorImpl) Update(user *pb.User) (bool, error) {

	// ユーザー登録
	if err := i.userRepository.Update(user); err != nil {
		return false, err
	}

	return true, nil
}

// delete
func (i *UserInteractorImpl) Delete(param *pb.UserIdRequest) (bool, error) {

	if err := i.userRepository.Delete(param.Id); err != nil {
		return false, err
	}

	return true, nil
}

// GetById
func (i *UserInteractorImpl) GetById(param *pb.UserIdRequest) (*pb.User, error) {
	var (
		user *pb.User
		err  error
	)

	// ユーザー登録
	user, err = i.userRepository.GetById(param.Id)
	if err != nil {
		log.Println("error is:", err)
		return user, err
	}

	return user, nil
}

// GetByUuid
func (i *UserInteractorImpl) GetByUuid(param *pb.UserUuidRequest) (*pb.User, error) {
	var (
		user *pb.User
		err  error
	)

	// ユーザー登録
	user, err = i.userRepository.GetByUuid(param.Uuid)
	if err != nil {
		log.Println("error is:", err)
		return user, err
	}

	return user, nil
}

// SignIn
func (i *UserInteractorImpl) SignIn(param *pb.SignInRequest) (*pb.User, error) {
	var (
		user *pb.User
		err  error
	)

	firebaseId, err := i.firebase.VerifyIDToken(param.Token)
	if err != nil {
		return user, err
	}

	log.Println("exported firebaseToken is:", param.Token)
	log.Println("exported firebaseId is:", firebaseId)

	// ユーザー登録
	user, err = i.userRepository.GetByFirebaseId(firebaseId)
	if err != nil {
		return user, err
	}

	return user, nil

}

// GetLatestList
func (i *UserInteractorImpl) GetLatestList() ([]*pb.User, error) {
	var (
		users []*pb.User
		err   error
	)

	users, err = i.userRepository.GetLatestList()
	if err != nil {
		return users, err
	}

	return users, nil
}

// GetTrendList
func (i *UserInteractorImpl) GetTrendList() ([]*pb.User, error) {
	var (
		users []*pb.User
		err   error
	)

	users, err = i.userRepository.GetTrendList()
	if err != nil {
		return users, err
	}

	return users, nil
}

// GetListBySearch
func (i *UserInteractorImpl) GetListBySearch(param *pb.UserSearchRequest) ([]*pb.User, error) {
	var (
		users []*pb.User
		err   error
	)

	users, err = i.userRepository.GetListBySearch(param.FreeWord)
	if err != nil {
		return users, err
	}

	return users, nil
}

// GetAll
func (i *UserInteractorImpl) GetAll() ([]*pb.User, error) {
	var (
		users []*pb.User
		err   error
	)

	users, err = i.userRepository.GetAll()
	if err != nil {
		return users, err
	}

	// i.userClient.DetectTextFromImage()

	return users, nil
}
