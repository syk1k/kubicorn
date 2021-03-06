package devices

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// AccessRights enumerates the values for access rights.
type AccessRights string

const (
	// DeviceConnect specifies the device connect state for access rights.
	DeviceConnect AccessRights = "DeviceConnect"
	// RegistryRead specifies the registry read state for access rights.
	RegistryRead AccessRights = "RegistryRead"
	// RegistryReadDeviceConnect specifies the registry read device connect state for access rights.
	RegistryReadDeviceConnect AccessRights = "RegistryRead, DeviceConnect"
	// RegistryReadRegistryWrite specifies the registry read registry write state for access rights.
	RegistryReadRegistryWrite AccessRights = "RegistryRead, RegistryWrite"
	// RegistryReadRegistryWriteDeviceConnect specifies the registry read registry write device connect state for access
	// rights.
	RegistryReadRegistryWriteDeviceConnect AccessRights = "RegistryRead, RegistryWrite, DeviceConnect"
	// RegistryReadRegistryWriteServiceConnect specifies the registry read registry write service connect state for access
	// rights.
	RegistryReadRegistryWriteServiceConnect AccessRights = "RegistryRead, RegistryWrite, ServiceConnect"
	// RegistryReadRegistryWriteServiceConnectDeviceConnect specifies the registry read registry write service connect
	// device connect state for access rights.
	RegistryReadRegistryWriteServiceConnectDeviceConnect AccessRights = "RegistryRead, RegistryWrite, ServiceConnect, DeviceConnect"
	// RegistryReadServiceConnect specifies the registry read service connect state for access rights.
	RegistryReadServiceConnect AccessRights = "RegistryRead, ServiceConnect"
	// RegistryReadServiceConnectDeviceConnect specifies the registry read service connect device connect state for access
	// rights.
	RegistryReadServiceConnectDeviceConnect AccessRights = "RegistryRead, ServiceConnect, DeviceConnect"
	// RegistryWrite specifies the registry write state for access rights.
	RegistryWrite AccessRights = "RegistryWrite"
	// RegistryWriteDeviceConnect specifies the registry write device connect state for access rights.
	RegistryWriteDeviceConnect AccessRights = "RegistryWrite, DeviceConnect"
	// RegistryWriteServiceConnect specifies the registry write service connect state for access rights.
	RegistryWriteServiceConnect AccessRights = "RegistryWrite, ServiceConnect"
	// RegistryWriteServiceConnectDeviceConnect specifies the registry write service connect device connect state for
	// access rights.
	RegistryWriteServiceConnectDeviceConnect AccessRights = "RegistryWrite, ServiceConnect, DeviceConnect"
	// ServiceConnect specifies the service connect state for access rights.
	ServiceConnect AccessRights = "ServiceConnect"
	// ServiceConnectDeviceConnect specifies the service connect device connect state for access rights.
	ServiceConnectDeviceConnect AccessRights = "ServiceConnect, DeviceConnect"
)

// Capabilities enumerates the values for capabilities.
type Capabilities string

const (
	// DeviceManagement specifies the device management state for capabilities.
	DeviceManagement Capabilities = "DeviceManagement"
	// None specifies the none state for capabilities.
	None Capabilities = "None"
)

// IotHubNameUnavailabilityReason enumerates the values for iot hub name unavailability reason.
type IotHubNameUnavailabilityReason string

const (
	// AlreadyExists specifies the already exists state for iot hub name unavailability reason.
	AlreadyExists IotHubNameUnavailabilityReason = "AlreadyExists"
	// Invalid specifies the invalid state for iot hub name unavailability reason.
	Invalid IotHubNameUnavailabilityReason = "Invalid"
)

// IotHubScaleType enumerates the values for iot hub scale type.
type IotHubScaleType string

const (
	// IotHubScaleTypeAutomatic specifies the iot hub scale type automatic state for iot hub scale type.
	IotHubScaleTypeAutomatic IotHubScaleType = "Automatic"
	// IotHubScaleTypeManual specifies the iot hub scale type manual state for iot hub scale type.
	IotHubScaleTypeManual IotHubScaleType = "Manual"
	// IotHubScaleTypeNone specifies the iot hub scale type none state for iot hub scale type.
	IotHubScaleTypeNone IotHubScaleType = "None"
)

