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
