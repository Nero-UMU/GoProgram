package model

type ArticleTag struct {
	*Model
	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

// 文章标签
func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