// IotHubSku enumerates the values for iot hub sku.
type IotHubSku string

const (
	// F1 specifies the f1 state for iot hub sku.
	F1 IotHubSku = "F1"
	// S1 specifies the s1 state for iot hub sku.
	S1 IotHubSku = "S1"
	// S2 specifies the s2 state for iot hub sku.
	S2 IotHubSku = "S2"
	// S3 specifies the s3 state for iot hub sku.
	S3 IotHubSku = "S3"
)

// IotHubSkuTier enumerates the values for iot hub sku tier.
type IotHubSkuTier string

const (
	// Free specifies the free state for iot hub sku tier.
	Free IotHubSkuTier = "Free"
	// Standard specifies the standard state for iot hub sku tier.
	Standard IotHubSkuTier = "Standard"
)

// IPFilterActionType enumerates the values for ip filter action type.
type IPFilterActionType string

const (
	// Accept specifies the accept state for ip filter action type.
	Accept IPFilterActionType = "Accept"
	// Reject specifies the reject state for ip filter action type.
	Reject IPFilterActionType = "Reject"
)

// JobStatus enumerates the values for job status.
type JobStatus string

const (
	// Cancelled specifies the cancelled state for job status.
	Cancelled JobStatus = "cancelled"
	// Completed specifies the completed state for job status.
	Completed JobStatus = "completed"
	// Enqueued specifies the enqueued state for job status.
	Enqueued JobStatus = "enqueued"
	// Failed specifies the failed state for job status.
	Failed JobStatus = "failed"
	// Running specifies the running state for job status.
	Running JobStatus = "running"
	// Unknown specifies the unknown state for job status.
	Unknown JobStatus = "unknown"
)

// JobType enumerates the values for job type.
type JobType string

const (
	// JobTypeBackup specifies the job type backup state for job type.
	JobTypeBackup JobType = "backup"
	// JobTypeExport specifies the job type export state for job type.
	JobTypeExport JobType = "export"
	// JobTypeFactoryResetDevice specifies the job type factory reset device state for job type.
	JobTypeFactoryResetDevice JobType = "factoryResetDevice"
	// JobTypeFirmwareUpdate specifies the job type firmware update state for job type.
	JobTypeFirmwareUpdate JobType = "firmwareUpdate"
	// JobTypeImport specifies the job type import state for job type.
	JobTypeImport JobType = "import"
	// JobTypeReadDeviceProperties specifies the job type read device properties state for job type.
	JobTypeReadDeviceProperties JobType = "readDeviceProperties"
	// JobTypeRebootDevice specifies the job type reboot device state for job type.
	JobTypeRebootDevice JobType = "rebootDevice"
	// JobTypeUnknown specifies the job type unknown state for job type.
	JobTypeUnknown JobType = "unknown"
	// JobTypeUpdateDeviceConfiguration specifies the job type update device configuration state for job type.
	JobTypeUpdateDeviceConfiguration JobType = "updateDeviceConfiguration"
	// JobTypeWriteDeviceProperties specifies the job type write device properties state for job type.
	JobTypeWriteDeviceProperties JobType = "writeDeviceProperties"
)

// OperationMonitoringLevel enumerates the values for operation monitoring level.
type OperationMonitoringLevel string

const (
	// OperationMonitoringLevelError specifies the operation monitoring level error state for operation monitoring level.
	OperationMonitoringLevelError OperationMonitoringLevel = "Error"
	// OperationMonitoringLevelErrorInformation specifies the operation monitoring level error information state for
	// operation monitoring level.
	OperationMonitoringLevelErrorInformation OperationMonitoringLevel = "Error, Information"
	// OperationMonitoringLevelInformation specifies the operation monitoring level information state for operation
	// monitoring level.
	OperationMonitoringLevelInformation OperationMonitoringLevel = "Information"
	// OperationMonitoringLevelNone specifies the operation monitoring level none state for operation monitoring level.
	OperationMonitoringLevelNone OperationMonitoringLevel = "None"
)

