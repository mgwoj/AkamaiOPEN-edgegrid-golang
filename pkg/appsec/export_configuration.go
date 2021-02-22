package appsec

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"time"
)

// ExportConfiguration represents a collection of ExportConfiguration
//
// See: ExportConfiguration.GetExportConfiguration()
// API Docs: // appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html

type (
	// ExportConfiguration  contains operations available on ExportConfiguration  resource
	// See: // appsec v1
	//
	// https://developer.akamai.com/api/cloud_security/application_security/v1.html#getexportconfiguration
	ExportConfiguration interface {
		GetExportConfigurations(ctx context.Context, params GetExportConfigurationsRequest) (*GetExportConfigurationsResponse, error)
	}

	ConditionsValue []string
	StringSlice     []string

	GetExportConfigurationsRequest struct {
		ConfigID int `json:"configId"`
		Version  int `json:"version"`
	}

	GetExportConfigurationsResponse struct {
		ConfigID   int    `json:"configId"`
		ConfigName string `json:"configName"`
		Version    int    `json:"version"`
		BasedOn    int    `json:"basedOn"`
		Staging    struct {
			Status string `json:"status"`
		} `json:"staging"`
		Production struct {
			Status string `json:"status"`
		} `json:"production"`
		CreateDate      time.Time `json:"createDate"`
		CreatedBy       string    `json:"createdBy"`
		SelectedHosts   []string  `json:"selectedHosts"`
		SelectableHosts []string  `json:"selectableHosts"`
		RatePolicies    []struct {
			AdditionalMatchOptions []struct {
				PositiveMatch bool     `json:"positiveMatch"`
				Type          string   `json:"type"`
				Values        []string `json:"values"`
			} `json:"additionalMatchOptions"`
			AllTraffic       bool                      `json:"allTraffic,omitempty"`
			AverageThreshold int                       `json:"averageThreshold"`
			BurstThreshold   int                       `json:"burstThreshold"`
			ClientIdentifier string                    `json:"clientIdentifier"`
			CreateDate       time.Time                 `json:"createDate"`
			Description      string                    `json:"description"`
			FileExtensions   *RatePolicyFileExtensions `json:"fileExtensions"`
			ID               int                       `json:"id"`
			MatchType        string                    `json:"matchType"`
			Name             string                    `json:"name"`
			Path             struct {
				PositiveMatch bool                    `json:"positiveMatch"`
				Values        *RatePoliciesPathValues `json:"values,omitempty"`
			} `json:"path"`
			PathMatchType         string                       `json:"pathMatchType"`
			PathURIPositiveMatch  bool                         `json:"pathUriPositiveMatch"`
			QueryParameters       *RatePoliciesQueryParameters `json:"queryParameters,omitempty"`
			RequestType           string                       `json:"requestType"`
			SameActionOnIpv6      bool                         `json:"sameActionOnIpv6"`
			Type                  string                       `json:"type"`
			UpdateDate            time.Time                    `json:"updateDate"`
			UseXForwardForHeaders bool                         `json:"useXForwardForHeaders"`
			Used                  bool                         `json:"used"`
		} `json:"ratePolicies"`
		ReputationProfiles []struct {
			Condition        *ConditionReputationProfile `json:"condition,omitempty"`
			Context          string                      `json:"context"`
			ContextReadable  string                      `json:"contextReadable"`
			Enabled          bool                        `json:"enabled"`
			ID               int                         `json:"id"`
			Name             string                      `json:"name"`
			SharedIPHandling string                      `json:"sharedIpHandling"`
			Threshold        int                         `json:"threshold"`
		} `json:"reputationProfiles"`
		CustomRules []struct {
			Conditions    *ConditionsExp `json:"conditions,omitempty"`
			Description   string         `json:"description,omitempty"`
			ID            int            `json:"id"`
			Name          string         `json:"name"`
			RuleActivated bool           `json:"ruleActivated"`
			Structured    bool           `json:"structured"`
			Tag           []string       `json:"tag"`
			Version       int            `json:"version"`
		} `json:"customRules"`
		Rulesets []struct {
			ID               int       `json:"id"`
			RulesetVersionID int       `json:"rulesetVersionId"`
			Type             string    `json:"type"`
			ReleaseDate      time.Time `json:"releaseDate"`
			Rules            []struct {
				ID                  int      `json:"id"`
				InspectRequestBody  bool     `json:"inspectRequestBody"`
				InspectResponseBody bool     `json:"inspectResponseBody"`
				Outdated            bool     `json:"outdated"`
				RuleVersion         int      `json:"ruleVersion"`
				Score               int      `json:"score"`
				Tag                 string   `json:"tag"`
				Title               string   `json:"title"`
				AttackGroups        []string `json:"attackGroups,omitempty"`
			} `json:"rules"`
			AttackGroups []struct {
				Group     string `json:"group"`
				GroupName string `json:"groupName"`
				Threshold int    `json:"threshold"`
			} `json:"attackGroups,omitempty"`
		} `json:"rulesets"`
		MatchTargets struct {
			APITargets []struct {
				Sequence int    `json:"sequence,omitempty"`
				ID       int    `json:"id,omitempty"`
				Type     string `json:"type,omitempty"`
				Apis     []struct {
					ID   int    `json:"id,omitempty"`
					Name string `json:"name,omitempty"`
				} `json:"apis,omitempty"`
				EffectiveSecurityControls struct {
					ApplyAPIConstraints           bool `json:"applyApiConstraints"`
					ApplyApplicationLayerControls bool `json:"applyApplicationLayerControls"`
					ApplyBotmanControls           bool `json:"applyBotmanControls,omitempty"`
					ApplyNetworkLayerControls     bool `json:"applyNetworkLayerControls"`
					ApplyRateControls             bool `json:"applyRateControls"`
					ApplyReputationControls       bool `json:"applyReputationControls"`
					ApplySlowPostControls         bool `json:"applySlowPostControls"`
				} `json:"effectiveSecurityControls,omitempty"`
				SecurityPolicy struct {
					PolicyID string `json:"policyId,omitempty"`
				} `json:"securityPolicy,omitempty"`
				BypassNetworkLists []struct {
					Name string `json:"name,omitempty"`
					ID   string `json:"id,omitempty"`
				} `json:"bypassNetworkLists,omitempty"`
			} `json:"apiTargets,omitempty"`
			WebsiteTargets []struct {
				Type               string `json:"type"`
				BypassNetworkLists []struct {
					ID   string `json:"id"`
					Name string `json:"name"`
				} `json:"bypassNetworkLists,omitempty"`
				DefaultFile               string `json:"defaultFile"`
				EffectiveSecurityControls struct {
					ApplyApplicationLayerControls bool `json:"applyApplicationLayerControls"`
					ApplyNetworkLayerControls     bool `json:"applyNetworkLayerControls"`
					ApplyRateControls             bool `json:"applyRateControls"`
					ApplyReputationControls       bool `json:"applyReputationControls"`
					ApplySlowPostControls         bool `json:"applySlowPostControls"`
				} `json:"effectiveSecurityControls"`
				FilePaths                    []string `json:"filePaths"`
				FileExtensions               []string `json:"fileExtensions,omitempty"`
				Hostnames                    []string `json:"hostnames,omitempty"`
				ID                           int      `json:"id"`
				IsNegativeFileExtensionMatch bool     `json:"isNegativeFileExtensionMatch"`
				IsNegativePathMatch          bool     `json:"isNegativePathMatch"`
				SecurityPolicy               struct {
					PolicyID string `json:"policyId"`
				} `json:"securityPolicy"`
				Sequence int `json:"sequence"`
			} `json:"websiteTargets"`
		} `json:"matchTargets"`
		SecurityPolicies []struct {
			ID                      string `json:"id"`
			Name                    string `json:"name"`
			HasRatePolicyWithAPIKey bool   `json:"hasRatePolicyWithApiKey"`
			SecurityControls        struct {
				ApplyAPIConstraints           bool `json:"applyApiConstraints"`
				ApplyApplicationLayerControls bool `json:"applyApplicationLayerControls"`
				ApplyBotmanControls           bool `json:"applyBotmanControls"`
				ApplyNetworkLayerControls     bool `json:"applyNetworkLayerControls"`
				ApplyRateControls             bool `json:"applyRateControls"`
				ApplyReputationControls       bool `json:"applyReputationControls"`
				ApplySlowPostControls         bool `json:"applySlowPostControls"`
			} `json:"securityControls"`
			WebApplicationFirewall struct {
				RuleActions []struct {
					Action           string `json:"action"`
					ID               int    `json:"id"`
					RulesetVersionID int    `json:"rulesetVersionId"`
					Conditions       []struct {
						Type          string   `json:"type,omitempty"`
						Extensions    []string `json:"extensions,omitempty"`
						PositiveMatch bool     `json:"positiveMatch"`
						Filenames     []string `json:"filenames,omitempty"`
						Hosts         []string `json:"hosts,omitempty"`
						Ips           []string `json:"ips,omitempty"`
						UseHeaders    bool     `json:"useHeaders,omitempty"`
						CaseSensitive bool     `json:"caseSensitive,omitempty"`
						Name          string   `json:"name,omitempty"`
						NameCase      bool     `json:"nameCase,omitempty"`
						Value         string   `json:"value,omitempty"`
						Wildcard      bool     `json:"wildcard,omitempty"`
						Header        string   `json:"header,omitempty"`
						ValueCase     bool     `json:"valueCase,omitempty"`
						ValueWildcard bool     `json:"valueWildcard,omitempty"`
						Methods       []string `json:"methods,omitempty"`
						Paths         []string `json:"paths,omitempty"`
					} `json:"conditions,omitempty"`
					Exception *ExceptioneExpruleaction `json:"exception,omitempty"`
				} `json:"ruleActions"`
				AttackGroupActions []struct {
					Action             string                 `json:"action"`
					Group              string                 `json:"group"`
					RulesetVersionID   int                    `json:"rulesetVersionId"`
					AdvancedExceptions *AdvancedExceptionsexp `json:"advancedExceptions,omitempty"`
					Exception          *AttackGroupException  `json:"exception,omitempty"`
				} `json:"attackGroupActions"`
				Evaluation *Evaluationexp `json:"evaluation,omitempty"`
			} `json:"webApplicationFirewall"`
			CustomRuleActions []struct {
				Action string `json:"action"`
				ID     int    `json:"id"`
			} `json:"customRuleActions,omitempty"`
			APIRequestConstraints *APIRequestConstraintsexp `json:"apiRequestConstraints,omitempty"`
			ClientReputation      struct {
				ReputationProfileActions []struct {
					Action string `json:"action"`
					ID     int    `json:"id"`
				} `json:"reputationProfileActions"`
			} `json:"clientReputation"`
			RatePolicyActions []struct {
				ID         int    `json:"id"`
				Ipv4Action string `json:"ipv4Action"`
				Ipv6Action string `json:"ipv6Action"`
			} `json:"ratePolicyActions"`
			IPGeoFirewall struct {
				Block       string `json:"block"`
				GeoControls struct {
					BlockedIPNetworkLists struct {
						NetworkList []string `json:"networkList"`
					} `json:"blockedIPNetworkLists"`
				} `json:"geoControls"`
				IPControls struct {
					AllowedIPNetworkLists struct {
						NetworkList []string `json:"networkList"`
					} `json:"allowedIPNetworkLists"`
					BlockedIPNetworkLists struct {
						NetworkList []string `json:"networkList"`
					} `json:"blockedIPNetworkLists"`
				} `json:"ipControls"`
			} `json:"ipGeoFirewall,omitempty"`
			PenaltyBox struct {
				Action               string `json:"action"`
				PenaltyBoxProtection bool   `json:"penaltyBoxProtection"`
			} `json:"penaltyBox,omitempty"`
			SlowPost         *SlowPostexp         `json:"slowPost,omitempty"`
			LoggingOverrides *LoggingOverridesexp `json:"loggingOverrides,omitempty"`
		} `json:"securityPolicies"`
		Siem            *Siemexp            `json:"siem,omitempty"`
		AdvancedOptions *AdvancedOptionsexp `json:"advancedOptions,omitempty"`
		CustomDenyList  *CustomDenyListexp  `json:"customDenyList,omitempty"`
	}

	//TODO fails export
	ReputationProfileActionsexp []struct {
		Action string `json:"action"`
		ID     int    `json:"id"`
	}
	//TODO fails export
	RatePolicyActionsexp []struct {
		ID         int    `json:"id"`
		Ipv4Action string `json:"ipv4Action"`
		Ipv6Action string `json:"ipv6Action"`
	}
	SlowPostexp struct {
		Action            string `json:"action"`
		SlowRateThreshold struct {
			Period int `json:"period"`
			Rate   int `json:"rate"`
		} `json:"slowRateThreshold"`
		DurationThreshold struct {
			Timeout int `json:"timeout"`
		} `json:"durationThreshold"`
	}
	AdvancedOptionsexp struct {
		Logging  *Loggingexp `json:"logging"`
		Prefetch struct {
			AllExtensions      bool     `json:"allExtensions"`
			EnableAppLayer     bool     `json:"enableAppLayer"`
			EnableRateControls bool     `json:"enableRateControls"`
			Extensions         []string `json:"extensions"`
		} `json:"prefetch"`
	}
	CustomDenyListexp []struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Parameters []struct {
			DisplayName string `json:"displayName"`
			Name        string `json:"name"`
			Value       string `json:"value"`
		} `json:"parameters"`
	}
	//TODO breaks export
	CustomRuleActionsexp []struct {
		Action string `json:"action"`
		ID     int    `json:"id"`
	}
	Siemexp struct {
		EnableForAllPolicies    bool     `json:"enableForAllPolicies,omitempty"`
		EnableSiem              bool     `json:"enableSiem"`
		EnabledBotmanSiemEvents bool     `json:"enabledBotmanSiemEvents,omitempty"`
		FirewallPolicyIds       []string `json:"firewallPolicyIds,omitempty"`
		SiemDefinitionID        int      `json:"siemDefinitionId,omitempty"`
	}
	//TODO nfails export
	PenaltyBoxexp struct {
		Action               string `json:"action"`
		PenaltyBoxProtection bool   `json:"penaltyBoxProtection"`
	}
	APIRequestConstraintsexp struct {
		Action       string `json:"action"`
		APIEndpoints []struct {
			Action string `json:"action"`
			ID     int    `json:"id"`
		} `json:"apiEndpoints"`
	}
	Evaluationexp struct {
		AttackGroupActions []struct {
			Action string `json:"action"`
			Group  string `json:"group"`
		} `json:"attackGroupActions"`
		EvaluationID      int `json:"evaluationId"`
		EvaluationVersion int `json:"evaluationVersion"`
		RuleActions       []struct {
			Action     string `json:"action"`
			ID         int    `json:"id"`
			Conditions []struct {
				Type          string   `json:"type,omitempty"`
				Extensions    []string `json:"extensions,omitempty"`
				PositiveMatch bool     `json:"positiveMatch"`
				Filenames     []string `json:"filenames,omitempty"`
				Hosts         []string `json:"hosts,omitempty"`
				Ips           []string `json:"ips,omitempty"`
				UseHeaders    bool     `json:"useHeaders,omitempty"`
				CaseSensitive bool     `json:"caseSensitive,omitempty"`
				Name          string   `json:"name,omitempty"`
				NameCase      bool     `json:"nameCase,omitempty"`
				Value         string   `json:"value,omitempty"`
				Wildcard      bool     `json:"wildcard,omitempty"`
				Header        string   `json:"header,omitempty"`
				ValueCase     bool     `json:"valueCase,omitempty"`
				ValueWildcard bool     `json:"valueWildcard,omitempty"`
				Methods       []string `json:"methods,omitempty"`
				Paths         []string `json:"paths,omitempty"`
			} `json:"conditions,omitempty"`
			Exception *ExceptioneExpEvalruleaction `json:"exception,omitempty"`
		} `json:"ruleActions"`
		RulesetVersionID int `json:"rulesetVersionId"`
	}
	ConditionReputationProfile struct {
		AtomicConditions *AtomicConditionsexp `json:"atomicConditions,omitempty"`
		CanDelete        bool                 `json:"-"`
		ConfigVersionID  int                  `json:"-"`
		ID               int                  `json:"-"`
		Name             string               `json:"-"`
		PositiveMatch    bool                 `json:"positiveMatch"`
		UUID             string               `json:"-"`
		Version          int64                `json:"-"`
	}
	ExceptioneExpEvalruleaction struct {
		AnyHeaderCookieOrParam           []string `json:"anyHeaderCookieOrParam,omitempty"`
		HeaderCookieOrParamValues        []string `json:"headerCookieOrParamValues,omitempty"`
		SpecificHeaderCookieOrParamNames []struct {
			Names    []string `json:"names,omitempty"`
			Selector string   `json:"selector,omitempty"`
		} `json:"specificHeaderCookieOrParamNames,omitempty"`
		SpecificHeaderCookieOrParamPrefix    *AttackGroupSpecificHeaderCookieOrParamPrefix `json:"specificHeaderCookieOrParamPrefix,omitempty"`
		SpecificHeaderCookieOrParamNameValue *SpecificHeaderCookieOrParamNameValueexp      `json:"specificHeaderCookieOrParamNameValue,omitempty"`
	}

	ExceptioneExpruleaction struct {
		AnyHeaderCookieOrParam               []string                                      `json:"anyHeaderCookieOrParam,omitempty"`
		HeaderCookieOrParamValues            []string                                      `json:"headerCookieOrParamValues,omitempty"`
		SpecificHeaderCookieOrParamNames     *AttackGroupSpecificHeaderCookieOrParamNames  `json:"specificHeaderCookieOrParamNames,omitempty"`
		SpecificHeaderCookieOrParamPrefix    *AttackGroupSpecificHeaderCookieOrParamPrefix `json:"specificHeaderCookieOrParamPrefix,omitempty"`
		SpecificHeaderCookieOrParamNameValue *SpecificHeaderCookieOrParamNameValueexp      `json:"specificHeaderCookieOrParamNameValue,omitempty"`
	}

	HeaderCookieOrParamValuesattackgroup []struct {
		Criteria []struct {
			Hostnames []string `json:"hostnames,omitempty"`
			Paths     []string `json:"paths,omitempty"`
			Values    []string `json:"values,omitempty"`
		} `json:"criteria"`
		ValueWildcard bool     `json:"valueWildcard,omitempty"`
		Values        []string `json:"values,omitempty"`
	}

	AdvancedExceptionsexp struct {
		Conditions                              *ConditionsAttackGroup                              `json:"conditions,omitempty"`
		HeaderCookieOrParamValues               *HeaderCookieOrParamValuesattackgroup               `json:"headerCookieOrParamValues,omitempty"`
		SpecificHeaderCookieOrParamNameValue    *SpecificHeaderCookieOrParamNameValueAttackGroups   `json:"specificHeaderCookieOrParamNameValue,omitempty"`
		SpecificHeaderCookieParamXMLOrJSONNames *AttackGroupSpecificHeaderCookieParamXMLOrJSONNames `json:"specificHeaderCookieParamXmlOrJsonNames,omitempty"`
	}

	SpecificHeaderCookieOrParamNameValueAttackGroups []struct {
		NamesValues []struct {
			Names  []string `json:"names"`
			Values []string `json:"values"`
		} `json:"namesValues"`
		Criteria      *AttackGroupCriteria `json:"criteria,omitempty"`
		Selector      string               `json:"selector"`
		ValueWildcard bool                 `json:"valueWildcard"`
		Wildcard      bool                 `json:"wildcard"`
	}
	SpecificHeaderCookieOrParamNameValueexp struct {
		Name     *json.RawMessage `json:"name,omitempty"`
		Selector string           `json:"selector,omitempty"`
		Value    *json.RawMessage `json:"value,omitempty"`
	}

	AtomicConditionsexp []struct {
		CheckIps      string           `json:"checkIps,omitempty"`
		ClassName     string           `json:"className"`
		Index         int              `json:"index"`
		PositiveMatch bool             `json:"positiveMatch,omitempty"`
		Value         []string         `json:"value,omitempty"`
		Name          *json.RawMessage `json:"name,omitempty"`
		NameCase      bool             `json:"nameCase,omitempty"`
		NameWildcard  bool             `json:"nameWildcard,omitempty"`
		ValueCase     bool             `json:"valueCase,omitempty"`
		ValueWildcard bool             `json:"valueWildcard,omitempty"`
		Host          []string         `json:"host,omitempty"`
	}

	Loggingexp struct {
		AllowSampling bool `json:"allowSampling"`
		Cookies       struct {
			Type   string   `json:"type"`
			Values []string `json:"values,omitempty"`
		} `json:"cookies"`
		CustomHeaders struct {
			Type   string   `json:"type"`
			Values []string `json:"values,omitempty"`
		} `json:"customHeaders"`
		StandardHeaders struct {
			Type   string   `json:"type"`
			Values []string `json:"values,omitempty"`
		} `json:"standardHeaders"`
	}

	LoggingOverridesexp struct {
		AllowSampling bool `json:"allowSampling"`
		Cookies       struct {
			Type   string   `json:"type"`
			Values []string `json:"values,omitempty"`
		} `json:"cookies"`
		CustomHeaders struct {
			Type   string   `json:"type"`
			Values []string `json:"values,omitempty"`
		} `json:"customHeaders"`
		Override        bool `json:"override"`
		StandardHeaders struct {
			Type   string   `json:"type"`
			Values []string `json:"values,omitempty"`
		} `json:"standardHeaders"`
	}

	ConditionsExp []struct {
		Type          string           `json:"type,omitempty"`
		PositiveMatch bool             `json:"positiveMatch"`
		Name          *json.RawMessage `json:"name,omitempty"`
		Value         *json.RawMessage `json:"value,omitempty"`
		ValueCase     bool             `json:"valueCase,omitempty"`
		ValueWildcard bool             `json:"valueWildcard,omitempty"`
	}

	RatePoliciesFileExtensionsValues []string

	RatePoliciesPathValues []string

	RatePoliciesQueryParameters []struct {
		Name          string                             `json:"name"`
		PositiveMatch bool                               `json:"positiveMatch"`
		ValueInRange  bool                               `json:"valueInRange"`
		Values        *RatePoliciesQueryParametersValues `json:"values,omitempty"`
	}

	RatePoliciesQueryParametersValues []string
)

