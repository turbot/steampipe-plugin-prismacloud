package prismacloud

type ComplianceDetails struct {
	AssignedPolicies                     int32  `json:"assignedPolicies"`
	CriticalSeverityFailedResources      int64  `json:"criticalSeverityFailedResources"`
	Default                              bool   `json:"default"`
	Description                          string `json:"description"`
	FailedResources                      int64  `json:"failedResources"`
	HighSeverityFailedResources          int64  `json:"highSeverityFailedResources"`
	ID                                   string `json:"id"`
	InformationalSeverityFailedResources int64  `json:"informationalSeverityFailedResources"`
	LowSeverityFailedResources           int64  `json:"lowSeverityFailedResources"`
	MediumSeverityFailedResources        int64  `json:"mediumSeverityFailedResources"`
	Name                                 string `json:"name"`
	PassedResources                      int64  `json:"passedResources"`
	TotalResources                       int64  `json:"totalResources"`
}

type ComplianceData struct {
	ComplianceDetails    []ComplianceDetails  `json:"complianceDetails"`
	RequestedTimestamp   int64                `json:"requestedTimestamp"`
	RequirementSummaries []RequirementSummary `json:"requirementSummaries"`
	Summary              ComplianceSummary    `json:"summary"`
}

type RequirementSummary struct {
	ID               string           `json:"id"`
	Name             string           `json:"name"`
	SectionSummaries []SectionSummary `json:"sectionSummaries"`
}

type SectionSummary struct {
	FailedResources int64  `json:"failedResources"`
	ID              string `json:"id"`
	Name            string `json:"name"`
	PassedResources int64  `json:"passedResources"`
	TotalResources  int64  `json:"totalResources"`
}

type ComplianceSummary struct {
	CriticalSeverityFailedResources      int64 `json:"criticalSeverityFailedResources"`
	FailedResources                      int64 `json:"failedResources"`
	HighSeverityFailedResources          int64 `json:"highSeverityFailedResources"`
	InformationalSeverityFailedResources int64 `json:"informationalSeverityFailedResources"`
	LowSeverityFailedResources           int64 `json:"lowSeverityFailedResources"`
	MediumSeverityFailedResources        int64 `json:"mediumSeverityFailedResources"`
	PassedResources                      int64 `json:"passedResources"`
	Timestamp                            int64 `json:"timestamp"`
	TotalResources                       int64 `json:"totalResources"`
}
