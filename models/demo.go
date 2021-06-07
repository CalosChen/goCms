package models

type Respondent struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type RoleRespondents struct {
	ID           int
	RoleID       int
	RespondentID int
}

type Permission struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type RolePermissions struct {
	ID           int
	RoleID       int
	PermissionId int
}
