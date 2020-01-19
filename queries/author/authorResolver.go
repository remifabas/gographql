package queries

import (
	"github.com/gographql/types"

	author "github.com/gographql/gateway/author"
)

// GetAllAuthor return all tutorials available
func GetAllAuthor() ([]types.Author, error) {
	author.TestRem()
	return author.FindAllAuthor(), nil
}
