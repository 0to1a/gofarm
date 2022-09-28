package exampleModule

type UserAccess struct {
	UserId int64  `json:"uid"`
	Email  string `json:"email"`
	Name   string `json:"name"`
}
