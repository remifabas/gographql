package queries

import (
	"github.com/gographql/types"

	author "github.com/gographql/gateway/author"
)

// GetAllAuthor return all tutorials available
func GetAllAuthor() ([]types.Author, error) {
	return author.FindAllAuthor(), nil
}
