/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Apr 30 14:54:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HsmfUpdateData struct {
	UnknownN1SmInfo                      *RefToBinaryData              `json:"unknownN1SmInfo,omitempty"`
	SupportedFeatures                    string                        `json:"supportedFeatures,omitempty"`
	SecondaryRatUsageDataReportContainer []string                      `json:"secondaryRatUsageDataReportContainer,omitempty"`
	Cause                                Cause                         `json:"cause,omitempty"`
	UlclBpInfo                           *UlclBpInformation            `json:"ulclBpInfo,omitempty"`
	SecurityResult                       *SecurityResult               `json:"securityResult,omitempty"`
	SatelliteBackhaulCat                 SatelliteBackhaulCategory     `json:"satelliteBackhaulCat,omitempty"`
	QosFlowsRelNotifyList                []QosFlowItem                 `json:"qosFlowsRelNotifyList,omitempty"`
	DlServingPlmnRateCtl                 *int                          `json:"dlServingPlmnRateCtl,omitempty"`
	RoamingChargingProfile               *RoamingChargingProfile       `json:"roamingChargingProfile,omitempty"`
	VplmnQos                             *VplmnQos                     `json:"vplmnQos,omitempty"`
	UpSecurityInfo                       *UpSecurityInfo               `json:"upSecurityInfo,omitempty"`
	VcnTunnelInfo                        *TunnelInfo                   `json:"vcnTunnelInfo,omitempty"`
	SecondaryRatUsageReport              []SecondaryRatUsageReport     `json:"secondaryRatUsageReport,omitempty"`
	IsmfId                               string                        `json:"ismfId,omitempty"`
	Pei                                  string                        `json:"pei,omitempty"`
	MaRequestInd                         *bool                         `json:"maRequestInd,omitempty"`
	VSmfServiceInstanceId                string                        `json:"vSmfServiceInstanceId,omitempty"`
	Guami                                *Guami                        `json:"guami,omitempty"`
	MaxIntegrityProtectedDataRateDl      MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	NgApCause                            *NgApCause                    `json:"ngApCause,omitempty"`
	SecondaryRatUsageInfo                []SecondaryRatUsageInfo       `json:"secondaryRatUsageInfo,omitempty"`
	N4InfoExt2                           *N4Information                `json:"n4InfoExt2,omitempty"`
	VsmfPduSessionUri                    string                        `json:"vsmfPduSessionUri,omitempty"`
	RatType                              RatType                       `json:"ratType,omitempty"`
	AnTypeCanBeChanged                   *bool                         `json:"anTypeCanBeChanged,omitempty"`
	N4Info                               *N4Information                `json:"n4Info,omitempty"`
	PresenceInLadn                       PresenceState                 `json:"presenceInLadn,omitempty"`
	DnaiList                             []string                      `json:"dnaiList,omitempty"`
	AmfNfId                              string                        `json:"amfNfId,omitempty"`
	FiveGMmCauseValue                    *int                          `json:"5gMmCauseValue,omitempty"`
	N4InfoExt1                           *N4Information                `json:"n4InfoExt1,omitempty"`
	SmPolicyNotifyInd                    *SmPolicyNotifyInd            `json:"smPolicyNotifyInd,omitempty"`
	Pti                                  *int                          `json:"pti,omitempty"`
	VsmfId                               string                        `json:"vsmfId,omitempty"`
	ISmfServiceInstanceId                string                        `json:"iSmfServiceInstanceId,omitempty"`
	RequestIndication                    RequestIndication             `json:"requestIndication"`
	PauseCharging                        *bool                         `json:"pauseCharging,omitempty"`
	RevokeEbiList                        []int                         `json:"revokeEbiList,omitempty"`
	PcfUeCallbackInfo                    *PcfUeCallbackInfo            `json:"pcfUeCallbackInfo,omitempty"`
	AdditionalAnType                     AccessType                    `json:"additionalAnType,omitempty"`
	MaReleaseInd                         MaReleaseIndication           `json:"maReleaseInd,omitempty"`
	IsmfPduSessionUri                    string                        `json:"ismfPduSessionUri,omitempty"`
	DisasterRoamingInd                   *DisasterRoamingInd           `json:"disasterRoamingInd,omitempty"`
	IcnTunnelInfo                        *TunnelInfo                   `json:"icnTunnelInfo,omitempty"`
	NotifyList                           []PduSessionNotifyItem        `json:"NotifyList,omitempty"`
	EpsBearerId                          []int                         `json:"epsBearerId,omitempty"`
	AdditionalCnTunnelInfo               *TunnelInfo                   `json:"additionalCnTunnelInfo,omitempty"`
	QosFlowsNotifyList                   []QosFlowNotifyItem           `json:"qosFlowsNotifyList,omitempty"`
	AlwaysOnRequested                    *bool                         `json:"alwaysOnRequested,omitempty"`
	EpsInterworkingInd                   EpsInterworkingIndication     `json:"epsInterworkingInd,omitempty"`
	MaNwUpgradeInd                       *bool                         `json:"maNwUpgradeInd,omitempty"`
	PsaInfo                              []PsaInformation              `json:"psaInfo,omitempty"`
	UeTimeZone                           string                        `json:"ueTimeZone,omitempty"`
	UnavailableAccessInd                 UnavailableAccessIndication   `json:"unavailableAccessInd,omitempty"`
	MaxIntegrityProtectedDataRateUl      MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	UpCnxState                           UpCnxState                    `json:"upCnxState,omitempty"`
	UeLocation                           *UserLocation                 `json:"ueLocation,omitempty"`
	ServingNetwork                       *PlmnIdNid                    `json:"servingNetwork,omitempty"`
	AnType                               AccessType                    `json:"anType,omitempty"`
	AddUeLocation                        *UserLocation                 `json:"addUeLocation,omitempty"`
	N1SmInfoFromUe                       *RefToBinaryData              `json:"n1SmInfoFromUe,omitempty"`
	HoPreparationIndication              *bool                         `json:"hoPreparationIndication,omitempty"`
	MoExpDataCounter                     *MoExpDataCounter             `json:"moExpDataCounter,omitempty"`
}
