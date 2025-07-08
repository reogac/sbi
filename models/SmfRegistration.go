/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmfRegistration struct {
	PcscfRestorationCallbackUri string             `json:"pcscfRestorationCallbackUri,omitempty"`
	ResetIds                    []string           `json:"resetIds,omitempty"`
	UdrRestartInd               *bool              `json:"udrRestartInd,omitempty"`
	SmfInstanceId               string             `json:"smfInstanceId"`
	SupportedFeatures           string             `json:"supportedFeatures,omitempty"`
	PlmnId                      PlmnId             `json:"plmnId"`
	DeregCallbackUri            string             `json:"deregCallbackUri,omitempty"`
	ContextInfo                 *ContextInfo       `json:"contextInfo,omitempty"`
	PcfId                       string             `json:"pcfId,omitempty"`
	SmfSetId                    string             `json:"smfSetId,omitempty"`
	Dnn                         string             `json:"dnn,omitempty"`
	EpdgInd                     *bool              `json:"epdgInd,omitempty"`
	RegistrationReason          RegistrationReason `json:"registrationReason,omitempty"`
	LastSynchronizationTime     string             `json:"lastSynchronizationTime,omitempty"`
	PduSessionId                int                `json:"pduSessionId"`
	SingleNssai                 Snssai             `json:"singleNssai"`
	EmergencyServices           *bool              `json:"emergencyServices,omitempty"`
	PgwFqdn                     string             `json:"pgwFqdn,omitempty"`
	PgwIpAddr                   *IpAddress         `json:"pgwIpAddr,omitempty"`
	RegistrationTime            string             `json:"registrationTime,omitempty"`
	DataRestorationCallbackUri  string             `json:"dataRestorationCallbackUri,omitempty"`
}
