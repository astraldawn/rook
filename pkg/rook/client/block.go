package client

import (
	"bytes"
	"encoding/json"
	"path"

	"github.com/rook/rook/pkg/model"
)

const (
	imageQueryName        = "image"
	imageMapInfoQueryName = "mapinfo"
)

func (c *RookNetworkRestClient) GetBlockImages() ([]model.BlockImage, error) {
	body, err := c.DoGet(imageQueryName)
	if err != nil {
		return nil, err
	}

	var images []model.BlockImage
	err = json.Unmarshal(body, &images)
	if err != nil {
		return nil, err
	}

	return images, nil
}

func (c *RookNetworkRestClient) CreateBlockImage(newImage model.BlockImage) (string, error) {
	body, err := json.Marshal(newImage)
	if err != nil {
		return "", err
	}

	resp, err := c.DoPost(imageQueryName, bytes.NewReader(body))
	if err != nil {
		return "", err
	}

	return string(resp), nil
}

func (c *RookNetworkRestClient) GetBlockImageMapInfo() (model.BlockImageMapInfo, error) {
	body, err := c.DoGet(path.Join(imageQueryName, imageMapInfoQueryName))
	if err != nil {
		return model.BlockImageMapInfo{}, err
	}

	var imageMapInfo model.BlockImageMapInfo
	err = json.Unmarshal(body, &imageMapInfo)
	if err != nil {
		return model.BlockImageMapInfo{}, err
	}

	return imageMapInfo, nil
}