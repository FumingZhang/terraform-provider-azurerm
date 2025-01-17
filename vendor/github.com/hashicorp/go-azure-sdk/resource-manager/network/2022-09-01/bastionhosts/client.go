package bastionhosts

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BastionHostsClient struct {
	Client *resourcemanager.Client
}

func NewBastionHostsClientWithBaseURI(api environments.Api) (*BastionHostsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "bastionhosts", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BastionHostsClient: %+v", err)
	}

	return &BastionHostsClient{
		Client: client,
	}, nil
}
