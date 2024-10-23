package dtos

type UpdateSongsDTO struct {
	Name        string `form:"name" binding:"omitempty,max=100"`
	GroupId     string `json:"group_id" binding:"omitempty"`
	ReleaseDate string `json:"release_date" binding:"omitempty" validate:"DateValidation" format:"date" example:"2006-01-02"`
	Text        string `json:"text" binding:"omitempty,max=10000"`
	Link        string `form:"link" binding:"omitempty,url"`
}
