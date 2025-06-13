/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:51 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmfRegistration struct {
	PduSessionId                int                `json:"pduSessionId"`
	DeregCallbackUri            string             `json:"deregCallbackUri,omitempty"`
	PcfId                       string             `json:"pcfId,omitempty"`
	SmfSetId                    string             `json:"smfSetId,omitempty"`
	SupportedFeatures           string             `json:"supportedFeatures,omitempty"`
	PcscfRestorationCallbackUri string             `json:"pcscfRestorationCallbackUri,omitempty"`
	EpdgInd                     *bool              `json:"epdgInd,omitempty"`
	DataRestorationCallbackUri  string             `json:"dataRestorationCallbackUri,omitempty"`
	LastSynchronizationTime     string             `json:"lastSynchronizationTime,omitempty"`
	SingleNssai                 Snssai             `json:"singleNssai"`
	PgwFqdn                     string             `json:"pgwFqdn,omitempty"`
	PgwIpAddr                   *IpAddress         `json:"pgwIpAddr,omitempty"`
	RegistrationReason          RegistrationReason `json:"registrationReason,omitempty"`
	ContextInfo                 *ContextInfo       `json:"contextInfo,omitempty"`
	SmfInstanceId               string             `json:"smfInstanceId"`
	Dnn                         string             `json:"dnn,omitempty"`
	EmergencyServices           *bool              `json:"emergencyServices,omitempty"`
	PlmnId                      PlmnId             `json:"plmnId"`
	RegistrationTime            string             `json:"registrationTime,omitempty"`
	ResetIds                    []string           `json:"resetIds,omitempty"`
	UdrRestartInd               *bool              `json:"udrRestartInd,omitempty"`
}
