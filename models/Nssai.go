/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Nssai struct {
	ProvisioningTime     string                          `json:"provisioningTime,omitempty"`
	AdditionalSnssaiData map[string]AdditionalSnssaiData `json:"additionalSnssaiData,omitempty"`
	SuppressNssrgInd     *bool                           `json:"suppressNssrgInd,omitempty"`
	SupportedFeatures    string                          `json:"supportedFeatures,omitempty"`
	DefaultSingleNssais  []Snssai                        `json:"defaultSingleNssais"`
	SingleNssais         []Snssai                        `json:"singleNssais,omitempty"`
}
