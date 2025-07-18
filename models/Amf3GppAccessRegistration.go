/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Amf3GppAccessRegistration struct {
	SorSnpnSiSupported          *bool                `json:"sorSnpnSiSupported,omitempty"`
	LastSynchronizationTime     string               `json:"lastSynchronizationTime,omitempty"`
	SupportedFeatures           string               `json:"supportedFeatures,omitempty"`
	DataRestorationCallbackUri  string               `json:"dataRestorationCallbackUri,omitempty"`
	AmfEeSubscriptionId         string               `json:"amfEeSubscriptionId,omitempty"`
	Supi                        string               `json:"supi,omitempty"`
	UdrRestartInd               *bool                `json:"udrRestartInd,omitempty"`
	AmfServiceNamePcscfRest     ServiceName          `json:"amfServiceNamePcscfRest,omitempty"`
	BackupAmfInfo               []BackupAmfInfo      `json:"backupAmfInfo,omitempty"`
	UrrpIndicator               *bool                `json:"urrpIndicator,omitempty"`
	EpsInterworkingInfo         *EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`
	AdminDeregSubWithdrawn      *bool                `json:"adminDeregSubWithdrawn,omitempty"`
	ResetIds                    []string             `json:"resetIds,omitempty"`
	UeMINTCapability            *bool                `json:"ueMINTCapability,omitempty"`
	Pei                         string               `json:"pei,omitempty"`
	DrFlag                      *bool                `json:"drFlag,omitempty"`
	InitialRegistrationInd      *bool                `json:"initialRegistrationInd,omitempty"`
	EmergencyRegistrationInd    *bool                `json:"emergencyRegistrationInd,omitempty"`
	RegistrationTime            string               `json:"registrationTime,omitempty"`
	DeregCallbackUri            string               `json:"deregCallbackUri"`
	AmfServiceNameDereg         ServiceName          `json:"amfServiceNameDereg,omitempty"`
	VgmlcAddress                *VgmlcAddress        `json:"vgmlcAddress,omitempty"`
	ContextInfo                 *ContextInfo         `json:"contextInfo,omitempty"`
	AmfInstanceId               string               `json:"amfInstanceId"`
	Guami                       Guami                `json:"guami"`
	UeReachableInd              UeReachableInd       `json:"ueReachableInd,omitempty"`
	ImsVoPs                     ImsVoPs              `json:"imsVoPs,omitempty"`
	RatType                     RatType              `json:"ratType"`
	ReRegistrationRequired      *bool                `json:"reRegistrationRequired,omitempty"`
	PcscfRestorationCallbackUri string               `json:"pcscfRestorationCallbackUri,omitempty"`
	NoEeSubscriptionInd         *bool                `json:"noEeSubscriptionInd,omitempty"`
	DisasterRoamingInd          *bool                `json:"disasterRoamingInd,omitempty"`
	PurgeFlag                   *bool                `json:"purgeFlag,omitempty"`
	UeSrvccCapability           *bool                `json:"ueSrvccCapability,omitempty"`
}
