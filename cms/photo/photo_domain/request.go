package photo_domain

type PhotoCreateRequest struct {
	CategoryId uint64 `json:"category_id" form:"category_id" validate:"required"`
}
