package azureerrors

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"errors"
	"net/http"
	"testing"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// The tests in this file contain verbatim copies of errors returned from Azure
// transcribed using github.com/shurcooL/go-goon.  Consider these errors
// immutable, but feel free to add additional examples.  Rationale: it is really
// easy to introduce regressions here.

func TestHasAuthorizationFailedError(t *testing.T) {
	for _, tt := range []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "Another error",
			err:  errors.New("something happened"),
		},
		{
			name: "Authorization Failed",
			err: autorest.DetailedError{
				Original: &azure.ServiceError{
					Code:    "AuthorizationFailed",
					Message: "The client 'a0f3c32d-647d-416c-8997-fb2463b1dcd5' with object id 'a0f3c32d-647d-416c-8997-fb2463b1dcd5' does not have authorization to perform action 'Microsoft.Resources/deployments/write' over scope '/subscriptions/447cf33b-a19b-42f7-ab5e-b0b6f7be7525/resourcegroups/jmintertest/providers/Microsoft.Resources/deployments/deployment' or the scope is invalid. If access was recently granted, please refresh your credentials.",
				},
				PackageType: "resources.DeploymentsClient",
				Method:      "CreateOrUpdate",
				StatusCode:  http.StatusForbidden,
				Message:     "Failure sending request",
				// Response omitted for brevity
			},
			want: true,
		},
		{
			name: "Nested authorization failed",
			err: &azure.ServiceError{
				Code:    "DeploymentFailed",
				Message: "At least one resource deployment operation failed. Please list deployment operations for details. Please see https://aka.ms/DeployOperations for usage details.",
				Details: []map[string]interface{}{
					{
						"code":    "Forbidden",
						"message": "{\r\n  \"error\": {\r\n    \"code\": \"AuthorizationFailed\",\r\n    \"message\": \"The client 'a0f3c32d-647d-416c-8997-fb2463b1dcd5' with object id 'a0f3c32d-647d-416c-8997-fb2463b1dcd5' does not have authorization to perform action 'Microsoft.Storage/storageAccounts/write' over scope '/subscriptions/225e02bc-43d0-43d1-a01a-17e584a4ef69/resourceGroups/test' or the scope is invalid. If access was recently granted, please refresh your credentials.\"\r\n  }\r\n}",
					},
				},
			},
			want: true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := HasAuthorizationFailedError(tt.err)
			if got != tt.want {
				t.Error(got)
			}
		})
	}
}

func TestIsDeploymentActiveError(t *testing.T) {
	for _, tt := range []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "Another error",
			err:  errors.New("something happened"),
		},
		{
			name: "Deployment active",
			err: autorest.DetailedError{
				Original: azure.RequestError{
					ServiceError: &azure.ServiceError{
						Code:    "DeploymentActive",
						Message: "Unable to edit or replace deployment 'deployment': previous deployment from '4/4/2020 2:17:07 AM' is still active (expiration time is '4/11/2020 2:17:01 AM'). Please see https://aka.ms/arm-deploy for usage details.",
					},
				},
				PackageType: "resources.DeploymentsClient",
				Method:      "CreateOrUpdate",
				Message:     "Failure sending request",
			},
			want: true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := IsDeploymentActiveError(autorest.NewErrorWithError(tt.err, "", "", nil, ""))
			if got != tt.want {
				t.Error(got)
			}
		})
	}
}
