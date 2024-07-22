package model


// ComplianceStandard represents the structure of the JSON data
type ComplianceStandard struct {
	CloudType            []string `json:"cloudType"`
	CreatedBy            string   `json:"createdBy"`
	CreatedOn            int64    `json:"createdOn"`
	Description          string   `json:"description"`
	ID                   string   `json:"id"`
	LastModifiedBy       string   `json:"lastModifiedBy"`
	LastModifiedOn       int64    `json:"lastModifiedOn"`
	Name                 string   `json:"name"`
	PoliciesAssignedCount int     `json:"policiesAssignedCount"`
	SystemDefault        bool     `json:"systemDefault"`
}