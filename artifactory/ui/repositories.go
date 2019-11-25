package ui

import "encoding/json"

// RepositoryDetails ..
type RepositoriesService Service

func (r RepositoryDetails) String() string {
	res, _ := json.MarshalIndent(r, "", "    ")
	return string(res)
}

// RepositoryDetails ...
type RepositoryDetails struct {
	RepoKey      *string `json:"repoKey,omitempty"`
	Type         *string `json:"type,omitempty"`
	IsLocal      *bool   `json:"isLocal,omitempty"`
	IsRemote     *bool   `json:"isRemote,omitempty"`
	IsVirtual    *bool   `json:"isVirtual,omitempty"`
	Distribution *bool   `json:"distribution,omitempty"`
	RepoType     *string `json:"repoType,omitempty"`
}
