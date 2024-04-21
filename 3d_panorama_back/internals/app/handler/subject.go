package handler

import (
	"3d_panorama_back/internals/app/processor"
	"errors"
	"fmt"
	"net/http"
)

type SubjectHandler struct {
	processor *processor.SubjectProcessor
}

func NewSubjectHandler(processor *processor.SubjectProcessor) *SubjectHandler {
	handler := new(SubjectHandler)
	handler.processor = processor
	return handler
}

func (handler *SubjectHandler) Image(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	imageName := queryParams.Get("id")
	if imageName != "" {
		imagePath := "./image/" + imageName
		WrapOKImage(w, imagePath)
	}
	WrapError(w, errors.New("Имя изображения не указано"))
}


func (handler *SubjectHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	err, data := handler.processor.GetAll()
	if err != nil {
		WrapError(w,err)
		return
	}
	m := map[string]interface{}{
		"result": "OK",
		"data": data,
	}
	WrapOK(w,m)
}


func (handler *SubjectHandler) Upload(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Добавление записи в бд")
	owner_value, name_value, file_url := Upload(w,r);

	handler.processor.AddImage(name_value, file_url,owner_value)

}