package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type StageProductInput struct {
	ProductName    string `json:"name"`
	ProductVersion string `json:"product_version"`
}

type StagedProductsOutput struct {
	Products []StagedProduct
}

type StagedProduct struct {
	GUID string
	Type string
}

type UnstageProductInput struct {
	ProductName string `json:"name"`
}

type UpdateStagedProductPropertiesInput struct {
	GUID       string
	Properties string
}

type UpdateStagedProductNetworksAndAZsInput struct {
	GUID           string
	NetworksAndAZs string
}

type ResponseProperty struct {
	Value        interface{} `json:"value"`
	Configurable bool        `json:"configurable"`
	IsCredential bool        `json:"credential"`
}

type UpgradeRequest struct {
	ToVersion string `json:"to_version"`
}

type ConfigurationRequest struct {
	Method        string
	URL           string
	Configuration string
}

// TODO: extract to helper package?
func (a Api) Stage(input StageProductInput, deployedGUID string) error {
	stagedGUID, err := a.checkStagedProducts(input.ProductName)
	if err != nil {
		return err
	}

	var stReq *http.Request
	if deployedGUID == "" && stagedGUID == "" {
		stagedProductBody, err := json.Marshal(input)
		if err != nil {
			return err
		}

		stReq, err = http.NewRequest("POST", "/api/v0/staged/products", bytes.NewBuffer(stagedProductBody))
		if err != nil {
			return err
		}
	} else if deployedGUID != "" {
		upgradeReq := UpgradeRequest{
			ToVersion: input.ProductVersion,
		}

		upgradeReqBody, err := json.Marshal(upgradeReq)
		if err != nil {
			return err
		}

		stReq, err = http.NewRequest("PUT", fmt.Sprintf("/api/v0/staged/products/%s", deployedGUID), bytes.NewBuffer(upgradeReqBody))
		if err != nil {
			return err
		}
	} else if stagedGUID != "" {
		upgradeReq := UpgradeRequest{
			ToVersion: input.ProductVersion,
		}

		upgradeReqBody, err := json.Marshal(upgradeReq)
		if err != nil {
			return err
		}

		stReq, err = http.NewRequest("PUT", fmt.Sprintf("/api/v0/staged/products/%s", stagedGUID), bytes.NewBuffer(upgradeReqBody))
		if err != nil {
			return err
		}
	}

	stReq.Header.Set("Content-Type", "application/json")
	stResp, err := a.client.Do(stReq)
	if err != nil {
		return fmt.Errorf("could not make %s api request to staged products endpoint: %s", stReq.Method, err)
	}
	defer stResp.Body.Close()

	if err = validateStatusOK(stResp); err != nil {
		return err
	}

	return nil
}

func (a Api) DeleteStagedProduct(input UnstageProductInput) error {
	stagedGUID, err := a.checkStagedProducts(input.ProductName)
	if err != nil {
		return err
	}

	if len(stagedGUID) == 0 {
		return fmt.Errorf("product is not staged: %s", input.ProductName)
	}

	var req *http.Request
	req, err = http.NewRequest("DELETE", fmt.Sprintf("/api/v0/staged/products/%s", stagedGUID), strings.NewReader("{}"))

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.client.Do(req)
	if err != nil {
		return fmt.Errorf("could not make %s api request to staged products endpoint: %s", req.Method, err)
	}
	defer resp.Body.Close()

	if err = validateStatusOK(resp); err != nil {
		return err
	}

	return nil
}

