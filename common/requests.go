package anchorutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

func AnchorRequest[TRequest any, TResponse any](
	client *http.Client,
	method, url string,
	requestBody *TRequest,
	urlParams url.Values,
	bearer string,
) (*TResponse, error) {
	var responseBody TResponse
	var req *http.Request
	var err error

	// Create the request with or without request body
	if requestBody != nil {
		requestData, err := json.Marshal(requestBody)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, url, bytes.NewBuffer(requestData))
		req.Header.Set("Content-type", "application/json")
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return nil, err
	}

	// Add URL parameters if provided
	if len(urlParams) > 0 {
		req.URL.RawQuery = urlParams.Encode()
	}

	// Set default headers
	if bearer != "" {
		req.Header.Set("Authorization", "Bearer "+bearer)
	}
	req.Header.Set("Accept", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Read the entire response body for debugging

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Check if the response status code is successful
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var anchorErr AnchorError
		if err := json.Unmarshal(bodyBytes, &anchorErr); err != nil {
			return nil, fmt.Errorf("request failed with status code: %d, body: %s", resp.StatusCode, string(bodyBytes))
		}
		return nil, &anchorErr
	}

	// Create a new reader with the body bytes for JSON decoding
	bodyReader := bytes.NewReader(bodyBytes)

	// Decode the response body
	decoder := json.NewDecoder(bodyReader)
	if err := decoder.Decode(&responseBody); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v, body: %s", err, string(bodyBytes))
	}

	return &responseBody, nil
}

// AnchorFileUpload handles file upload requests as defined in SEP-12
func AnchorFileUpload[TResponse any](
	client *http.Client,
	url string,
	file io.Reader,
	filename string,
	bearer string,
) (*TResponse, error) {
	var responseBody TResponse
	// Create a buffer to store the multipart form data
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	// Create a form file field
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return nil, err
	}

	// Copy the file content to the form field
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	// Close the multipart writer to finalize the form data
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, &body)
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Accept", "application/json")
	if bearer != "" {
		req.Header.Set("Authorization", "Bearer "+bearer)
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Check if the response status code is successful
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var anchorErr AnchorError
		if err := json.Unmarshal(bodyBytes, &anchorErr); err != nil {
			return nil, fmt.Errorf("request failed with status code: %d, body: %s", resp.StatusCode, string(bodyBytes))
		}
		return nil, &anchorErr
	}

	// Decode the response body
	if err := json.Unmarshal(bodyBytes, &responseBody); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v, body: %s", err, string(bodyBytes))
	}

	return &responseBody, nil
}
