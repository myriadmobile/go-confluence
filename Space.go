package confluence

import (
	"encoding/json"
	"strings"

	"github.com/google/go-querystring/query"
	log "github.com/sirupsen/logrus"
)

// GetSpaces Returns all spaces in a Confluence instance.
// https://developer.atlassian.com/cloud/confluence/rest/#api-spaces-get
func (client Client) GetSpaces(qp *GetSpaceQueryParameters) ([]Space, error) {
	qp.ExpandString = strings.Join(qp.Expand, ",")
	v, _ := query.Values(qp)
	queryParams := v.Encode()

	body, err := client.request("GET", "/rest/api/space", queryParams, "")
	if err != nil {
		return nil, err
	}
	var spaceResponse SpaceResponse
	err = json.Unmarshal(body, &spaceResponse)
	if err != nil {
		log.Error("Unable to unmarshal SpaceResponse. Received: '", string(body), "'")
	}
	return spaceResponse.Results, err
}

// GetSpaceQueryParameters query parameters for GetSpaces
type GetSpaceQueryParameters struct {
	QueryParameters
	Expand           []string `url:"-"`
	ExpandString     string   `url:"expand,omitempty"`
	SpaceKey         string   `url:"spaceKey,omitempty"`
	Type             string   `url:"type,omitempty"`
	Status           string   `url:"status,omitempty"`
	Label            []string `url:"label,omitempty"`
	Favourite        bool     `url:"favourite,omitempty"`
	FavouriteUserKey string   `url:"favouriteUserKey,omitempty"`
	Start            int      `url:"start,omitempty"`
	Limit            int      `url:"limit,omitempty"`
}

// SpaceResponse represents the data returned from the Confluence API
type SpaceResponse struct {
	Results []Space `json:"results"`
}

// Space represents the data exchanged with the Confluence API
type Space struct {
	ID   int    `json:"id,omitempty"`
	Key  string `json:"key,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}
