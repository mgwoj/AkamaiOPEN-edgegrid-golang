package appsec

import (
	"context"
	"fmt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	// The WAFMode interface supports retrieving and modifying the mode setting that determines how
	// rule sets are upgraded.
	//
	// https://developer.akamai.com/api/cloud_security/application_security/v1.html#mode
	WAFMode interface {
		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#getmode
		// Deprecated: this method will be removed in a future release. Use GetWAFMode instead.
		GetWAFModes(ctx context.Context, params GetWAFModesRequest) (*GetWAFModesResponse, error)

		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#getmode
		GetWAFMode(ctx context.Context, params GetWAFModeRequest) (*GetWAFModeResponse, error)

		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#putmode
		UpdateWAFMode(ctx context.Context, params UpdateWAFModeRequest) (*UpdateWAFModeResponse, error)
	}

	// GetWAFModesRequest is used to retrieve the setting that determines this mode how rules will be kept up to date.
	// Deprecated: this struct will be removed in a future release.
	GetWAFModesRequest struct {
		ConfigID int    `json:"-"`
		Version  int    `json:"-"`
		PolicyID string `json:"-"`
		Current  string `json:"current"`
		Mode     string `json:"mode"`
		Eval     string `json:"eval"`
	}

	// GetWAFModesResponse is returned from a call to GetWAFModes.
	// Deprecated: this struct will be removed in a future release.
	GetWAFModesResponse struct {
		Current    string `json:"current,omitempty"`
		Mode       string `json:"mode,omitempty"`
		Eval       string `json:"eval,omitempty"`
		Evaluating string `json:"evaluating,omitempty"`
		Expires    string `json:"expires,omitempty"`
	}

	// GetWAFModeRequest is used to retrieve the setting that determines this mode how rules will be kept up to date.
	GetWAFModeRequest struct {
		ConfigID int    `json:"-"`
		Version  int    `json:"-"`
		PolicyID string `json:"-"`
		Current  string `json:"current"`
		Mode     string `json:"mode"`
		Eval     string `json:"eval"`
	}

	// GetWAFModeResponse is returned from a call to GetWAFMode.
	GetWAFModeResponse struct {
		Current    string `json:"current,omitempty"`
		Mode       string `json:"mode,omitempty"`
		Eval       string `json:"eval,omitempty"`
		Evaluating string `json:"evaluating,omitempty"`
		Expires    string `json:"expires,omitempty"`
	}

	// UpdateWAFModeRequest is used to modify the setting that determines this mode how rules will be kept up to date.
	UpdateWAFModeRequest struct {
		ConfigID int    `json:"-"`
		Version  int    `json:"-"`
		PolicyID string `json:"-"`
		Current  string `json:"-"`
		Mode     string `json:"mode"`
		Eval     string `json:"-"`
	}

	// UpdateWAFModeResponse is returned from a call to UpdateWAFMode.
	UpdateWAFModeResponse struct {
		Current string `json:"current"`
		Mode    string `json:"mode"`
	}
)

// Validate validates a GetWAFModeRequest.
func (v GetWAFModeRequest) Validate() error {
	return validation.Errors{
		"ConfigID": validation.Validate(v.ConfigID, validation.Required),
		"Version":  validation.Validate(v.Version, validation.Required),
		"PolicyID": validation.Validate(v.PolicyID, validation.Required),
	}.Filter()
}

// Validate validates a GetWAFModesRequest.
// Deprecated: this method will be removed in a future release.
func (v GetWAFModesRequest) Validate() error {
	return validation.Errors{
		"ConfigID": validation.Validate(v.ConfigID, validation.Required),
		"Version":  validation.Validate(v.Version, validation.Required),
		"PolicyID": validation.Validate(v.PolicyID, validation.Required),
	}.Filter()
}

// Validate validates an UpdateWAFModeRequest.
func (v UpdateWAFModeRequest) Validate() error {
	return validation.Errors{
		"ConfigID": validation.Validate(v.ConfigID, validation.Required),
		"Version":  validation.Validate(v.Version, validation.Required),
		"PolicyID": validation.Validate(v.PolicyID, validation.Required),
	}.Filter()
}

func (p *appsec) GetWAFMode(ctx context.Context, params GetWAFModeRequest) (*GetWAFModeResponse, error) {
	logger := p.Log(ctx)
	logger.Debug("GetWAFMode")

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	uri := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/security-policies/%s/mode",
		params.ConfigID,
		params.Version,
		params.PolicyID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GetWAFMode request: %w", err)
	}

	var result GetWAFModeResponse
	resp, err := p.Exec(req, &result)
	if err != nil {
		return nil, fmt.Errorf("get WAF mode request failed: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	return &result, nil
}

// Deprecated: this method will be removed in a future release.
func (p *appsec) GetWAFModes(ctx context.Context, params GetWAFModesRequest) (*GetWAFModesResponse, error) {
	logger := p.Log(ctx)
	logger.Debug("GetWAFModes")

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	uri := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/security-policies/%s/mode",
		params.ConfigID,
		params.Version,
		params.PolicyID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GetWAFModes request: %w", err)
	}

	var result GetWAFModesResponse
	resp, err := p.Exec(req, &result)
	if err != nil {
		return nil, fmt.Errorf("get WAF Modes request failed: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	return &result, nil
}

func (p *appsec) UpdateWAFMode(ctx context.Context, params UpdateWAFModeRequest) (*UpdateWAFModeResponse, error) {
	logger := p.Log(ctx)
	logger.Debug("UpdateWAFMode")

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	uri := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/security-policies/%s/mode",
		params.ConfigID,
		params.Version,
		params.PolicyID,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create UpdateWAFMode request: %w", err)
	}

	var result UpdateWAFModeResponse
	resp, err := p.Exec(req, &result, params)
	if err != nil {
		return nil, fmt.Errorf("update WAF mode request failed: %w", err)
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, p.Error(resp)
	}

	return &result, nil
}
