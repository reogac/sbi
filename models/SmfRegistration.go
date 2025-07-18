/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmfRegistration struct {
	RegistrationTime            string             `json:"registrationTime,omitempty"`
	ResetIds                    []string           `json:"resetIds,omitempty"`
	LastSynchronizationTime     string             `json:"lastSynchronizationTime,omitempty"`
	SupportedFeatures           string             `json:"supportedFeatures,omitempty"`
	PduSessionId                int                `json:"pduSessionId"`
	PcscfRestorationCallbackUri string             `json:"pcscfRestorationCallbackUri,omitempty"`
	RegistrationReason          RegistrationReason `json:"registrationReason,omitempty"`
	SingleNssai                 Snssai             `json:"singleNssai"`
	PlmnId                      PlmnId             `json:"plmnId"`
	PgwIpAddr                   *IpAddress         `json:"pgwIpAddr,omitempty"`
	ContextInfo                 *ContextInfo       `json:"contextInfo,omitempty"`
	SmfSetId                    string             `json:"smfSetId,omitempty"`
	PgwFqdn                     string             `json:"pgwFqdn,omitempty"`
	DeregCallbackUri            string             `json:"deregCallbackUri,omitempty"`
	UdrRestartInd               *bool              `json:"udrRestartInd,omitempty"`
	PcfId                       string             `json:"pcfId,omitempty"`
	DataRestorationCallbackUri  string             `json:"dataRestorationCallbackUri,omitempty"`
	SmfInstanceId               string             `json:"smfInstanceId"`
	Dnn                         string             `json:"dnn,omitempty"`
	EmergencyServices           *bool              `json:"emergencyServices,omitempty"`
	EpdgInd                     *bool              `json:"epdgInd,omitempty"`
}
