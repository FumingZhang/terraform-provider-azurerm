// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2022-07-01-preview/configurationassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2022-07-01-preview/maintenanceconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2022-07-01-preview/publicmaintenanceconfigurations"
	"github.com/hashicorp/terraform-provider-azurerm/internal/common"
)

type Client struct {
	ConfigurationsClient           *maintenanceconfigurations.MaintenanceConfigurationsClient
	ConfigurationAssignmentsClient *configurationassignments.ConfigurationAssignmentsClient
	PublicConfigurationsClient     *publicmaintenanceconfigurations.PublicMaintenanceConfigurationsClient
}

func NewClient(o *common.ClientOptions) (*Client, error) {
	configurationsClient, err := maintenanceconfigurations.NewMaintenanceConfigurationsClientWithBaseURI(o.Environment.ResourceManager)
	if err != nil {
		return nil, fmt.Errorf("building Configurations Client: %+v", err)
	}
	o.Configure(configurationsClient.Client, o.Authorizers.ResourceManager)

	configurationAssignmentsClient, err := configurationassignments.NewConfigurationAssignmentsClientWithBaseURI(o.Environment.ResourceManager)
	if err != nil {
		return nil, fmt.Errorf("building Configuration Assignments Client: %+v", err)
	}
	o.Configure(configurationAssignmentsClient.Client, o.Authorizers.ResourceManager)

	publicConfigurationsClient, err := publicmaintenanceconfigurations.NewPublicMaintenanceConfigurationsClientWithBaseURI(o.Environment.ResourceManager)
	if err != nil {
		return nil, fmt.Errorf("building Public Configurations Client: %+v", err)
	}
	o.Configure(publicConfigurationsClient.Client, o.Authorizers.ResourceManager)

	return &Client{
		ConfigurationsClient:           configurationsClient,
		ConfigurationAssignmentsClient: configurationAssignmentsClient,
		PublicConfigurationsClient:     publicConfigurationsClient,
	}, nil
}
