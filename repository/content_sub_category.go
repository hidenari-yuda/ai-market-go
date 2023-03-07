package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/hidenari-yuda/ai-market-go/domain/utils"
	"github.com/hidenari-yuda/ai-market-go/pb"
	"github.com/hidenari-yuda/ai-market-go/usecase"
)

type ContentSubCategoryRepositoryImpl struct {
	Name     string
	executer SQLExecuter
}

func NewContentSubCategoryRepositoryImpl(ex SQLExecuter) usecase.ContentSubCategoryRepository {
	return &ContentSubCategoryRepositoryImpl{
		Name:     "ContentSubCategoryRepository",
		executer: ex,
	}
}

// create
func (r *ContentSubCategoryRepositoryImpl) Create(param *pb.ContentSubCategory) error {
	now := time.Now()

	_, err := r.executer.Exec(
		r.Name+"Create",
		`INSERT INTO content_sub_categories (
			uuid,
			content_id,
			sub_category,
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
		param.SubCategory,
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
func (r *ContentSubCategoryRepositoryImpl) Update(param *pb.ContentSubCategory) error {
	_, err := r.executer.Exec(
		r.Name+"Update",
		`UPDATE content_sub_categories SET
			subcategory = ?,
			updated_at = ?
		WHERE id = ?`,
		param.SubCategory,
		time.Now(),
		param.Id,
	)

	if err != nil {
		err = fmt.Errorf("failed to update contentSubCategory: %w", err)
		log.Println(err)
		return err
	}

	return nil
}

// delete
func (r *ContentSubCategoryRepositoryImpl) Delete(id int64) error {
	_, err := r.executer.Exec(
		r.Name+"Delete",
		"UPDATE content_sub_categories SET is_deleted = ? WHERE id = ?",
		true,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

// get
func (r *ContentSubCategoryRepositoryImpl) GetById(id int64) (*pb.ContentSubCategory, error) {
	var (
		contentSubCategory pb.ContentSubCategory
	)

	err := r.executer.Get(
		r.Name+"GetById",
		&contentSubCategory,
		"SELECT * FROM content_sub_categories WHERE id = ?",
		id,
	)

	if err != nil {
		return nil, err
	}

	return &contentSubCategory, nil
}

// getByContent
func (r *ContentSubCategoryRepositoryImpl) GetListByContent(contentId int64) ([]*pb.ContentSubCategory, error) {
	var (
		content_sub_categories []*pb.ContentSubCategory
	)

	err := r.executer.Select(
		r.Name+"GetListByContent",
		&content_sub_categories,
		"SELECT * FROM content_sub_categories WHERE content_id = ?",
		contentId,
	)

	if err != nil {
		return nil, err
	}

	return content_sub_categories, nil
}

// get list by id list
func (r *ContentSubCategoryRepositoryImpl) GetListByIdList(idList []int64) ([]*pb.ContentSubCategory, error) {
	var (
		content_sub_categories []*pb.ContentSubCategory
	)

	err := r.executer.Select(
		r.Name+"GetListByIdList",
		&content_sub_categories,
		"SELECT * FROM content_sub_categories WHERE id IN (?)",
		idList,
	)

	if err != nil {
		return nil, err
	}

	return content_sub_categories, nil
}

// admin
// getAll
func (r *ContentSubCategoryRepositoryImpl) GetAll() ([]*pb.ContentSubCategory, error) {
	var (
		content_sub_categories []*pb.ContentSubCategory
	)
	err := r.executer.Select(
		r.Name+"GetAll",
		&content_sub_categories,
		"SELECT * FROM content_sub_categories ORDER BY id DESC",
	)

	if err != nil {
		err = fmt.Errorf("failed to get all content_sub_categories: %w", err)
		log.Println(err)
		return nil, err
	}

	return content_sub_categories, nil
}
