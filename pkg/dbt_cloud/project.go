package dbt_cloud

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Project struct {
	ID                     *int    `json:"id"`
	Name                   string  `json:"name"`
	DbtProjectSubdirectory *string `json:"dbt_project_subdirectory"`
	ConnectionID           *int    `json:"connection_id"`
	RepositoryID           *int    `json:"repository_id"`
	State                  int     `json:"state"`
}

type ProjectResponse struct {
	Data   Project        `json:"data"`
	Status ResponseStatus `json:"status"`
}

func (c *Client) GetProject(projectID string) (*Project, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/accounts/%s/v3/projects/%s/", c.HostURL, strconv.Itoa(c.AccountID), projectID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	projectResponse := ProjectResponse{}
	err = json.Unmarshal(body, &projectResponse)
	if err != nil {
		return nil, err
	}

	return &projectResponse.Data, nil
}

func (c *Client) CreateProject(name string, dbtProjectSubdirectory string, connectionID int, repositoryID int) (*Project, error) {
	newProject := Project{
		Name:                   name,
		DbtProjectSubdirectory: &dbtProjectSubdirectory,
		ConnectionID:           &connectionID,
		RepositoryID:           &repositoryID,
		State:                  1,
	}
	newProjectData, err := json.Marshal(newProject)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/accounts/%s/v3/projects/", c.HostURL, strconv.Itoa(c.AccountID)), strings.NewReader(string(newProjectData)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	projectResponse := ProjectResponse{}
	err = json.Unmarshal(body, &projectResponse)
	if err != nil {
		return nil, err
	}

	return &projectResponse.Data, nil
}

func (c *Client) UpdateProject(projectID string, project Project) (*Project, error) {
	projectData, err := json.Marshal(project)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/accounts/%s/v3/projects/%s/", c.HostURL, strconv.Itoa(c.AccountID), projectID), strings.NewReader(string(projectData)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	projectResponse := ProjectResponse{}
	err = json.Unmarshal(body, &projectResponse)
	if err != nil {
		return nil, err
	}

	return &projectResponse.Data, nil
}
