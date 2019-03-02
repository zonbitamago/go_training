package book

// Book 本
type Book interface {
	ShowISBN() string
}

// Novel 小説
type Novel struct {
	ISBN string
}

// ShowISBN ISBN情報取得
func (novel *Novel) ShowISBN() string {
	return novel.ISBN
}
