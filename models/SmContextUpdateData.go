/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextUpdateData struct {
	UpCnxState                           UpCnxState                         `json:"upCnxState,omitempty"`
	FiveGMmCauseValue                    *int                               `json:"5gMmCauseValue,omitempty"`
	ForwardingFTeid                      string                             `json:"forwardingFTeid,omitempty"`
	SmPolicyNotifyInd                    *SmPolicyNotifyInd                 `json:"smPolicyNotifyInd,omitempty"`
	AnTypeCanBeChanged                   *bool                              `json:"anTypeCanBeChanged,omitempty"`
	MoExpDataCounter                     *MoExpDataCounter                  `json:"moExpDataCounter,omitempty"`
	MaReleaseInd                         MaReleaseIndication                `json:"maReleaseInd,omitempty"`
	SupportedFeatures                    string                             `json:"supportedFeatures,omitempty"`
	UeLocation                           *UserLocation                      `json:"ueLocation,omitempty"`
	TargetId                             *NgRanTargetId                     `json:"targetId,omitempty"`
	N9UlForwardingTnlList                []IndirectDataForwardingTunnelInfo `json:"n9UlForwardingTnlList,omitempty"`
	N9InactivityTimer                    *int                               `json:"n9InactivityTimer,omitempty"`
	NgApCause                            *NgApCause                         `json:"ngApCause,omitempty"`
	N2SmInfoTypeExt1                     N2SmInfoType                       `json:"n2SmInfoTypeExt1,omitempty"`
	DdnFailureSubs                       *DdnFailureSubs                    `json:"ddnFailureSubs,omitempty"`
	BackupAmfInfo                        []BackupAmfInfo                    `json:"backupAmfInfo,omitempty"`
	AdditionalAnType                     AccessType                         `json:"additionalAnType,omitempty"`
	HoState                              HoState                            `json:"hoState,omitempty"`
	FailedToBeSwitched                   *bool                              `json:"failedToBeSwitched,omitempty"`
	N1SmMsg                              *RefToBinaryData                   `json:"n1SmMsg,omitempty"`
	Guami                                *Guami                             `json:"guami,omitempty"`
	RanUeInfo                            *RanUeInfo                         `json:"ranUeInfo,omitempty"`
	N2SmInfoExt1                         *RefToBinaryData                   `json:"n2SmInfoExt1,omitempty"`
	ExtendedNasSmTimerInd                *bool                              `json:"extendedNasSmTimerInd,omitempty"`
	Pei                                  string                             `json:"pei,omitempty"`
	N9ForwardingTunnel                   *TunnelInfo                        `json:"n9ForwardingTunnel,omitempty"`
	N9DlForwardingTnlList                []IndirectDataForwardingTunnelInfo `json:"n9DlForwardingTnlList,omitempty"`
	SkipN2PduSessionResRelInd            *bool                              `json:"skipN2PduSessionResRelInd,omitempty"`
	SatelliteBackhaulCat                 SatelliteBackhaulCategory          `json:"satelliteBackhaulCat,omitempty"`
	RanTransportNets                     []string                           `json:"ranTransportNets,omitempty"`
	N2SmInfo                             *RefToBinaryData                   `json:"n2SmInfo,omitempty"`
	RevokeEbiList                        []int                              `json:"revokeEbiList,omitempty"`
	EpsInterworkingInd                   EpsInterworkingIndication          `json:"epsInterworkingInd,omitempty"`
	ExemptionInd                         *ExemptionInd                      `json:"exemptionInd,omitempty"`
	DataForwarding                       *bool                              `json:"dataForwarding,omitempty"`
	EpsBearerSetup                       []string                           `json:"epsBearerSetup,omitempty"`
	Cause                                Cause                              `json:"cause,omitempty"`
	ForwardingBearerContexts             []string                           `json:"forwardingBearerContexts,omitempty"`
	AnType                               AccessType                         `json:"anType,omitempty"`
	AnTypeToReactivate                   AccessType                         `json:"anTypeToReactivate,omitempty"`
	AddUeLocation                        *UserLocation                      `json:"addUeLocation,omitempty"`
	SmContextStatusUri                   string                             `json:"smContextStatusUri,omitempty"`
	TargetServingNfId                    string                             `json:"targetServingNfId,omitempty"`
	TraceData                            *TraceData                         `json:"traceData,omitempty"`
	MaNwUpgradeInd                       *bool                              `json:"maNwUpgradeInd,omitempty"`
	MaRequestInd                         *bool                              `json:"maRequestInd,omitempty"`
	PcfUeCallbackInfo                    *PcfUeCallbackInfo                 `json:"pcfUeCallbackInfo,omitempty"`
	ServingNfId                          string                             `json:"servingNfId,omitempty"`
	RatType                              RatType                            `json:"ratType,omitempty"`
	PresenceInLadn                       PresenceState                      `json:"presenceInLadn,omitempty"`
	N9DlForwardingTunnel                 *TunnelInfo                        `json:"n9DlForwardingTunnel,omitempty"`
	Release                              *bool                              `json:"release,omitempty"`
	N2SmInfoType                         N2SmInfoType                       `json:"n2SmInfoType,omitempty"`
	ToBeSwitched                         *bool                              `json:"toBeSwitched,omitempty"`
	SNssai                               *Snssai                            `json:"sNssai,omitempty"`
	ServingNetwork                       *PlmnIdNid                         `json:"servingNetwork,omitempty"`
	UeTimeZone                           string                             `json:"ueTimeZone,omitempty"`
	SecondaryRatUsageDataReportContainer []string                           `json:"secondaryRatUsageDataReportContainer,omitempty"`
}
