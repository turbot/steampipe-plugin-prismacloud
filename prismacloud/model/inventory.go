package model

type PathRiskFactors struct {
	ResponseSensitiveData  []string `json:"responseSensitiveData"`
	RequestSensitiveData   []string `json:"requestSensitiveData"`
	RequiresAuthentication bool     `json:"requiresAuthentication"`
	InternetExposed        bool     `json:"internetExposed"`
	OwaspAPIAttacks        []string `json:"owaspAPIAttacks"`
}

type InventoryDiscoveredAPIMember struct {
	AssetID             string          `json:"assetId"`
	APIPath             string          `json:"apiPath"`
	HTTPMethod          string          `json:"httpMethod"`
	APIServer           string          `json:"apiServer"`
	Hits                int             `json:"hits"`
	Workloads           []string        `json:"workloads"`
	ServiceName         string          `json:"serviceName"`
	CloudType           string          `json:"cloudType"`
	Region              string          `json:"region"`
	AccountID           string          `json:"accountId"`
	AccountName         string          `json:"accountName"`
	PathRiskFactors     PathRiskFactors `json:"pathRiskFactors"`
	DiscoveryMethod     string          `json:"discoveryMethod"`
	InspectionType      string          `json:"inspectionType"`
	LastChanged         int64           `json:"lastChanged"`
	LastObserved        int64           `json:"lastObserved"`
	RequestContentType  []string        `json:"requestContentType"`
	ResponseContentType []string        `json:"responseContentType"`
}

type InventoryDiscoveredAPIResponse struct {
	Members       []InventoryDiscoveredAPIMember `json:"members"`
	NextPageToken *string                        `json:"nextPageToken"`
	Total         int                            `json:"_total"`
	Count         int                            `json:"_count"`
}

// // WORKLOAD
type InventoryWorkload struct {
	ContainerImages ContainerImages `json:"containerImages"`
	Hosts           Hosts           `json:"hosts"`
}

type ContainerImages struct {
	Stages struct {
		Build  int `json:"build"`
		Deploy int `json:"deploy"`
		Run    int `json:"run"`
	} `json:"stages"`
	Vulnerable     int         `json:"vulnerable"`
	CloudProviders interface{} `json:"cloudProviders"`
}

type Hosts struct {
	Total          int      `json:"total"`
	Vulnerable     int      `json:"vulnerable"`
	CloudProviders []string `json:"cloudProviders"`
}

//// WOrkload Container Images

type Stage struct {
	Build  int `json:"build"`
	Deploy int `json:"deploy"`
	Run    int `json:"run"`
}
type VulnFunnel struct {
	Total        int `json:"total"`
	Urgent       int `json:"urgent"`
	Exploitable  int `json:"exploitable"`
	Patchable    int `json:"patchable"`
	PackageInUse int `json:"packageInUse"`
}

type WorkloadContainerImage struct {
	Name              string     `json:"name"`
	UaiID             string     `json:"uaiID"`
	Stages            Stage      `json:"stages"`
	RunningContainers int        `json:"runningContainers"`
	VulnFunnel        VulnFunnel `json:"vulnFunnel"`
	ScanPassed        bool       `json:"scanPassed"`
	Base              bool       `json:"base"`
	RelatedImages     int        `json:"relatedImages"`
}

type WorkloadContainerImagesResponse struct {
	Value         []WorkloadContainerImage `json:"value"`
	Total         int                      `json:"total"`
	NextPageToken string                   `json:"nextPageToken"`
}

//// WORKLOAD HOST

type WorkloadContainerHostResponse struct {
	Value         []WorkLoadInventoryHost `json:"value"`
	Total         int                     `json:"total"`
	NextPageToken string                  `json:"nextPageToken"`
}

type WorkLoadInventoryHost struct {
	Name       string     `json:"name"`
	ID         string     `json:"id"`
	UaiID      string     `json:"uaiID"`
	VulnFunnel VulnFunnel `json:"vulnFunnel"`
}

//// GROUPED AGGREGATED ASSET

type GroupedAggregateAssetResponse struct {
	GroupedAggregate []GroupedAggregateAsset `json:"groupedAggregates"`
}

type GroupedAggregateAsset struct {
	AccountId                         string `json:"accountId"`
	AccountName                       string `json:"accountName"`
	AllowDrillDown                    bool   `json:"allowDrillDown"`
	CloudTypeName                     string `json:"cloudTypeName"`
	CriticalSeverityFailedResources   int64  `json:"criticalSeverityFailedResources"`
	CriticalVulnerabilityFailedResources int64 `json:"criticalVulnerabilityFailedResources"`
	FailedResources                   int64  `json:"failedResources"`
	HighSeverityFailedResources       int64  `json:"highSeverityFailedResources"`
	HighVulnerabilityFailedResources  int64  `json:"highVulnerabilityFailedResources"`
	InformationalSeverityFailedResources int64 `json:"informationalSeverityFailedResources"`
	LowSeverityFailedResources        int64  `json:"lowSeverityFailedResources"`
	LowVulnerabilityFailedResources   int64  `json:"lowVulnerabilityFailedResources"`
	MediumSeverityFailedResources     int64  `json:"mediumSeverityFailedResources"`
	MediumVulnerabilityFailedResources int64  `json:"mediumVulnerabilityFailedResources"`
	PassedResources                   int64  `json:"passedResources"`
	RegionName                        string `json:"regionName"`
	ResourceTypeName                  string `json:"resourceTypeName"`
	ServiceName                       string `json:"serviceName"`
	TotalResources                    int64  `json:"totalResources"`
	TotalVulnerabilityFailedResources int64  `json:"totalVulnerabilityFailedResources"`
	UnscannedResources                int64  `json:"unscannedResources"`
}
