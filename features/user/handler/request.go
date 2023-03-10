package handler

import (
	"campyuk-api/features/user"
	"mime/multipart"
)

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type RegisterRequest struct {
	Username string `json:"username" form:"username"`
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
}

type UpdateRequest struct {
	Username   string `json:"username" form:"username"`
	Fullname   string `json:"fullname" form:"fullname"`
	Email      string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
	Role       string `json:"role" form:"role"`
	UserImage  string `json:"user_image" form:"user_image"`
	FileHeader multipart.FileHeader
}

func ReqToCore(data interface{}) *user.Core {
	res := user.Core{}

	switch data.(type) {
	case LoginRequest:
		cnv := data.(LoginRequest)
		res.Username = cnv.Username
		res.Password = cnv.Password
	case RegisterRequest:
		cnv := data.(RegisterRequest)
		res.Username = cnv.Username
		res.Fullname = cnv.Fullname
		res.Email = cnv.Email
		res.Password = cnv.Password
		res.Role = cnv.Role
	case UpdateRequest:
		cnv := data.(UpdateRequest)
		res.Username = cnv.Username
		res.Fullname = cnv.Fullname
		res.Email = cnv.Email
		res.Password = cnv.Password
		res.UserImage = cnv.UserImage
		res.Role = cnv.Role
	default:
		return nil
	}

	return &res
}
