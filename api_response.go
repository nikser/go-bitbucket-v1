/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package bitbucketv1

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/mitchellh/mapstructure"
)

// APIResponse contains generic data from API Response
type APIResponse struct {
	*http.Response `json:"-"`
	Message        string `json:"message,omitempty"`
	// Operation is the name of the swagger operation.
	Operation string `json:"operation,omitempty"`
	// RequestURL is the request URL. This value is always available, even if the
	// embedded *http.Response is nil.
	RequestURL string `json:"url,omitempty"`
	// Method is the HTTP method used for the request.  This value is always
	// available, even if the embedded *http.Response is nil.
	Method string `json:"method,omitempty"`
	// Payload holds the contents of the response body (which may be nil or empty).
	// This is provided here as the raw response.Body() reader will have already
	// been drained.
	Payload []byte `json:"-"`
	Values  map[string]interface{}
}

type SelfLink struct {
	Href string `json:"href"`
}

type CloneLink struct {
	Href string `json:"href"`
	Name string `json:"name"`
}

type Links struct {
	Self []SelfLink `json:"self"`
}

type Project struct {
	Key         string `json:"key"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Public      bool   `json:"public"`
	Type        string `json:"type"`
	Links       Links  `json:"links"`
}

// Repository contains data from a BitBucket Repository
type Repository struct {
	Slug          string  `json:"slug"`
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	ScmID         string  `json:"scmId"`
	State         string  `json:"state"`
	StatusMessage string  `json:"statusMessage"`
	Forkable      bool    `json:"forkable"`
	Project       Project `json:"project"`
	Public        bool    `json:"public"`
	Links         struct {
		Clone []CloneLink `json:"clone"`
		Self  []SelfLink  `json:"self"`
	} `json:"links"`
}

type UserWithNameEmail struct {
	Name         string `json:"name"`
	EmailAddress string `json:"emailAddress"`
}

type User struct {
	Name        string `json:"name"`
	Email       string `json:"emailAddress"`
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	Active      bool   `json:"active"`
	Slug        string `json:"slug"`
	Type        string `json:"type"`
	Links       Links  `json:"links"`
}

type UserWithMetadata struct {
	User     User   `json:"user"`
	Role     string `json:"role"`
	Approved bool   `json:"approved"`
	Status   string `json:"status"`
}

type MergeResult struct {
	Outcome string `json:"outcome"`
	Current bool   `json:"current"`
}

type PullRequestRef struct {
	ID           string     `json:"id"`
	DisplayID    string     `json:"displayId"`
	LatestCommit string     `json:"latestCommit"`
	Repository   Repository `json:"repository"`
}

type PullRequest struct {
	ID           int                `json:"id"`
	Version      int                `json:"version"`
	Title        string             `json:"title"`
	State        string             `json:"state"`
	Open         bool               `json:"open"`
	Closed       bool               `json:"closed"`
	CreatedDate  int64              `json:"createdDate"`
	UpdatedDate  int64              `json:"updatedDate"`
	FromRef      PullRequestRef     `json:"fromRef"`
	ToRef        PullRequestRef     `json:"toRef"`
	Locked       bool               `json:"locked"`
	Author       UserWithMetadata   `json:"author"`
	Reviewers    []UserWithMetadata `json:"reviewers"`
	Participants []UserWithMetadata `json:"participants"`
	Properties   struct {
		MergeResult       MergeResult `json:"mergeResult"`
		ResolvedTaskCount int         `json:"resolvedTaskCount"`
		OpenTaskCount     int         `json:"openTaskCount"`
	} `json:"properties"`
	Links Links `json:"links"`
}

// SSHKey contains data from a SSHKey in the BitBucket Server
type SSHKey struct {
	ID    int    `json:"id"`
	Text  string `json:"text"`
	Label string `json:"label"`
}

// Commit contains data from a commit in BitBucket
type Commit struct {
	ID                 string            `json:"id"`
	DisplayID          string            `json:"displayId"`
	Author             UserWithNameEmail `json:"author"`
	AuthorTimestamp    int64             `json:"authorTimestamp"`
	Committer          UserWithNameEmail `json:"committer"`
	CommitterTimestamp int64             `json:"committerTimestamp"`
	Message            string            `json:"message"`
	Parents            []struct {
		ID        string `json:"id"`
		DisplayID string `json:"displayId"`
	} `json:"parents"`
}

type Diff struct {
	Diffs []struct {
		Source struct {
			Components []string `json:"components"`
			Parent     string   `json:"parent"`
			Name       string   `json:"name"`
			Extension  string   `json:"extension"`
			ToString   string   `json:"toString"`
		} `json:"source"`
		Destination struct {
			Components []string `json:"components"`
			Parent     string   `json:"parent"`
			Name       string   `json:"name"`
			Extension  string   `json:"extension"`
			ToString   string   `json:"toString"`
		} `json:"destination"`
		Hunks []struct {
			SourceLine      int `json:"sourceLine"`
			SourceSpan      int `json:"sourceSpan"`
			DestinationLine int `json:"destinationLine"`
			DestinationSpan int `json:"destinationSpan"`
			Segments        []struct {
				Type  string `json:"type"`
				Lines []struct {
					Destination int    `json:"destination"`
					Source      int    `json:"source"`
					Line        string `json:"line"`
					Truncated   bool   `json:"truncated"`
				} `json:"lines"`
				Truncated bool `json:"truncated"`
			} `json:"segments"`
			Truncated bool `json:"truncated"`
		} `json:"hunks"`
		Truncated    bool    `json:"truncated"`
		ContextLines float64 `json:"contextLines"`
		FromHash     string  `json:"fromHash"`
		ToHash       string  `json:"toHash"`
		WhiteSpace   string  `json:"whiteSpace"`
	} `json:"diffs"`
	Truncated    bool    `json:"truncated"`
	ContextLines float64 `json:"contextLines"`
	FromHash     string  `json:"fromHash"`
	ToHash       string  `json:"toHash"`
	WhiteSpace   string  `json:"whiteSpace"`
}

// Tag contaings git Tag information
type Tag struct {
	ID              string `json:"id"`
	DisplayID       string `json:"displayId"`
	Type            string `json:"type"`
	LatestCommit    string `json:"latestCommit"`
	LatestChangeset string `json:"latestChangeset"`
	Hash            string `json:"hash"`
}

// Branch contains git Branch information
type Branch struct {
	ID              string `json:"id"`
	DisplayID       string `json:"displayId"`
	Type            string `json:"type"`
	LatestCommit    string `json:"latestCommit"`
	LatestChangeset string `json:"latestChangeset"`
	IsDefault       bool   `json:"isDefault"`
}

func (k *SSHKey) String() string {
	parts := make([]string, 1, 2)
	parts[0] = strings.TrimSpace(k.Text)
	return strings.Join(parts, " ")
}

// GetCommitsResponse cast Commits into structure
func GetCommitsResponse(r *APIResponse) ([]Commit, error) {
	var m []Commit
	err := mapstructure.Decode(r.Values["values"], &m)
	return m, err
}

// GetTagsResponse cast Tags into structure
func GetTagsResponse(r *APIResponse) ([]Tag, error) {
	var m []Tag
	err := mapstructure.Decode(r.Values["values"], &m)
	return m, err
}

// GetBranchesResponse cast Tags into structure
func GetBranchesResponse(r *APIResponse) ([]Branch, error) {
	var m []Branch
	err := mapstructure.Decode(r.Values["values"], &m)
	return m, err
}

// GetRepositoriesResponse cast Repositories into structure
func GetRepositoriesResponse(r *APIResponse) ([]Repository, error) {
	var m []Repository
	err := mapstructure.Decode(r.Values["values"], &m)
	return m, err
}

// GetRepositoryResponse cast Repositories into structure
func GetRepositoryResponse(r *APIResponse) (Repository, error) {
	var m Repository
	err := mapstructure.Decode(r.Values, &m)
	return m, err
}

// GetDiffResponse cast Diff into structure
func GetDiffResponse(r *APIResponse) (Diff, error) {
	var m Diff
	err := mapstructure.Decode(r.Values, &m)
	return m, err
}

// GetSSHKeysResponse cast SSHKeys into structure
func GetSSHKeysResponse(r *APIResponse) ([]SSHKey, error) {
	var m []SSHKey
	err := mapstructure.Decode(r.Values["values"], &m)
	return m, err
}

// NewAPIResponse create new APIResponse from http.Response
func NewAPIResponse(r *http.Response) *APIResponse {

	response := &APIResponse{Response: r}
	return response
}

// NewAPIResponseWithError create new erroneous API response from http.response and error
func NewAPIResponseWithError(r *http.Response, err error) (*APIResponse, error) {

	response := &APIResponse{Response: r, Message: err.Error()}
	return response, err
}

// NewBitbucketAPIResponse create new API response from http.response
func NewBitbucketAPIResponse(r *http.Response) (*APIResponse, error) {
	response := &APIResponse{Response: r}
	err := json.NewDecoder(r.Body).Decode(&response.Values)
	if err != nil {
		return nil, err
	}
	return response, err
}
