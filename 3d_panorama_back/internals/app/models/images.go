package models

type Image struct {
	Id         int    `json:"id"`
	ImageName  string `json:"image_name"`
	ImageUrl   string `json:"image_url"`
	ImageOwner string `json:"image_owner"`
}
