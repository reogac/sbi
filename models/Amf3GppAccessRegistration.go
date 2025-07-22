/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Amf3GppAccessRegistration struct {
	SupportedFeatures           string               `json:"supportedFeatures,omitempty"`
	PcscfRestorationCallbackUri string               `json:"pcscfRestorationCallbackUri,omitempty"`
	RatType                     RatType              `json:"ratType"`
	EpsInterworkingInfo         *EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`
	UeSrvccCapability           *bool                `json:"ueSrvccCapability,omitempty"`
	LastSynchronizationTime     string               `json:"lastSynchronizationTime,omitempty"`
	PurgeFlag                   *bool                `json:"purgeFlag,omitempty"`
	AmfServiceNameDereg         ServiceName          `json:"amfServiceNameDereg,omitempty"`
	DrFlag                      *bool                `json:"drFlag,omitempty"`
	UrrpIndicator               *bool                `json:"urrpIndicator,omitempty"`
	DeregCallbackUri            string               `json:"deregCallbackUri"`
	InitialRegistrationInd      *bool                `json:"initialRegistrationInd,omitempty"`
	DataRestorationCallbackUri  string               `json:"dataRestorationCallbackUri,omitempty"`
	VgmlcAddress                *VgmlcAddress        `json:"vgmlcAddress,omitempty"`
	AdminDeregSubWithdrawn      *bool                `json:"adminDeregSubWithdrawn,omitempty"`
	EmergencyRegistrationInd    *bool                `json:"emergencyRegistrationInd,omitempty"`
	UdrRestartInd               *bool                `json:"udrRestartInd,omitempty"`
	SorSnpnSiSupported          *bool                `json:"sorSnpnSiSupported,omitempty"`
	ImsVoPs                     ImsVoPs              `json:"imsVoPs,omitempty"`
	BackupAmfInfo               []BackupAmfInfo      `json:"backupAmfInfo,omitempty"`
	Supi                        string               `json:"supi,omitempty"`
	ResetIds                    []string             `json:"resetIds,omitempty"`
	DisasterRoamingInd          *bool                `json:"disasterRoamingInd,omitempty"`
	UeMINTCapability            *bool                `json:"ueMINTCapability,omitempty"`
	AmfServiceNamePcscfRest     ServiceName          `json:"amfServiceNamePcscfRest,omitempty"`
	Guami                       Guami                `json:"guami"`
	AmfEeSubscriptionId         string               `json:"amfEeSubscriptionId,omitempty"`
	ReRegistrationRequired      *bool                `json:"reRegistrationRequired,omitempty"`
	NoEeSubscriptionInd         *bool                `json:"noEeSubscriptionInd,omitempty"`
	UeReachableInd              UeReachableInd       `json:"ueReachableInd,omitempty"`
	AmfInstanceId               string               `json:"amfInstanceId"`
	Pei                         string               `json:"pei,omitempty"`
	RegistrationTime            string               `json:"registrationTime,omitempty"`
	ContextInfo                 *ContextInfo         `json:"contextInfo,omitempty"`
}
