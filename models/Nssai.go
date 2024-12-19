/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:49:54 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Nssai struct {
	SupportedFeatures    string                          `json:"supportedFeatures,omitempty"`
	DefaultSingleNssais  []Snssai                        `json:"defaultSingleNssais"`
	SingleNssais         []Snssai                        `json:"singleNssais,omitempty"`
	ProvisioningTime     string                          `json:"provisioningTime,omitempty"`
	AdditionalSnssaiData map[string]AdditionalSnssaiData `json:"additionalSnssaiData,omitempty"`
	SuppressNssrgInd     *bool                           `json:"suppressNssrgInd,omitempty"`
}
