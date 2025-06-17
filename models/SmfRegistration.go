/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:02 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmfRegistration struct {
	ContextInfo                 *ContextInfo       `json:"contextInfo,omitempty"`
	PcscfRestorationCallbackUri string             `json:"pcscfRestorationCallbackUri,omitempty"`
	PgwFqdn                     string             `json:"pgwFqdn,omitempty"`
	EpdgInd                     *bool              `json:"epdgInd,omitempty"`
	DeregCallbackUri            string             `json:"deregCallbackUri,omitempty"`
	PlmnId                      PlmnId             `json:"plmnId"`
	RegistrationReason          RegistrationReason `json:"registrationReason,omitempty"`
	UdrRestartInd               *bool              `json:"udrRestartInd,omitempty"`
	SupportedFeatures           string             `json:"supportedFeatures,omitempty"`
	PduSessionId                int                `json:"pduSessionId"`
	Dnn                         string             `json:"dnn,omitempty"`
	EmergencyServices           *bool              `json:"emergencyServices,omitempty"`
	SmfInstanceId               string             `json:"smfInstanceId"`
	RegistrationTime            string             `json:"registrationTime,omitempty"`
	DataRestorationCallbackUri  string             `json:"dataRestorationCallbackUri,omitempty"`
	ResetIds                    []string           `json:"resetIds,omitempty"`
	LastSynchronizationTime     string             `json:"lastSynchronizationTime,omitempty"`
	SmfSetId                    string             `json:"smfSetId,omitempty"`
	SingleNssai                 Snssai             `json:"singleNssai"`
	PgwIpAddr                   *IpAddress         `json:"pgwIpAddr,omitempty"`
	PcfId                       string             `json:"pcfId,omitempty"`
}
