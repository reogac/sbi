/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:02 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfNon3GppAccessRegistration struct {
	Pei                         string          `json:"pei,omitempty"`
	AmfServiceNameDereg         ServiceName     `json:"amfServiceNameDereg,omitempty"`
	AmfServiceNamePcscfRest     ServiceName     `json:"amfServiceNamePcscfRest,omitempty"`
	Guami                       Guami           `json:"guami"`
	BackupAmfInfo               []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	ContextInfo                 *ContextInfo    `json:"contextInfo,omitempty"`
	DataRestorationCallbackUri  string          `json:"dataRestorationCallbackUri,omitempty"`
	AmfInstanceId               string          `json:"amfInstanceId"`
	ImsVoPs                     ImsVoPs         `json:"imsVoPs"`
	VgmlcAddress                *VgmlcAddress   `json:"vgmlcAddress,omitempty"`
	AdminDeregSubWithdrawn      *bool           `json:"adminDeregSubWithdrawn,omitempty"`
	PurgeFlag                   *bool           `json:"purgeFlag,omitempty"`
	UrrpIndicator               *bool           `json:"urrpIndicator,omitempty"`
	ResetIds                    []string        `json:"resetIds,omitempty"`
	DisasterRoamingInd          *bool           `json:"disasterRoamingInd,omitempty"`
	RatType                     RatType         `json:"ratType"`
	RegistrationTime            string          `json:"registrationTime,omitempty"`
	Supi                        string          `json:"supi,omitempty"`
	DeregCallbackUri            string          `json:"deregCallbackUri"`
	NoEeSubscriptionInd         *bool           `json:"noEeSubscriptionInd,omitempty"`
	LastSynchronizationTime     string          `json:"lastSynchronizationTime,omitempty"`
	SupportedFeatures           string          `json:"supportedFeatures,omitempty"`
	AmfEeSubscriptionId         string          `json:"amfEeSubscriptionId,omitempty"`
	ReRegistrationRequired      *bool           `json:"reRegistrationRequired,omitempty"`
	SorSnpnSiSupported          *bool           `json:"sorSnpnSiSupported,omitempty"`
	PcscfRestorationCallbackUri string          `json:"pcscfRestorationCallbackUri,omitempty"`
	UdrRestartInd               *bool           `json:"udrRestartInd,omitempty"`
}
