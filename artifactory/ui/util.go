package ui

import "github.com/rickardl/go-artifactory/v2/artifactory/client"

func String(v string) *string { return &v }

func NewUI(client *client.Client) *UI {
	v := &UI{}
	v.common.client = client

	v.Repositories = (*RepositoriesService)(&v.common)
	v.Security = (*SecurityService)(&v.common)

	return v
}
