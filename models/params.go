package models

//用户模块
type ParamLogin struct {
	Mobile string `json:"mobile" binding:"required,VerifyMobileFormat"`
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

//商品模块
type ParamGetProductList struct {
	Name string `json:"name" form:"name"`
	Page int `json:"page" form:"page" binding:"required"`
	Limit int `json:"limit" form:"limit"  binding:"required"`
}
type ParamGetProductDetail struct {
	Id int `json:"id" form:"id"`
}
type ParamGetProductCategory struct {
	Page int `json:"page" form:"page" binding:"required"`
	Limit int `json:"limit" form:"limit"  binding:"required"`
	CategoryId int `json:"category_id" form:"category_id" binding:"required"`
}







