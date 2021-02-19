package restvirt

import (
	"bytes"
	"encoding/json"
	"strconv"
)

type PortForwarding struct {
	SourcePort uint16 `json:"source_port"`
	TargetPort uint16 `json:"target_port"`
	TargetIP   string `json:"target_ip"`
}

var forwardingPath = "forwardings"

func (c *Client) CreatePortForwarding(forwarding PortForwarding) (string, error) {
	forwardingJSON, err := json.Marshal(forwarding)
	if err != nil {
		return "", err
	}

	resp, err := c.doRequest("POST", bytes.NewBuffer(forwardingJSON), forwardingPath)
	if err != nil {
		return "", err
	}

	err = checkForErrors(resp)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(int(forwarding.SourcePort)), nil
}

func (c *Client) GetPortForwarding(id string) (*PortForwarding, error) {
	resp, err := c.doRequest("GET", nil, forwardingPath, id)
	if err != nil {
		return nil, err
	}

	err = checkForErrors(resp)
	if err != nil {
		return nil, err
	}

	var forwarding PortForwarding
	err = json.NewDecoder(resp.Body).Decode(&forwarding)
	if err != nil {
		return nil, err
	}

	return &forwarding, nil
}

func (c *Client) DeletePortForwarding(id string) error {
	resp, err := c.doRequest("DELETE", nil, forwardingPath, id)
	if err != nil {
		return err
	}

	err = checkForErrors(resp)
	if err != nil {
		return err
	}

	return nil
}
