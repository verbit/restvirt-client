package restvirt

import (
	"bytes"
	"encoding/json"
)

type Domain struct {
	UUID      string `json:"uuid,omitempty"`
	Name      string `json:"name"`
	VCPU      int    `json:"vcpu"`
	MemoryMiB int    `json:"memory"`
	PrivateIP string `json:"private_ip"`
	UserData  string `json:"user_data,omitempty"`
}

var domainPath = "domains"

func (c *Client) CreateDomain(domain Domain) (string, error) {
	domainJSON, err := json.Marshal(domain)
	if err != nil {
		return "", err
	}

	resp, err := c.doRequest("POST", bytes.NewBuffer(domainJSON), domainPath)
	if err != nil {
		return "", err
	}

	err = checkForErrors(resp)
	if err != nil {
		return "", err
	}

	var uuidDomain Domain
	err = json.NewDecoder(resp.Body).Decode(&uuidDomain)
	if err != nil {
		return "", err
	}

	return uuidDomain.UUID, nil
}

func (c *Client) GetDomain(id string) (*Domain, error) {
	resp, err := c.doRequest("GET", nil, domainPath, id)
	if err != nil {
		return nil, err
	}

	err = checkForErrors(resp)
	if err != nil {
		return nil, err
	}
	var domain Domain
	err = json.NewDecoder(resp.Body).Decode(&domain)
	if err != nil {
		return nil, err
	}

	return &domain, nil
}

func (c *Client) DeleteDomain(id string) error {
	resp, err := c.doRequest("DELETE", nil, domainPath, id)
	if err != nil {
		return err
	}

	err = checkForErrors(resp)
	if err != nil {
		return err
	}

	return nil
}
