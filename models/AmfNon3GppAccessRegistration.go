/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfNon3GppAccessRegistration struct {
	ResetIds                    []string        `json:"resetIds,omitempty"`
	SorSnpnSiSupported          *bool           `json:"sorSnpnSiSupported,omitempty"`
	ReRegistrationRequired      *bool           `json:"reRegistrationRequired,omitempty"`
	AdminDeregSubWithdrawn      *bool           `json:"adminDeregSubWithdrawn,omitempty"`
	DisasterRoamingInd          *bool           `json:"disasterRoamingInd,omitempty"`
	LastSynchronizationTime     string          `json:"lastSynchronizationTime,omitempty"`
	RatType                     RatType         `json:"ratType"`
	UrrpIndicator               *bool           `json:"urrpIndicator,omitempty"`
	ContextInfo                 *ContextInfo    `json:"contextInfo,omitempty"`
	Supi                        string          `json:"supi,omitempty"`
	AmfInstanceId               string          `json:"amfInstanceId"`
	PurgeFlag                   *bool           `json:"purgeFlag,omitempty"`
	AmfServiceNameDereg         ServiceName     `json:"amfServiceNameDereg,omitempty"`
	VgmlcAddress                *VgmlcAddress   `json:"vgmlcAddress,omitempty"`
	DataRestorationCallbackUri  string          `json:"dataRestorationCallbackUri,omitempty"`
	SupportedFeatures           string          `json:"supportedFeatures,omitempty"`
	ImsVoPs                     ImsVoPs         `json:"imsVoPs"`
	PcscfRestorationCallbackUri string          `json:"pcscfRestorationCallbackUri,omitempty"`
	NoEeSubscriptionInd         *bool           `json:"noEeSubscriptionInd,omitempty"`
	Pei                         string          `json:"pei,omitempty"`
	UdrRestartInd               *bool           `json:"udrRestartInd,omitempty"`
	Guami                       Guami           `json:"guami"`
	AmfEeSubscriptionId         string          `json:"amfEeSubscriptionId,omitempty"`
	RegistrationTime            string          `json:"registrationTime,omitempty"`
	DeregCallbackUri            string          `json:"deregCallbackUri"`
	AmfServiceNamePcscfRest     ServiceName     `json:"amfServiceNamePcscfRest,omitempty"`
	BackupAmfInfo               []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
}
