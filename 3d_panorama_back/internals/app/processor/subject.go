package processor

import (
	"3d_panorama_back/internals/app/db"
	"3d_panorama_back/internals/app/models"
	"errors"
)

type SubjectProcessor struct {
	storage *db.SubjectStorage
}

func NewSubjectProcessor(storage *db.SubjectStorage) *SubjectProcessor {
	processor := new(SubjectProcessor)
	processor.storage = storage
	return processor
}

func (process *SubjectProcessor) AddImage(file_name, file_url, owner string) (error) {
	return process.storage.AddImage(file_name,file_url,owner)
}

func (processor *SubjectProcessor) GetAll() (error, []models.Image) { 
	data := processor.storage.GetAll()
	if len(data) == 0 {
		return errors.New("данные не найдены"), data
	}
	return nil, data
}