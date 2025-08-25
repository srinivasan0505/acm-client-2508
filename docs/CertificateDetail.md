

# CertificateDetail

Contains metadata about an ACM certificate. This structure is returned in the response to a <a>DescribeCertificate</a> request. 

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**certificateArn** | [**String**](String.md) |  |  [optional] |
|**domainName** | [**String**](String.md) |  |  [optional] |
|**subjectAlternativeNames** | [**List**](List.md) |  |  [optional] |
|**domainValidationOptions** | [**List**](List.md) |  |  [optional] |
|**serial** | [**String**](String.md) |  |  [optional] |
|**subject** | [**String**](String.md) |  |  [optional] |
|**issuer** | [**String**](String.md) |  |  [optional] |
|**createdAt** | [**OffsetDateTime**](OffsetDateTime.md) |  |  [optional] |
|**issuedAt** | [**OffsetDateTime**](OffsetDateTime.md) |  |  [optional] |
|**importedAt** | [**OffsetDateTime**](OffsetDateTime.md) |  |  [optional] |
|**status** | [**CertificateStatus**](CertificateStatus.md) |  |  [optional] |
|**revokedAt** | [**OffsetDateTime**](OffsetDateTime.md) |  |  [optional] |
|**revocationReason** | [**RevocationReason**](RevocationReason.md) |  |  [optional] |
|**notBefore** | [**OffsetDateTime**](OffsetDateTime.md) |  |  [optional] |
|**notAfter** | [**OffsetDateTime**](OffsetDateTime.md) |  |  [optional] |
|**keyAlgorithm** | [**KeyAlgorithm**](KeyAlgorithm.md) |  |  [optional] |
|**signatureAlgorithm** | [**String**](String.md) |  |  [optional] |
|**inUseBy** | [**List**](List.md) |  |  [optional] |
|**failureReason** | [**FailureReason**](FailureReason.md) |  |  [optional] |
|**type** | [**CertificateType**](CertificateType.md) |  |  [optional] |
|**renewalSummary** | [**CertificateDetailRenewalSummary**](CertificateDetailRenewalSummary.md) |  |  [optional] |
|**keyUsages** | [**List**](List.md) |  |  [optional] |
|**extendedKeyUsages** | [**List**](List.md) |  |  [optional] |
|**certificateAuthorityArn** | [**String**](String.md) |  |  [optional] |
|**renewalEligibility** | [**RenewalEligibility**](RenewalEligibility.md) |  |  [optional] |
|**options** | [**CertificateDetailOptions**](CertificateDetailOptions.md) |  |  [optional] |



