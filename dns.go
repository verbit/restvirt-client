package restvirt

import (
	"context"
	"fmt"
	"strings"

	"github.com/verbit/restvirt-client/pb"
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

func parseDNSID(id string) (string, string, error) {
	s := strings.Split(id, "-")
	if len(s) != 2 {
		return "", "", fmt.Errorf("ID looks fishy")
	}
	return s[0], s[1], nil
}

func (c *Client) SetDNSRecord(record DNSRecord) error {
	_, err := c.DNSClient.PutDNSRecord(context.Background(), &pb.PutDNSRecordRequest{DnsRecord: &pb.DNSRecord{
		Name:    record.Name,
		Type:    record.Type,
		Ttl:     uint64(record.TTL),
		Records: record.Records,
	}})
	return err
}

func (c *Client) GetDNSRecordByID(id string) (*DNSRecord, error) {
	name, dtype, err := parseDNSID(id)
	if err != nil {
		return nil, err
	}
	return c.GetDNSRecord(name, dtype)
}

func (c *Client) GetDNSRecord(name string, recordType string) (*DNSRecord, error) {
	record, err := c.DNSClient.GetDNSRecord(context.Background(), &pb.DNSRecordIdentifier{
		Name: name,
		Type: recordType,
	})
	if err != nil {
		return nil, err
	}

	return &DNSRecord{
		Name:    record.Name,
		Type:    record.Type,
		TTL:     int(record.Ttl),
		Records: record.Records,
	}, err
}

func (c *Client) DeleteDNSRecordByID(id string) error {
	name, dtype, err := parseDNSID(id)
	if err != nil {
		return err
	}
	return c.DeleteDNSRecord(name, dtype)
}

func (c *Client) DeleteDNSRecord(name string, recordType string) error {
	_, err := c.DNSClient.DeleteDNSRecord(context.Background(), &pb.DNSRecordIdentifier{
		Name: name,
		Type: recordType,
	})
	return err
}
