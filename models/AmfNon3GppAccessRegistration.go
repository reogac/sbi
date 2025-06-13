/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:51 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfNon3GppAccessRegistration struct {
	AmfServiceNamePcscfRest     ServiceName     `json:"amfServiceNamePcscfRest,omitempty"`
	DisasterRoamingInd          *bool           `json:"disasterRoamingInd,omitempty"`
	ImsVoPs                     ImsVoPs         `json:"imsVoPs"`
	AmfEeSubscriptionId         string          `json:"amfEeSubscriptionId,omitempty"`
	AmfInstanceId               string          `json:"amfInstanceId"`
	SupportedFeatures           string          `json:"supportedFeatures,omitempty"`
	Pei                         string          `json:"pei,omitempty"`
	DeregCallbackUri            string          `json:"deregCallbackUri"`
	AdminDeregSubWithdrawn      *bool           `json:"adminDeregSubWithdrawn,omitempty"`
	DataRestorationCallbackUri  string          `json:"dataRestorationCallbackUri,omitempty"`
	PurgeFlag                   *bool           `json:"purgeFlag,omitempty"`
	Guami                       Guami           `json:"guami"`
	RatType                     RatType         `json:"ratType"`
	ContextInfo                 *ContextInfo    `json:"contextInfo,omitempty"`
	Supi                        string          `json:"supi,omitempty"`
	RegistrationTime            string          `json:"registrationTime,omitempty"`
	SorSnpnSiSupported          *bool           `json:"sorSnpnSiSupported,omitempty"`
	NoEeSubscriptionInd         *bool           `json:"noEeSubscriptionInd,omitempty"`
	ReRegistrationRequired      *bool           `json:"reRegistrationRequired,omitempty"`
	PcscfRestorationCallbackUri string          `json:"pcscfRestorationCallbackUri,omitempty"`
	BackupAmfInfo               []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	UrrpIndicator               *bool           `json:"urrpIndicator,omitempty"`
	VgmlcAddress                *VgmlcAddress   `json:"vgmlcAddress,omitempty"`
	UdrRestartInd               *bool           `json:"udrRestartInd,omitempty"`
	LastSynchronizationTime     string          `json:"lastSynchronizationTime,omitempty"`
	ResetIds                    []string        `json:"resetIds,omitempty"`
	AmfServiceNameDereg         ServiceName     `json:"amfServiceNameDereg,omitempty"`
}