// CloudToDeviceProperties is the IoT hub cloud-to-device messaging properties.
type CloudToDeviceProperties struct {
	MaxDeliveryCount    *int32              `json:"maxDeliveryCount,omitempty"`
	DefaultTTLAsIso8601 *string             `json:"defaultTtlAsIso8601,omitempty"`
	Feedback            *FeedbackProperties `json:"feedback,omitempty"`
}

// ErrorDetails is error details.
type ErrorDetails struct {
	Code           *string `json:"Code,omitempty"`
	HTTPStatusCode *string `json:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty"`
	Details        *string `json:"Details,omitempty"`
}

// EventHubConsumerGroupInfo is the properties of the EventHubConsumerGroupInfo object.
type EventHubConsumerGroupInfo struct {
	autorest.Response `json:"-"`
	Tags              *map[string]*string `json:"tags,omitempty"`
	ID                *string             `json:"id,omitempty"`
	Name              *string             `json:"name,omitempty"`
}

// EventHubConsumerGroupsListResult is the JSON-serialized array of Event Hub-compatible consumer group names with a
// next link.
type EventHubConsumerGroupsListResult struct {
	autorest.Response `json:"-"`
	Value             *[]string `json:"value,omitempty"`
	NextLink          *string   `json:"nextLink,omitempty"`
}

// EventHubConsumerGroupsListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client EventHubConsumerGroupsListResult) EventHubConsumerGroupsListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// EventHubProperties is the properties of the provisioned Event Hub-compatible endpoint used by the IoT hub.
type EventHubProperties struct {
	RetentionTimeInDays *int64    `json:"retentionTimeInDays,omitempty"`
	PartitionCount      *int32    `json:"partitionCount,omitempty"`
	PartitionIds        *[]string `json:"partitionIds,omitempty"`
	Path                *string   `json:"path,omitempty"`
	Endpoint            *string   `json:"endpoint,omitempty"`
}

// ExportDevicesRequest is use to provide parameters when requesting an export of all devices in the IoT hub.
type ExportDevicesRequest struct {
	ExportBlobContainerURI *string `json:"ExportBlobContainerUri,omitempty"`
	ExcludeKeys            *bool   `json:"ExcludeKeys,omitempty"`
}

// FeedbackProperties is the properties of the feedback queue for cloud-to-device messages.
type FeedbackProperties struct {
	LockDurationAsIso8601 *string `json:"lockDurationAsIso8601,omitempty"`
	TTLAsIso8601          *string `json:"ttlAsIso8601,omitempty"`
	MaxDeliveryCount      *int32  `json:"maxDeliveryCount,omitempty"`
}

// ImportDevicesRequest is use to provide parameters when requesting an import of all devices in the hub.
type ImportDevicesRequest struct {
	InputBlobContainerURI  *string `json:"InputBlobContainerUri,omitempty"`
	OutputBlobContainerURI *string `json:"OutputBlobContainerUri,omitempty"`
}

// IotHubCapacity is ioT Hub capacity information.
type IotHubCapacity struct {
	Minimum   *int64          `json:"minimum,omitempty"`
	Maximum   *int64          `json:"maximum,omitempty"`
	Default   *int64          `json:"default,omitempty"`
	ScaleType IotHubScaleType `json:"scaleType,omitempty"`
}

