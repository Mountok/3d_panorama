package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func WrapError(w http.ResponseWriter, err error) {
	WrapErrorWithStatus(w, err, http.StatusBadRequest)
}

func WrapErrorWithStatus(w http.ResponseWriter, err error, httpStatus int) {
	var m = map[string]string{
		"result": "error",
		"data":   err.Error(),
	}

	res, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.Header().Set("X-Content-Type-Options", "nosniff") //даем понять что ответ приходит в формате json
	w.WriteHeader(httpStatus)
	fmt.Fprintln(w, string(res))
}

func WrapOK(w http.ResponseWriter, m map[string]interface{}) {
	res, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(res))
}
func WrapOKImage(w http.ResponseWriter, m string) {
	fileBytes, err := os.ReadFile(m)
	if err != nil {
		WrapError(w, fmt.Errorf("файл не найден"))
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
}

func Upload(w http.ResponseWriter, r *http.Request) (owner_value,name_value,file_url string) {
	// Парсим мультипарт форму, содержащую файл изображения
	err := r.ParseMultipartForm(20 << 20)// Максимальный размер файла - 10MB
	if err != nil {
		WrapError(w, err)
		return
	}

	fmt.Println(r)

	file, handler, err := r.FormFile("image")
	owner_value = r.FormValue("owner")
	name_value = r.FormValue("name")

	file_url = handler.Filename


	if err != nil {
		WrapError(w, fmt.Errorf("Не удалось получить файл изображения"))
		return
	}
	defer file.Close()

	

	// Создаем файл в папке "image" и копируем в него содержимое полученного файла
	imagePath := "./image/" + handler.Filename
	newFile, err := os.Create(imagePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("файл успешно сохранен")

	return owner_value,name_value,file_url

}
