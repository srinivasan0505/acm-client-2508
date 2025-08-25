package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// RenewCertificateRequest represents the RenewCertificateRequest schema from the OpenAPI specification
type RenewCertificateRequest struct {
	Certificatearn interface{} `json:"CertificateArn"`
}

// UpdateCertificateOptionsRequestOptions represents the UpdateCertificateOptionsRequestOptions schema from the OpenAPI specification
type UpdateCertificateOptionsRequestOptions struct {
	Certificatetransparencyloggingpreference interface{} `json:"CertificateTransparencyLoggingPreference,omitempty"`
}

// AddTagsToCertificateRequest represents the AddTagsToCertificateRequest schema from the OpenAPI specification
type AddTagsToCertificateRequest struct {
	Certificatearn interface{} `json:"CertificateArn"`
	Tags interface{} `json:"Tags"`
}

// GetAccountConfigurationResponse represents the GetAccountConfigurationResponse schema from the OpenAPI specification
type GetAccountConfigurationResponse struct {
	Expiryevents GetAccountConfigurationResponseExpiryEvents `json:"ExpiryEvents,omitempty"`
}

// KeyUsage represents the KeyUsage schema from the OpenAPI specification
type KeyUsage struct {
	Name interface{} `json:"Name,omitempty"`
}

// CertificateDetailOptions represents the CertificateDetailOptions schema from the OpenAPI specification
type CertificateDetailOptions struct {
	Certificatetransparencyloggingpreference interface{} `json:"CertificateTransparencyLoggingPreference,omitempty"`
}

// DeleteCertificateRequest represents the DeleteCertificateRequest schema from the OpenAPI specification
type DeleteCertificateRequest struct {
	Certificatearn interface{} `json:"CertificateArn"`
}

// UpdateCertificateOptionsRequest represents the UpdateCertificateOptionsRequest schema from the OpenAPI specification
type UpdateCertificateOptionsRequest struct {
	Certificatearn interface{} `json:"CertificateArn"`
	Options UpdateCertificateOptionsRequestOptions `json:"Options"`
}

// ListCertificatesRequest represents the ListCertificatesRequest schema from the OpenAPI specification
type ListCertificatesRequest struct {
	Certificatestatuses interface{} `json:"CertificateStatuses,omitempty"`
	Includes ListCertificatesRequestIncludes `json:"Includes,omitempty"`
	Maxitems interface{} `json:"MaxItems,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
}

// CertificateDetailRenewalSummary represents the CertificateDetailRenewalSummary schema from the OpenAPI specification
type CertificateDetailRenewalSummary struct {
	Renewalstatusreason interface{} `json:"RenewalStatusReason,omitempty"`
	Updatedat interface{} `json:"UpdatedAt"`
	Domainvalidationoptions interface{} `json:"DomainValidationOptions"`
	Renewalstatus interface{} `json:"RenewalStatus"`
}

// CertificateSummary represents the CertificateSummary schema from the OpenAPI specification
type CertificateSummary struct {
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Extendedkeyusages interface{} `json:"ExtendedKeyUsages,omitempty"`
	Importedat interface{} `json:"ImportedAt,omitempty"`
	Renewaleligibility interface{} `json:"RenewalEligibility,omitempty"`
	Notbefore interface{} `json:"NotBefore,omitempty"`
	Domainname interface{} `json:"DomainName,omitempty"`
	Hasadditionalsubjectalternativenames interface{} `json:"HasAdditionalSubjectAlternativeNames,omitempty"`
	Revokedat interface{} `json:"RevokedAt,omitempty"`
	Issuedat interface{} `json:"IssuedAt,omitempty"`
	Keyusages interface{} `json:"KeyUsages,omitempty"`
	Subjectalternativenamesummaries interface{} `json:"SubjectAlternativeNameSummaries,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Keyalgorithm interface{} `json:"KeyAlgorithm,omitempty"`
	Inuse interface{} `json:"InUse,omitempty"`
	Notafter interface{} `json:"NotAfter,omitempty"`
	Exported interface{} `json:"Exported,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Certificatearn interface{} `json:"CertificateArn,omitempty"`
}

// RequestCertificateResponse represents the RequestCertificateResponse schema from the OpenAPI specification
type RequestCertificateResponse struct {
	Certificatearn interface{} `json:"CertificateArn,omitempty"`
}

// ImportCertificateRequest represents the ImportCertificateRequest schema from the OpenAPI specification
type ImportCertificateRequest struct {
	Privatekey interface{} `json:"PrivateKey"`
	Tags interface{} `json:"Tags,omitempty"`
	Certificate interface{} `json:"Certificate"`
	Certificatearn interface{} `json:"CertificateArn,omitempty"`
	Certificatechain interface{} `json:"CertificateChain,omitempty"`
}

// ListCertificatesRequestIncludes represents the ListCertificatesRequestIncludes schema from the OpenAPI specification
type ListCertificatesRequestIncludes struct {
	Keytypes interface{} `json:"keyTypes,omitempty"`
	Keyusage interface{} `json:"keyUsage,omitempty"`
	Extendedkeyusage interface{} `json:"extendedKeyUsage,omitempty"`
}

// RequestCertificateRequest represents the RequestCertificateRequest schema from the OpenAPI specification
type RequestCertificateRequest struct {
	Validationmethod interface{} `json:"ValidationMethod,omitempty"`
	Keyalgorithm interface{} `json:"KeyAlgorithm,omitempty"`
	Subjectalternativenames interface{} `json:"SubjectAlternativeNames,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Certificateauthorityarn interface{} `json:"CertificateAuthorityArn,omitempty"`
	Domainname interface{} `json:"DomainName"`
	Domainvalidationoptions interface{} `json:"DomainValidationOptions,omitempty"`
	Idempotencytoken interface{} `json:"IdempotencyToken,omitempty"`
	Options RequestCertificateRequestOptions `json:"Options,omitempty"`
}