// IotHubDescription is the description of the IoT hub.
type IotHubDescription struct {
	autorest.Response `json:"-"`
	ID                *string             `json:"id,omitempty"`
	Name              *string             `json:"name,omitempty"`
	Type              *string             `json:"type,omitempty"`
	Location          *string             `json:"location,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
	Subscriptionid    *string             `json:"subscriptionid,omitempty"`
	Resourcegroup     *string             `json:"resourcegroup,omitempty"`
	Etag              *string             `json:"etag,omitempty"`
	Properties        *IotHubProperties   `json:"properties,omitempty"`
	Sku               *IotHubSkuInfo      `json:"sku,omitempty"`
}

// IotHubDescriptionListResult is the JSON-serialized array of IotHubDescription objects with a next link.
type IotHubDescriptionListResult struct {
	autorest.Response `json:"-"`
	Value             *[]IotHubDescription `json:"value,omitempty"`
	NextLink          *string              `json:"nextLink,omitempty"`
}

// IotHubDescriptionListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client IotHubDescriptionListResult) IotHubDescriptionListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// IotHubNameAvailabilityInfo is the properties indicating whether a given IoT hub name is available.
type IotHubNameAvailabilityInfo struct {
	autorest.Response `json:"-"`
	NameAvailable     *bool                          `json:"nameAvailable,omitempty"`
	Reason            IotHubNameUnavailabilityReason `json:"reason,omitempty"`
	Message           *string                        `json:"message,omitempty"`
}

// IotHubProperties is the properties of an IoT hub.
type IotHubProperties struct {
	AuthorizationPolicies          *[]SharedAccessSignatureAuthorizationRule `json:"authorizationPolicies,omitempty"`
	IPFilterRules                  *[]IPFilterRule                           `json:"ipFilterRules,omitempty"`
	ProvisioningState              *string                                   `json:"provisioningState,omitempty"`
	HostName                       *string                                   `json:"hostName,omitempty"`
	EventHubEndpoints              *map[string]*EventHubProperties           `json:"eventHubEndpoints,omitempty"`
	StorageEndpoints               *map[string]*StorageEndpointProperties    `json:"storageEndpoints,omitempty"`
	MessagingEndpoints             *map[string]*MessagingEndpointProperties  `json:"messagingEndpoints,omitempty"`
	EnableFileUploadNotifications  *bool                                     `json:"enableFileUploadNotifications,omitempty"`
	CloudToDevice                  *CloudToDeviceProperties                  `json:"cloudToDevice,omitempty"`
	Comments                       *string                                   `json:"comments,omitempty"`
	OperationsMonitoringProperties *OperationsMonitoringProperties           `json:"operationsMonitoringProperties,omitempty"`
	Features                       Capabilities                              `json:"features,omitempty"`
}

// IotHubQuotaMetricInfo is quota metrics properties.
type IotHubQuotaMetricInfo struct {
	Name         *string `json:"Name,omitempty"`
	CurrentValue *int64  `json:"CurrentValue,omitempty"`
	MaxValue     *int64  `json:"MaxValue,omitempty"`
}

// IotHubQuotaMetricInfoListResult is the JSON-serialized array of IotHubQuotaMetricInfo objects with a next link.
type IotHubQuotaMetricInfoListResult struct {
	autorest.Response `json:"-"`
	Value             *[]IotHubQuotaMetricInfo `json:"value,omitempty"`
	NextLink          *string                  `json:"nextLink,omitempty"`
}

// IotHubQuotaMetricInfoListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client IotHubQuotaMetricInfoListResult) IotHubQuotaMetricInfoListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// IotHubSkuDescription is SKU properties.
type IotHubSkuDescription struct {
	ResourceType *string         `json:"resourceType,omitempty"`
	Sku          *IotHubSkuInfo  `json:"sku,omitempty"`
	Capacity     *IotHubCapacity `json:"capacity,omitempty"`
}

// IotHubSkuDescriptionListResult is the JSON-serialized array of IotHubSkuDescription objects with a next link.
type IotHubSkuDescriptionListResult struct {
	autorest.Response `json:"-"`
	Value             *[]IotHubSkuDescription `json:"value,omitempty"`
	NextLink          *string                 `json:"nextLink,omitempty"`
}

// IotHubSkuDescriptionListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client IotHubSkuDescriptionListResult) IotHubSkuDescriptionListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// IotHubSkuInfo is information about the SKU of the IoT hub.
type IotHubSkuInfo struct {
	Name     IotHubSku     `json:"name,omitempty"`
	Tier     IotHubSkuTier `json:"tier,omitempty"`
	Capacity *int64        `json:"capacity,omitempty"`
}

// IPFilterRule is the IP filter rules for the IoT hub.
type IPFilterRule struct {
	FilterName *string            `json:"filterName,omitempty"`
	Action     IPFilterActionType `json:"action,omitempty"`
	IPMask     *string            `json:"ipMask,omitempty"`
}

// JobResponse is the properties of the Job Response object.
type JobResponse struct {
	autorest.Response `json:"-"`
	JobID             *string           `json:"jobId,omitempty"`
	StartTimeUtc      *date.TimeRFC1123 `json:"startTimeUtc,omitempty"`
	EndTimeUtc        *date.TimeRFC1123 `json:"endTimeUtc,omitempty"`
	Type              JobType           `json:"type,omitempty"`
	Status            JobStatus         `json:"status,omitempty"`
	FailureReason     *string           `json:"failureReason,omitempty"`
	StatusMessage     *string           `json:"statusMessage,omitempty"`
	ParentJobID       *string           `json:"parentJobId,omitempty"`
}

// JobResponseListResult is the JSON-serialized array of JobResponse objects with a next link.
type JobResponseListResult struct {
	autorest.Response `json:"-"`
	Value             *[]JobResponse `json:"value,omitempty"`
	NextLink          *string        `json:"nextLink,omitempty"`
}

// JobResponseListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client JobResponseListResult) JobResponseListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// MessagingEndpointProperties is the properties of the messaging endpoints used by this IoT hub.
type MessagingEndpointProperties struct {
	LockDurationAsIso8601 *string `json:"lockDurationAsIso8601,omitempty"`
	TTLAsIso8601          *string `json:"ttlAsIso8601,omitempty"`
	MaxDeliveryCount      *int32  `json:"maxDeliveryCount,omitempty"`
}

// OperationInputs is input values.
type OperationInputs struct {
	Name *string `json:"Name,omitempty"`
}

// OperationsMonitoringProperties is the operations monitoring properties for the IoT hub. The possible keys to the
// dictionary are Connections, DeviceTelemetry, C2DCommands, DeviceIdentityOperations, FileUploadOperations.
type OperationsMonitoringProperties struct {
	Events *map[string]*OperationMonitoringLevel `json:"events,omitempty"`
}

// RegistryStatistics is identity registry statistics.
type RegistryStatistics struct {
	autorest.Response   `json:"-"`
	TotalDeviceCount    *int64 `json:"totalDeviceCount,omitempty"`
	EnabledDeviceCount  *int64 `json:"enabledDeviceCount,omitempty"`
	DisabledDeviceCount *int64 `json:"disabledDeviceCount,omitempty"`
}

// Resource is the common properties of an Azure resource.
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// SetObject is
type SetObject struct {
	autorest.Response `json:"-"`
	Value             *map[string]interface{} `json:"value,omitempty"`
}

// SharedAccessSignatureAuthorizationRule is the properties of an IoT hub shared access policy.
type SharedAccessSignatureAuthorizationRule struct {
	autorest.Response `json:"-"`
	KeyName           *string      `json:"keyName,omitempty"`
	PrimaryKey        *string      `json:"primaryKey,omitempty"`
	SecondaryKey      *string      `json:"secondaryKey,omitempty"`
	Rights            AccessRights `json:"rights,omitempty"`
}

// SharedAccessSignatureAuthorizationRuleListResult is the list of shared access policies with a next link.
type SharedAccessSignatureAuthorizationRuleListResult struct {
	autorest.Response `json:"-"`
	Value             *[]SharedAccessSignatureAuthorizationRule `json:"value,omitempty"`
	NextLink          *string                                   `json:"nextLink,omitempty"`
}

// SharedAccessSignatureAuthorizationRuleListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client SharedAccessSignatureAuthorizationRuleListResult) SharedAccessSignatureAuthorizationRuleListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// StorageEndpointProperties is the properties of the Azure Storage endpoint for file upload.
type StorageEndpointProperties struct {
	SasTTLAsIso8601  *string `json:"sasTtlAsIso8601,omitempty"`
	ConnectionString *string `json:"connectionString,omitempty"`
	ContainerName    *string `json:"containerName,omitempty"`
}
