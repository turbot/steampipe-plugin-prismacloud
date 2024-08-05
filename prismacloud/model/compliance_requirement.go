package model

// ComplianceRequirement represents the structure of the JSON response
// https://pan.dev/prisma-cloud/api/cspm/get-requirement-by-id/#responses
type ComplianceRequirement struct {
	ComplianceID          string `json:"complianceId"`
	CreatedBy             string `json:"createdBy"`
	CreatedOn             int64  `json:"createdOn"`
	Description           string `json:"description"`
	ID                    string `json:"id"`
	LastModifiedBy        string `json:"lastModifiedBy"`
	LastModifiedOn        int64  `json:"lastModifiedOn"`
	Name                  string `json:"name"`
	PoliciesAssignedCount int    `json:"policiesAssignedCount"`
	RequirementID         string `json:"requirementId"`
	StandardName          string `json:"standardName"`
	SystemDefault         bool   `json:"systemDefault"`
	ViewOrder             int    `json:"viewOrder"`
}

// ComplianceRequirementSection represents the structure of the JSON data
// https://pan.dev/prisma-cloud/api/cspm/get-sections/#responses
type ComplianceRequirementSection struct {
	AssociatedPolicyIDs   []string `json:"associatedPolicyIds"`
	CreatedBy             string   `json:"createdBy"`
	CreatedOn             int64    `json:"createdOn"`
	Description           string   `json:"description"`
	ID                    string   `json:"id"`
	Label                 string   `json:"label"`
	LastModifiedBy        string   `json:"lastModifiedBy"`
	LastModifiedOn        int64    `json:"lastModifiedOn"`
	PoliciesAssignedCount int      `json:"policiesAssignedCount"`
	RequirementID         string   `json:"requirementId"`
	RequirementName       string   `json:"requirementName"`
	SectionID             string   `json:"sectionId"`
	StandardName          string   `json:"standardName"`
	SystemDefault         bool     `json:"systemDefault"`
	ViewOrder             int      `json:"viewOrder"`
}
