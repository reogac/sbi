/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmfRegistration struct {
	RegistrationTime            string             `json:"registrationTime,omitempty"`
	PcfId                       string             `json:"pcfId,omitempty"`
	SmfSetId                    string             `json:"smfSetId,omitempty"`
	EmergencyServices           *bool              `json:"emergencyServices,omitempty"`
	PcscfRestorationCallbackUri string             `json:"pcscfRestorationCallbackUri,omitempty"`
	UdrRestartInd               *bool              `json:"udrRestartInd,omitempty"`
	Dnn                         string             `json:"dnn,omitempty"`
	RegistrationReason          RegistrationReason `json:"registrationReason,omitempty"`
	ContextInfo                 *ContextInfo       `json:"contextInfo,omitempty"`
	SingleNssai                 Snssai             `json:"singleNssai"`
	PlmnId                      PlmnId             `json:"plmnId"`
	DeregCallbackUri            string             `json:"deregCallbackUri,omitempty"`
	DataRestorationCallbackUri  string             `json:"dataRestorationCallbackUri,omitempty"`
	ResetIds                    []string           `json:"resetIds,omitempty"`
	SmfInstanceId               string             `json:"smfInstanceId"`
	SupportedFeatures           string             `json:"supportedFeatures,omitempty"`
	PduSessionId                int                `json:"pduSessionId"`
	LastSynchronizationTime     string             `json:"lastSynchronizationTime,omitempty"`
	PgwFqdn                     string             `json:"pgwFqdn,omitempty"`
	PgwIpAddr                   *IpAddress         `json:"pgwIpAddr,omitempty"`
	EpdgInd                     *bool              `json:"epdgInd,omitempty"`
}
