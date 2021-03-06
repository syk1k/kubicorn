// +build go1.9

// Copyright 2017 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder
// commit ID: dab57ee609fffdc578f48519d5cdc980efd8cc00

package policy

import original "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2017-06-01-preview/policy"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type DefinitionsClient = original.DefinitionsClient
type Mode = original.Mode

const (
	All          Mode = original.All
	Indexed      Mode = original.Indexed
	NotSpecified Mode = original.NotSpecified
)

type Type = original.Type

const (
	TypeBuiltIn      Type = original.TypeBuiltIn
	TypeCustom       Type = original.TypeCustom
	TypeNotSpecified Type = original.TypeNotSpecified
)

type Assignment = original.Assignment
type AssignmentListResult = original.AssignmentListResult
type AssignmentProperties = original.AssignmentProperties
type Definition = original.Definition
type DefinitionListResult = original.DefinitionListResult
type DefinitionProperties = original.DefinitionProperties
type DefinitionReference = original.DefinitionReference
type ErrorResponse = original.ErrorResponse
type SetDefinition = original.SetDefinition
type SetDefinitionListResult = original.SetDefinitionListResult
type SetDefinitionProperties = original.SetDefinitionProperties
type Sku = original.Sku
type SetDefinitionsClient = original.SetDefinitionsClient
type AssignmentsClient = original.AssignmentsClient

func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewDefinitionsClient(subscriptionID string) DefinitionsClient {
	return original.NewDefinitionsClient(subscriptionID)
}
func NewDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) DefinitionsClient {
	return original.NewDefinitionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSetDefinitionsClient(subscriptionID string) SetDefinitionsClient {
	return original.NewSetDefinitionsClient(subscriptionID)
}
func NewSetDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) SetDefinitionsClient {
	return original.NewSetDefinitionsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
func NewAssignmentsClient(subscriptionID string) AssignmentsClient {
	return original.NewAssignmentsClient(subscriptionID)
}
func NewAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) AssignmentsClient {
	return original.NewAssignmentsClientWithBaseURI(baseURI, subscriptionID)
}
