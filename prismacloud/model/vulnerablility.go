package model

type VulnerableOverviewSummary struct {
	OverviewSummary OverviewSummary `json:"overviewSummary"`
	Values          []Value         `json:"values"`
}

// OverviewSummary represents the overview summary part of the JSON structure
type OverviewSummary struct {
	TotalVulnerableRuntimeAssets  TotalVulnerableRuntimeAssets  `json:"totalVulnerableRuntimeAssets"`
	TotalVulnerabilitiesinRuntime TotalVulnerabilitiesInRuntime `json:"totalVulnerabilitiesinRuntime"`
	TotalRemediatedInRuntime      TotalRemediatedInRuntime      `json:"totalRemediatedinRuntime"`
}

// TotalVulnerableRuntimeAssets represents the total vulnerable runtime assets part of the JSON structure
type TotalVulnerableRuntimeAssets struct {
	TotalCount              int `json:"totalCount"`
	DeployedImageCount      int `json:"deployedImageCount"`
	ServerlessFunctionCount int `json:"serverlessFunctionCount"`
	HostCount               int `json:"hostCount"`
}

// TotalVulnerabilitiesInRuntime represents the total vulnerabilities in runtime part of the JSON structure
type TotalVulnerabilitiesInRuntime struct {
	TotalCount    int `json:"totalCount"`
	CriticalCount int `json:"criticalCount"`
	HighCount     int `json:"highCount"`
	MediumCount   int `json:"mediumCount"`
	LowCount      int `json:"lowCount"`
}

// TotalRemediatedInRuntime represents the total remediated in runtime part of the JSON structure
type TotalRemediatedInRuntime struct {
	TotalCount    int `json:"totalCount"`
	CriticalCount int `json:"criticalCount"`
	HighCount     int `json:"highCount"`
	MediumCount   int `json:"mediumCount"`
	LowCount      int `json:"lowCount"`
}

// Value represents each value in the values array of the JSON structure
type Value struct {
	LastUpdatedDateTime     int `json:"lastUpdatedDateTime"`
	TotalVulnerabilityCount int `json:"totalVulnerabilityCount"`
	TotalVulnerableAsset    int `json:"totalVulnerableAsset"`
	TotalRemediationCount   int `json:"totalRemediationCount"`
}

type PrioritizedVulnerabilitySummary struct {
	LastUpdatedDateTime  int                  `json:"lastUpdatedDateTime"`
	TotalVulnerabilities int                  `json:"totalVulnerabilities"`
	Urgent               VulnerabilityDetails `json:"urgent"`
	Patchable            VulnerabilityDetails `json:"patchable"`
	Exploitable          VulnerabilityDetails `json:"exploitable"`
	InternetExposed      VulnerabilityDetails `json:"internetExposed"`
	PackageInUse         VulnerabilityDetails `json:"packageInUse"`
}

type VulnerabilityDetails struct {
	VulnerabilityCount int `json:"vulnerability_count"`
	AssetCount         int `json:"asset_count"`
}