package ui

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type SecurityService Service

type UserDetails struct {
	Name  *string `json:"name,omitempty"`
	Uri   *string `json:"uri,omitempty"`
	Realm *string `json:"realm,omitempty"`
}

func (r UserDetails) String() string {
	res, _ := json.MarshalIndent(r, "", "    ")
	return string(res)
}

// Get the users list
// Since: 2.4.0
// Notes: Requires Artifactory Pro
// Security: Requires an admin user
func (s *SecurityService) ListUsers(ctx context.Context) (*[]UserDetails, *http.Response, error) {
	path := "/ui/users"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Accept", mediaTypeJSON)

	users := new([]UserDetails)
	resp, err := s.client.Do(ctx, req, users)
	return users, resp, err
}

// application/vnd.org.jfrog.artifactory.security.User+json
type NewUser struct {
	Name                     *string  `json:"name,omitempty"`                     // Optional element in create/replace queries
	Email                    *string  `json:"email,omitempty"`                    // Mandatory element in create/replace queries, optional in "update" queries
	Password                 *string  `json:"password,omitempty"`                 // Mandatory element in create/replace queries, optional in "update" queries
	RetypePassword           *string  `json:"retypePassword,omitempty"`           // Mandatory element in create/replace queries, optional in "update" queries
	Admin                    *bool    `json:"admin,omitempty"`                    // Optional element in create/replace queries; Default: false
	ProfileUpdatable         *bool    `json:"profileUpdatable,omitempty"`         // Optional element in create/replace queries; Default: true
	DisableUIAccess          *bool    `json:"disableUIAccess,omitempty"`          // Optional element in create/replace queries; Default: false
	InternalPasswordDisabled *bool    `json:"internalPasswordDisabled,omitempty"` // Optional element in create/replace queries; Default: false
	Realm                    *string  `json:"realm,omitempty"`                    // Read-only element
	UserGroups               *[]Group `json:"groups,omitempty"`                   // Optional element in create/replace queries (requires groupName and Realm)
}

type User struct {
	Name                     *string   `json:"name,omitempty"`                     // Optional element in create/replace queries
	Email                    *string   `json:"email,omitempty"`                    // Mandatory element in create/replace queries, optional in "update" queries
	Password                 *string   `json:"password,omitempty"`                 // Mandatory element in create/replace queries, optional in "update" queries
	Admin                    *bool     `json:"admin,omitempty"`                    // Optional element in create/replace queries; Default: false
	ProfileUpdatable         *bool     `json:"profileUpdatable,omitempty"`         // Optional element in create/replace queries; Default: true
	DisableUIAccess          *bool     `json:"disableUIAccess,omitempty"`          // Optional element in create/replace queries; Default: false
	InternalPasswordDisabled *bool     `json:"internalPasswordDisabled,omitempty"` // Optional element in create/replace queries; Default: false
	LastLoggedIn             *string   `json:"lastLoggedIn,omitempty"`             // Read-only element
	Realm                    *string   `json:"realm,omitempty"`                    // Read-only element
	Groups                   *[]string `json:"groups,omitempty"`                   // Optional element in create/replace queries
}

func (r User) String() string {
	res, _ := json.MarshalIndent(r, "", "    ")
	return string(res)
}

// Get the details of an Artifactory user
// Since: 2.4.0
// Notes: Requires Artifactory Pro
// Security: Requires an admin user
func (s *SecurityService) GetUser(ctx context.Context, username string) (*User, *http.Response, error) {
	path := fmt.Sprintf("/ui/users/%s", username)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Accept", mediaTypeJSON)

	user := new(User)
	resp, err := s.client.Do(ctx, req, user)
	return user, resp, err
}

