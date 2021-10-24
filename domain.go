package restvirt

import (
	"context"

	"github.com/verbit/restvirt-client/pb"
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
	d, err := c.DomainServiceClient.CreateDomain(context.Background(), &pb.CreateDomainRequest{Domain: &pb.Domain{
		Uuid:      domain.UUID,
		Name:      domain.Name,
		Vcpu:      uint32(domain.VCPU),
		Memory:    uint64(domain.MemoryMiB),
		PrivateIp: domain.PrivateIP,
		UserData:  domain.UserData,
	}})
	if err != nil {
		return "", err
	}
	return d.Uuid, err
}

func (c *Client) GetDomain(id string) (*Domain, error) {
	domain, err := c.DomainServiceClient.GetDomain(context.Background(), &pb.GetDomainRequest{Uuid: id})
	if err != nil {
		return nil, err
	}
	return &Domain{
		UUID:      domain.Uuid,
		Name:      domain.Name,
		VCPU:      int(domain.Vcpu),
		MemoryMiB: int(domain.Memory),
		PrivateIP: domain.PrivateIp,
		UserData:  domain.UserData,
	}, nil
}

func (c *Client) DeleteDomain(id string) error {
	_, err := c.DomainServiceClient.DeleteDomain(context.Background(), &pb.DeleteDomainRequest{Uuid: id})
	return err
}
