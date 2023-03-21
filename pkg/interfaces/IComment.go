package interfaces

type ICreateComment struct {
	BlogID   uint   `json:"blog_id" validate:"required"`
	Comment  string `json:"comment" validate:"required,min=1,max=1000"`
	ReplayTo uint   `json:"replay_to,omitempty"`
}
