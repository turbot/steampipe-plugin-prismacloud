package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	prismacloud "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/model"
)

// This API is not documented.
// It was obtained by inspecting the Prisma Cloud console.
func ListInventoryDiscoveredAPI(c *prismacloud.Client, req map[string]interface{}) (*model.InventoryDiscoveredAPIResponse, error) {
	// https://api.anz.prismacloud.io/waas-api-discovery/api/v1/discovered-api
	c.Log(prismacloud.LogAction, "list of %s", "inventory api endpoints")
	var apis model.InventoryDiscoveredAPIResponse
	if _, err := c.Communicate("POST", []string{"waas-api-discovery", "api", "v1", "discovered-api"}, nil, req, &apis); err != nil {
		return nil, err
	}

	return &apis, nil
}

// Asset Inventory View V3 - GET
// https://pan.dev/prisma-cloud/api/cspm/asset-inventory-v-3/
func ListInventoryAsset(c *prismacloud.Client, query url.Values) (*model.GroupedAggregateAssetResponse, error) {
	c.Log(prismacloud.LogAction, "list of %s", "inventory api endpoints")

	var assets model.GroupedAggregateAssetResponse
	if _, err := c.Communicate("GET", []string{"v3", "inventory"}, query, nil, &assets); err != nil {
		return nil, err
	}

	return &assets, nil
}

func GetInventoryWorkloads(authToken string) (*model.InventoryWorkload, error) {
	// Define the URL and headers
	url := "https://asia-northeast1.cloud.twistlock.com/anz-3001938/api/v1/bff/assets/summary"
	contentType := "application/json"
	accept := "application/json"

	// Create a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Accept", accept)
	req.Header.Set("x-redlock-auth", authToken)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Unmarshal the response body into the ContainerImages struct
	var workload model.InventoryWorkload
	err = json.Unmarshal(body, &workload)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &workload, nil
}

func GetInventoryWorkloadContainerImages(authToken string, nextPageToken string) (*model.WorkloadContainerImagesResponse, error) {
	// Define the URL and headers
	url := "https://asia-northeast1.cloud.twistlock.com/anz-3001938/api/v1/bff/images/collated"
	contentType := "application/json"
	accept := "application/json"

	// Create the payload
	payload := map[string]interface{}{
		"stage":         "all",
		"sort":          "vulnerabilities",
		"limit":         30,
		"nextPageToken": nextPageToken,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	// Create a new request with the payload
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Accept", accept)
	req.Header.Set("x-redlock-auth", authToken)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Unmarshal the response body into the ContainerImages struct
	var cImages model.WorkloadContainerImagesResponse
	err = json.Unmarshal(body, &cImages)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &cImages, nil
}

func GetInventoryWorkloadHosts(authToken string, nextPageToken string) (*model.WorkloadContainerHostResponse, error) {
	// Define the URL and headers
	url := "https://asia-northeast1.cloud.twistlock.com/anz-3001938/api/v1/bff/hosts"
	contentType := "application/json"
	accept := "application/json"

	// Create the payload
	payload := map[string]interface{}{
		"sort":          "vulnerabilities",
		"limit":         30,
		"nextPageToken": nextPageToken,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	// Create a new request with the payload
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Accept", accept)
	req.Header.Set("x-redlock-auth", authToken)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Unmarshal the response body into the ContainerImages struct
	var hosts model.WorkloadContainerHostResponse
	err = json.Unmarshal(body, &hosts)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &hosts, nil
}
