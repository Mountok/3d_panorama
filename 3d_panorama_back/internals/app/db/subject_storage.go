package db

import (
	"3d_panorama_back/internals/app/models"
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

type SubjectStorage struct {
	databasePool *pgxpool.Pool
}

func NewSubjectStorage(pool *pgxpool.Pool) *SubjectStorage {
	storage := new(SubjectStorage)
	storage.databasePool = pool
	return storage
}

func (db *SubjectStorage) AddImage(file_name, file_url, owner string) (err error) {
	query := "insert into images (image_name,image_url,image_owner) values ($1,$2,$3);"
	_, err = db.databasePool.Exec(context.Background(), query, file_name, file_url, owner)
	if err != nil {
		log.Println("Ошибка sql запроса")
		return err
	}
	return nil
}

func (db *SubjectStorage) GetAll() []models.Image {
	var result []models.Image
	query := "select id,image_name,image_url,image_owner from images;"
	err := pgxscan.Select(context.Background(), db.databasePool, &result, query)
	if err != nil {
		log.Fatalln(err)
	}
	return result
}
