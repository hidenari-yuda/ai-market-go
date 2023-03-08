package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/hidenari-yuda/ai-market-go/domain/utils"
	"github.com/hidenari-yuda/ai-market-go/pb"
	"github.com/hidenari-yuda/ai-market-go/usecase"
)

type ContentRepositoryImpl struct {
	Name     string
	executer SQLExecuter
}

func NewContentRepositoryImpl(ex SQLExecuter) usecase.ContentRepository {
	return &ContentRepositoryImpl{
		Name:     "ContentRepository",
		executer: ex,
	}
}

// create
func (r *ContentRepositoryImpl) Create(param *pb.Content) error {
	now := time.Now().UTC()

	lastId, err := r.executer.Exec(
		r.Name+"Create",
		`INSERT INTO contents (
			uuid,
			user_id,
			title,
			description,
			thumbnail,
			price,
			asp_rate,
			request_progress,
			is_open,
			like_count,
			view_count,
			review_count,
			purchase_count,
			created_at,
			updated_at,
			is_deleted
			) VALUES (
				?,
				?,
				?, 
				?,
				?,
				?,
				?,
				?,
				?,
				?,
				?,
				?,
				?,
				?,
				?,
				?
		)`,
		utils.CreateUuid(),
		param.UserId,
		param.Title,
		param.Description,
		param.Thumbnail,
		param.Price,
		param.AspRate,
		param.RequestProgress,
		param.IsOpen,
		param.LikeCount,
		param.ViewCount,
		param.ReviewCount,
		param.PurchaseCount,
		now,
		now,
		false,
	)

	if err != nil {
		return err
	}

	param.Id = lastId

	return nil
}

// update
func (r *ContentRepositoryImpl) Update(param *pb.Content) error {
		now := time.Now().UTC()

	lastId, err := r.executer.Exec(
		r.Name+"Update",
		`UPDATE contents SET
			title = ?,
			description = ?,
			thumbnail = ?,
			price = ?,
			asp_rate = ?,
			request_progress = ?,
			is_open = ?,
			like_count = ?,
			view_count = ?,
			review_count = ?,
			purchase_count = ?,
			updated_at = ?
		WHERE id = ?`,
		param.Title,
		param.Description,
		param.Thumbnail,
		param.Price,
		param.AspRate,
		param.RequestProgress,
		param.IsOpen,
		param.LikeCount,
		param.ViewCount,
		param.ReviewCount,
		param.PurchaseCount,
		now,
		param.Id,
	)

	if err != nil {
		err = fmt.Errorf("failed to update content: %w", err)
		log.Println(err)
		return err
	}

	param.Id = lastId

	return nil
}

