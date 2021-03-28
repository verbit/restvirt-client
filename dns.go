package restvirt

import (
	"bytes"
	"encoding/json"
)

type DNSMapping struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
}

var dnsPath = "dns"

func (c *Client) SetDNSMapping(mapping DNSMapping) error {
	forwardingJSON, err := json.Marshal(mapping)
	if err != nil {
		return err
	}

	resp, err := c.doRequest("PUT", bytes.NewBuffer(forwardingJSON), dnsPath, mapping.Name)
	if err != nil {
		return err
	}

	err = checkForErrors(resp)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetDNSMapping(name string) (*DNSMapping, error) {
	resp, err := c.doRequest("GET", nil, dnsPath, name)
	if err != nil {
		return nil, err
	}

	err = checkForErrors(resp)
	if err != nil {
		return nil, err
	}

	var forwarding DNSMapping
	err = json.NewDecoder(resp.Body).Decode(&forwarding)
	if err != nil {
		return nil, err
	}

	forwarding.Name = name
	return &forwarding, nil
}

func (c *Client) DeleteDNSMapping(name string) error {
	resp, err := c.doRequest("DELETE", nil, dnsPath, name)
	if err != nil {
		return err
	}

	err = checkForErrors(resp)
	if err != nil {
		return err
	}

	return nil
}
