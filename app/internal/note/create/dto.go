package create

type Input struct {
	Title   string `json:"title"   binding:"required" validate:"required,gt=0"`
	Content string `json:"content" binding:"required" validate:"required,gt=0"`
}

type Output struct {
	ID string
}
