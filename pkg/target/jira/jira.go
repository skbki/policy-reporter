package jira

import (
	"encoding/json"

	"github.com/skbki/policy-reporter/pkg/crd/api/policyreport/v1alpha2"
	"github.com/skbki/policy-reporter/pkg/target"
	"github.com/skbki/policy-reporter/pkg/target/http"
)

// Options to configure the Discord target
type Options struct {
	target.ClientOptions
	Host         string
	Headers      map[string]string
	CustomFields map[string]string
	HTTPClient   http.Client
}

type client struct {
	target.BaseClient
	host         string
	headers      map[string]string
	customFields map[string]string
	client       http.Client
}

func EmbedJSONInData(inputStruct interface{}) (map[string]interface{}) {
	// Marshal the input struct to JSON.
	inputJSON, err := json.Marshal(inputStruct)
	if err != nil {
		return nil
	}

	// Parse the JSON into a map.
	var jsonData map[string]interface{}
	if err := json.Unmarshal(inputJSON, &jsonData); err != nil {
		return nil
	}

	// Create a new map with "data" key and embed the parsed JSON in it.
	result := map[string]interface{}{
		"data": jsonData,
	}

	return result
}


func (e *client) Send(result v1alpha2.PolicyReportResult) {
	if len(e.customFields) > 0 {
		props := make(map[string]string, 0)

		for property, value := range e.customFields {
			props[property] = value
		}

		for property, value := range result.Properties {
			props[property] = value
		}

		result.Properties = props
	}

	req, err := http.CreateJSONRequest(e.Name(), "POST", e.host, EmbedJSONInData(http.NewJSONResult(result)))
	if err != nil {
		return
	}

	for header, value := range e.headers {
		req.Header.Set(header, value)
	}

	resp, err := e.client.Do(req)
	http.ProcessHTTPResponse(e.Name(), resp, err)
}

// NewClient creates a new client to send Results to Jira
func NewClient(options Options) target.Client {
	return &client{
		target.NewBaseClient(options.ClientOptions),
		options.Host,
		options.Headers,
		options.CustomFields,
		options.HTTPClient,
	}
}
