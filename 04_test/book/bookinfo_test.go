package book

import (
	"go_training/04_test/book/mock_book"
	"go_training/04_test/testutils"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestBookinfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book := mock_book.NewMockBook(ctrl)
	book.EXPECT().ShowISBN().Return("mockedISBN")

	test := []struct {
		name string
		a    Book
		want string
	}{
		{"mock_sample", book, "ISBN is mockedISBN"},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bookinfo(tt.a); got != tt.want {
				testutils.ErrorfHandler(t, tt.want, got)
			}
		})
	}
}
