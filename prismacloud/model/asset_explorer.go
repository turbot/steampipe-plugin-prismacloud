package model

type AlertStatus struct {
	Critical      int `json:"critical"`
	High          int `json:"high"`
	Informational int `json:"informational"`
	Low           int `json:"low"`
	Medium        int `json:"medium"`
}

type ScannedPolicy struct {
	Id       string   `json:"id"`
	Labels   []string `json:"labels"`
	Name     string   `json:"name"`
	Passed   bool     `json:"passed"`
	Severity string   `json:"severity"`
}

type Resource struct {
	AccountId                   string          `json:"accountId"`
	AccountName                 string          `json:"accountName"`
	AlertStatus                 AlertStatus     `json:"alertStatus"`
	AppNames                    []string        `json:"appNames"`
	AssetType                   string          `json:"assetType"`
	CloudType                   string          `json:"cloudType"`
	Id                          string          `json:"id"`
	Name                        string          `json:"name"`
	OverallPassed               bool            `json:"overallPassed"`
	RegionId                    string          `json:"regionId"`
	RegionName                  string          `json:"regionName"`
	ResourceConfigJsonAvailable bool            `json:"resourceConfigJsonAvailable"`
	ResourceDetailsAvailable    bool            `json:"resourceDetailsAvailable"`
	Rrn                         string          `json:"rrn"`
	ScannedPolicies             []ScannedPolicy `json:"scannedPolicies"`
	UnifiedAssetId              string          `json:"unifiedAssetId"`
	VulnerabilityStatus         AlertStatus     `json:"vulnerabilityStatus"`
}

type PrismaCloudAssetExplorer struct {
	NextPageToken     string     `json:"nextPageToken"`
	PageSize          int        `json:"pageSize"`
	Resources         []Resource `json:"resources"`
	Timestamp         int64      `json:"timestamp"`
	TotalMatchedCount int        `json:"totalMatchedCount"`
}
