// Copyright 2013 Google. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or at
// https://developers.google.com/open-source/licenses/bsd

package github

import (
	"fmt"
)

// OrganizationsService provides access to the organization related functions
// in the GitHub API.
//
// GitHub API docs: http://developer.github.com/v3/orgs/
type OrganizationsService struct {
	client *Client
}

type Organization struct {
	Login     string `json:"login,omitempty"`
	ID        int    `json:"id,omitempty"`
	URL       string `json:"url,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
	Location  string `json:"location,omitempty"`
}

type Team struct {
	ID           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	URL          string `json:"url,omitempty"`
	Slug         string `json:"slug,omitempty"`
	Permission   string `json:"permission,omitempty"`
	MembersCount int    `json:"members_count,omitempty"`
	ReposCount   int    `json:"repos_count,omitempty"`
}

// List the organizations for a user.  Passing the empty string will list
// organizations for the authenticated user.
func (s *OrganizationsService) List(user string) ([]Organization, error) {
	var url string
	if user != "" {
		url = fmt.Sprintf("users/%v/orgs", user)
	} else {
		url = "user/orgs"
	}
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	orgs := new([]Organization)
	_, err = s.client.Do(req, orgs)
	return *orgs, err
}

// Get an organization.
func (s *OrganizationsService) Get(org string) (*Organization, error) {
	url := fmt.Sprintf("orgs/%v", org)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	organization := new(Organization)
	_, err = s.client.Do(req, organization)
	return organization, err
}

// Edit an organization.
func (s *OrganizationsService) Edit(name string, org *Organization) (*Organization, error) {
	url := fmt.Sprintf("orgs/%v", name)
	req, err := s.client.NewRequest("PATCH", url, org)
	if err != nil {
		return nil, err
	}

	updatedOrg := new(Organization)
	_, err = s.client.Do(req, updatedOrg)
	return updatedOrg, err
}

// List the members for an organization.  If the authenticated user is an owner
// of the organization, this will return concealed and public members,
// otherwise it will only return public members.
func (s *OrganizationsService) ListMembers(org string) ([]User, error) {
	url := fmt.Sprintf("orgs/%v/members", org)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	members := new([]User)
	_, err = s.client.Do(req, members)
	return *members, err
}

// List the public members for an organization.
func (s *OrganizationsService) ListPublicMembers(org string) ([]User, error) {
	url := fmt.Sprintf("orgs/%v/public_members", org)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	members := new([]User)
	_, err = s.client.Do(req, members)
	return *members, err
}

// List the teams for an organization.
func (s *OrganizationsService) ListTeams(org string) ([]Team, error) {
	url := fmt.Sprintf("orgs/%v/teams", org)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	teams := new([]Team)
	_, err = s.client.Do(req, teams)
	return *teams, err
}

// Add a user to a team.
func (s *OrganizationsService) AddTeamMember(team int, user string) error {
	url := fmt.Sprintf("teams/%v/members/%v", team, user)
	req, err := s.client.NewRequest("PUT", url, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}

// Remove a user from a team.
func (s *OrganizationsService) RemoveTeamMember(team int, user string) error {
	url := fmt.Sprintf("teams/%v/members/%v", team, user)
	req, err := s.client.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}

// Publicize a user's membership in an organization.
func (s *OrganizationsService) PublicizeMembership(org, user string) error {
	url := fmt.Sprintf("orgs/%v/public_members/%v", org, user)
	req, err := s.client.NewRequest("PUT", url, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}

// Conceal a user's membership in an organization.
func (s *OrganizationsService) ConcealMembership(org, user string) error {
	url := fmt.Sprintf("orgs/%v/public_members/%v", org, user)
	req, err := s.client.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}