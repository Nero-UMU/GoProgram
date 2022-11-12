package model

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"conver_image_url"`
	State         uint8  `json:"state"`
}

// 文章
func (a Article) TableName() string {
	return "blog_article"
}
