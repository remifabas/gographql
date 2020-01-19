package queries

import (
	"github.com/gographql/types"

	tuto "github.com/gographql/gateway/tutorials"
)

// GetTutorialByID return a list with only the tutorial with the specify id
/*func GetTutorialByID(id string) ([]types.Tutorial, error) {
	var tutoResponse []types.Tutorial
	// Parse our tutorial array for the matching id
	for _, tutorial := range tutorials {
		if tutorial.ID == id {
			// return our tutorial implemente database access
			tutoResponse = append(tutoResponse, tutorial)
			return tutoResponse, nil
		}
	}
	return tutoResponse, nil
}*/

// GetAllTutorial return all tutorials alviable
func GetAllTutorial() ([]types.Tutorial, error) {
	return tuto.FindAllTutorials(), nil
}
