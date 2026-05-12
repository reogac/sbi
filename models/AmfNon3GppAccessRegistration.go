/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfNon3GppAccessRegistration struct {
	NoEeSubscriptionInd         *bool           `json:"noEeSubscriptionInd,omitempty"`
	ReRegistrationRequired      *bool           `json:"reRegistrationRequired,omitempty"`
	AdminDeregSubWithdrawn      *bool           `json:"adminDeregSubWithdrawn,omitempty"`
	SupportedFeatures           string          `json:"supportedFeatures,omitempty"`
	AmfEeSubscriptionId         string          `json:"amfEeSubscriptionId,omitempty"`
	RegistrationTime            string          `json:"registrationTime,omitempty"`
	VgmlcAddress                *VgmlcAddress   `json:"vgmlcAddress,omitempty"`
	LastSynchronizationTime     string          `json:"lastSynchronizationTime,omitempty"`
	ImsVoPs                     ImsVoPs         `json:"imsVoPs"`
	PcscfRestorationCallbackUri string          `json:"pcscfRestorationCallbackUri,omitempty"`
	DataRestorationCallbackUri  string          `json:"dataRestorationCallbackUri,omitempty"`
	ResetIds                    []string        `json:"resetIds,omitempty"`
	DisasterRoamingInd          *bool           `json:"disasterRoamingInd,omitempty"`
	SorSnpnSiSupported          *bool           `json:"sorSnpnSiSupported,omitempty"`
	AmfInstanceId               string          `json:"amfInstanceId"`
	Pei                         string          `json:"pei,omitempty"`
	AmfServiceNameDereg         ServiceName     `json:"amfServiceNameDereg,omitempty"`
	Guami                       Guami           `json:"guami"`
	BackupAmfInfo               []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	ContextInfo                 *ContextInfo    `json:"contextInfo,omitempty"`
	Supi                        string          `json:"supi,omitempty"`
	UdrRestartInd               *bool           `json:"udrRestartInd,omitempty"`
	PurgeFlag                   *bool           `json:"purgeFlag,omitempty"`
	DeregCallbackUri            string          `json:"deregCallbackUri"`
	AmfServiceNamePcscfRest     ServiceName     `json:"amfServiceNamePcscfRest,omitempty"`
	RatType                     RatType         `json:"ratType"`
	UrrpIndicator               *bool           `json:"urrpIndicator,omitempty"`
}
