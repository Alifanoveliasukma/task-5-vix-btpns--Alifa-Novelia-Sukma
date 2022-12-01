package dto

//ImageupdateDYO is a model that client use when update a image
type ImageUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Title    string `json:"title" form:"title" binding:"required"`
	Caption  string `json:"caption" form:"caption" binding:"required"`
	PhotoUrl string `json:"photoulr" form:"photourl" binding:"required"`
	UserID   uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}

//Book CreatedDTO is a model that client use when create a new image
type ImageCreateDTO struct {
	Title    string `json:"title" form:"title" binding:"required"`
	Caption  string `json:"caption" form:"caption" binding:"required"`
	PhotoUrl string `json:"photoulr" form:"photourl" binding:"required"`
	UserID   uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}
