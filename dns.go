package restvirt

import (
	"bytes"
	"encoding/json"
)

type DNSRecord struct {
	Name    string   `json:"name"`
	Type    string   `json:"type"`
	TTL     int      `json:"ttl"`
	Records []string `json:"records"`
}

func id(name string, recordType string) string {
	return name + "-" + recordType
}

func (m DNSRecord) ID() string {
	return id(m.Name, m.Type)
}

var dnsPath = "dns"

func (c *Client) SetDNSRecord(record DNSRecord) error {
	recordJSON, err := json.Marshal(record)
	if err != nil {
		return err
	}

	resp, err := c.doRequest("PUT", bytes.NewBuffer(recordJSON), dnsPath, record.ID())
	if err != nil {
		return err
	}

	err = checkForErrors(resp)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetDNSRecordByID(id string) (*DNSRecord, error) {
	resp, err := c.doRequest("GET", nil, dnsPath, id)
	if err != nil {
		return nil, err
	}

	err = checkForErrors(resp)
	if err != nil {
		return nil, err
	}

	var record DNSRecord
	err = json.NewDecoder(resp.Body).Decode(&record)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (c *Client) GetDNSRecord(name string, recordType string) (*DNSRecord, error) {
	return c.GetDNSRecordByID(id(name, recordType))
}

func (c *Client) DeleteDNSRecordByID(id string) error {
	resp, err := c.doRequest("DELETE", nil, dnsPath, id)
	if err != nil {
		return err
	}

	err = checkForErrors(resp)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteDNSRecord(name string, recordType string) error {
	return c.DeleteDNSRecordByID(id(name, recordType))
}
