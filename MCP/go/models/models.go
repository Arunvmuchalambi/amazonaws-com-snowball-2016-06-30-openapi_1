package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// UpdateLongTermPricingResult represents the UpdateLongTermPricingResult schema from the OpenAPI specification
type UpdateLongTermPricingResult struct {
}

// GetSoftwareUpdatesResult represents the GetSoftwareUpdatesResult schema from the OpenAPI specification
type GetSoftwareUpdatesResult struct {
	Updatesuri interface{} `json:"UpdatesURI,omitempty"`
}

// CreateClusterRequest represents the CreateClusterRequest schema from the OpenAPI specification
type CreateClusterRequest struct {
	Initialclustersize interface{} `json:"InitialClusterSize,omitempty"`
	Forwardingaddressid interface{} `json:"ForwardingAddressId,omitempty"`
	Kmskeyarn interface{} `json:"KmsKeyARN,omitempty"`
	Jobtype interface{} `json:"JobType"`
	Notification interface{} `json:"Notification,omitempty"`
	Snowballcapacitypreference interface{} `json:"SnowballCapacityPreference,omitempty"`
	Addressid interface{} `json:"AddressId"`
	Forcecreatejobs interface{} `json:"ForceCreateJobs,omitempty"`
	Longtermpricingids interface{} `json:"LongTermPricingIds,omitempty"`
	Snowballtype interface{} `json:"SnowballType"`
	Taxdocuments interface{} `json:"TaxDocuments,omitempty"`
	Resources interface{} `json:"Resources,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Shippingoption interface{} `json:"ShippingOption"`
	Ondeviceserviceconfiguration interface{} `json:"OnDeviceServiceConfiguration,omitempty"`
	Remotemanagement interface{} `json:"RemoteManagement,omitempty"`
	Rolearn interface{} `json:"RoleARN,omitempty"`
}

// DescribeReturnShippingLabelResult represents the DescribeReturnShippingLabelResult schema from the OpenAPI specification
type DescribeReturnShippingLabelResult struct {
	Status interface{} `json:"Status,omitempty"`
	Expirationdate interface{} `json:"ExpirationDate,omitempty"`
	Returnshippinglabeluri interface{} `json:"ReturnShippingLabelURI,omitempty"`
}

// ListLongTermPricingResult represents the ListLongTermPricingResult schema from the OpenAPI specification
type ListLongTermPricingResult struct {
	Longtermpricingentries interface{} `json:"LongTermPricingEntries,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CancelJobResult represents the CancelJobResult schema from the OpenAPI specification
type CancelJobResult struct {
}

// DescribeAddressesResult represents the DescribeAddressesResult schema from the OpenAPI specification
type DescribeAddressesResult struct {
	Addresses interface{} `json:"Addresses,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CreateLongTermPricingResult represents the CreateLongTermPricingResult schema from the OpenAPI specification
type CreateLongTermPricingResult struct {
	Longtermpricingid interface{} `json:"LongTermPricingId,omitempty"`
}

// UpdateJobResult represents the UpdateJobResult schema from the OpenAPI specification
type UpdateJobResult struct {
}

// DeviceConfiguration represents the DeviceConfiguration schema from the OpenAPI specification
type DeviceConfiguration struct {
	Snowconedeviceconfiguration interface{} `json:"SnowconeDeviceConfiguration,omitempty"`
}

// DescribeClusterResult represents the DescribeClusterResult schema from the OpenAPI specification
type DescribeClusterResult struct {
	Clustermetadata interface{} `json:"ClusterMetadata,omitempty"`
}

// WirelessConnection represents the WirelessConnection schema from the OpenAPI specification
type WirelessConnection struct {
	Iswifienabled interface{} `json:"IsWifiEnabled,omitempty"`
}

// CreateReturnShippingLabelRequest represents the CreateReturnShippingLabelRequest schema from the OpenAPI specification
type CreateReturnShippingLabelRequest struct {
	Jobid interface{} `json:"JobId"`
	Shippingoption interface{} `json:"ShippingOption,omitempty"`
}

// Address represents the Address schema from the OpenAPI specification
type Address struct {
	Street2 interface{} `json:"Street2,omitempty"`
	Phonenumber interface{} `json:"PhoneNumber,omitempty"`
	Prefectureordistrict interface{} `json:"PrefectureOrDistrict,omitempty"`
	Isrestricted interface{} `json:"IsRestricted,omitempty"`
	Postalcode interface{} `json:"PostalCode,omitempty"`
	Stateorprovince interface{} `json:"StateOrProvince,omitempty"`
	Street1 interface{} `json:"Street1,omitempty"`
	Company interface{} `json:"Company,omitempty"`
	Country interface{} `json:"Country,omitempty"`
	Landmark interface{} `json:"Landmark,omitempty"`
	Street3 interface{} `json:"Street3,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	City interface{} `json:"City,omitempty"`
	Addressid interface{} `json:"AddressId,omitempty"`
}

// JobResource represents the JobResource schema from the OpenAPI specification
type JobResource struct {
	Ec2amiresources interface{} `json:"Ec2AmiResources,omitempty"`
	Lambdaresources interface{} `json:"LambdaResources,omitempty"`
	S3resources interface{} `json:"S3Resources,omitempty"`
}

// DescribeClusterRequest represents the DescribeClusterRequest schema from the OpenAPI specification
type DescribeClusterRequest struct {
	Clusterid interface{} `json:"ClusterId"`
}

// ListJobsResult represents the ListJobsResult schema from the OpenAPI specification
type ListJobsResult struct {
	Joblistentries interface{} `json:"JobListEntries,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribeAddressesRequest represents the DescribeAddressesRequest schema from the OpenAPI specification
type DescribeAddressesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListPickupLocationsResult represents the ListPickupLocationsResult schema from the OpenAPI specification
type ListPickupLocationsResult struct {
	Addresses interface{} `json:"Addresses,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// INDTaxDocuments represents the INDTaxDocuments schema from the OpenAPI specification
type INDTaxDocuments struct {
	Gstin interface{} `json:"GSTIN,omitempty"`
}

// GetSoftwareUpdatesRequest represents the GetSoftwareUpdatesRequest schema from the OpenAPI specification
type GetSoftwareUpdatesRequest struct {
	Jobid interface{} `json:"JobId"`
}

// SnowconeDeviceConfiguration represents the SnowconeDeviceConfiguration schema from the OpenAPI specification
type SnowconeDeviceConfiguration struct {
	Wirelessconnection interface{} `json:"WirelessConnection,omitempty"`
}

// GetJobManifestResult represents the GetJobManifestResult schema from the OpenAPI specification
type GetJobManifestResult struct {
	Manifesturi interface{} `json:"ManifestURI,omitempty"`
}

// CancelClusterResult represents the CancelClusterResult schema from the OpenAPI specification
type CancelClusterResult struct {
}

// CreateAddressRequest represents the CreateAddressRequest schema from the OpenAPI specification
type CreateAddressRequest struct {
	Address interface{} `json:"Address"`
}

// CreateLongTermPricingRequest represents the CreateLongTermPricingRequest schema from the OpenAPI specification
type CreateLongTermPricingRequest struct {
	Islongtermpricingautorenew interface{} `json:"IsLongTermPricingAutoRenew,omitempty"`
	Longtermpricingtype interface{} `json:"LongTermPricingType"`
	Snowballtype interface{} `json:"SnowballType"`
}

// ListClusterJobsResult represents the ListClusterJobsResult schema from the OpenAPI specification
type ListClusterJobsResult struct {
	Joblistentries interface{} `json:"JobListEntries,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CreateJobResult represents the CreateJobResult schema from the OpenAPI specification
type CreateJobResult struct {
	Jobid interface{} `json:"JobId,omitempty"`
}

// PickupDetails represents the PickupDetails schema from the OpenAPI specification
type PickupDetails struct {
	Email interface{} `json:"Email,omitempty"`
	Identificationexpirationdate interface{} `json:"IdentificationExpirationDate,omitempty"`
	Identificationissuingorg interface{} `json:"IdentificationIssuingOrg,omitempty"`
	Identificationnumber interface{} `json:"IdentificationNumber,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Phonenumber interface{} `json:"PhoneNumber,omitempty"`
	Devicepickupid interface{} `json:"DevicePickupId,omitempty"`
}

// EKSOnDeviceServiceConfiguration represents the EKSOnDeviceServiceConfiguration schema from the OpenAPI specification
type EKSOnDeviceServiceConfiguration struct {
	Eksanywhereversion interface{} `json:"EKSAnywhereVersion,omitempty"`
	Kubernetesversion interface{} `json:"KubernetesVersion,omitempty"`
}

// ListJobsRequest represents the ListJobsRequest schema from the OpenAPI specification
type ListJobsRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// Shipment represents the Shipment schema from the OpenAPI specification
type Shipment struct {
	Trackingnumber interface{} `json:"TrackingNumber,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// S3Resource represents the S3Resource schema from the OpenAPI specification
type S3Resource struct {
	Keyrange interface{} `json:"KeyRange,omitempty"`
	Targetondeviceservices interface{} `json:"TargetOnDeviceServices,omitempty"`
	Bucketarn interface{} `json:"BucketArn,omitempty"`
}

// EventTriggerDefinition represents the EventTriggerDefinition schema from the OpenAPI specification
type EventTriggerDefinition struct {
	Eventresourcearn interface{} `json:"EventResourceARN,omitempty"`
}

// TGWOnDeviceServiceConfiguration represents the TGWOnDeviceServiceConfiguration schema from the OpenAPI specification
type TGWOnDeviceServiceConfiguration struct {
	Storageunit interface{} `json:"StorageUnit,omitempty"`
	Storagelimit interface{} `json:"StorageLimit,omitempty"`
}

// ListCompatibleImagesRequest represents the ListCompatibleImagesRequest schema from the OpenAPI specification
type ListCompatibleImagesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CancelJobRequest represents the CancelJobRequest schema from the OpenAPI specification
type CancelJobRequest struct {
	Jobid interface{} `json:"JobId"`
}

// GetJobManifestRequest represents the GetJobManifestRequest schema from the OpenAPI specification
type GetJobManifestRequest struct {
	Jobid interface{} `json:"JobId"`
}

// JobLogs represents the JobLogs schema from the OpenAPI specification
type JobLogs struct {
	Jobfailureloguri interface{} `json:"JobFailureLogURI,omitempty"`
	Jobsuccessloguri interface{} `json:"JobSuccessLogURI,omitempty"`
	Jobcompletionreporturi interface{} `json:"JobCompletionReportURI,omitempty"`
}

// GetJobUnlockCodeRequest represents the GetJobUnlockCodeRequest schema from the OpenAPI specification
type GetJobUnlockCodeRequest struct {
	Jobid interface{} `json:"JobId"`
}

// GetJobUnlockCodeResult represents the GetJobUnlockCodeResult schema from the OpenAPI specification
type GetJobUnlockCodeResult struct {
	Unlockcode interface{} `json:"UnlockCode,omitempty"`
}

// S3OnDeviceServiceConfiguration represents the S3OnDeviceServiceConfiguration schema from the OpenAPI specification
type S3OnDeviceServiceConfiguration struct {
	Storageunit interface{} `json:"StorageUnit,omitempty"`
	Faulttolerance interface{} `json:"FaultTolerance,omitempty"`
	Servicesize interface{} `json:"ServiceSize,omitempty"`
	Storagelimit interface{} `json:"StorageLimit,omitempty"`
}

// DataTransfer represents the DataTransfer schema from the OpenAPI specification
type DataTransfer struct {
	Bytestransferred interface{} `json:"BytesTransferred,omitempty"`
	Objectstransferred interface{} `json:"ObjectsTransferred,omitempty"`
	Totalbytes interface{} `json:"TotalBytes,omitempty"`
	Totalobjects interface{} `json:"TotalObjects,omitempty"`
}

// CancelClusterRequest represents the CancelClusterRequest schema from the OpenAPI specification
type CancelClusterRequest struct {
	Clusterid interface{} `json:"ClusterId"`
}

// CreateJobRequest represents the CreateJobRequest schema from the OpenAPI specification
type CreateJobRequest struct {
	Longtermpricingid interface{} `json:"LongTermPricingId,omitempty"`
	Remotemanagement interface{} `json:"RemoteManagement,omitempty"`
	Shippingoption interface{} `json:"ShippingOption,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Addressid interface{} `json:"AddressId,omitempty"`
	Clusterid interface{} `json:"ClusterId,omitempty"`
	Deviceconfiguration interface{} `json:"DeviceConfiguration,omitempty"`
	Impactlevel interface{} `json:"ImpactLevel,omitempty"`
	Jobtype interface{} `json:"JobType,omitempty"`
	Forwardingaddressid interface{} `json:"ForwardingAddressId,omitempty"`
	Rolearn interface{} `json:"RoleARN,omitempty"`
	Kmskeyarn interface{} `json:"KmsKeyARN,omitempty"`
	Taxdocuments interface{} `json:"TaxDocuments,omitempty"`
	Notification interface{} `json:"Notification,omitempty"`
	Resources interface{} `json:"Resources,omitempty"`
	Ondeviceserviceconfiguration interface{} `json:"OnDeviceServiceConfiguration,omitempty"`
	Pickupdetails interface{} `json:"PickupDetails,omitempty"`
	Snowballcapacitypreference interface{} `json:"SnowballCapacityPreference,omitempty"`
	Snowballtype interface{} `json:"SnowballType,omitempty"`
}

// UpdateJobShipmentStateRequest represents the UpdateJobShipmentStateRequest schema from the OpenAPI specification
type UpdateJobShipmentStateRequest struct {
	Jobid interface{} `json:"JobId"`
	Shipmentstate interface{} `json:"ShipmentState"`
}

// ListClustersResult represents the ListClustersResult schema from the OpenAPI specification
type ListClustersResult struct {
	Clusterlistentries interface{} `json:"ClusterListEntries,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CreateReturnShippingLabelResult represents the CreateReturnShippingLabelResult schema from the OpenAPI specification
type CreateReturnShippingLabelResult struct {
	Status interface{} `json:"Status,omitempty"`
}

// ListServiceVersionsResult represents the ListServiceVersionsResult schema from the OpenAPI specification
type ListServiceVersionsResult struct {
	Servicename interface{} `json:"ServiceName"`
	Serviceversions interface{} `json:"ServiceVersions"`
	Dependentservices interface{} `json:"DependentServices,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// JobListEntry represents the JobListEntry schema from the OpenAPI specification
type JobListEntry struct {
	Jobstate interface{} `json:"JobState,omitempty"`
	Jobtype interface{} `json:"JobType,omitempty"`
	Snowballtype interface{} `json:"SnowballType,omitempty"`
	Creationdate interface{} `json:"CreationDate,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Ismaster interface{} `json:"IsMaster,omitempty"`
	Jobid interface{} `json:"JobId,omitempty"`
}

// UpdateJobShipmentStateResult represents the UpdateJobShipmentStateResult schema from the OpenAPI specification
type UpdateJobShipmentStateResult struct {
}

// DescribeJobResult represents the DescribeJobResult schema from the OpenAPI specification
type DescribeJobResult struct {
	Jobmetadata interface{} `json:"JobMetadata,omitempty"`
	Subjobmetadata interface{} `json:"SubJobMetadata,omitempty"`
}

// CreateAddressResult represents the CreateAddressResult schema from the OpenAPI specification
type CreateAddressResult struct {
	Addressid interface{} `json:"AddressId,omitempty"`
}

// ListClusterJobsRequest represents the ListClusterJobsRequest schema from the OpenAPI specification
type ListClusterJobsRequest struct {
	Clusterid interface{} `json:"ClusterId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// LongTermPricingListEntry represents the LongTermPricingListEntry schema from the OpenAPI specification
type LongTermPricingListEntry struct {
	Longtermpricingstatus interface{} `json:"LongTermPricingStatus,omitempty"`
	Longtermpricingtype interface{} `json:"LongTermPricingType,omitempty"`
	Jobids interface{} `json:"JobIds,omitempty"`
	Longtermpricingstartdate interface{} `json:"LongTermPricingStartDate,omitempty"`
	Islongtermpricingautorenew interface{} `json:"IsLongTermPricingAutoRenew,omitempty"`
	Longtermpricingenddate interface{} `json:"LongTermPricingEndDate,omitempty"`
	Longtermpricingid interface{} `json:"LongTermPricingId,omitempty"`
	Snowballtype interface{} `json:"SnowballType,omitempty"`
	Currentactivejob interface{} `json:"CurrentActiveJob,omitempty"`
	Replacementjob interface{} `json:"ReplacementJob,omitempty"`
}

// UpdateClusterResult represents the UpdateClusterResult schema from the OpenAPI specification
type UpdateClusterResult struct {
}

// CompatibleImage represents the CompatibleImage schema from the OpenAPI specification
type CompatibleImage struct {
	Amiid interface{} `json:"AmiId,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// DescribeJobRequest represents the DescribeJobRequest schema from the OpenAPI specification
type DescribeJobRequest struct {
	Jobid interface{} `json:"JobId"`
}

// UpdateLongTermPricingRequest represents the UpdateLongTermPricingRequest schema from the OpenAPI specification
type UpdateLongTermPricingRequest struct {
	Islongtermpricingautorenew interface{} `json:"IsLongTermPricingAutoRenew,omitempty"`
	Longtermpricingid interface{} `json:"LongTermPricingId"`
	Replacementjob interface{} `json:"ReplacementJob,omitempty"`
}

// ShippingDetails represents the ShippingDetails schema from the OpenAPI specification
type ShippingDetails struct {
	Outboundshipment interface{} `json:"OutboundShipment,omitempty"`
	Shippingoption interface{} `json:"ShippingOption,omitempty"`
	Inboundshipment interface{} `json:"InboundShipment,omitempty"`
}

// ListPickupLocationsRequest represents the ListPickupLocationsRequest schema from the OpenAPI specification
type ListPickupLocationsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// NFSOnDeviceServiceConfiguration represents the NFSOnDeviceServiceConfiguration schema from the OpenAPI specification
type NFSOnDeviceServiceConfiguration struct {
	Storagelimit interface{} `json:"StorageLimit,omitempty"`
	Storageunit interface{} `json:"StorageUnit,omitempty"`
}

// ListClustersRequest represents the ListClustersRequest schema from the OpenAPI specification
type ListClustersRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// TaxDocuments represents the TaxDocuments schema from the OpenAPI specification
type TaxDocuments struct {
	Ind INDTaxDocuments `json:"IND,omitempty"` // The tax documents required in Amazon Web Services Region in India.
}

// ServiceVersion represents the ServiceVersion schema from the OpenAPI specification
type ServiceVersion struct {
	Version interface{} `json:"Version,omitempty"`
}

// ListServiceVersionsRequest represents the ListServiceVersionsRequest schema from the OpenAPI specification
type ListServiceVersionsRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Servicename interface{} `json:"ServiceName"`
	Dependentservices interface{} `json:"DependentServices,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// TargetOnDeviceService represents the TargetOnDeviceService schema from the OpenAPI specification
type TargetOnDeviceService struct {
	Servicename interface{} `json:"ServiceName,omitempty"`
	Transferoption interface{} `json:"TransferOption,omitempty"`
}

// DescribeReturnShippingLabelRequest represents the DescribeReturnShippingLabelRequest schema from the OpenAPI specification
type DescribeReturnShippingLabelRequest struct {
	Jobid interface{} `json:"JobId"`
}

// CreateClusterResult represents the CreateClusterResult schema from the OpenAPI specification
type CreateClusterResult struct {
	Clusterid interface{} `json:"ClusterId,omitempty"`
	Joblistentries interface{} `json:"JobListEntries,omitempty"`
}

// OnDeviceServiceConfiguration represents the OnDeviceServiceConfiguration schema from the OpenAPI specification
type OnDeviceServiceConfiguration struct {
	Eksondeviceservice interface{} `json:"EKSOnDeviceService,omitempty"`
	Nfsondeviceservice interface{} `json:"NFSOnDeviceService,omitempty"`
	S3ondeviceservice interface{} `json:"S3OnDeviceService,omitempty"`
	Tgwondeviceservice interface{} `json:"TGWOnDeviceService,omitempty"`
}

// GetSnowballUsageResult represents the GetSnowballUsageResult schema from the OpenAPI specification
type GetSnowballUsageResult struct {
	Snowballlimit interface{} `json:"SnowballLimit,omitempty"`
	Snowballsinuse interface{} `json:"SnowballsInUse,omitempty"`
}

// KeyRange represents the KeyRange schema from the OpenAPI specification
type KeyRange struct {
	Beginmarker interface{} `json:"BeginMarker,omitempty"`
	Endmarker interface{} `json:"EndMarker,omitempty"`
}

// DependentService represents the DependentService schema from the OpenAPI specification
type DependentService struct {
	Servicename interface{} `json:"ServiceName,omitempty"`
	Serviceversion interface{} `json:"ServiceVersion,omitempty"`
}

// LambdaResource represents the LambdaResource schema from the OpenAPI specification
type LambdaResource struct {
	Eventtriggers interface{} `json:"EventTriggers,omitempty"`
	Lambdaarn interface{} `json:"LambdaArn,omitempty"`
}

// UpdateClusterRequest represents the UpdateClusterRequest schema from the OpenAPI specification
type UpdateClusterRequest struct {
	Clusterid interface{} `json:"ClusterId"`
	Notification interface{} `json:"Notification,omitempty"`
	Shippingoption interface{} `json:"ShippingOption,omitempty"`
	Forwardingaddressid interface{} `json:"ForwardingAddressId,omitempty"`
	Resources interface{} `json:"Resources,omitempty"`
	Rolearn interface{} `json:"RoleARN,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Addressid interface{} `json:"AddressId,omitempty"`
	Ondeviceserviceconfiguration interface{} `json:"OnDeviceServiceConfiguration,omitempty"`
}

// DescribeAddressResult represents the DescribeAddressResult schema from the OpenAPI specification
type DescribeAddressResult struct {
	Address interface{} `json:"Address,omitempty"`
}

// ListCompatibleImagesResult represents the ListCompatibleImagesResult schema from the OpenAPI specification
type ListCompatibleImagesResult struct {
	Compatibleimages interface{} `json:"CompatibleImages,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ClusterMetadata represents the ClusterMetadata schema from the OpenAPI specification
type ClusterMetadata struct {
	Clusterid interface{} `json:"ClusterId,omitempty"`
	Ondeviceserviceconfiguration interface{} `json:"OnDeviceServiceConfiguration,omitempty"`
	Jobtype interface{} `json:"JobType,omitempty"`
	Rolearn interface{} `json:"RoleARN,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Shippingoption interface{} `json:"ShippingOption,omitempty"`
	Kmskeyarn interface{} `json:"KmsKeyARN,omitempty"`
	Snowballtype interface{} `json:"SnowballType,omitempty"`
	Taxdocuments interface{} `json:"TaxDocuments,omitempty"`
	Resources interface{} `json:"Resources,omitempty"`
	Creationdate interface{} `json:"CreationDate,omitempty"`
	Forwardingaddressid interface{} `json:"ForwardingAddressId,omitempty"`
	Notification interface{} `json:"Notification,omitempty"`
	Addressid interface{} `json:"AddressId,omitempty"`
	Clusterstate interface{} `json:"ClusterState,omitempty"`
}

// ClusterListEntry represents the ClusterListEntry schema from the OpenAPI specification
type ClusterListEntry struct {
	Creationdate interface{} `json:"CreationDate,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Clusterid interface{} `json:"ClusterId,omitempty"`
	Clusterstate interface{} `json:"ClusterState,omitempty"`
}

// UpdateJobRequest represents the UpdateJobRequest schema from the OpenAPI specification
type UpdateJobRequest struct {
	Snowballcapacitypreference interface{} `json:"SnowballCapacityPreference,omitempty"`
	Forwardingaddressid interface{} `json:"ForwardingAddressId,omitempty"`
	Shippingoption interface{} `json:"ShippingOption,omitempty"`
	Resources interface{} `json:"Resources,omitempty"`
	Addressid interface{} `json:"AddressId,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Jobid interface{} `json:"JobId"`
	Notification interface{} `json:"Notification,omitempty"`
	Pickupdetails PickupDetails `json:"PickupDetails,omitempty"` // Information identifying the person picking up the device.
	Ondeviceserviceconfiguration interface{} `json:"OnDeviceServiceConfiguration,omitempty"`
	Rolearn interface{} `json:"RoleARN,omitempty"`
}

// GetSnowballUsageRequest represents the GetSnowballUsageRequest schema from the OpenAPI specification
type GetSnowballUsageRequest struct {
}

// JobMetadata represents the JobMetadata schema from the OpenAPI specification
type JobMetadata struct {
	Pickupdetails interface{} `json:"PickupDetails,omitempty"`
	Datatransferprogress interface{} `json:"DataTransferProgress,omitempty"`
	Clusterid interface{} `json:"ClusterId,omitempty"`
	Jobid interface{} `json:"JobId,omitempty"`
	Longtermpricingid interface{} `json:"LongTermPricingId,omitempty"`
	Addressid interface{} `json:"AddressId,omitempty"`
	Jobloginfo interface{} `json:"JobLogInfo,omitempty"`
	Jobtype interface{} `json:"JobType,omitempty"`
	Snowballtype interface{} `json:"SnowballType,omitempty"`
	Snowballid interface{} `json:"SnowballId,omitempty"`
	Ondeviceserviceconfiguration interface{} `json:"OnDeviceServiceConfiguration,omitempty"`
	Taxdocuments interface{} `json:"TaxDocuments,omitempty"`
	Kmskeyarn interface{} `json:"KmsKeyARN,omitempty"`
	Creationdate interface{} `json:"CreationDate,omitempty"`
	Impactlevel interface{} `json:"ImpactLevel,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Notification interface{} `json:"Notification,omitempty"`
	Remotemanagement interface{} `json:"RemoteManagement,omitempty"`
	Resources interface{} `json:"Resources,omitempty"`
	Shippingdetails interface{} `json:"ShippingDetails,omitempty"`
	Deviceconfiguration DeviceConfiguration `json:"DeviceConfiguration,omitempty"` // The container for <code>SnowconeDeviceConfiguration</code>.
	Jobstate interface{} `json:"JobState,omitempty"`
	Rolearn interface{} `json:"RoleARN,omitempty"`
	Snowballcapacitypreference interface{} `json:"SnowballCapacityPreference,omitempty"`
	Forwardingaddressid interface{} `json:"ForwardingAddressId,omitempty"`
}

// DescribeAddressRequest represents the DescribeAddressRequest schema from the OpenAPI specification
type DescribeAddressRequest struct {
	Addressid interface{} `json:"AddressId"`
}

// ListLongTermPricingRequest represents the ListLongTermPricingRequest schema from the OpenAPI specification
type ListLongTermPricingRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// Ec2AmiResource represents the Ec2AmiResource schema from the OpenAPI specification
type Ec2AmiResource struct {
	Amiid interface{} `json:"AmiId"`
	Snowballamiid interface{} `json:"SnowballAmiId,omitempty"`
}

// Notification represents the Notification schema from the OpenAPI specification
type Notification struct {
	Notifyall interface{} `json:"NotifyAll,omitempty"`
	Snstopicarn interface{} `json:"SnsTopicARN,omitempty"`
	Devicepickupsnstopicarn interface{} `json:"DevicePickupSnsTopicARN,omitempty"`
	Jobstatestonotify interface{} `json:"JobStatesToNotify,omitempty"`
}
