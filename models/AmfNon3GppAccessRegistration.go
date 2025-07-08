/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfNon3GppAccessRegistration struct {
	PurgeFlag                   *bool           `json:"purgeFlag,omitempty"`
	ImsVoPs                     ImsVoPs         `json:"imsVoPs"`
	UdrRestartInd               *bool           `json:"udrRestartInd,omitempty"`
	BackupAmfInfo               []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	AmfEeSubscriptionId         string          `json:"amfEeSubscriptionId,omitempty"`
	AdminDeregSubWithdrawn      *bool           `json:"adminDeregSubWithdrawn,omitempty"`
	DisasterRoamingInd          *bool           `json:"disasterRoamingInd,omitempty"`
	AmfServiceNameDereg         ServiceName     `json:"amfServiceNameDereg,omitempty"`
	UrrpIndicator               *bool           `json:"urrpIndicator,omitempty"`
	NoEeSubscriptionInd         *bool           `json:"noEeSubscriptionInd,omitempty"`
	DataRestorationCallbackUri  string          `json:"dataRestorationCallbackUri,omitempty"`
	VgmlcAddress                *VgmlcAddress   `json:"vgmlcAddress,omitempty"`
	SupportedFeatures           string          `json:"supportedFeatures,omitempty"`
	PcscfRestorationCallbackUri string          `json:"pcscfRestorationCallbackUri,omitempty"`
	AmfInstanceId               string          `json:"amfInstanceId"`
	RegistrationTime            string          `json:"registrationTime,omitempty"`
	SorSnpnSiSupported          *bool           `json:"sorSnpnSiSupported,omitempty"`
	ResetIds                    []string        `json:"resetIds,omitempty"`
	Pei                         string          `json:"pei,omitempty"`
	AmfServiceNamePcscfRest     ServiceName     `json:"amfServiceNamePcscfRest,omitempty"`
	ContextInfo                 *ContextInfo    `json:"contextInfo,omitempty"`
	Supi                        string          `json:"supi,omitempty"`
	LastSynchronizationTime     string          `json:"lastSynchronizationTime,omitempty"`
	DeregCallbackUri            string          `json:"deregCallbackUri"`
	Guami                       Guami           `json:"guami"`
	RatType                     RatType         `json:"ratType"`
	ReRegistrationRequired      *bool           `json:"reRegistrationRequired,omitempty"`
}
