/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Amf3GppAccessRegistration struct {
	AmfInstanceId               string               `json:"amfInstanceId"`
	ImsVoPs                     ImsVoPs              `json:"imsVoPs,omitempty"`
	AmfServiceNamePcscfRest     ServiceName          `json:"amfServiceNamePcscfRest,omitempty"`
	InitialRegistrationInd      *bool                `json:"initialRegistrationInd,omitempty"`
	UrrpIndicator               *bool                `json:"urrpIndicator,omitempty"`
	DisasterRoamingInd          *bool                `json:"disasterRoamingInd,omitempty"`
	EpsInterworkingInfo         *EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`
	ResetIds                    []string             `json:"resetIds,omitempty"`
	UdrRestartInd               *bool                `json:"udrRestartInd,omitempty"`
	EmergencyRegistrationInd    *bool                `json:"emergencyRegistrationInd,omitempty"`
	AdminDeregSubWithdrawn      *bool                `json:"adminDeregSubWithdrawn,omitempty"`
	SorSnpnSiSupported          *bool                `json:"sorSnpnSiSupported,omitempty"`
	DrFlag                      *bool                `json:"drFlag,omitempty"`
	RatType                     RatType              `json:"ratType"`
	AmfEeSubscriptionId         string               `json:"amfEeSubscriptionId,omitempty"`
	Supi                        string               `json:"supi,omitempty"`
	ReRegistrationRequired      *bool                `json:"reRegistrationRequired,omitempty"`
	RegistrationTime            string               `json:"registrationTime,omitempty"`
	BackupAmfInfo               []BackupAmfInfo      `json:"backupAmfInfo,omitempty"`
	PurgeFlag                   *bool                `json:"purgeFlag,omitempty"`
	PcscfRestorationCallbackUri string               `json:"pcscfRestorationCallbackUri,omitempty"`
	LastSynchronizationTime     string               `json:"lastSynchronizationTime,omitempty"`
	Pei                         string               `json:"pei,omitempty"`
	DeregCallbackUri            string               `json:"deregCallbackUri"`
	AmfServiceNameDereg         ServiceName          `json:"amfServiceNameDereg,omitempty"`
	Guami                       Guami                `json:"guami"`
	VgmlcAddress                *VgmlcAddress        `json:"vgmlcAddress,omitempty"`
	NoEeSubscriptionInd         *bool                `json:"noEeSubscriptionInd,omitempty"`
	DataRestorationCallbackUri  string               `json:"dataRestorationCallbackUri,omitempty"`
	UeMINTCapability            *bool                `json:"ueMINTCapability,omitempty"`
	SupportedFeatures           string               `json:"supportedFeatures,omitempty"`
	UeReachableInd              UeReachableInd       `json:"ueReachableInd,omitempty"`
	UeSrvccCapability           *bool                `json:"ueSrvccCapability,omitempty"`
	ContextInfo                 *ContextInfo         `json:"contextInfo,omitempty"`
}
