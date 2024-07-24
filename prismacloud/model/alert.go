package model

type ComplianceMetadata struct {
	StandardName           string `json:"standardName"`
	StandardDescription    string `json:"standardDescription"`
	RequirementId          string `json:"requirementId"`
	RequirementName        string `json:"requirementName"`
	RequirementDescription string `json:"requirementDescription"`
	SectionId              string `json:"sectionId"`
	SectionDescription     string `json:"sectionDescription"`
	PolicyId               string `json:"policyId"`
	ComplianceId           string `json:"complianceId"`
	SectionLabel           string `json:"sectionLabel"`
	SectionViewOrder       int    `json:"sectionViewOrder"`
	RequirementViewOrder   int    `json:"requirementViewOrder"`
	SystemDefault          bool   `json:"systemDefault"`
	PolicyName             string `json:"policyName"`
	CustomAssigned         bool   `json:"customAssigned"`
}

type Policy struct {
	AlertCount            int                `json:"alertCount"`
	PolicyId              string             `json:"policyId"`
	PolicyName            string             `json:"policyName"`
	PolicyType            string             `json:"policyType"`
	Severity              string             `json:"severity"`
	PolicyLabels          []string           `json:"policyLabels"`
	ComplianceMetadata    []ComplianceMetadata `json:"complianceMetadata"`
	ResourceType          string             `json:"resourceType"`
	Remediable            bool               `json:"remediable"`
	CloudType             string             `json:"cloudType"`
	MittreAttacks         []string           `json:"mittreAttacks"`
	FindingTypes          []string           `json:"findingTypes"`
	RestrictAlertDismissal bool              `json:"restrictAlertDismissal"`
}

type CountDetails struct {
	TotalAlerts   int `json:"totalAlerts"`
	TotalPolicies int `json:"totalPolicies"`
}

type AlertCount struct {
	Policies      []Policy    `json:"policies"`
	CountDetails  CountDetails `json:"countDetails"`
	NextPageToken string      `json:"nextPageToken"`
}