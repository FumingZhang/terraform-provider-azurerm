package cdn

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// AFDOriginGroupsClient is the cdn Management Client
type AFDOriginGroupsClient struct {
	BaseClient
}

// NewAFDOriginGroupsClient creates an instance of the AFDOriginGroupsClient client.
func NewAFDOriginGroupsClient(subscriptionID string) AFDOriginGroupsClient {
	return NewAFDOriginGroupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAFDOriginGroupsClientWithBaseURI creates an instance of the AFDOriginGroupsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewAFDOriginGroupsClientWithBaseURI(baseURI string, subscriptionID string) AFDOriginGroupsClient {
	return AFDOriginGroupsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a new origin group within the specified profile.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// profileName - name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique
// within the resource group.
// originGroupName - name of the origin group which is unique within the endpoint.
// originGroup - origin group properties
func (client AFDOriginGroupsClient) Create(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originGroup AFDOriginGroup) (result AFDOriginGroupsCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AFDOriginGroupsClient.Create")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: originGroup,
			Constraints: []validation.Constraint{{Target: "originGroup.AFDOriginGroupProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "originGroup.AFDOriginGroupProperties.HealthProbeSettings", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "originGroup.AFDOriginGroupProperties.HealthProbeSettings.ProbeIntervalInSeconds", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "originGroup.AFDOriginGroupProperties.HealthProbeSettings.ProbeIntervalInSeconds", Name: validation.InclusiveMaximum, Rule: int64(255), Chain: nil},
							{Target: "originGroup.AFDOriginGroupProperties.HealthProbeSettings.ProbeIntervalInSeconds", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
						}},
					}},
					{Target: "originGroup.AFDOriginGroupProperties.TrafficRestorationTimeToHealedOrNewEndpointsInMinutes", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "originGroup.AFDOriginGroupProperties.TrafficRestorationTimeToHealedOrNewEndpointsInMinutes", Name: validation.InclusiveMaximum, Rule: int64(50), Chain: nil},
							{Target: "originGroup.AFDOriginGroupProperties.TrafficRestorationTimeToHealedOrNewEndpointsInMinutes", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("cdn.AFDOriginGroupsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, profileName, originGroupName, originGroup)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.AFDOriginGroupsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.AFDOriginGroupsClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client AFDOriginGroupsClient) CreatePreparer(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originGroup AFDOriginGroup) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"originGroupName":   autorest.Encode("path", originGroupName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/originGroups/{originGroupName}", pathParameters),
		autorest.WithJSON(originGroup),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client AFDOriginGroupsClient) CreateSender(req *http.Request) (future AFDOriginGroupsCreateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client AFDOriginGroupsClient) CreateResponder(resp *http.Response) (result AFDOriginGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an existing origin group within a profile.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// profileName - name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique
// within the resource group.
// originGroupName - name of the origin group which is unique within the profile.
func (client AFDOriginGroupsClient) Delete(ctx context.Context, resourceGroupName string, profileName string, originGroupName string) (result AFDOriginGroupsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AFDOriginGroupsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cdn.AFDOriginGroupsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, profileName, originGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.AFDOriginGroupsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.AFDOriginGroupsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AFDOriginGroupsClient) DeletePreparer(ctx context.Context, resourceGroupName string, profileName string, originGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"originGroupName":   autorest.Encode("path", originGroupName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/originGroups/{originGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AFDOriginGroupsClient) DeleteSender(req *http.Request) (future AFDOriginGroupsDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AFDOriginGroupsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets an existing origin group within a profile.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// profileName - name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique
// within the resource group.
// originGroupName - name of the origin group which is unique within the endpoint.
func (client AFDOriginGroupsClient) Get(ctx context.Context, resourceGroupName string, profileName string, originGroupName string) (result AFDOriginGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AFDOriginGroupsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cdn.AFDOriginGroupsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, profileName, originGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.AFDOriginGroupsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.AFDOriginGroupsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.AFDOriginGroupsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client AFDOriginGroupsClient) GetPreparer(ctx context.Context, resourceGroupName string, profileName string, originGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"originGroupName":   autorest.Encode("path", originGroupName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/originGroups/{originGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AFDOriginGroupsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AFDOriginGroupsClient) GetResponder(resp *http.Response) (result AFDOriginGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByProfile lists all of the existing origin groups within a profile.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// profileName - name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique
// within the resource group.
func (client AFDOriginGroupsClient) ListByProfile(ctx context.Context, resourceGroupName string, profileName string) (result AFDOriginGroupListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AFDOriginGroupsClient.ListByProfile")
		defer func() {
			sc := -1
			if result.aoglr.Response.Response != nil {
				sc = result.aoglr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cdn.AFDOriginGroupsClient", "ListByProfile", err.Error())
	}

	result.fn = client.listByProfileNextResults
	req, err := client.ListByProfilePreparer(ctx, resourceGroupName, profileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.AFDOriginGroupsClient", "ListByProfile", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByProfileSender(req)
	if err != nil {
		result.aoglr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.AFDOriginGroupsClient", "ListByProfile", resp, "Failure sending request")
		return
	}

	result.aoglr, err = client.ListByProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.AFDOriginGroupsClient", "ListByProfile", resp, "Failure responding to request")
		return
	}
	if result.aoglr.hasNextLink() && result.aoglr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByProfilePreparer prepares the ListByProfile request.
func (client AFDOriginGroupsClient) ListByProfilePreparer(ctx context.Context, resourceGroupName string, profileName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/originGroups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByProfileSender sends the ListByProfile request. The method will close the
// http.Response Body if it receives an error.
func (client AFDOriginGroupsClient) ListByProfileSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByProfileResponder handles the response to the ListByProfile request. The method always
// closes the http.Response Body.
func (client AFDOriginGroupsClient) ListByProfileResponder(resp *http.Response) (result AFDOriginGroupListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByProfileNextResults retrieves the next set of results, if any.
func (client AFDOriginGroupsClient) listByProfileNextResults(ctx context.Context, lastResults AFDOriginGroupListResult) (result AFDOriginGroupListResult, err error) {
	req, err := lastResults.aFDOriginGroupListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "cdn.AFDOriginGroupsClient", "listByProfileNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByProfileSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "cdn.AFDOriginGroupsClient", "listByProfileNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.AFDOriginGroupsClient", "listByProfileNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByProfileComplete enumerates all values, automatically crossing page boundaries as required.
func (client AFDOriginGroupsClient) ListByProfileComplete(ctx context.Context, resourceGroupName string, profileName string) (result AFDOriginGroupListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AFDOriginGroupsClient.ListByProfile")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByProfile(ctx, resourceGroupName, profileName)
	return
}

// ListResourceUsage checks the quota and actual usage of endpoints under the given CDN profile.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// profileName - name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique
// within the resource group.
// originGroupName - name of the origin group which is unique within the endpoint.
func (client AFDOriginGroupsClient) ListResourceUsage(ctx context.Context, resourceGroupName string, profileName string, originGroupName string) (result UsagesListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AFDOriginGroupsClient.ListResourceUsage")
		defer func() {
			sc := -1
			if result.ulr.Response.Response != nil {
				sc = result.ulr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cdn.AFDOriginGroupsClient", "ListResourceUsage", err.Error())
	}

	result.fn = client.listResourceUsageNextResults
	req, err := client.ListResourceUsagePreparer(ctx, resourceGroupName, profileName, originGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.AFDOriginGroupsClient", "ListResourceUsage", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListResourceUsageSender(req)
	if err != nil {
		result.ulr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.AFDOriginGroupsClient", "ListResourceUsage", resp, "Failure sending request")
		return
	}

	result.ulr, err = client.ListResourceUsageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.AFDOriginGroupsClient", "ListResourceUsage", resp, "Failure responding to request")
		return
	}
	if result.ulr.hasNextLink() && result.ulr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListResourceUsagePreparer prepares the ListResourceUsage request.
func (client AFDOriginGroupsClient) ListResourceUsagePreparer(ctx context.Context, resourceGroupName string, profileName string, originGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"originGroupName":   autorest.Encode("path", originGroupName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/originGroups/{originGroupName}/usages", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListResourceUsageSender sends the ListResourceUsage request. The method will close the
// http.Response Body if it receives an error.
func (client AFDOriginGroupsClient) ListResourceUsageSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResourceUsageResponder handles the response to the ListResourceUsage request. The method always
// closes the http.Response Body.
func (client AFDOriginGroupsClient) ListResourceUsageResponder(resp *http.Response) (result UsagesListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listResourceUsageNextResults retrieves the next set of results, if any.
func (client AFDOriginGroupsClient) listResourceUsageNextResults(ctx context.Context, lastResults UsagesListResult) (result UsagesListResult, err error) {
	req, err := lastResults.usagesListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "cdn.AFDOriginGroupsClient", "listResourceUsageNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListResourceUsageSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "cdn.AFDOriginGroupsClient", "listResourceUsageNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResourceUsageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.AFDOriginGroupsClient", "listResourceUsageNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListResourceUsageComplete enumerates all values, automatically crossing page boundaries as required.
func (client AFDOriginGroupsClient) ListResourceUsageComplete(ctx context.Context, resourceGroupName string, profileName string, originGroupName string) (result UsagesListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AFDOriginGroupsClient.ListResourceUsage")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListResourceUsage(ctx, resourceGroupName, profileName, originGroupName)
	return
}

// Update updates an existing origin group within a profile.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// profileName - name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique
// within the resource group.
// originGroupName - name of the origin group which is unique within the profile.
// originGroupUpdateProperties - origin group properties
func (client AFDOriginGroupsClient) Update(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originGroupUpdateProperties AFDOriginGroupUpdateParameters) (result AFDOriginGroupsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AFDOriginGroupsClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cdn.AFDOriginGroupsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, profileName, originGroupName, originGroupUpdateProperties)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.AFDOriginGroupsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.AFDOriginGroupsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client AFDOriginGroupsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originGroupUpdateProperties AFDOriginGroupUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"originGroupName":   autorest.Encode("path", originGroupName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/originGroups/{originGroupName}", pathParameters),
		autorest.WithJSON(originGroupUpdateProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client AFDOriginGroupsClient) UpdateSender(req *http.Request) (future AFDOriginGroupsUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client AFDOriginGroupsClient) UpdateResponder(resp *http.Response) (result AFDOriginGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}