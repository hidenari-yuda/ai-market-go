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
		content_categories []*pb.ContentCategory
	)

	err := r.executer.Select(
		r.Name+"GetListByContent",
		&content_categories,
		"SELECT * FROM content_categories WHERE content_id = ?",
		contentId,
	)

	if err != nil {
		return nil, err
	}

	return content_categories, nil
}

// get list by id list
func (r *ContentCategoryRepositoryImpl) GetListByIdList(idList []int64) ([]*pb.ContentCategory, error) {
	var (
		content_categories []*pb.ContentCategory
	)

	err := r.executer.Select(
		r.Name+"GetListByIdList",
		&content_categories,
		"SELECT * FROM content_categories WHERE id IN (?)",
		idList,
	)

	if err != nil {
		return nil, err
	}

	return content_categories, nil
}

// admin
// getAll
func (r *ContentCategoryRepositoryImpl) GetAll() ([]*pb.ContentCategory, error) {
	var (
		content_categories []*pb.ContentCategory
	)
	err := r.executer.Select(
		r.Name+"GetAll",
		&content_categories,
		"SELECT * FROM content_categories ORDER BY id DESC",
	)

	if err != nil {
		err = fmt.Errorf("failed to get all content_categories: %w", err)
		log.Println(err)
		return nil, err
	}

	return content_categories, nil
}