// ExportCertificateRequest represents the ExportCertificateRequest schema from the OpenAPI specification
type ExportCertificateRequest struct {
	Certificatearn interface{} `json:"CertificateArn"`
	Passphrase interface{} `json:"Passphrase"`
}

// ListCertificatesResponse represents the ListCertificatesResponse schema from the OpenAPI specification
type ListCertificatesResponse struct {
	Certificatesummarylist interface{} `json:"CertificateSummaryList,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CertificateDetail represents the CertificateDetail schema from the OpenAPI specification
type CertificateDetail struct {
	TypeField interface{} `json:"Type,omitempty"`
	Certificateauthorityarn interface{} `json:"CertificateAuthorityArn,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Revocationreason interface{} `json:"RevocationReason,omitempty"`
	Revokedat interface{} `json:"RevokedAt,omitempty"`
	Domainname interface{} `json:"DomainName,omitempty"`
	Extendedkeyusages interface{} `json:"ExtendedKeyUsages,omitempty"`
	Importedat interface{} `json:"ImportedAt,omitempty"`
	Signaturealgorithm interface{} `json:"SignatureAlgorithm,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Options CertificateDetailOptions `json:"Options,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Domainvalidationoptions interface{} `json:"DomainValidationOptions,omitempty"`
	Issuedat interface{} `json:"IssuedAt,omitempty"`
	Notbefore interface{} `json:"NotBefore,omitempty"`
	Certificatearn interface{} `json:"CertificateArn,omitempty"`
	Inuseby interface{} `json:"InUseBy,omitempty"`
	Keyusages interface{} `json:"KeyUsages,omitempty"`
	Serial interface{} `json:"Serial,omitempty"`
	Notafter interface{} `json:"NotAfter,omitempty"`
	Renewaleligibility interface{} `json:"RenewalEligibility,omitempty"`
	Subject interface{} `json:"Subject,omitempty"`
	Keyalgorithm interface{} `json:"KeyAlgorithm,omitempty"`
	Renewalsummary CertificateDetailRenewalSummary `json:"RenewalSummary,omitempty"`
	Issuer interface{} `json:"Issuer,omitempty"`
	Subjectalternativenames interface{} `json:"SubjectAlternativeNames,omitempty"`
}

