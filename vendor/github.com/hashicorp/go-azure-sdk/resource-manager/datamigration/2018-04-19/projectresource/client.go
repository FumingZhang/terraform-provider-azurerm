package projectresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProjectResourceClient struct {
	Client *resourcemanager.Client
}

func NewProjectResourceClientWithBaseURI(api environments.Api) (*ProjectResourceClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "projectresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProjectResourceClient: %+v", err)
	}

	return &ProjectResourceClient{
		Client: client,
	}, nil
}
