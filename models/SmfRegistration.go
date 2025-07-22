/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmfRegistration struct {
	EpdgInd                     *bool              `json:"epdgInd,omitempty"`
	DeregCallbackUri            string             `json:"deregCallbackUri,omitempty"`
	DataRestorationCallbackUri  string             `json:"dataRestorationCallbackUri,omitempty"`
	PduSessionId                int                `json:"pduSessionId"`
	PcscfRestorationCallbackUri string             `json:"pcscfRestorationCallbackUri,omitempty"`
	PcfId                       string             `json:"pcfId,omitempty"`
	PgwIpAddr                   *IpAddress         `json:"pgwIpAddr,omitempty"`
	RegistrationTime            string             `json:"registrationTime,omitempty"`
	UdrRestartInd               *bool              `json:"udrRestartInd,omitempty"`
	SmfSetId                    string             `json:"smfSetId,omitempty"`
	SingleNssai                 Snssai             `json:"singleNssai"`
	Dnn                         string             `json:"dnn,omitempty"`
	PlmnId                      PlmnId             `json:"plmnId"`
	PgwFqdn                     string             `json:"pgwFqdn,omitempty"`
	RegistrationReason          RegistrationReason `json:"registrationReason,omitempty"`
	ContextInfo                 *ContextInfo       `json:"contextInfo,omitempty"`
	ResetIds                    []string           `json:"resetIds,omitempty"`
	SmfInstanceId               string             `json:"smfInstanceId"`
	SupportedFeatures           string             `json:"supportedFeatures,omitempty"`
	EmergencyServices           *bool              `json:"emergencyServices,omitempty"`
	LastSynchronizationTime     string             `json:"lastSynchronizationTime,omitempty"`
}