// Creates a new user in Artifactory
// Since: 2.4.0
// Notes: Requires Artifactory Pro
// Missing values will be set to the default values as defined by the consumed type.
// Security: Requires an admin user
func (s *SecurityService) CreateUser(ctx context.Context, username string, user *NewUser) (*http.Response, error) {
	path := fmt.Sprintf("/ui/users")
	req, err := s.client.NewJSONEncodedRequest("POST", path, user)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// Updates an exiting user in Artifactory with the provided user details.
// Since: 2.4.0
// Notes: Requires Artifactory Pro
// Missing values will be set to the default values as defined by the consumed type
// Security: Requires an admin user
func (s *SecurityService) UpdateUser(ctx context.Context, username string, user *User) (*http.Response, error) {
	path := fmt.Sprintf("/ui/users/%s", username)
	req, err := s.client.NewJSONEncodedRequest("PUT", path, user)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// Removes an Artifactory user.
// Since: 2.4.0
// Notes: Requires Artifactory Pro
// Security: Requires an admin user
func (s *SecurityService) DeleteUser(ctx context.Context, username string) (*string, *http.Response, error) {
	path := fmt.Sprintf("/ui/users/%v", username)
	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, nil, err
	}

	buf := new(bytes.Buffer)
	resp, err := s.client.Do(ctx, req, buf)
	if err != nil {
		return nil, resp, err
	}
	return String(buf.String()), resp, nil
}

// application/vnd.org.jfrog.artifactory.security.Group+json
type Group struct {
	Name            *string   `json:"name,omitempty"`            // Optional element in create/replace queries
	GroupName       *string   `json:"groupName,omitempty"`       // Optional element in create/replace queries
	Description     *string   `json:"description,omitempty"`     // Optional element in create/replace queries
	AutoJoin        *bool     `json:"autoJoin,omitempty"`        // Optional element in create/replace queries; default: false (must be false if adminPrivileges is true)
	AdminPrivileges *bool     `json:"adminPrivileges,omitempty"` // Optional element in create/replace queries; default: false
	Realm           *string   `json:"realm,omitempty"`           // Optional element in create/replace queries
	UsersInGroup    *[]string `json:"usersInGroup,omitempty"`    // Optional element in create/replace queries
	Permissions     *[]string `json:"permissions,omitempty"`     // Optional element in create/replace queries
	External        *bool     `json:"external,omitempty"`        // Optional element in create/replace queries
	NewUsersDefault *bool     `json:"newUserDefault,omitempty"`  // Optional element in create/replace queries
}

func (r Group) String() string {
	res, _ := json.MarshalIndent(r, "", "    ")
	return string(res)
}

// GetGroup Get the details of an Artifactory Group
// Since: 2.4.0
// Notes: Requires Artifactory Pro
// Security: Requires an admin user
func (s *SecurityService) GetGroup(ctx context.Context, groupName string) (*Group, *http.Response, error) {
	path := fmt.Sprintf("/ui/groups/%s", groupName)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Accept", mediaTypeJSON)

	group := new(Group)
	resp, err := s.client.Do(ctx, req, group)
	return group, resp, err
}

// CreateOrReplaceGroup Creates a new group in Artifactory or replaces an existing group
// Since: 2.4.0
// Notes: Requires Artifactory Pro
// Missing values will be set to the default values as defined by the consumed type.
// Security: Requires an admin user
func (s *SecurityService) CreateOrReplaceGroup(ctx context.Context, groupName string, group *Group) (*http.Response, error) {
	url := fmt.Sprintf("/ui/groups/%s", groupName)
	req, err := s.client.NewJSONEncodedRequest("PUT", url, group)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// UpdateGroup Updates an exiting group in Artifactory with the provided group details.
// Since: 2.4.0
// Notes: Requires Artifactory Pro
// Security: Requires an admin user
func (s *SecurityService) UpdateGroup(ctx context.Context, groupName string, group *Group) (*http.Response, error) {
	path := fmt.Sprintf("/ui/groups/%s", groupName)
	req, err := s.client.NewJSONEncodedRequest("POST", path, group)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// Removes an Artifactory group.
// Since: 2.4.0
// Notes: Requires Artifactory Pro
// Security: Requires an admin user
func (s *SecurityService) DeleteGroup(ctx context.Context, groupName string) (*string, *http.Response, error) {
	path := fmt.Sprintf("/ui/groups/%v", groupName)
	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, nil, err
	}

	buf := new(bytes.Buffer)
	resp, err := s.client.Do(ctx, req, buf)
	if err != nil {
		return nil, resp, err
	}
	return String(buf.String()), resp, nil
}

// Get the permission targets list
// Since: 2.4.0
// Notes: Requires Artifactory Pro
// Security: Requires an admin user
func (s *SecurityService) ListPermissionTargets(ctx context.Context) ([]*PermissionTargetsDetails, *http.Response, error) {
	path := "/ui/permissiontargets"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Accept", mediaTypeJSON)

	var permissionTargets []*PermissionTargetsDetails
	resp, err := s.client.Do(ctx, req, &permissionTargets)
	return permissionTargets, resp, err
}

// application/vnd.org.jfrog.artifactory.security.PermissionTarget+json
// Permissions are set/returned according to the following conventions:
//     m=admin; d=delete; w=deploy; n=annotate; r=read
type PermissionTargetsDetails struct {
	Name         *string              `json:"name,omitempty"`   // Optional element in create/replace queries
	Repositories *[]RepositoryDetails `json:"repos,omitempty"`  // Mandatory element in create/replace queries, optional in "update" queries
	Groups       *[]string            `json:"groups,omitempty"` // Optional element in create/replace queries
	Users        *[]string            `json:"users,omitempty"`  // Optional element in create/replace queries

}

func (r *PermissionTargetsDetails) String() string {
	res, _ := json.MarshalIndent(r, "", "    ")
	return string(res)
}

type PermissionTargets struct {
	UserPermissionActions  *[]UserPermissionAction  `json:"userPermissionActions"`
	GroupPermissionActions *[]GroupPermissionAction `json:"groupPermissionActions"`
	RepoKeys               *[]string                `json:"repoKeys"`
	IncludePatterns        *[]string                `json:"includePatterns"`
	ExcludePatterns        *[]string                `json:"excludePatterns"`
}

type GroupPermissionAction struct {
	Principal *string  `json:"principal,omitempty"`
	Actions   []string `json:"actions"`
}

type UserPermissionAction struct {
	Principal *string  `json:"principal,omitempty"`
	Actions   []string `json:"actions"`
}

// Get the details of an Artifactory Permission Target
// Since: 2.4.0
// Notes: Requires Artifactory Pro
// Security: Requires an admin user
func (s *SecurityService) GetPermissionTargets(ctx context.Context, permissionName string) (*PermissionTargets, *http.Response, error) {
	path := fmt.Sprintf("/ui/permissiontargets/%s", permissionName)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Accept", mediaTypeJSON)

	permission := new(PermissionTargets)
	resp, err := s.client.Do(ctx, req, permission)
	return permission, resp, err
}

// Creates a new permission target in Artifactory or replaces an existing permission target
// Since: 2.4.0
// Notes: Requires Artifactory Pro
// Missing values will be set to the default values as defined by the consumed type.
// Security: Requires an admin user
func (s *SecurityService) CreateOrReplacePermissionTargets(ctx context.Context, permissionName string, permissionTargets *PermissionTargets) (*http.Response, error) {
	path := fmt.Sprintf("/ui/permissiontargets/%s", permissionName)
	req, err := s.client.NewJSONEncodedRequest("PUT", path, permissionTargets)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// Deletes an Artifactory permission target.
// Since: 2.4.0
// Notes: Requires Artifactory Pro
// Security: Requires an admin user
func (s *SecurityService) DeletePermissionTargets(ctx context.Context, permissionName string) (*string, *http.Response, error) {
	path := fmt.Sprintf("/ui/permissiontargets/%s", permissionName)
	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, nil, err
	}

	buf := new(bytes.Buffer)
	resp, err := s.client.Do(ctx, req, buf)
	if err != nil {
		return nil, resp, err
	}
	return String(buf.String()), resp, nil
}
