/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmfRegistration struct {
	PcfId                       string             `json:"pcfId,omitempty"`
	ResetIds                    []string           `json:"resetIds,omitempty"`
	EmergencyServices           *bool              `json:"emergencyServices,omitempty"`
	PgwFqdn                     string             `json:"pgwFqdn,omitempty"`
	DeregCallbackUri            string             `json:"deregCallbackUri,omitempty"`
	ContextInfo                 *ContextInfo       `json:"contextInfo,omitempty"`
	PduSessionId                int                `json:"pduSessionId"`
	SingleNssai                 Snssai             `json:"singleNssai"`
	PcscfRestorationCallbackUri string             `json:"pcscfRestorationCallbackUri,omitempty"`
	UdrRestartInd               *bool              `json:"udrRestartInd,omitempty"`
	PgwIpAddr                   *IpAddress         `json:"pgwIpAddr,omitempty"`
	EpdgInd                     *bool              `json:"epdgInd,omitempty"`
	RegistrationReason          RegistrationReason `json:"registrationReason,omitempty"`
	RegistrationTime            string             `json:"registrationTime,omitempty"`
	SmfInstanceId               string             `json:"smfInstanceId"`
	SmfSetId                    string             `json:"smfSetId,omitempty"`
	SupportedFeatures           string             `json:"supportedFeatures,omitempty"`
	Dnn                         string             `json:"dnn,omitempty"`
	DataRestorationCallbackUri  string             `json:"dataRestorationCallbackUri,omitempty"`
	PlmnId                      PlmnId             `json:"plmnId"`
	LastSynchronizationTime     string             `json:"lastSynchronizationTime,omitempty"`
}
