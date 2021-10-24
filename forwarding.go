package restvirt

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/verbit/restvirt-client/pb"
)

type PortForwarding struct {
	SourcePort uint16 `json:"source_port"`
	TargetPort uint16 `json:"target_port"`
	TargetIP   string `json:"target_ip"`
	Protocol   string `json:"protocol"`
}

func parsePortForwardingID(id string) (uint16, string, error) {
	s := strings.Split(id, "-")
	if len(s) != 2 {
		return 0, "", fmt.Errorf("ID looks fishy")
	}
	port, err := strconv.ParseUint(s[0], 10, 16)
	if err != nil {
		return 0, "", err
	}
	return uint16(port), s[1], nil
}

func (c *Client) CreatePortForwarding(forwarding PortForwarding) (string, error) {
	f, err := c.PortForwardingServiceClient.PutPortForwarding(context.Background(), &pb.PutPortForwardingRequest{PortForwarding: &pb.PortForwarding{
		Protocol:   forwarding.Protocol,
		SourcePort: uint32(forwarding.SourcePort),
		TargetIp:   forwarding.TargetIP,
		TargetPort: uint32(forwarding.TargetPort),
	}})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d-%s", f.SourcePort, f.Protocol), nil
}

func (c *Client) GetPortForwarding(id string) (*PortForwarding, error) {
	sourcePort, protocol, err := parsePortForwardingID(id)
	if err != nil {
		return nil, err
	}
	f, err := c.PortForwardingServiceClient.GetPortForwarding(context.Background(), &pb.PortForwardingIdentifier{
		Protocol:   protocol,
		SourcePort: uint32(sourcePort),
	})
	if err != nil {
		return nil, err
	}
	return &PortForwarding{
		SourcePort: uint16(f.SourcePort),
		TargetPort: uint16(f.TargetPort),
		TargetIP:   f.TargetIp,
		Protocol:   f.Protocol,
	}, nil
}

func (c *Client) DeletePortForwarding(id string) error {
	sourcePort, protocol, err := parsePortForwardingID(id)
	if err != nil {
		return err
	}
	_, err = c.PortForwardingServiceClient.DeletePortForwarding(context.Background(), &pb.PortForwardingIdentifier{
		Protocol:   protocol,
		SourcePort: uint32(sourcePort),
	})
	return err
}
