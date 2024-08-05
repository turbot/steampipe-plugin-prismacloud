package model

type UserProfile struct {
	AccessKeysAllowed bool      `json:"accessKeysAllowed"`
	ActiveRole        Role      `json:"activeRole"`
	DefaultRoleID     string    `json:"defaultRoleId"`
	DisplayName       string    `json:"displayName"`
	Email             string    `json:"email"`
	Enabled           bool      `json:"enabled"`
	FirstName         string    `json:"firstName"`
	LastLoginTs       int64     `json:"lastLoginTs"`
	LastModifiedBy    string    `json:"lastModifiedBy"`
	LastModifiedTs    int64     `json:"lastModifiedTs"`
	LastName          string    `json:"lastName"`
	RoleIDs           []string  `json:"roleIds"`
	Roles             []Role    `json:"roles"`
	TimeZone          string    `json:"timeZone"`
}

type Role struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	OnlyAllowCIAccess   bool   `json:"onlyAllowCIAccess"`
	OnlyAllowComputeAccess bool `json:"onlyAllowComputeAccess"`
	OnlyAllowReadAccess bool   `json:"onlyAllowReadAccess"`
	Type                string `json:"type"`
}