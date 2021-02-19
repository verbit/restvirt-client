package restvirt

import (
	"encoding/json"
)

type VolumeAttachment struct {
	DiskAddress string `json:"disk_address"`
}

func (c *Client) CreateAttachment(domainID string, volumeID string) (*VolumeAttachment, error) {
	resp, err := c.doRequest("PUT", nil, domainPath, domainID, volumePath, volumeID)
	if err != nil {
		return nil, err
	}

	err = checkForErrors(resp)
	if err != nil {
		return nil, err
	}

	var attachment VolumeAttachment
	err = json.NewDecoder(resp.Body).Decode(&attachment)
	if err != nil {
		return nil, err
	}

	return &attachment, nil
}

func (c *Client) GetAttachment(domainID string, volumeID string) (*VolumeAttachment, error) {
	resp, err := c.doRequest("GET", nil, domainPath, domainID, volumePath, volumeID)
	if err != nil {
		return nil, err
	}

	err = checkForErrors(resp)
	if err != nil {
		return nil, err
	}

	var attachment VolumeAttachment
	err = json.NewDecoder(resp.Body).Decode(&attachment)
	if err != nil {
		return nil, err
	}

	return &attachment, nil
}

func (c *Client) DeleteAttachment(domainID string, volumeID string) error {
	resp, err := c.doRequest("DELETE", nil, domainPath, domainID, volumePath, volumeID)
	if err != nil {
		return err
	}

	err = checkForErrors(resp)
	if err != nil {
		return err
	}

	return nil
}
