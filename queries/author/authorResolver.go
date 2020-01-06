package queries

import "github.com/gographql/types"

/* let's mock some datas */
var authors = []types.Author{
	{
		ID:        1,
		Name:      "el kabach",
		Tutorials: []int{1, 2},
	},
	{
		ID:        1,
		Name:      "djibbril",
		Tutorials: []int{1, 2},
	},
}

// GetAllTutorial return all tutorials alviable
func GetAllAuthor() ([]types.Author, error) {
	// implement database acces
	return authors, nil
}
