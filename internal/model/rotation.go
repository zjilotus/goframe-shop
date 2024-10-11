package model

type RotationCreateUpdateBase struct {
	PicUrl string
	Link   string
	Sort   int
}

type RotationCreateInput struct {
	RotationCreateUpdateBase
}

type RotationCreateOutput struct {
	RotationId int `json:"rotation_id"`
}

type RotationUpdateInput struct {
	RotationCreateUpdateBase
	Id uint
}