// delete
func (r *ContentRepositoryImpl) Delete(id int64) error {
	_, err := r.executer.Exec(
		r.Name+"Delete",
		"UPDATE contents SET is_deleted = ? WHERE id = ?",
		true,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

// get
func (r *ContentRepositoryImpl) GetById(id int64) (*pb.Content, error) {
	var (
		content pb.Content
	)

	if err := r.executer.Get(
		r.Name+"GetById",
		&content,
		"SELECT * FROM contents WHERE id = ?",
		id,
	);err != nil {
		return nil, err
	}

	return &content, nil
}

// get by uuid
func (r *ContentRepositoryImpl) GetByUuid(uuid string) (*pb.Content, error) {
	var (
		content pb.Content
	)

	if err := r.executer.Get(
		r.Name+"GetByUuid",
		&content,
		"SELECT * FROM contents WHERE uuid = ?",
		uuid,
	); err != nil {
		return nil, err
	}

	return &content, nil
}

// getByUser
func (r *ContentRepositoryImpl) GetListByUser(userId int64) ([]*pb.Content, error) {
	var (
		contents []*pb.Content
	)

	if err := r.executer.Select(
		r.Name+"GetListByUser",
		&contents,
		"SELECT * FROM contents WHERE user_id = ?",
		userId,
	); err != nil {
		return nil, err
	}

	return contents, nil
}

// get list by category
func (r *ContentRepositoryImpl) GetListByFreeWord(freeWord string) ([]*pb.Content, error) {
	var (
		contents []*pb.Content
	)

	if err := r.executer.Select(
		r.Name+"GetListByFreeWord",
		&contents,
		"SELECT * FROM contents WHERE title LIKE ? OR description LIKE ?",
		"%"+freeWord+"%",
		"%"+freeWord+"%",
	); err != nil {
		return nil, err
	}

	return contents, nil
}

// get by latest id=user_id
func (r *ContentRepositoryImpl) GetLatestList() ([]*pb.Content, error) {
	var (
		contents []*pb.Content
	)

	if err := r.executer.Select(
		r.Name+"GetLatestList",
		&contents,
		"SELECT * FROM contents ORDER BY created_at DESC",
	); err != nil {
		return nil, err
	}

	return contents, nil
}

// get by trend id=user_id
func (r *ContentRepositoryImpl) GetTrendList() ([]*pb.Content, error) {
	var (
		contents []*pb.Content
	)

	if err := r.executer.Select(
		r.Name+"GetTrendList",
		&contents,
		"SELECT * FROM contents ORDER BY view_count DESC",
	); err != nil {
		return nil, err
	}

	return contents, nil
}

// get recommended list
func (r *ContentRepositoryImpl) GetRecommendedListByUser(userId int64) ([]*pb.Content, error) {
	var (
		contents []*pb.Content
	)

	if err := r.executer.Select(
		r.Name+"GetRecommendedListByUser",
		&contents,
		"SELECT * FROM contents WHERE user_id = ? ORDER BY view_count DESC",
		userId,
	); err != nil {
		return nil, err
	}

	return contents, nil
}

// get by sold id=user_id
func (r *ContentRepositoryImpl) GetSoldListByUser(userId int64) ([]*pb.Content, error) {
	var (
		contents []*pb.Content
	)

	if err := r.executer.Select(
		r.Name+"GetSoldListByUser",
		&contents,
		"SELECT * FROM contents WHERE user_id = ? ORDER BY created_at DESC",
		userId,
	); err != nil {
		return nil, err
	}

	return contents, nil
}

// get by purchased id=user_id
// GetPurchasedListByUser(userId int64) ([]*pb.Content, error)
func (r *ContentRepositoryImpl) GetPurchasedListByUser(userId int64) ([]*pb.Content, error) {
	var (
		contents []*pb.Content
	)

	if err := r.executer.Select(
		r.Name+"GetPurchasedListByUser",
		&contents,
		"SELECT * FROM contents WHERE user_id = ? ORDER BY created_at DESC",
		userId,
	); err != nil {
		return nil, err
	}

	return contents, nil
}

// get by liked id=user_id
// GetLikedListByUser(userId int64) ([]*pb.Content, error)
func (r *ContentRepositoryImpl) GetLikedListByUser(userId int64) ([]*pb.Content, error) {
	var (
		contents []*pb.Content
	)

	if err := r.executer.Select(
		r.Name+"GetLikedListByUser",
		&contents,
		"SELECT * FROM contents WHERE user_id = ? ORDER BY created_at DESC",
		userId,
	); err != nil {
		return nil, err
	}

	return contents, nil
}

// get list by id list
//
//	GetListByIdList(idList []int64) ([]*pb.Content, error)
func (r *ContentRepositoryImpl) GetListByIdList(idList []int64) ([]*pb.Content, error) {
	var (
		contents []*pb.Content
	)

	if err := r.executer.Select(
		r.Name+"GetListByIdList",
		&contents,
		"SELECT * FROM contents WHERE id IN (?)",
		idList,
	); err != nil {
		return nil, err
	}

	return contents, nil
}

// admin
// getAll
func (r *ContentRepositoryImpl) GetAll() ([]*pb.Content, error) {
	var (
		contents []*pb.Content
	)

	if err := r.executer.Select(
		r.Name+"GetAll",
		&contents,
		"SELECT * FROM contents ORDER BY id DESC",
	); err != nil {
		err = fmt.Errorf("failed to get all contents: %w", err)
		log.Println(err)
		return nil, err
	}

	return contents, nil
}
