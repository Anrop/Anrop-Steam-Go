package steam

import (
	"github.com/aoisensi/steam4go"
)

func QueryFiles(q string) (*steam4go.PublishedFileService, error) {
	query := steam4go.NewQueryFilesArgs()
	query.SetAppID(107410)
	query.SetNumPerPage(50)
	query.SetSearchText(q)

	steamApi := steam4go.NewSteamAPI(WebApiKey)
	result, err := steamApi.QueryFiles(query)
	return result, err
}
