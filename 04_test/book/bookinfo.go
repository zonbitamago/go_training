package book

import "fmt"

// Bookinfo 書籍情報取得
func Bookinfo(book Book) string {
	return fmt.Sprintf("ISBN is %v", book.ShowISBN())
}
