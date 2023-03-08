package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/hidenari-yuda/ai-market-go/domain/utils"
	"github.com/hidenari-yuda/ai-market-go/pb"
	"github.com/hidenari-yuda/ai-market-go/usecase"
)

type ContentCategoryRepositoryImpl struct {
	Name     string
	executer SQLExecuter
}

func NewContentCategoryRepositoryImpl(ex SQLExecuter) usecase.ContentCategoryRepository {
	return &ContentCategoryRepositoryImpl{
		Name:     "ContentCategoryRepository",
		executer: ex,
	}
}

// create
func (r *ContentCategoryRepositoryImpl) Create(param *pb.ContentCategory) error {
	now := time.Now().UTC()

	_, err := r.executer.Exec(
		r.Name+"Create",
		`INSERT INTO content_categories (
			uuid,
			content_id,
			category,
			created_at,
			updated_at
			) VALUES (
				?,
				?,
				?, 
				?,
				?
		)`,
		utils.CreateUuid(),
		param.ContentId,
		param.Category,
		now,
		now,
		false,
	)

	if err != nil {
		return err
	}

	return nil
}

// update
func (r *ContentCategoryRepositoryImpl) Update(param *pb.ContentCategory) error {
		now := time.Now().UTC()

	_, err := r.executer.Exec(
		r.Name+"Update",
		`UPDATE content_categories SET
			category = ?,
			updated_at = ?
		WHERE id = ?`,
		param.Category,
		now,
		param.Id,
	)

	if err != nil {
		err = fmt.Errorf("failed to update contentCategory: %w", err)
		log.Println(err)
		return err
	}

	return nil
}

// delete
func (r *ContentCategoryRepositoryImpl) Delete(id int64) error {
	_, err := r.executer.Exec(
		r.Name+"Delete",
		"UPDATE content_categories SET is_deleted = ? WHERE id = ?",
		true,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

// delete by content
func (r *ContentCategoryRepositoryImpl) DeleteByContent(contentId int64) error {
	_, err := r.executer.Exec(
		r.Name+"DeleteByContent",
		"DELETE FROM content_categories WHERE content_id = ?",
		contentId,
	)

	if err != nil {
		return err
	}

	return nil
}

// get
func (r *ContentCategoryRepositoryImpl) GetById(id int64) (*pb.ContentCategory, error) {
	var (
		contentCategory pb.ContentCategory
	)

	err := r.executer.Get(
		r.Name+"GetById",
		&contentCategory,
		"SELECT * FROM content_categories WHERE id = ?",
		id,
	)

	if err != nil {
		return nil, err
	}

	return &contentCategory, nil
}

// getByContent
func (r *ContentCategoryRepositoryImpl) GetListByContent(contentId int64) ([]*pb.ContentCategory, error) {
	var (
		contentCategories []*pb.ContentCategory
	)

	err := r.executer.Select(
		r.Name+"GetListByContent",
		&contentCategories,
		"SELECT * FROM content_categories WHERE content_id = ?",
		contentId,
	)

	if err != nil {
		return nil, err
	}

	return contentCategories, nil
}

// get list by user
func (r *ContentCategoryRepositoryImpl) GetListByUser(userId int64) ([]*pb.ContentCategory, error) {
	var (
		contentCategories []*pb.ContentCategory
	)

	err := r.executer.Select(
		r.Name+"GetListByUser",
		&contentCategories,
		`SELECT * FROM content_categories WHERE content_id IN (
			SELECT id FROM contents WHERE user_id = ?
		)`,
		userId,
	)

	if err != nil {
		return nil, err
	}

	return contentCategories, nil
}

// get list by id list
func (r *ContentCategoryRepositoryImpl) GetListByIdList(idList []int64) ([]*pb.ContentCategory, error) {
	var (
		contentCategories []*pb.ContentCategory
	)

	err := r.executer.Select(
		r.Name+"GetListByIdList",
		&contentCategories,
		"SELECT * FROM content_categories WHERE id IN (?)",
		idList,
	)

	if err != nil {
		return nil, err
	}

	return contentCategories, nil
}

// get list by content id list
func (r *ContentCategoryRepositoryImpl) GetListByContentIdList(contentIdList []int64) ([]*pb.ContentCategory, error) {
	var (
		contentCategories []*pb.ContentCategory
	)

	err := r.executer.Select(
		r.Name+"GetListByContentIdList",
		&contentCategories,
		"SELECT * FROM content_categories WHERE content_id IN (?)",
		contentIdList,
	)

	if err != nil {
		return nil, err
	}

	return contentCategories, nil
}

// admin
// getAll
func (r *ContentCategoryRepositoryImpl) GetAll() ([]*pb.ContentCategory, error) {
	var (
		contentCategories []*pb.ContentCategory
	)
	err := r.executer.Select(
		r.Name+"GetAll",
		&contentCategories,
		"SELECT * FROM content_categories ORDER BY id DESC",
	)

	if err != nil {
		err = fmt.Errorf("failed to get all content_categories: %w", err)
		log.Println(err)
		return nil, err
	}

	return contentCategories, nil
}
