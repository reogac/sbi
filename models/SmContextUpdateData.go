/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextUpdateData struct {
	AddUeLocation                        *UserLocation                      `json:"addUeLocation,omitempty"`
	MaRequestInd                         *bool                              `json:"maRequestInd,omitempty"`
	RanUeInfo                            *RanUeInfo                         `json:"ranUeInfo,omitempty"`
	SupportedFeatures                    string                             `json:"supportedFeatures,omitempty"`
	ToBeSwitched                         *bool                              `json:"toBeSwitched,omitempty"`
	DataForwarding                       *bool                              `json:"dataForwarding,omitempty"`
	N9ForwardingTunnel                   *TunnelInfo                        `json:"n9ForwardingTunnel,omitempty"`
	MaNwUpgradeInd                       *bool                              `json:"maNwUpgradeInd,omitempty"`
	MoExpDataCounter                     *MoExpDataCounter                  `json:"moExpDataCounter,omitempty"`
	ExtendedNasSmTimerInd                *bool                              `json:"extendedNasSmTimerInd,omitempty"`
	RanTransportNets                     []string                           `json:"ranTransportNets,omitempty"`
	FailedToBeSwitched                   *bool                              `json:"failedToBeSwitched,omitempty"`
	SmContextStatusUri                   string                             `json:"smContextStatusUri,omitempty"`
	N9DlForwardingTnlList                []IndirectDataForwardingTunnelInfo `json:"n9DlForwardingTnlList,omitempty"`
	N9InactivityTimer                    *int                               `json:"n9InactivityTimer,omitempty"`
	N2SmInfoExt1                         *RefToBinaryData                   `json:"n2SmInfoExt1,omitempty"`
	ForwardingBearerContexts             []string                           `json:"forwardingBearerContexts,omitempty"`
	TargetId                             *NgRanTargetId                     `json:"targetId,omitempty"`
	ForwardingFTeid                      string                             `json:"forwardingFTeid,omitempty"`
	SkipN2PduSessionResRelInd            *bool                              `json:"skipN2PduSessionResRelInd,omitempty"`
	AnType                               AccessType                         `json:"anType,omitempty"`
	UpCnxState                           UpCnxState                         `json:"upCnxState,omitempty"`
	NgApCause                            *NgApCause                         `json:"ngApCause,omitempty"`
	Guami                                *Guami                             `json:"guami,omitempty"`
	SatelliteBackhaulCat                 SatelliteBackhaulCategory          `json:"satelliteBackhaulCat,omitempty"`
	SNssai                               *Snssai                            `json:"sNssai,omitempty"`
	DdnFailureSubs                       *DdnFailureSubs                    `json:"ddnFailureSubs,omitempty"`
	HoState                              HoState                            `json:"hoState,omitempty"`
	N2SmInfo                             *RefToBinaryData                   `json:"n2SmInfo,omitempty"`
	N9DlForwardingTunnel                 *TunnelInfo                        `json:"n9DlForwardingTunnel,omitempty"`
	TraceData                            *TraceData                         `json:"traceData,omitempty"`
	SmPolicyNotifyInd                    *SmPolicyNotifyInd                 `json:"smPolicyNotifyInd,omitempty"`
	UeLocation                           *UserLocation                      `json:"ueLocation,omitempty"`
	N2SmInfoTypeExt1                     N2SmInfoType                       `json:"n2SmInfoTypeExt1,omitempty"`
	SecondaryRatUsageDataReportContainer []string                           `json:"secondaryRatUsageDataReportContainer,omitempty"`
	BackupAmfInfo                        []BackupAmfInfo                    `json:"backupAmfInfo,omitempty"`
	AdditionalAnType                     AccessType                         `json:"additionalAnType,omitempty"`
	RatType                              RatType                            `json:"ratType,omitempty"`
	N1SmMsg                              *RefToBinaryData                   `json:"n1SmMsg,omitempty"`
	TargetServingNfId                    string                             `json:"targetServingNfId,omitempty"`
	MaReleaseInd                         MaReleaseIndication                `json:"maReleaseInd,omitempty"`
	ServingNfId                          string                             `json:"servingNfId,omitempty"`
	UeTimeZone                           string                             `json:"ueTimeZone,omitempty"`
	Cause                                Cause                              `json:"cause,omitempty"`
	Pei                                  string                             `json:"pei,omitempty"`
	ServingNetwork                       *PlmnIdNid                         `json:"servingNetwork,omitempty"`
	AnTypeToReactivate                   AccessType                         `json:"anTypeToReactivate,omitempty"`
	FiveGMmCauseValue                    *int                               `json:"5gMmCauseValue,omitempty"`
	EpsInterworkingInd                   EpsInterworkingIndication          `json:"epsInterworkingInd,omitempty"`
	ExemptionInd                         *ExemptionInd                      `json:"exemptionInd,omitempty"`
	PcfUeCallbackInfo                    *PcfUeCallbackInfo                 `json:"pcfUeCallbackInfo,omitempty"`
	N2SmInfoType                         N2SmInfoType                       `json:"n2SmInfoType,omitempty"`
	N9UlForwardingTnlList                []IndirectDataForwardingTunnelInfo `json:"n9UlForwardingTnlList,omitempty"`
	EpsBearerSetup                       []string                           `json:"epsBearerSetup,omitempty"`
	RevokeEbiList                        []int                              `json:"revokeEbiList,omitempty"`
	AnTypeCanBeChanged                   *bool                              `json:"anTypeCanBeChanged,omitempty"`
	PresenceInLadn                       PresenceState                      `json:"presenceInLadn,omitempty"`
	Release                              *bool                              `json:"release,omitempty"`
}
