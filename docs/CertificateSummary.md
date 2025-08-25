

# CertificateSummary

This structure is returned in the response object of <a>ListCertificates</a> action. 

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**certificateArn** | [**String**](String.md) |  |  [optional] |
|**domainName** | [**String**](String.md) |  |  [optional] |
|**subjectAlternativeNameSummaries** | [**List**](List.md) |  |  [optional] |
|**hasAdditionalSubjectAlternativeNames** | [**Boolean**](Boolean.md) |  |  [optional] |
|**status** | [**CertificateStatus**](CertificateStatus.md) |  |  [optional] |
|**type** | [**CertificateType**](CertificateType.md) |  |  [optional] |
|**keyAlgorithm** | [**KeyAlgorithm**](KeyAlgorithm.md) |  |  [optional] |
|**keyUsages** | [**List**](List.md) |  |  [optional] |
|**extendedKeyUsages** | [**List**](List.md) |  |  [optional] |
|**inUse** | [**Boolean**](Boolean.md) |  |  [optional] |
|**exported** | [**Boolean**](Boolean.md) |  |  [optional] |
|**renewalEligibility** | [**RenewalEligibility**](RenewalEligibility.md) |  |  [optional] |
|**notBefore** | [**OffsetDateTime**](OffsetDateTime.md) |  |  [optional] |
|**notAfter** | [**OffsetDateTime**](OffsetDateTime.md) |  |  [optional] |
|**createdAt** | [**OffsetDateTime**](OffsetDateTime.md) |  |  [optional] |
|**issuedAt** | [**OffsetDateTime**](OffsetDateTime.md) |  |  [optional] |
|**importedAt** | [**OffsetDateTime**](OffsetDateTime.md) |  |  [optional] |
|**revokedAt** | [**OffsetDateTime**](OffsetDateTime.md) |  |  [optional] |



