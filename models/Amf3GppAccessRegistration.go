/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Amf3GppAccessRegistration struct {
	PcscfRestorationCallbackUri string               `json:"pcscfRestorationCallbackUri,omitempty"`
	RegistrationTime            string               `json:"registrationTime,omitempty"`
	ResetIds                    []string             `json:"resetIds,omitempty"`
	DeregCallbackUri            string               `json:"deregCallbackUri"`
	DrFlag                      *bool                `json:"drFlag,omitempty"`
	AdminDeregSubWithdrawn      *bool                `json:"adminDeregSubWithdrawn,omitempty"`
	LastSynchronizationTime     string               `json:"lastSynchronizationTime,omitempty"`
	AmfServiceNamePcscfRest     ServiceName          `json:"amfServiceNamePcscfRest,omitempty"`
	UeMINTCapability            *bool                `json:"ueMINTCapability,omitempty"`
	SorSnpnSiSupported          *bool                `json:"sorSnpnSiSupported,omitempty"`
	ImsVoPs                     ImsVoPs              `json:"imsVoPs,omitempty"`
	VgmlcAddress                *VgmlcAddress        `json:"vgmlcAddress,omitempty"`
	NoEeSubscriptionInd         *bool                `json:"noEeSubscriptionInd,omitempty"`
	ReRegistrationRequired      *bool                `json:"reRegistrationRequired,omitempty"`
	BackupAmfInfo               []BackupAmfInfo      `json:"backupAmfInfo,omitempty"`
	UrrpIndicator               *bool                `json:"urrpIndicator,omitempty"`
	RatType                     RatType              `json:"ratType"`
	Guami                       Guami                `json:"guami"`
	AmfEeSubscriptionId         string               `json:"amfEeSubscriptionId,omitempty"`
	EpsInterworkingInfo         *EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`
	AmfInstanceId               string               `json:"amfInstanceId"`
	PurgeFlag                   *bool                `json:"purgeFlag,omitempty"`
	AmfServiceNameDereg         ServiceName          `json:"amfServiceNameDereg,omitempty"`
	EmergencyRegistrationInd    *bool                `json:"emergencyRegistrationInd,omitempty"`
	ContextInfo                 *ContextInfo         `json:"contextInfo,omitempty"`
	UeReachableInd              UeReachableInd       `json:"ueReachableInd,omitempty"`
	DataRestorationCallbackUri  string               `json:"dataRestorationCallbackUri,omitempty"`
	DisasterRoamingInd          *bool                `json:"disasterRoamingInd,omitempty"`
	SupportedFeatures           string               `json:"supportedFeatures,omitempty"`
	UdrRestartInd               *bool                `json:"udrRestartInd,omitempty"`
	InitialRegistrationInd      *bool                `json:"initialRegistrationInd,omitempty"`
	UeSrvccCapability           *bool                `json:"ueSrvccCapability,omitempty"`
	Supi                        string               `json:"supi,omitempty"`
	Pei                         string               `json:"pei,omitempty"`
}
