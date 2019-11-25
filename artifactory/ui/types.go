package ui

import "github.com/rickardl/go-artifactory/v2/artifactory/client"

const (
	mediaTypeJSON = "application/json"
)

type Service struct {
	client *client.Client
}

type UI struct {
	common Service

	// Services used for talking to different parts of the Artifactory UI API.
	Repositories *RepositoriesService
	Security     *SecurityService
}
