/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:33 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmfRegistration struct {
	RegistrationReason          RegistrationReason `json:"registrationReason,omitempty"`
	PcfId                       string             `json:"pcfId,omitempty"`
	SingleNssai                 Snssai             `json:"singleNssai"`
	PlmnId                      PlmnId             `json:"plmnId"`
	SupportedFeatures           string             `json:"supportedFeatures,omitempty"`
	PduSessionId                int                `json:"pduSessionId"`
	EmergencyServices           *bool              `json:"emergencyServices,omitempty"`
	PgwIpAddr                   *IpAddress         `json:"pgwIpAddr,omitempty"`
	EpdgInd                     *bool              `json:"epdgInd,omitempty"`
	DataRestorationCallbackUri  string             `json:"dataRestorationCallbackUri,omitempty"`
	SmfInstanceId               string             `json:"smfInstanceId"`
	SmfSetId                    string             `json:"smfSetId,omitempty"`
	ResetIds                    []string           `json:"resetIds,omitempty"`
	UdrRestartInd               *bool              `json:"udrRestartInd,omitempty"`
	PgwFqdn                     string             `json:"pgwFqdn,omitempty"`
	ContextInfo                 *ContextInfo       `json:"contextInfo,omitempty"`
	LastSynchronizationTime     string             `json:"lastSynchronizationTime,omitempty"`
	Dnn                         string             `json:"dnn,omitempty"`
	PcscfRestorationCallbackUri string             `json:"pcscfRestorationCallbackUri,omitempty"`
	DeregCallbackUri            string             `json:"deregCallbackUri,omitempty"`
	RegistrationTime            string             `json:"registrationTime,omitempty"`
}
