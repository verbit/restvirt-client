package restvirt

import (
	"context"

	"github.com/verbit/restvirt-client/pb"
)

type Volume struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
	Size int64  `json:"size"`
}

func (c *Client) CreateVolume(volume Volume) (string, error) {
	v, err := c.VolumeServiceClient.CreateVolume(context.Background(), &pb.CreateVolumeRequest{Volume: &pb.Volume{
		Name: volume.Name,
		Size: uint64(volume.Size),
	}})
	if err != nil {
		return "", err
	}
	return v.Id, err
}

func (c *Client) GetVolume(id string) (*Volume, error) {
	volume, err := c.VolumeServiceClient.GetVolume(context.Background(), &pb.GetVolumeRequest{Uuid: id})
	if err != nil {
		return nil, err
	}
	return &Volume{
		ID:   volume.Id,
		Name: volume.Name,
		Size: int64(volume.Size),
	}, nil
}

func (c *Client) DeleteVolume(id string) error {
	_, err := c.VolumeServiceClient.DeleteVolume(context.Background(), &pb.DeleteVolumeRequest{Uuid: id})
	return err
}