func (a Api) ListStagedProducts() (StagedProductsOutput, error) {
	req, err := http.NewRequest("GET", "/api/v0/staged/products", nil)
	if err != nil {
		return StagedProductsOutput{}, err
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return StagedProductsOutput{}, fmt.Errorf("could not make request to staged-products endpoint: %s", err)
	}
	defer resp.Body.Close()

	if err = validateStatusOK(resp); err != nil {
		return StagedProductsOutput{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return StagedProductsOutput{}, err
	}

	var stagedProducts []StagedProduct
	err = json.Unmarshal(respBody, &stagedProducts)
	if err != nil {
		return StagedProductsOutput{}, fmt.Errorf("could not unmarshal staged products response: %s", err)
	}

	return StagedProductsOutput{
		Products: stagedProducts,
	}, nil
}

func (a Api) UpdateStagedProductProperties(input UpdateStagedProductPropertiesInput) error {
	body := bytes.NewBufferString(fmt.Sprintf(`{"properties": %s}`, input.Properties))
	req, err := http.NewRequest("PUT", fmt.Sprintf("/api/v0/staged/products/%s/properties", input.GUID), body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return fmt.Errorf("could not make api request to staged product properties endpoint: %s", err)
	}
	defer resp.Body.Close()

	if err = validateStatusOK(resp); err != nil {
		return err
	}

	return nil
}

func (a Api) UpdateStagedProductNetworksAndAZs(input UpdateStagedProductNetworksAndAZsInput) error {
	body := bytes.NewBufferString(fmt.Sprintf(`{"networks_and_azs": %s}`, input.NetworksAndAZs))
	req, err := http.NewRequest("PUT", fmt.Sprintf("/api/v0/staged/products/%s/networks_and_azs", input.GUID), body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return fmt.Errorf("could not make api request to staged product networks_and_azs endpoint: %s", err)
	}
	defer resp.Body.Close()

	if err = validateStatusOK(resp); err != nil {
		return err
	}

	return nil
}

//TODO consider refactoring to use fetchProductResource
func (a Api) GetStagedProductManifest(guid string) (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/v0/staged/products/%s/manifest", guid), nil)
	if err != nil {
		return "", err
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("could not make api request to staged products manifest endpoint: %s", err)
	}

	if err = validateStatusOK(resp); err != nil {
		return "", err
	}

	defer resp.Body.Close()
	var contents struct {
		Manifest interface{} `json:"manifest"`
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if err = yaml.Unmarshal(body, &contents); err != nil {
		return "", fmt.Errorf("could not parse json: %s", err)
	}

	manifest, err := yaml.Marshal(contents.Manifest)
	if err != nil {
		return "", err // this should never happen, all valid json can be marshalled
	}

	return string(manifest), nil
}

func (a Api) GetStagedProductProperties(product string) (map[string]ResponseProperty, error) {
	respBody, err := a.fetchProductResource(product, "properties")
	if err != nil {
		return nil, err
	}
	defer respBody.Close()

	propertiesResponse := struct {
		Properties map[string]ResponseProperty `json:"properties"`
	}{}
	if err = json.NewDecoder(respBody).Decode(&propertiesResponse); err != nil {
		return nil, fmt.Errorf("could not parse json: %s", err)
	}

	return propertiesResponse.Properties, nil
}

func (a Api) GetStagedProductNetworksAndAZs(product string) (map[string]interface{}, error) {
	respBody, err := a.fetchProductResource(product, "networks_and_azs")
	if err != nil {
		return nil, err
	}
	defer respBody.Close()

	networksResponse := struct {
		Networks map[string]interface{} `json:"networks_and_azs"`
	}{}
	if err = json.NewDecoder(respBody).Decode(&networksResponse); err != nil {
		return nil, fmt.Errorf("could not parse json: %s", err)
	}

	return networksResponse.Networks, nil
}

func (a Api) fetchProductResource(guid, endpoint string) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/v0/staged/products/%s/%s", guid, endpoint), nil)
	if err != nil {
		return nil, err // un-tested
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return nil,
			fmt.Errorf("could not make api request to staged product properties endpoint: %s", err)
	}

	if err = validateStatusOK(resp); err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func (a Api) checkStagedProducts(productName string) (string, error) {
	stagedProductsOutput, err := a.ListStagedProducts()
	if err != nil {
		return "", err
	}

	for _, product := range stagedProductsOutput.Products {
		if productName == product.Type {
			return product.GUID, nil
		}
	}

	return "", nil
}
