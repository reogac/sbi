/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HsmfUpdateData struct {
	ServingNetwork                       *PlmnIdNid                    `json:"servingNetwork,omitempty"`
	Cause                                Cause                         `json:"cause,omitempty"`
	NgApCause                            *NgApCause                    `json:"ngApCause,omitempty"`
	VcnTunnelInfo                        *TunnelInfo                   `json:"vcnTunnelInfo,omitempty"`
	UeTimeZone                           string                        `json:"ueTimeZone,omitempty"`
	ISmfServiceInstanceId                string                        `json:"iSmfServiceInstanceId,omitempty"`
	VsmfId                               string                        `json:"vsmfId,omitempty"`
	SatelliteBackhaulCat                 SatelliteBackhaulCategory     `json:"satelliteBackhaulCat,omitempty"`
	MaNwUpgradeInd                       *bool                         `json:"maNwUpgradeInd,omitempty"`
	UnavailableAccessInd                 UnavailableAccessIndication   `json:"unavailableAccessInd,omitempty"`
	PresenceInLadn                       PresenceState                 `json:"presenceInLadn,omitempty"`
	VsmfPduSessionUri                    string                        `json:"vsmfPduSessionUri,omitempty"`
	AmfNfId                              string                        `json:"amfNfId,omitempty"`
	DisasterRoamingInd                   *DisasterRoamingInd           `json:"disasterRoamingInd,omitempty"`
	IcnTunnelInfo                        *TunnelInfo                   `json:"icnTunnelInfo,omitempty"`
	Pti                                  *int                          `json:"pti,omitempty"`
	QosFlowsRelNotifyList                []QosFlowItem                 `json:"qosFlowsRelNotifyList,omitempty"`
	DnaiList                             []string                      `json:"dnaiList,omitempty"`
	AdditionalAnType                     AccessType                    `json:"additionalAnType,omitempty"`
	N1SmInfoFromUe                       *RefToBinaryData              `json:"n1SmInfoFromUe,omitempty"`
	N4InfoExt1                           *N4Information                `json:"n4InfoExt1,omitempty"`
	AdditionalCnTunnelInfo               *TunnelInfo                   `json:"additionalCnTunnelInfo,omitempty"`
	IsmfPduSessionUri                    string                        `json:"ismfPduSessionUri,omitempty"`
	Guami                                *Guami                        `json:"guami,omitempty"`
	RequestIndication                    RequestIndication             `json:"requestIndication"`
	IsmfId                               string                        `json:"ismfId,omitempty"`
	MoExpDataCounter                     *MoExpDataCounter             `json:"moExpDataCounter,omitempty"`
	SecurityResult                       *SecurityResult               `json:"securityResult,omitempty"`
	PauseCharging                        *bool                         `json:"pauseCharging,omitempty"`
	QosFlowsNotifyList                   []QosFlowNotifyItem           `json:"qosFlowsNotifyList,omitempty"`
	RevokeEbiList                        []int                         `json:"revokeEbiList,omitempty"`
	N4Info                               *N4Information                `json:"n4Info,omitempty"`
	HoPreparationIndication              *bool                         `json:"hoPreparationIndication,omitempty"`
	PsaInfo                              []PsaInformation              `json:"psaInfo,omitempty"`
	UlclBpInfo                           *UlclBpInformation            `json:"ulclBpInfo,omitempty"`
	PcfUeCallbackInfo                    *PcfUeCallbackInfo            `json:"pcfUeCallbackInfo,omitempty"`
	EpsBearerId                          []int                         `json:"epsBearerId,omitempty"`
	AlwaysOnRequested                    *bool                         `json:"alwaysOnRequested,omitempty"`
	N4InfoExt2                           *N4Information                `json:"n4InfoExt2,omitempty"`
	UpSecurityInfo                       *UpSecurityInfo               `json:"upSecurityInfo,omitempty"`
	MaxIntegrityProtectedDataRateUl      MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	MaxIntegrityProtectedDataRateDl      MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	UnknownN1SmInfo                      *RefToBinaryData              `json:"unknownN1SmInfo,omitempty"`
	SecondaryRatUsageReport              []SecondaryRatUsageReport     `json:"secondaryRatUsageReport,omitempty"`
	VSmfServiceInstanceId                string                        `json:"vSmfServiceInstanceId,omitempty"`
	RoamingChargingProfile               *RoamingChargingProfile       `json:"roamingChargingProfile,omitempty"`
	RatType                              RatType                       `json:"ratType,omitempty"`
	VplmnQos                             *VplmnQos                     `json:"vplmnQos,omitempty"`
	AnType                               AccessType                    `json:"anType,omitempty"`
	SupportedFeatures                    string                        `json:"supportedFeatures,omitempty"`
	SmPolicyNotifyInd                    *SmPolicyNotifyInd            `json:"smPolicyNotifyInd,omitempty"`
	UpCnxState                           UpCnxState                    `json:"upCnxState,omitempty"`
	AddUeLocation                        *UserLocation                 `json:"addUeLocation,omitempty"`
	MaReleaseInd                         MaReleaseIndication           `json:"maReleaseInd,omitempty"`
	Pei                                  string                        `json:"pei,omitempty"`
	UeLocation                           *UserLocation                 `json:"ueLocation,omitempty"`
	AnTypeCanBeChanged                   *bool                         `json:"anTypeCanBeChanged,omitempty"`
	SecondaryRatUsageDataReportContainer []string                      `json:"secondaryRatUsageDataReportContainer,omitempty"`
	MaRequestInd                         *bool                         `json:"maRequestInd,omitempty"`
	DlServingPlmnRateCtl                 *int                          `json:"dlServingPlmnRateCtl,omitempty"`
	NotifyList                           []PduSessionNotifyItem        `json:"NotifyList,omitempty"`
	FiveGMmCauseValue                    *int                          `json:"5gMmCauseValue,omitempty"`
	EpsInterworkingInd                   EpsInterworkingIndication     `json:"epsInterworkingInd,omitempty"`
	SecondaryRatUsageInfo                []SecondaryRatUsageInfo       `json:"secondaryRatUsageInfo,omitempty"`
}
