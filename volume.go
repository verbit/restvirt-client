package restvirt

import (
	"bytes"
	"encoding/json"
)

type Volume struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
	Size int64  `json:"size"`
}

var volumePath = "volumes"

func (c *Client) CreateVolume(volume Volume) (string, error) {
	volumeJSON, err := json.Marshal(volume)
	if err != nil {
		return "", err
	}

	resp, err := c.doRequest("POST", bytes.NewBuffer(volumeJSON), volumePath)
	if err != nil {
		return "", err
	}

	err = checkForErrors(resp)
	if err != nil {
		return "", err
	}

	var vol Volume
	err = json.NewDecoder(resp.Body).Decode(&vol)
	if err != nil {
		return "", err
	}

	return vol.ID, nil
}

func (c *Client) GetVolume(id string) (*Volume, error) {
	resp, err := c.doRequest("GET", nil, volumePath, id)
	if err != nil {
		return nil, err
	}

	err = checkForErrors(resp)
	if err != nil {
		return nil, err
	}
	var volume Volume
	err = json.NewDecoder(resp.Body).Decode(&volume)
	if err != nil {
		return nil, err
	}

	return &volume, nil
}

func (c *Client) GetVolumes() ([]Volume, error) {
	resp, err := c.doRequest("GET", nil, volumePath)
	if err != nil {
		return nil, err
	}

	err = checkForErrors(resp)
	if err != nil {
		return nil, err
	}
	var volumes map[string][]Volume
	err = json.NewDecoder(resp.Body).Decode(&volumes)
	if err != nil {
		return nil, err
	}

	return volumes["volumes"], nil
}

func (c *Client) DeleteVolume(id string) error {
	resp, err := c.doRequest("DELETE", nil, volumePath, id)
	if err != nil {
		return err
	}

	err = checkForErrors(resp)
	if err != nil {
		return err
	}

	return nil
}