func (c *ConditionsValue) UnmarshalJSON(data []byte) error {
	var nums interface{}
	err := json.Unmarshal(data, &nums)
	if err != nil {
		return err
	}

	items := reflect.ValueOf(nums)
	switch items.Kind() {
	case reflect.String:
		*c = append(*c, items.String())

	case reflect.Slice:
		*c = make(ConditionsValue, 0, items.Len())
		for i := 0; i < items.Len(); i++ {
			item := items.Index(i)
			switch item.Kind() {
			case reflect.String:
				*c = append(*c, item.String())
			case reflect.Interface:
				*c = append(*c, item.Interface().(string))
			}
		}
	}
	return nil
}

func (p *appsec) GetExportConfigurations(ctx context.Context, params GetExportConfigurationsRequest) (*GetExportConfigurationsResponse, error) {

	logger := p.Log(ctx)
	logger.Debug("GetExportConfigurations")

	var rval GetExportConfigurationsResponse

	uri := fmt.Sprintf(
		"/appsec/v1/export/configs/%d/versions/%d",
		params.ConfigID,
		params.Version)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create getexportconfigurations request: %w", err)
	}

	resp, err := p.Exec(req, &rval)
	if err != nil {
		return nil, fmt.Errorf("getexportconfigurations request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	return &rval, nil

}
