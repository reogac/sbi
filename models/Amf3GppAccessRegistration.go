/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Amf3GppAccessRegistration struct {
	VgmlcAddress                *VgmlcAddress        `json:"vgmlcAddress,omitempty"`
	PurgeFlag                   *bool                `json:"purgeFlag,omitempty"`
	EmergencyRegistrationInd    *bool                `json:"emergencyRegistrationInd,omitempty"`
	BackupAmfInfo               []BackupAmfInfo      `json:"backupAmfInfo,omitempty"`
	UeReachableInd              UeReachableInd       `json:"ueReachableInd,omitempty"`
	UdrRestartInd               *bool                `json:"udrRestartInd,omitempty"`
	AdminDeregSubWithdrawn      *bool                `json:"adminDeregSubWithdrawn,omitempty"`
	PcscfRestorationCallbackUri string               `json:"pcscfRestorationCallbackUri,omitempty"`
	AmfEeSubscriptionId         string               `json:"amfEeSubscriptionId,omitempty"`
	UeSrvccCapability           *bool                `json:"ueSrvccCapability,omitempty"`
	NoEeSubscriptionInd         *bool                `json:"noEeSubscriptionInd,omitempty"`
	ReRegistrationRequired      *bool                `json:"reRegistrationRequired,omitempty"`
	UeMINTCapability            *bool                `json:"ueMINTCapability,omitempty"`
	SorSnpnSiSupported          *bool                `json:"sorSnpnSiSupported,omitempty"`
	AmfInstanceId               string               `json:"amfInstanceId"`
	InitialRegistrationInd      *bool                `json:"initialRegistrationInd,omitempty"`
	ContextInfo                 *ContextInfo         `json:"contextInfo,omitempty"`
	Supi                        string               `json:"supi,omitempty"`
	DisasterRoamingInd          *bool                `json:"disasterRoamingInd,omitempty"`
	AmfServiceNamePcscfRest     ServiceName          `json:"amfServiceNamePcscfRest,omitempty"`
	RegistrationTime            string               `json:"registrationTime,omitempty"`
	EpsInterworkingInfo         *EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`
	LastSynchronizationTime     string               `json:"lastSynchronizationTime,omitempty"`
	Pei                         string               `json:"pei,omitempty"`
	ImsVoPs                     ImsVoPs              `json:"imsVoPs,omitempty"`
	DeregCallbackUri            string               `json:"deregCallbackUri"`
	RatType                     RatType              `json:"ratType"`
	UrrpIndicator               *bool                `json:"urrpIndicator,omitempty"`
	Guami                       Guami                `json:"guami"`
	DrFlag                      *bool                `json:"drFlag,omitempty"`
	SupportedFeatures           string               `json:"supportedFeatures,omitempty"`
	AmfServiceNameDereg         ServiceName          `json:"amfServiceNameDereg,omitempty"`
	DataRestorationCallbackUri  string               `json:"dataRestorationCallbackUri,omitempty"`
	ResetIds                    []string             `json:"resetIds,omitempty"`
}
