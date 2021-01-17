package models

type ParamLogin struct {
	Mobile string `json:"mobile" binding:"required,verifyMobileFormat"`
	Password string `json:"password"`
	Code string `json:"code"`
}

type ParamSignup struct {
	Mobile string `json:"mobile" binding:"required,VerifyMobileFormat"`
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	HeadImg string `json:"head_img"`

}

type ParamMobile struct {
	Mobile string `json:"mobile"`
	Code string `json:"code"`
}






