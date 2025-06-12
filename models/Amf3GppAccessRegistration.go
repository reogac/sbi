/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Amf3GppAccessRegistration struct {
	BackupAmfInfo               []BackupAmfInfo      `json:"backupAmfInfo,omitempty"`
	AmfInstanceId               string               `json:"amfInstanceId"`
	AmfServiceNameDereg         ServiceName          `json:"amfServiceNameDereg,omitempty"`
	InitialRegistrationInd      *bool                `json:"initialRegistrationInd,omitempty"`
	AmfEeSubscriptionId         string               `json:"amfEeSubscriptionId,omitempty"`
	UeSrvccCapability           *bool                `json:"ueSrvccCapability,omitempty"`
	Supi                        string               `json:"supi,omitempty"`
	SupportedFeatures           string               `json:"supportedFeatures,omitempty"`
	PcscfRestorationCallbackUri string               `json:"pcscfRestorationCallbackUri,omitempty"`
	AmfServiceNamePcscfRest     ServiceName          `json:"amfServiceNamePcscfRest,omitempty"`
	RatType                     RatType              `json:"ratType"`
	NoEeSubscriptionInd         *bool                `json:"noEeSubscriptionInd,omitempty"`
	UeReachableInd              UeReachableInd       `json:"ueReachableInd,omitempty"`
	EpsInterworkingInfo         *EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`
	RegistrationTime            string               `json:"registrationTime,omitempty"`
	ContextInfo                 *ContextInfo         `json:"contextInfo,omitempty"`
	UrrpIndicator               *bool                `json:"urrpIndicator,omitempty"`
	VgmlcAddress                *VgmlcAddress        `json:"vgmlcAddress,omitempty"`
	DataRestorationCallbackUri  string               `json:"dataRestorationCallbackUri,omitempty"`
	LastSynchronizationTime     string               `json:"lastSynchronizationTime,omitempty"`
	Pei                         string               `json:"pei,omitempty"`
	ImsVoPs                     ImsVoPs              `json:"imsVoPs,omitempty"`
	DeregCallbackUri            string               `json:"deregCallbackUri"`
	ReRegistrationRequired      *bool                `json:"reRegistrationRequired,omitempty"`
	DisasterRoamingInd          *bool                `json:"disasterRoamingInd,omitempty"`
	SorSnpnSiSupported          *bool                `json:"sorSnpnSiSupported,omitempty"`
	EmergencyRegistrationInd    *bool                `json:"emergencyRegistrationInd,omitempty"`
	Guami                       Guami                `json:"guami"`
	UeMINTCapability            *bool                `json:"ueMINTCapability,omitempty"`
	ResetIds                    []string             `json:"resetIds,omitempty"`
	UdrRestartInd               *bool                `json:"udrRestartInd,omitempty"`
	PurgeFlag                   *bool                `json:"purgeFlag,omitempty"`
	DrFlag                      *bool                `json:"drFlag,omitempty"`
	AdminDeregSubWithdrawn      *bool                `json:"adminDeregSubWithdrawn,omitempty"`
}
