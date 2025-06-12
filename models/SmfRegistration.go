/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmfRegistration struct {
	UdrRestartInd               *bool              `json:"udrRestartInd,omitempty"`
	PgwFqdn                     string             `json:"pgwFqdn,omitempty"`
	PgwIpAddr                   *IpAddress         `json:"pgwIpAddr,omitempty"`
	ContextInfo                 *ContextInfo       `json:"contextInfo,omitempty"`
	DataRestorationCallbackUri  string             `json:"dataRestorationCallbackUri,omitempty"`
	PlmnId                      PlmnId             `json:"plmnId"`
	EpdgInd                     *bool              `json:"epdgInd,omitempty"`
	SingleNssai                 Snssai             `json:"singleNssai"`
	Dnn                         string             `json:"dnn,omitempty"`
	EmergencyServices           *bool              `json:"emergencyServices,omitempty"`
	PcscfRestorationCallbackUri string             `json:"pcscfRestorationCallbackUri,omitempty"`
	DeregCallbackUri            string             `json:"deregCallbackUri,omitempty"`
	RegistrationReason          RegistrationReason `json:"registrationReason,omitempty"`
	SupportedFeatures           string             `json:"supportedFeatures,omitempty"`
	PduSessionId                int                `json:"pduSessionId"`
	RegistrationTime            string             `json:"registrationTime,omitempty"`
	ResetIds                    []string           `json:"resetIds,omitempty"`
	PcfId                       string             `json:"pcfId,omitempty"`
	LastSynchronizationTime     string             `json:"lastSynchronizationTime,omitempty"`
	SmfInstanceId               string             `json:"smfInstanceId"`
	SmfSetId                    string             `json:"smfSetId,omitempty"`
}
