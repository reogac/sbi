/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfNon3GppAccessRegistration struct {
	UrrpIndicator               *bool           `json:"urrpIndicator,omitempty"`
	SorSnpnSiSupported          *bool           `json:"sorSnpnSiSupported,omitempty"`
	DisasterRoamingInd          *bool           `json:"disasterRoamingInd,omitempty"`
	Guami                       Guami           `json:"guami"`
	DataRestorationCallbackUri  string          `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd               *bool           `json:"udrRestartInd,omitempty"`
	LastSynchronizationTime     string          `json:"lastSynchronizationTime,omitempty"`
	SupportedFeatures           string          `json:"supportedFeatures,omitempty"`
	AmfServiceNameDereg         ServiceName     `json:"amfServiceNameDereg,omitempty"`
	NoEeSubscriptionInd         *bool           `json:"noEeSubscriptionInd,omitempty"`
	Pei                         string          `json:"pei,omitempty"`
	RatType                     RatType         `json:"ratType"`
	BackupAmfInfo               []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	AmfEeSubscriptionId         string          `json:"amfEeSubscriptionId,omitempty"`
	RegistrationTime            string          `json:"registrationTime,omitempty"`
	VgmlcAddress                *VgmlcAddress   `json:"vgmlcAddress,omitempty"`
	ContextInfo                 *ContextInfo    `json:"contextInfo,omitempty"`
	PurgeFlag                   *bool           `json:"purgeFlag,omitempty"`
	ImsVoPs                     ImsVoPs         `json:"imsVoPs"`
	ResetIds                    []string        `json:"resetIds,omitempty"`
	DeregCallbackUri            string          `json:"deregCallbackUri"`
	AdminDeregSubWithdrawn      *bool           `json:"adminDeregSubWithdrawn,omitempty"`
	PcscfRestorationCallbackUri string          `json:"pcscfRestorationCallbackUri,omitempty"`
	Supi                        string          `json:"supi,omitempty"`
	ReRegistrationRequired      *bool           `json:"reRegistrationRequired,omitempty"`
	AmfInstanceId               string          `json:"amfInstanceId"`
	AmfServiceNamePcscfRest     ServiceName     `json:"amfServiceNamePcscfRest,omitempty"`
}
