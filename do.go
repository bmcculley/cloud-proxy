package main

import (
	"context"
	"fmt"
	"time"

	"github.com/digitalocean/godo"
	"github.com/jmcvetta/randutil"
)

func doRegions(client *godo.Client) ([]string, error) {
	var slugs []string
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	regions, _, err := client.Regions.List(ctx, &godo.ListOptions{})
	if err != nil {
		return slugs, err
	}
	for _, r := range regions {
		slugs = append(slugs, r.Slug)
	}
	return slugs, nil
}

func newDropLetMultiCreateRequest(prefix, region, keyID string, count int) *godo.DropletMultiCreateRequest {
	var names []string
	for i := 0; i < count; i++ {
		name, _ := randutil.AlphaString(8)
		names = append(names, fmt.Sprintf("%s-%s", prefix, name))
	}
	return &godo.DropletMultiCreateRequest{
		Names:  names,
		Region: region,
		Size:   "512mb",
		Image: godo.DropletCreateImage{
			Slug: "ubuntu-14-04-x64",
		},
		SSHKeys: []godo.DropletCreateSSHKey{
			godo.DropletCreateSSHKey{
				Fingerprint: keyID,
			},
		},
		Backups:           false,
		IPv6:              false,
		PrivateNetworking: false,
	}
}
