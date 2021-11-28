package dto

//VideoUpdateDTO used when PUT method impelement
type VideoUpdateDTO struct {
	ID       uint64 `binding:"required" form:"id" json:"id"`
	Title    string `binding:"required" form:"title" json:"title"`
	Semester int16  `binding:"required" form:"semester" json:"semester"`
	URL      string `binding:"required" form:"url" json:"url"`
	UserID   uint64 `binding:"required" form:"user_id,omitempty" json:"user_id,omitempty"`
}

//VideoCreateDTO used when create new Video
type VideoCreateDTO struct {
	Title    string `binding:"required" form:"title" json:"title"`
	Semester int16  `binding:"required" form:"semester" json:"semester"`
	URL      string `binding:"required" form:"url" json:"url"`
	UserID   uint64 `binding:"required" form:"user_id,omitempty" json:"user_id,omitempty"`
}
