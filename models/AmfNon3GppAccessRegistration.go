/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfNon3GppAccessRegistration struct {
	ResetIds                    []string        `json:"resetIds,omitempty"`
	RatType                     RatType         `json:"ratType"`
	VgmlcAddress                *VgmlcAddress   `json:"vgmlcAddress,omitempty"`
	SupportedFeatures           string          `json:"supportedFeatures,omitempty"`
	AmfServiceNamePcscfRest     ServiceName     `json:"amfServiceNamePcscfRest,omitempty"`
	ReRegistrationRequired      *bool           `json:"reRegistrationRequired,omitempty"`
	LastSynchronizationTime     string          `json:"lastSynchronizationTime,omitempty"`
	DeregCallbackUri            string          `json:"deregCallbackUri"`
	AmfEeSubscriptionId         string          `json:"amfEeSubscriptionId,omitempty"`
	SorSnpnSiSupported          *bool           `json:"sorSnpnSiSupported,omitempty"`
	AmfInstanceId               string          `json:"amfInstanceId"`
	PurgeFlag                   *bool           `json:"purgeFlag,omitempty"`
	Pei                         string          `json:"pei,omitempty"`
	PcscfRestorationCallbackUri string          `json:"pcscfRestorationCallbackUri,omitempty"`
	Supi                        string          `json:"supi,omitempty"`
	DataRestorationCallbackUri  string          `json:"dataRestorationCallbackUri,omitempty"`
	AmfServiceNameDereg         ServiceName     `json:"amfServiceNameDereg,omitempty"`
	Guami                       Guami           `json:"guami"`
	BackupAmfInfo               []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	UrrpIndicator               *bool           `json:"urrpIndicator,omitempty"`
	DisasterRoamingInd          *bool           `json:"disasterRoamingInd,omitempty"`
	ImsVoPs                     ImsVoPs         `json:"imsVoPs"`
	RegistrationTime            string          `json:"registrationTime,omitempty"`
	ContextInfo                 *ContextInfo    `json:"contextInfo,omitempty"`
	NoEeSubscriptionInd         *bool           `json:"noEeSubscriptionInd,omitempty"`
	AdminDeregSubWithdrawn      *bool           `json:"adminDeregSubWithdrawn,omitempty"`
	UdrRestartInd               *bool           `json:"udrRestartInd,omitempty"`
}
