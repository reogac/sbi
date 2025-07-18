/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfNon3GppAccessRegistration struct {
	Supi                        string          `json:"supi,omitempty"`
	ResetIds                    []string        `json:"resetIds,omitempty"`
	UdrRestartInd               *bool           `json:"udrRestartInd,omitempty"`
	LastSynchronizationTime     string          `json:"lastSynchronizationTime,omitempty"`
	DeregCallbackUri            string          `json:"deregCallbackUri"`
	ReRegistrationRequired      *bool           `json:"reRegistrationRequired,omitempty"`
	PurgeFlag                   *bool           `json:"purgeFlag,omitempty"`
	SupportedFeatures           string          `json:"supportedFeatures,omitempty"`
	PcscfRestorationCallbackUri string          `json:"pcscfRestorationCallbackUri,omitempty"`
	Guami                       Guami           `json:"guami"`
	AmfInstanceId               string          `json:"amfInstanceId"`
	RegistrationTime            string          `json:"registrationTime,omitempty"`
	RatType                     RatType         `json:"ratType"`
	VgmlcAddress                *VgmlcAddress   `json:"vgmlcAddress,omitempty"`
	DisasterRoamingInd          *bool           `json:"disasterRoamingInd,omitempty"`
	SorSnpnSiSupported          *bool           `json:"sorSnpnSiSupported,omitempty"`
	ImsVoPs                     ImsVoPs         `json:"imsVoPs"`
	BackupAmfInfo               []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	UrrpIndicator               *bool           `json:"urrpIndicator,omitempty"`
	AmfEeSubscriptionId         string          `json:"amfEeSubscriptionId,omitempty"`
	NoEeSubscriptionInd         *bool           `json:"noEeSubscriptionInd,omitempty"`
	AdminDeregSubWithdrawn      *bool           `json:"adminDeregSubWithdrawn,omitempty"`
	AmfServiceNamePcscfRest     ServiceName     `json:"amfServiceNamePcscfRest,omitempty"`
	Pei                         string          `json:"pei,omitempty"`
	AmfServiceNameDereg         ServiceName     `json:"amfServiceNameDereg,omitempty"`
	ContextInfo                 *ContextInfo    `json:"contextInfo,omitempty"`
	DataRestorationCallbackUri  string          `json:"dataRestorationCallbackUri,omitempty"`
}
