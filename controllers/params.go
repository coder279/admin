package controllers

type ParamLogin struct {
	Mobile string `json:"mobile" binding:"required,verifyMobileFormat"`
	Password string `json:"password"`
	Captcha string `json:"captcha" binding:"required"`
}

type ParamSignup struct {
	Mobile string `json:"mobile" binding:"required,verifyMobileFormat"`
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	HeadImg string `json:"head_img"`
}






