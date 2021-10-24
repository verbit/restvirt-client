package restvirt

import (
	"context"

	"github.com/verbit/restvirt-client/pb"
)

type VolumeAttachment struct {
	DiskAddress string `json:"disk_address"`
}

func (c *Client) CreateAttachment(domainID string, volumeID string) (*VolumeAttachment, error) {
	attachment, err := c.VolumeServiceClient.AttachVolume(context.Background(), &pb.VolumeAttachmentIdentifier{
		DomainId: domainID,
		VolumeId: volumeID,
	})
	if err != nil {
		return nil, err
	}
	return &VolumeAttachment{DiskAddress: attachment.DiskAddress}, nil

}

func (c *Client) GetAttachment(domainID string, volumeID string) (*VolumeAttachment, error) {
	attachment, err := c.VolumeServiceClient.GetVolumeAttachment(context.Background(), &pb.VolumeAttachmentIdentifier{
		DomainId: domainID,
		VolumeId: volumeID,
	})
	if err != nil {
		return nil, err
	}
	return &VolumeAttachment{DiskAddress: attachment.DiskAddress}, nil
}

func (c *Client) DeleteAttachment(domainID string, volumeID string) error {
	_, err := c.VolumeServiceClient.DetachVolume(context.Background(), &pb.VolumeAttachmentIdentifier{
		DomainId: domainID,
		VolumeId: volumeID,
	})
	return err
}
