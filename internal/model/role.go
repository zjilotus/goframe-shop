package model

type RoleCreateUpdateBase struct {
	Name string
	Desc string
}

type RoleCreateInput struct {
	RoleCreateUpdateBase
}

type RoleCreateOutput struct {
	RoleId int `json:"role_id"`
}
