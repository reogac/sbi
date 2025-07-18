/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HsmfUpdateData struct {
	Pei                                  string                        `json:"pei,omitempty"`
	RatType                              RatType                       `json:"ratType,omitempty"`
	QosFlowsNotifyList                   []QosFlowNotifyItem           `json:"qosFlowsNotifyList,omitempty"`
	SecondaryRatUsageReport              []SecondaryRatUsageReport     `json:"secondaryRatUsageReport,omitempty"`
	N4Info                               *N4Information                `json:"n4Info,omitempty"`
	UpSecurityInfo                       *UpSecurityInfo               `json:"upSecurityInfo,omitempty"`
	SecondaryRatUsageInfo                []SecondaryRatUsageInfo       `json:"secondaryRatUsageInfo,omitempty"`
	MaReleaseInd                         MaReleaseIndication           `json:"maReleaseInd,omitempty"`
	MaRequestInd                         *bool                         `json:"maRequestInd,omitempty"`
	AnType                               AccessType                    `json:"anType,omitempty"`
	AddUeLocation                        *UserLocation                 `json:"addUeLocation,omitempty"`
	RoamingChargingProfile               *RoamingChargingProfile       `json:"roamingChargingProfile,omitempty"`
	AmfNfId                              string                        `json:"amfNfId,omitempty"`
	Guami                                *Guami                        `json:"guami,omitempty"`
	PcfUeCallbackInfo                    *PcfUeCallbackInfo            `json:"pcfUeCallbackInfo,omitempty"`
	DnaiList                             []string                      `json:"dnaiList,omitempty"`
	SmPolicyNotifyInd                    *SmPolicyNotifyInd            `json:"smPolicyNotifyInd,omitempty"`
	N1SmInfoFromUe                       *RefToBinaryData              `json:"n1SmInfoFromUe,omitempty"`
	HoPreparationIndication              *bool                         `json:"hoPreparationIndication,omitempty"`
	VsmfId                               string                        `json:"vsmfId,omitempty"`
	SatelliteBackhaulCat                 SatelliteBackhaulCategory     `json:"satelliteBackhaulCat,omitempty"`
	MaxIntegrityProtectedDataRateDl      MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	RequestIndication                    RequestIndication             `json:"requestIndication"`
	RevokeEbiList                        []int                         `json:"revokeEbiList,omitempty"`
	AnTypeCanBeChanged                   *bool                         `json:"anTypeCanBeChanged,omitempty"`
	UnavailableAccessInd                 UnavailableAccessIndication   `json:"unavailableAccessInd,omitempty"`
	N4InfoExt1                           *N4Information                `json:"n4InfoExt1,omitempty"`
	ISmfServiceInstanceId                string                        `json:"iSmfServiceInstanceId,omitempty"`
	DlServingPlmnRateCtl                 *int                          `json:"dlServingPlmnRateCtl,omitempty"`
	IcnTunnelInfo                        *TunnelInfo                   `json:"icnTunnelInfo,omitempty"`
	ServingNetwork                       *PlmnIdNid                    `json:"servingNetwork,omitempty"`
	PsaInfo                              []PsaInformation              `json:"psaInfo,omitempty"`
	SupportedFeatures                    string                        `json:"supportedFeatures,omitempty"`
	QosFlowsRelNotifyList                []QosFlowItem                 `json:"qosFlowsRelNotifyList,omitempty"`
	Cause                                Cause                         `json:"cause,omitempty"`
	FiveGMmCauseValue                    *int                          `json:"5gMmCauseValue,omitempty"`
	UlclBpInfo                           *UlclBpInformation            `json:"ulclBpInfo,omitempty"`
	SecondaryRatUsageDataReportContainer []string                      `json:"secondaryRatUsageDataReportContainer,omitempty"`
	DisasterRoamingInd                   *DisasterRoamingInd           `json:"disasterRoamingInd,omitempty"`
	AlwaysOnRequested                    *bool                         `json:"alwaysOnRequested,omitempty"`
	MaxIntegrityProtectedDataRateUl      MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	UeLocation                           *UserLocation                 `json:"ueLocation,omitempty"`
	Pti                                  *int                          `json:"pti,omitempty"`
	VsmfPduSessionUri                    string                        `json:"vsmfPduSessionUri,omitempty"`
	AdditionalAnType                     AccessType                    `json:"additionalAnType,omitempty"`
	EpsBearerId                          []int                         `json:"epsBearerId,omitempty"`
	IsmfId                               string                        `json:"ismfId,omitempty"`
	SecurityResult                       *SecurityResult               `json:"securityResult,omitempty"`
	EpsInterworkingInd                   EpsInterworkingIndication     `json:"epsInterworkingInd,omitempty"`
	PauseCharging                        *bool                         `json:"pauseCharging,omitempty"`
	N4InfoExt2                           *N4Information                `json:"n4InfoExt2,omitempty"`
	IsmfPduSessionUri                    string                        `json:"ismfPduSessionUri,omitempty"`
	VcnTunnelInfo                        *TunnelInfo                   `json:"vcnTunnelInfo,omitempty"`
	NgApCause                            *NgApCause                    `json:"ngApCause,omitempty"`
	VSmfServiceInstanceId                string                        `json:"vSmfServiceInstanceId,omitempty"`
	AdditionalCnTunnelInfo               *TunnelInfo                   `json:"additionalCnTunnelInfo,omitempty"`
	UnknownN1SmInfo                      *RefToBinaryData              `json:"unknownN1SmInfo,omitempty"`
	NotifyList                           []PduSessionNotifyItem        `json:"NotifyList,omitempty"`
	VplmnQos                             *VplmnQos                     `json:"vplmnQos,omitempty"`
	UpCnxState                           UpCnxState                    `json:"upCnxState,omitempty"`
	UeTimeZone                           string                        `json:"ueTimeZone,omitempty"`
	MaNwUpgradeInd                       *bool                         `json:"maNwUpgradeInd,omitempty"`
	PresenceInLadn                       PresenceState                 `json:"presenceInLadn,omitempty"`
	MoExpDataCounter                     *MoExpDataCounter             `json:"moExpDataCounter,omitempty"`
}
