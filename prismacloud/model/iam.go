package model

type Item struct {
	Id                            string      `json:"id"`
	SourcePublic                  bool        `json:"sourcePublic"`
	SourceCloudType               string      `json:"sourceCloudType"`
	SourceCloudAccount            string      `json:"sourceCloudAccount"`
	SourceCloudRegion             string      `json:"sourceCloudRegion"`
	SourceCloudServiceName        string      `json:"sourceCloudServiceName"`
	SourceResourceName            string      `json:"sourceResourceName"`
	SourceResourceType            string      `json:"sourceResourceType"`
	SourceResourceId              string      `json:"sourceResourceId"`
	SourceCloudResourceUai        string      `json:"sourceCloudResourceUai"`
	SourceIdpService              string      `json:"sourceIdpService"`
	SourceIdpDomain               string      `json:"sourceIdpDomain"`
	SourceIdpEmail                string      `json:"sourceIdpEmail"`
	SourceIdpUserId               string      `json:"sourceIdpUserId"`
	SourceIdpUsername             string      `json:"sourceIdpUsername"`
	SourceIdpGroup                string      `json:"sourceIdpGroup"`
	SourceIdpUai                  string      `json:"sourceIdpUai"`
	DestCloudType                 string      `json:"destCloudType"`
	DestCloudAccount              string      `json:"destCloudAccount"`
	DestCloudRegion               string      `json:"destCloudRegion"`
	DestCloudServiceName          string      `json:"destCloudServiceName"`
	DestResourceName              string      `json:"destResourceName"`
	DestResourceType              string      `json:"destResourceType"`
	DestResourceId                string      `json:"destResourceId"`
	DestCloudResourceUai          string      `json:"destCloudResourceUai"`
	GrantedByCloudType            string      `json:"grantedByCloudType"`
	GrantedByCloudPolicyId        string      `json:"grantedByCloudPolicyId"`
	GrantedByCloudPolicyName      string      `json:"grantedByCloudPolicyName"`
	GrantedByCloudPolicyType      string      `json:"grantedByCloudPolicyType"`
	GrantedByCloudPolicyUai       string      `json:"grantedByCloudPolicyUai"`
	GrantedByCloudPolicyAccount   string      `json:"grantedByCloudPolicyAccount"`
	GrantedByCloudEntityId        string      `json:"grantedByCloudEntityId"`
	GrantedByCloudEntityName      string      `json:"grantedByCloudEntityName"`
	GrantedByCloudEntityType      string      `json:"grantedByCloudEntityType"`
	GrantedByCloudEntityAccount   string      `json:"grantedByCloudEntityAccount"`
	GrantedByCloudEntityUai       string      `json:"grantedByCloudEntityUai"`
	GrantedByLevelType            string      `json:"grantedByLevelType"`
	GrantedByLevelId              string      `json:"grantedByLevelId"`
	GrantedByLevelName            string      `json:"grantedByLevelName"`
	GrantedByLevelUai             string      `json:"grantedByLevelUai"`
	LastAccessDate                string      `json:"lastAccessDate"`
	LastAccessStatus              string      `json:"lastAccessStatus"`
	AccessedResourcesCount        int         `json:"accessedResourcesCount"`
	EffectiveActionName           string      `json:"effectiveActionName"`
	Exceptions                    []Exception `json:"exceptions"`
	WildCardDestCloudResourceName bool        `json:"wildCardDestCloudResourceName"`
}

type Data struct {
	Items                          []Item   `json:"items"`
	NextPageToken                  string   `json:"nextPageToken"`
	TotalRows                      int      `json:"totalRows"`
	SearchedDestCloudResourceNames []string `json:"searchedDestCloudResourceNames"`
}

type Exception struct {
	MessageCode string `json:"messageCode"`
}

type PermissionResponse struct {
	Data        Data                   `json:"data"`
	Query       string                 `json:"query"`
	Id          string                 `json:"id"`
	Saved       bool                   `json:"saved"`
	Name        string                 `json:"name"`
	TimeRange   map[string]interface{} `json:"timeRange"`
	SearchType  string                 `json:"searchType"`
	Description string                 `json:"description"`
	CloudType   string                 `json:"cloudType"`
}

// Permission access structs

type PermissionAccessResponse struct {
	Data PermissionAccessData `json:"data"`
}

type PermissionAccessData struct {
	Items         []PermissionAccessItem `json:"items"`
	NextPageToken string                 `json:"nextPageToken"`
	TotalRows     int                    `json:"totalRows"`
}

type PermissionAccessItem struct {
	DestCloudResourceName string `json:"destCloudResourceName"`
	LastAccessDate        string `json:"lastAccessDate"`
	DestCloudRegion       string `json:"destCloudRegion"`
	DestCloudAccount      string `json:"destCloudAccount"`
}
