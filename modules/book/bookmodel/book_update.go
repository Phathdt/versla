package bookmodel

type BookUpdate struct {
	Title  *string `json:"title" gorm:"column:title"`
	Author *string `json:"author" gorm:"column:author"`
}

func (BookUpdate) TableName() string {
	return "books"
}
