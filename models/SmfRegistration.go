/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmfRegistration struct {
	PduSessionId                int                `json:"pduSessionId"`
	SingleNssai                 Snssai             `json:"singleNssai"`
	RegistrationTime            string             `json:"registrationTime,omitempty"`
	UdrRestartInd               *bool              `json:"udrRestartInd,omitempty"`
	DeregCallbackUri            string             `json:"deregCallbackUri,omitempty"`
	RegistrationReason          RegistrationReason `json:"registrationReason,omitempty"`
	SmfSetId                    string             `json:"smfSetId,omitempty"`
	SupportedFeatures           string             `json:"supportedFeatures,omitempty"`
	Dnn                         string             `json:"dnn,omitempty"`
	EmergencyServices           *bool              `json:"emergencyServices,omitempty"`
	PlmnId                      PlmnId             `json:"plmnId"`
	PgwFqdn                     string             `json:"pgwFqdn,omitempty"`
	SmfInstanceId               string             `json:"smfInstanceId"`
	PcscfRestorationCallbackUri string             `json:"pcscfRestorationCallbackUri,omitempty"`
	PgwIpAddr                   *IpAddress         `json:"pgwIpAddr,omitempty"`
	ContextInfo                 *ContextInfo       `json:"contextInfo,omitempty"`
	ResetIds                    []string           `json:"resetIds,omitempty"`
	LastSynchronizationTime     string             `json:"lastSynchronizationTime,omitempty"`
	EpdgInd                     *bool              `json:"epdgInd,omitempty"`
	PcfId                       string             `json:"pcfId,omitempty"`
	DataRestorationCallbackUri  string             `json:"dataRestorationCallbackUri,omitempty"`
}
