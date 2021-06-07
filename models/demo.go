package models

type Respondent struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Permission struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