// RemoveTagsFromCertificateRequest represents the RemoveTagsFromCertificateRequest schema from the OpenAPI specification
type RemoveTagsFromCertificateRequest struct {
	Certificatearn interface{} `json:"CertificateArn"`
	Tags interface{} `json:"Tags"`
}

// ResendValidationEmailRequest represents the ResendValidationEmailRequest schema from the OpenAPI specification
type ResendValidationEmailRequest struct {
	Certificatearn interface{} `json:"CertificateArn"`
	Domain interface{} `json:"Domain"`
	Validationdomain interface{} `json:"ValidationDomain"`
}

// GetAccountConfigurationResponseExpiryEvents represents the GetAccountConfigurationResponseExpiryEvents schema from the OpenAPI specification
type GetAccountConfigurationResponseExpiryEvents struct {
	Daysbeforeexpiry interface{} `json:"DaysBeforeExpiry,omitempty"`
}

// RequestCertificateRequestOptions represents the RequestCertificateRequestOptions schema from the OpenAPI specification
type RequestCertificateRequestOptions struct {
	Certificatetransparencyloggingpreference interface{} `json:"CertificateTransparencyLoggingPreference,omitempty"`
}

// DomainValidationResourceRecord represents the DomainValidationResourceRecord schema from the OpenAPI specification
type DomainValidationResourceRecord struct {
	TypeField interface{} `json:"Type"`
	Value interface{} `json:"Value"`
	Name interface{} `json:"Name"`
}

// DescribeCertificateResponse represents the DescribeCertificateResponse schema from the OpenAPI specification
type DescribeCertificateResponse struct {
	Certificate DescribeCertificateResponseCertificate `json:"Certificate,omitempty"`
}

// ExportCertificateResponse represents the ExportCertificateResponse schema from the OpenAPI specification
type ExportCertificateResponse struct {
	Privatekey interface{} `json:"PrivateKey,omitempty"`
	Certificate interface{} `json:"Certificate,omitempty"`
	Certificatechain interface{} `json:"CertificateChain,omitempty"`
}

// CertificateOptions represents the CertificateOptions schema from the OpenAPI specification
type CertificateOptions struct {
	Certificatetransparencyloggingpreference interface{} `json:"CertificateTransparencyLoggingPreference,omitempty"`
}

