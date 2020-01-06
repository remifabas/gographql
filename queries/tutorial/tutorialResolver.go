package queries

import "github.com/gographql/types"

/* let's mock some datas */
var tutorials = []types.Tutorial{
	{
		ID:    12,
		Title: "JO Wil Tsonga",
	},
	{
		ID:    26,
		Title: "Jack Jack el dragon",
	},
}

// GetTutorialByID return a list with only the tutorial with the specify id
func GetTutorialByID(id int) ([]types.Tutorial, error) {
	var tutoResponse []types.Tutorial
	// Parse our tutorial array for the matching id
	for _, tutorial := range tutorials {
		if int(tutorial.ID) == id {
			// return our tutorial implemente database access
			tutoResponse = append(tutoResponse, tutorial)
			return tutoResponse, nil
		}
	}
	return tutoResponse, nil
}

// GetAllTutorial return all tutorials alviable
func GetAllTutorial() ([]types.Tutorial, error) {
	// implement database acces
	return tutorials, nil
}
