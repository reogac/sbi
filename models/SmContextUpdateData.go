/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:30 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextUpdateData struct {
	AnType                               AccessType                         `json:"anType,omitempty"`
	RanUeInfo                            *RanUeInfo                         `json:"ranUeInfo,omitempty"`
	SmContextStatusUri                   string                             `json:"smContextStatusUri,omitempty"`
	N9ForwardingTunnel                   *TunnelInfo                        `json:"n9ForwardingTunnel,omitempty"`
	EpsBearerSetup                       []string                           `json:"epsBearerSetup,omitempty"`
	N2SmInfoTypeExt1                     N2SmInfoType                       `json:"n2SmInfoTypeExt1,omitempty"`
	Pei                                  string                             `json:"pei,omitempty"`
	ForwardingBearerContexts             []string                           `json:"forwardingBearerContexts,omitempty"`
	UeLocation                           *UserLocation                      `json:"ueLocation,omitempty"`
	AnTypeToReactivate                   AccessType                         `json:"anTypeToReactivate,omitempty"`
	FailedToBeSwitched                   *bool                              `json:"failedToBeSwitched,omitempty"`
	N2SmInfo                             *RefToBinaryData                   `json:"n2SmInfo,omitempty"`
	N9UlForwardingTnlList                []IndirectDataForwardingTunnelInfo `json:"n9UlForwardingTnlList,omitempty"`
	N9DlForwardingTunnel                 *TunnelInfo                        `json:"n9DlForwardingTunnel,omitempty"`
	N9InactivityTimer                    *int                               `json:"n9InactivityTimer,omitempty"`
	AnTypeCanBeChanged                   *bool                              `json:"anTypeCanBeChanged,omitempty"`
	SmPolicyNotifyInd                    *SmPolicyNotifyInd                 `json:"smPolicyNotifyInd,omitempty"`
	ServingNetwork                       *PlmnIdNid                         `json:"servingNetwork,omitempty"`
	BackupAmfInfo                        []BackupAmfInfo                    `json:"backupAmfInfo,omitempty"`
	RanTransportNets                     []string                           `json:"ranTransportNets,omitempty"`
	TraceData                            *TraceData                         `json:"traceData,omitempty"`
	MaRequestInd                         *bool                              `json:"maRequestInd,omitempty"`
	ExemptionInd                         *ExemptionInd                      `json:"exemptionInd,omitempty"`
	ExtendedNasSmTimerInd                *bool                              `json:"extendedNasSmTimerInd,omitempty"`
	MoExpDataCounter                     *MoExpDataCounter                  `json:"moExpDataCounter,omitempty"`
	DdnFailureSubs                       *DdnFailureSubs                    `json:"ddnFailureSubs,omitempty"`
	SatelliteBackhaulCat                 SatelliteBackhaulCategory          `json:"satelliteBackhaulCat,omitempty"`
	ServingNfId                          string                             `json:"servingNfId,omitempty"`
	Guami                                *Guami                             `json:"guami,omitempty"`
	TargetId                             *NgRanTargetId                     `json:"targetId,omitempty"`
	UeTimeZone                           string                             `json:"ueTimeZone,omitempty"`
	N2SmInfoType                         N2SmInfoType                       `json:"n2SmInfoType,omitempty"`
	N2SmInfoExt1                         *RefToBinaryData                   `json:"n2SmInfoExt1,omitempty"`
	ForwardingFTeid                      string                             `json:"forwardingFTeid,omitempty"`
	UpCnxState                           UpCnxState                         `json:"upCnxState,omitempty"`
	RatType                              RatType                            `json:"ratType,omitempty"`
	TargetServingNfId                    string                             `json:"targetServingNfId,omitempty"`
	Release                              *bool                              `json:"release,omitempty"`
	PresenceInLadn                       PresenceState                      `json:"presenceInLadn,omitempty"`
	ToBeSwitched                         *bool                              `json:"toBeSwitched,omitempty"`
	MaNwUpgradeInd                       *bool                              `json:"maNwUpgradeInd,omitempty"`
	RevokeEbiList                        []int                              `json:"revokeEbiList,omitempty"`
	MaReleaseInd                         MaReleaseIndication                `json:"maReleaseInd,omitempty"`
	SNssai                               *Snssai                            `json:"sNssai,omitempty"`
	AdditionalAnType                     AccessType                         `json:"additionalAnType,omitempty"`
	HoState                              HoState                            `json:"hoState,omitempty"`
	N9DlForwardingTnlList                []IndirectDataForwardingTunnelInfo `json:"n9DlForwardingTnlList,omitempty"`
	FiveGMmCauseValue                    *int                               `json:"5gMmCauseValue,omitempty"`
	EpsInterworkingInd                   EpsInterworkingIndication          `json:"epsInterworkingInd,omitempty"`
	SecondaryRatUsageDataReportContainer []string                           `json:"secondaryRatUsageDataReportContainer,omitempty"`
	N1SmMsg                              *RefToBinaryData                   `json:"n1SmMsg,omitempty"`
	DataForwarding                       *bool                              `json:"dataForwarding,omitempty"`
	NgApCause                            *NgApCause                         `json:"ngApCause,omitempty"`
	SupportedFeatures                    string                             `json:"supportedFeatures,omitempty"`
	PcfUeCallbackInfo                    *PcfUeCallbackInfo                 `json:"pcfUeCallbackInfo,omitempty"`
	AddUeLocation                        *UserLocation                      `json:"addUeLocation,omitempty"`
	Cause                                Cause                              `json:"cause,omitempty"`
	SkipN2PduSessionResRelInd            *bool                              `json:"skipN2PduSessionResRelInd,omitempty"`
}
