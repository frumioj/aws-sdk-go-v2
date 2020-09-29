// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AdditionalArtifact string

// Enum values for AdditionalArtifact
const (
	AdditionalArtifactRedshift   AdditionalArtifact = "REDSHIFT"
	AdditionalArtifactQuicksight AdditionalArtifact = "QUICKSIGHT"
	AdditionalArtifactAthena     AdditionalArtifact = "ATHENA"
)

type AWSRegion string

// Enum values for AWSRegion
const (
	AWSRegionUs_standard         AWSRegion = "us-east-1"
	AWSRegionNorthern_california AWSRegion = "us-west-1"
	AWSRegionOregon              AWSRegion = "us-west-2"
	AWSRegionFrankfurt           AWSRegion = "eu-central-1"
	AWSRegionIreland             AWSRegion = "eu-west-1"
	AWSRegionSingapore           AWSRegion = "ap-southeast-1"
	AWSRegionSydney              AWSRegion = "ap-southeast-2"
	AWSRegionTokyo               AWSRegion = "ap-northeast-1"
	AWSRegionStockholm           AWSRegion = "eu-north-1"
	AWSRegionOsaka               AWSRegion = "ap-northeast-3"
	AWSRegionHong_kong           AWSRegion = "ap-east-1"
)

type CompressionFormat string

// Enum values for CompressionFormat
const (
	CompressionFormatZip     CompressionFormat = "ZIP"
	CompressionFormatGzip    CompressionFormat = "GZIP"
	CompressionFormatParquet CompressionFormat = "Parquet"
)

type ReportFormat string

// Enum values for ReportFormat
const (
	ReportFormatCsv     ReportFormat = "textORcsv"
	ReportFormatParquet ReportFormat = "Parquet"
)

type ReportVersioning string

// Enum values for ReportVersioning
const (
	ReportVersioningCreate_new_report ReportVersioning = "CREATE_NEW_REPORT"
	ReportVersioningOverwrite_report  ReportVersioning = "OVERWRITE_REPORT"
)

type SchemaElement string

// Enum values for SchemaElement
const (
	SchemaElementResources SchemaElement = "RESOURCES"
)

type TimeUnit string

// Enum values for TimeUnit
const (
	TimeUnitHourly TimeUnit = "HOURLY"
	TimeUnitDaily  TimeUnit = "DAILY"
)