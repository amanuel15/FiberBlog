package interfaces

type ICreateBlog struct {
	Title      string       `json:"title" validate:"required,min=1,max=100"`
	Body       string       `json:"body" validate:"required,min=10"`
	Category   string       `json:"category,omitempty" validate:"omitempty"`
	Tags       []string     `json:"tags"`
	References []IReference `json:"references,omitempty" validate:"omitempty,dive"`
}

type IReference struct {
	Explanation string `json:"explanation" validate:"required,min=3,max=100"`
	Link        string `json:"link" validate:"required,url"`
	BlogID      uint   `json:"blog_id,omitempty"`
}
