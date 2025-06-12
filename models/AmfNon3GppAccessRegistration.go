/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfNon3GppAccessRegistration struct {
	RatType                     RatType         `json:"ratType"`
	AmfEeSubscriptionId         string          `json:"amfEeSubscriptionId,omitempty"`
	RegistrationTime            string          `json:"registrationTime,omitempty"`
	AmfInstanceId               string          `json:"amfInstanceId"`
	ImsVoPs                     ImsVoPs         `json:"imsVoPs"`
	SorSnpnSiSupported          *bool           `json:"sorSnpnSiSupported,omitempty"`
	LastSynchronizationTime     string          `json:"lastSynchronizationTime,omitempty"`
	SupportedFeatures           string          `json:"supportedFeatures,omitempty"`
	AmfServiceNameDereg         ServiceName     `json:"amfServiceNameDereg,omitempty"`
	Supi                        string          `json:"supi,omitempty"`
	PurgeFlag                   *bool           `json:"purgeFlag,omitempty"`
	VgmlcAddress                *VgmlcAddress   `json:"vgmlcAddress,omitempty"`
	ReRegistrationRequired      *bool           `json:"reRegistrationRequired,omitempty"`
	DisasterRoamingInd          *bool           `json:"disasterRoamingInd,omitempty"`
	Pei                         string          `json:"pei,omitempty"`
	UrrpIndicator               *bool           `json:"urrpIndicator,omitempty"`
	AdminDeregSubWithdrawn      *bool           `json:"adminDeregSubWithdrawn,omitempty"`
	DeregCallbackUri            string          `json:"deregCallbackUri"`
	DataRestorationCallbackUri  string          `json:"dataRestorationCallbackUri,omitempty"`
	Guami                       Guami           `json:"guami"`
	ContextInfo                 *ContextInfo    `json:"contextInfo,omitempty"`
	UdrRestartInd               *bool           `json:"udrRestartInd,omitempty"`
	BackupAmfInfo               []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	NoEeSubscriptionInd         *bool           `json:"noEeSubscriptionInd,omitempty"`
	ResetIds                    []string        `json:"resetIds,omitempty"`
	PcscfRestorationCallbackUri string          `json:"pcscfRestorationCallbackUri,omitempty"`
	AmfServiceNamePcscfRest     ServiceName     `json:"amfServiceNamePcscfRest,omitempty"`
}