// DescribeCertificateResponseCertificate represents the DescribeCertificateResponseCertificate schema from the OpenAPI specification
type DescribeCertificateResponseCertificate struct {
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Domainvalidationoptions interface{} `json:"DomainValidationOptions,omitempty"`
	Issuedat interface{} `json:"IssuedAt,omitempty"`
	Notbefore interface{} `json:"NotBefore,omitempty"`
	Certificatearn interface{} `json:"CertificateArn,omitempty"`
	Inuseby interface{} `json:"InUseBy,omitempty"`
	Keyusages interface{} `json:"KeyUsages,omitempty"`
	Serial interface{} `json:"Serial,omitempty"`
	Notafter interface{} `json:"NotAfter,omitempty"`
	Renewaleligibility interface{} `json:"RenewalEligibility,omitempty"`
	Subject interface{} `json:"Subject,omitempty"`
	Keyalgorithm interface{} `json:"KeyAlgorithm,omitempty"`
	Renewalsummary CertificateDetailRenewalSummary `json:"RenewalSummary,omitempty"`
	Issuer interface{} `json:"Issuer,omitempty"`
	Subjectalternativenames interface{} `json:"SubjectAlternativeNames,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Certificateauthorityarn interface{} `json:"CertificateAuthorityArn,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Revocationreason interface{} `json:"RevocationReason,omitempty"`
	Revokedat interface{} `json:"RevokedAt,omitempty"`
	Domainname interface{} `json:"DomainName,omitempty"`
	Extendedkeyusages interface{} `json:"ExtendedKeyUsages,omitempty"`
	Importedat interface{} `json:"ImportedAt,omitempty"`
	Signaturealgorithm interface{} `json:"SignatureAlgorithm,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Options CertificateDetailOptions `json:"Options,omitempty"`
}

// ExtendedKeyUsage represents the ExtendedKeyUsage schema from the OpenAPI specification
type ExtendedKeyUsage struct {
	Name interface{} `json:"Name,omitempty"`
	Oid interface{} `json:"OID,omitempty"`
}

// ResourceRecord represents the ResourceRecord schema from the OpenAPI specification
type ResourceRecord struct {
	Value interface{} `json:"Value"`
	Name interface{} `json:"Name"`
	TypeField interface{} `json:"Type"`
}

// DomainValidation represents the DomainValidation schema from the OpenAPI specification
type DomainValidation struct {
	Validationdomain interface{} `json:"ValidationDomain,omitempty"`
	Validationemails interface{} `json:"ValidationEmails,omitempty"`
	Validationmethod interface{} `json:"ValidationMethod,omitempty"`
	Validationstatus interface{} `json:"ValidationStatus,omitempty"`
	Domainname interface{} `json:"DomainName"`
	Resourcerecord DomainValidationResourceRecord `json:"ResourceRecord,omitempty"`
}

// DescribeCertificateRequest represents the DescribeCertificateRequest schema from the OpenAPI specification
type DescribeCertificateRequest struct {
	Certificatearn interface{} `json:"CertificateArn"`
}

// Filters represents the Filters schema from the OpenAPI specification
type Filters struct {
	Extendedkeyusage interface{} `json:"extendedKeyUsage,omitempty"`
	Keytypes interface{} `json:"keyTypes,omitempty"`
	Keyusage interface{} `json:"keyUsage,omitempty"`
}

// PutAccountConfigurationRequest represents the PutAccountConfigurationRequest schema from the OpenAPI specification
type PutAccountConfigurationRequest struct {
	Expiryevents PutAccountConfigurationRequestExpiryEvents `json:"ExpiryEvents,omitempty"`
	Idempotencytoken interface{} `json:"IdempotencyToken"`
}

// GetCertificateRequest represents the GetCertificateRequest schema from the OpenAPI specification
type GetCertificateRequest struct {
	Certificatearn interface{} `json:"CertificateArn"`
}

// GetCertificateResponse represents the GetCertificateResponse schema from the OpenAPI specification
type GetCertificateResponse struct {
	Certificate interface{} `json:"Certificate,omitempty"`
	Certificatechain interface{} `json:"CertificateChain,omitempty"`
}

// ListTagsForCertificateRequest represents the ListTagsForCertificateRequest schema from the OpenAPI specification
type ListTagsForCertificateRequest struct {
	Certificatearn interface{} `json:"CertificateArn"`
}

// RenewalSummary represents the RenewalSummary schema from the OpenAPI specification
type RenewalSummary struct {
	Updatedat interface{} `json:"UpdatedAt"`
	Domainvalidationoptions interface{} `json:"DomainValidationOptions"`
	Renewalstatus interface{} `json:"RenewalStatus"`
	Renewalstatusreason interface{} `json:"RenewalStatusReason,omitempty"`
}

// ExpiryEventsConfiguration represents the ExpiryEventsConfiguration schema from the OpenAPI specification
type ExpiryEventsConfiguration struct {
	Daysbeforeexpiry interface{} `json:"DaysBeforeExpiry,omitempty"`
}

// ImportCertificateResponse represents the ImportCertificateResponse schema from the OpenAPI specification
type ImportCertificateResponse struct {
	Certificatearn interface{} `json:"CertificateArn,omitempty"`
}

// DomainValidationOption represents the DomainValidationOption schema from the OpenAPI specification
type DomainValidationOption struct {
	Domainname interface{} `json:"DomainName"`
	Validationdomain interface{} `json:"ValidationDomain"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value,omitempty"`
}

// PutAccountConfigurationRequestExpiryEvents represents the PutAccountConfigurationRequestExpiryEvents schema from the OpenAPI specification
type PutAccountConfigurationRequestExpiryEvents struct {
	Daysbeforeexpiry interface{} `json:"DaysBeforeExpiry,omitempty"`
}

// ListTagsForCertificateResponse represents the ListTagsForCertificateResponse schema from the OpenAPI specification
type ListTagsForCertificateResponse struct {
	Tags interface{} `json:"Tags,omitempty"`
}
