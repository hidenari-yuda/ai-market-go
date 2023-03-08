package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/hidenari-yuda/ai-market-go/domain/utils"
	"github.com/hidenari-yuda/ai-market-go/pb"
	"github.com/hidenari-yuda/ai-market-go/usecase"
)

type ContentToolRepositoryImpl struct {
	Name     string
	executer SQLExecuter
}

func NewContentToolRepositoryImpl(ex SQLExecuter) usecase.ContentToolRepository {
	return &ContentToolRepositoryImpl{
		Name:     "ContentToolRepository",
		executer: ex,
	}
}

// create
func (r *ContentToolRepositoryImpl) Create(param *pb.ContentTool) error {
	now := time.Now().UTC()

	_, err := r.executer.Exec(
		r.Name+"Create",
		`INSERT INTO content_tools (
			uuid,
			content_id,
			tool,
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
		param.Tool,
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
func (r *ContentToolRepositoryImpl) Update(param *pb.ContentTool) error {
	now := time.Now().UTC()

	_, err := r.executer.Exec(
		r.Name+"Update",
		`UPDATE content_tools SET
			tool = ?,
			updated_at = ?
		WHERE id = ?`,
		param.Tool,
		now,
		param.Id,
	)

	if err != nil {
		err = fmt.Errorf("failed to update contentTool: %w", err)
		log.Println(err)
		return err
	}

	return nil
}

// delete
func (r *ContentToolRepositoryImpl) Delete(id int64) error {
	_, err := r.executer.Exec(
		r.Name+"Delete",
		`UPDATE content_tools SET is_deleted = ? WHERE id = ?`,
		true,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

// delete by content
func (r *ContentToolRepositoryImpl) DeleteByContent(contentId int64) error {
	_, err := r.executer.Exec(
		r.Name+"DeleteByContent",
		`DELETE FROM content_tools WHERE content_id = ?`,
		contentId,
	)

	if err != nil {
		return err
	}

	return nil
}

// get
func (r *ContentToolRepositoryImpl) GetById(id int64) (*pb.ContentTool, error) {
	var (
		contentTool pb.ContentTool
	)

	err := r.executer.Get(
		r.Name+"GetById",
		&contentTool,
		`SELECT * FROM content_tools WHERE id = ?`,
		id,
	)

	if err != nil {
		return nil, err
	}

	return &contentTool, nil
}

// getByContent
func (r *ContentToolRepositoryImpl) GetListByContent(contentId int64) ([]*pb.ContentTool, error) {
	var (
		contentTools []*pb.ContentTool
	)

	err := r.executer.Select(
		r.Name+"GetListByContent",
		&contentTools,
		`SELECT * FROM content_tools WHERE content_id = ?`,
		contentId,
	)

	if err != nil {
		return nil, err
	}

	return contentTools, nil
}

// get list by user
func (r *ContentToolRepositoryImpl) GetListByUser(userId int64) ([]*pb.ContentTool, error) {
	var (
		contentTools []*pb.ContentTool
	)

	err := r.executer.Select(
		r.Name+"GetListByUser",
		&contentTools,
		`
		SELECT * 
		FROM content_tools 
		WHERE content_id IN (
			SELECT id FROM contents WHERE user_id = ?
		)
		AND is_deleted = false
		ORDER BY id DESC
		`,
		userId,
	)

	if err != nil {
		return nil, err
	}

	return contentTools, nil
}

// get list by id list
func (r *ContentToolRepositoryImpl) GetListByIdList(idList []int64) ([]*pb.ContentTool, error) {
	var (
		contentTools []*pb.ContentTool
	)

	err := r.executer.Select(
		r.Name+"GetListByIdList",
		&contentTools,
		`SELECT * FROM content_tools WHERE id IN (?)`,
		idList,
	)

	if err != nil {
		return nil, err
	}

	return contentTools, nil
}

// get list by content id list
func (r *ContentToolRepositoryImpl) GetListByContentIdList(contentIdList []int64) ([]*pb.ContentTool, error) {
	var (
		contentTools []*pb.ContentTool
	)

	err := r.executer.Select(
		r.Name+"GetListByContentIdList",
		&contentTools,
		"SELECT * FROM content_tools WHERE content_id IN (?)",
		contentIdList,
	)

	if err != nil {
		return nil, err
	}

	return contentTools, nil
}

// admin
// getAll
func (r *ContentToolRepositoryImpl) GetAll() ([]*pb.ContentTool, error) {
	var (
		contentTools []*pb.ContentTool
	)
	err := r.executer.Select(
		r.Name+"GetAll",
		&contentTools,
		`SELECT * FROM content_tools ORDER BY id DESC`,
	)

	if err != nil {
		err = fmt.Errorf("failed to get all content_tools: %w", err)
		log.Println(err)
		return nil, err
	}

	return contentTools, nil
}
