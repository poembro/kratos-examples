package model

func init() {
	Register(&Article{})
}

// Article saas平台管理员
type Article struct {
	BaseModel
	Title   string `json:"title"`
	Content string `json:"content"`
	Like    int64  `json:"like"`
}

func (Article) TableName() string {
	return "t_article"
}

func (Article) TableOptions() string {
	return "ENGINE=InnoDB COMMENT='管理员用户'"
}
