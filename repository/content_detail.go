package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/hidenari-yuda/ai-market-go/domain/utils"
	"github.com/hidenari-yuda/ai-market-go/pb"
	"github.com/hidenari-yuda/ai-market-go/usecase"
)

type ContentDetailRepositoryImpl struct {
	Name     string
	executer SQLExecuter
}

func NewContentDetailRepositoryImpl(ex SQLExecuter) usecase.ContentDetailRepository {
	return &ContentDetailRepositoryImpl{
		Name:     "ContentDetailRepository",
		executer: ex,
	}
}

// create
func (r *ContentDetailRepositoryImpl) Create(param *pb.ContentDetail) error {
	now := time.Now()

	_, err := r.executer.Exec(
		r.Name+"Create",
		`INSERT INTO content_tools (
			uuid,
			content_id,
			title,
			description,
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
		param.Title,
		param.Description,
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
func (r *ContentDetailRepositoryImpl) Update(param *pb.ContentDetail) error {
	_, err := r.executer.Exec(
		r.Name+"Update",
		`UPDATE content_tools SET
			title = ?,
			description = ?,
			updated_at = ?
		WHERE id = ?`,
		param.Title,
		param.Description,
		time.Now(),
		param.Id,
	)

	if err != nil {
		err = fmt.Errorf("failed to update contentDetail: %w", err)
		log.Println(err)
		return err
	}

	return nil
}

// delete
func (r *ContentDetailRepositoryImpl) Delete(id int64) error {
	_, err := r.executer.Exec(
		r.Name+"Delete",
		"UPDATE content_tools SET is_deleted = ? WHERE id = ?",
		true,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

// get
func (r *ContentDetailRepositoryImpl) GetById(id int64) (*pb.ContentDetail, error) {
	var (
		contentDetail pb.ContentDetail
	)

	err := r.executer.Get(
		r.Name+"GetById",
		&contentDetail,
		"SELECT * FROM content_tools WHERE id = ?",
		id,
	)

	if err != nil {
		return nil, err
	}

	return &contentDetail, nil
}

// getByContent
func (r *ContentDetailRepositoryImpl) GetListByContent(contentId int64) ([]*pb.ContentDetail, error) {
	var (
		content_tools []*pb.ContentDetail
	)

	err := r.executer.Select(
		r.Name+"GetListByContent",
		&content_tools,
		"SELECT * FROM content_tools WHERE content_id = ?",
		contentId,
	)

	if err != nil {
		return nil, err
	}

	return content_tools, nil
}

// get list by id list
func (r *ContentDetailRepositoryImpl) GetListByIdList(idList []int64) ([]*pb.ContentDetail, error) {
	var (
		content_tools []*pb.ContentDetail
	)

	err := r.executer.Select(
		r.Name+"GetListByIdList",
		&content_tools,
		"SELECT * FROM content_tools WHERE id IN (?)",
		idList,
	)

	if err != nil {
		return nil, err
	}

	return content_tools, nil
}

// admin
// getAll
func (r *ContentDetailRepositoryImpl) GetAll() ([]*pb.ContentDetail, error) {
	var (
		content_tools []*pb.ContentDetail
	)
	err := r.executer.Select(
		r.Name+"GetAll",
		&content_tools,
		"SELECT * FROM content_tools ORDER BY id DESC",
	)

	if err != nil {
		err = fmt.Errorf("failed to get all content_tools: %w", err)
		log.Println(err)
		return nil, err
	}

	return content_tools, nil
}
