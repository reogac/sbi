/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:33 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfNon3GppAccessRegistration struct {
	AmfInstanceId               string          `json:"amfInstanceId"`
	Pei                         string          `json:"pei,omitempty"`
	SorSnpnSiSupported          *bool           `json:"sorSnpnSiSupported,omitempty"`
	DeregCallbackUri            string          `json:"deregCallbackUri"`
	ContextInfo                 *ContextInfo    `json:"contextInfo,omitempty"`
	SupportedFeatures           string          `json:"supportedFeatures,omitempty"`
	RatType                     RatType         `json:"ratType"`
	RegistrationTime            string          `json:"registrationTime,omitempty"`
	NoEeSubscriptionInd         *bool           `json:"noEeSubscriptionInd,omitempty"`
	UdrRestartInd               *bool           `json:"udrRestartInd,omitempty"`
	AmfServiceNamePcscfRest     ServiceName     `json:"amfServiceNamePcscfRest,omitempty"`
	UrrpIndicator               *bool           `json:"urrpIndicator,omitempty"`
	AdminDeregSubWithdrawn      *bool           `json:"adminDeregSubWithdrawn,omitempty"`
	DisasterRoamingInd          *bool           `json:"disasterRoamingInd,omitempty"`
	PcscfRestorationCallbackUri string          `json:"pcscfRestorationCallbackUri,omitempty"`
	VgmlcAddress                *VgmlcAddress   `json:"vgmlcAddress,omitempty"`
	ImsVoPs                     ImsVoPs         `json:"imsVoPs"`
	Guami                       Guami           `json:"guami"`
	Supi                        string          `json:"supi,omitempty"`
	ReRegistrationRequired      *bool           `json:"reRegistrationRequired,omitempty"`
	ResetIds                    []string        `json:"resetIds,omitempty"`
	PurgeFlag                   *bool           `json:"purgeFlag,omitempty"`
	AmfEeSubscriptionId         string          `json:"amfEeSubscriptionId,omitempty"`
	DataRestorationCallbackUri  string          `json:"dataRestorationCallbackUri,omitempty"`
	AmfServiceNameDereg         ServiceName     `json:"amfServiceNameDereg,omitempty"`
	BackupAmfInfo               []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	LastSynchronizationTime     string          `json:"lastSynchronizationTime,omitempty"`
}
