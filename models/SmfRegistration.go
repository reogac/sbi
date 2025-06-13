/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmfRegistration struct {
	RegistrationTime            string             `json:"registrationTime,omitempty"`
	SmfInstanceId               string             `json:"smfInstanceId"`
	Dnn                         string             `json:"dnn,omitempty"`
	PlmnId                      PlmnId             `json:"plmnId"`
	RegistrationReason          RegistrationReason `json:"registrationReason,omitempty"`
	EmergencyServices           *bool              `json:"emergencyServices,omitempty"`
	PgwFqdn                     string             `json:"pgwFqdn,omitempty"`
	PgwIpAddr                   *IpAddress         `json:"pgwIpAddr,omitempty"`
	LastSynchronizationTime     string             `json:"lastSynchronizationTime,omitempty"`
	PcfId                       string             `json:"pcfId,omitempty"`
	UdrRestartInd               *bool              `json:"udrRestartInd,omitempty"`
	PduSessionId                int                `json:"pduSessionId"`
	SingleNssai                 Snssai             `json:"singleNssai"`
	EpdgInd                     *bool              `json:"epdgInd,omitempty"`
	ContextInfo                 *ContextInfo       `json:"contextInfo,omitempty"`
	DataRestorationCallbackUri  string             `json:"dataRestorationCallbackUri,omitempty"`
	ResetIds                    []string           `json:"resetIds,omitempty"`
	SmfSetId                    string             `json:"smfSetId,omitempty"`
	SupportedFeatures           string             `json:"supportedFeatures,omitempty"`
	PcscfRestorationCallbackUri string             `json:"pcscfRestorationCallbackUri,omitempty"`
	DeregCallbackUri            string             `json:"deregCallbackUri,omitempty"`
}
