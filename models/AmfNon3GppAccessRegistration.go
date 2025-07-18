/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfNon3GppAccessRegistration struct {
	Pei                         string          `json:"pei,omitempty"`
	DataRestorationCallbackUri  string          `json:"dataRestorationCallbackUri,omitempty"`
	SorSnpnSiSupported          *bool           `json:"sorSnpnSiSupported,omitempty"`
	SupportedFeatures           string          `json:"supportedFeatures,omitempty"`
	BackupAmfInfo               []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	RegistrationTime            string          `json:"registrationTime,omitempty"`
	ContextInfo                 *ContextInfo    `json:"contextInfo,omitempty"`
	Supi                        string          `json:"supi,omitempty"`
	LastSynchronizationTime     string          `json:"lastSynchronizationTime,omitempty"`
	AmfInstanceId               string          `json:"amfInstanceId"`
	PurgeFlag                   *bool           `json:"purgeFlag,omitempty"`
	ImsVoPs                     ImsVoPs         `json:"imsVoPs"`
	PcscfRestorationCallbackUri string          `json:"pcscfRestorationCallbackUri,omitempty"`
	DeregCallbackUri            string          `json:"deregCallbackUri"`
	AmfServiceNameDereg         ServiceName     `json:"amfServiceNameDereg,omitempty"`
	UrrpIndicator               *bool           `json:"urrpIndicator,omitempty"`
	AmfEeSubscriptionId         string          `json:"amfEeSubscriptionId,omitempty"`
	UdrRestartInd               *bool           `json:"udrRestartInd,omitempty"`
	ResetIds                    []string        `json:"resetIds,omitempty"`
	NoEeSubscriptionInd         *bool           `json:"noEeSubscriptionInd,omitempty"`
	ReRegistrationRequired      *bool           `json:"reRegistrationRequired,omitempty"`
	AmfServiceNamePcscfRest     ServiceName     `json:"amfServiceNamePcscfRest,omitempty"`
	Guami                       Guami           `json:"guami"`
	VgmlcAddress                *VgmlcAddress   `json:"vgmlcAddress,omitempty"`
	AdminDeregSubWithdrawn      *bool           `json:"adminDeregSubWithdrawn,omitempty"`
	RatType                     RatType         `json:"ratType"`
	DisasterRoamingInd          *bool           `json:"disasterRoamingInd,omitempty"`
}
