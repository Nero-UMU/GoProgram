package model

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

// 标签
func (t Tag) TableName() string {
	return "blog_tag"
}
