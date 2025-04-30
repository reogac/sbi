/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Apr 30 14:54:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextUpdateData struct {
	TargetServingNfId                    string                             `json:"targetServingNfId,omitempty"`
	AnType                               AccessType                         `json:"anType,omitempty"`
	N9InactivityTimer                    *int                               `json:"n9InactivityTimer,omitempty"`
	AnTypeCanBeChanged                   *bool                              `json:"anTypeCanBeChanged,omitempty"`
	FailedToBeSwitched                   *bool                              `json:"failedToBeSwitched,omitempty"`
	RevokeEbiList                        []int                              `json:"revokeEbiList,omitempty"`
	PresenceInLadn                       PresenceState                      `json:"presenceInLadn,omitempty"`
	SmContextStatusUri                   string                             `json:"smContextStatusUri,omitempty"`
	Pei                                  string                             `json:"pei,omitempty"`
	RanTransportNets                     []string                           `json:"ranTransportNets,omitempty"`
	Cause                                Cause                              `json:"cause,omitempty"`
	SupportedFeatures                    string                             `json:"supportedFeatures,omitempty"`
	N2SmInfoExt1                         *RefToBinaryData                   `json:"n2SmInfoExt1,omitempty"`
	N2SmInfoTypeExt1                     N2SmInfoType                       `json:"n2SmInfoTypeExt1,omitempty"`
	PcfUeCallbackInfo                    *PcfUeCallbackInfo                 `json:"pcfUeCallbackInfo,omitempty"`
	Guami                                *Guami                             `json:"guami,omitempty"`
	Release                              *bool                              `json:"release,omitempty"`
	ExemptionInd                         *ExemptionInd                      `json:"exemptionInd,omitempty"`
	ExtendedNasSmTimerInd                *bool                              `json:"extendedNasSmTimerInd,omitempty"`
	ForwardingBearerContexts             []string                           `json:"forwardingBearerContexts,omitempty"`
	AdditionalAnType                     AccessType                         `json:"additionalAnType,omitempty"`
	AnTypeToReactivate                   AccessType                         `json:"anTypeToReactivate,omitempty"`
	DataForwarding                       *bool                              `json:"dataForwarding,omitempty"`
	SatelliteBackhaulCat                 SatelliteBackhaulCategory          `json:"satelliteBackhaulCat,omitempty"`
	N2SmInfo                             *RefToBinaryData                   `json:"n2SmInfo,omitempty"`
	N9DlForwardingTunnel                 *TunnelInfo                        `json:"n9DlForwardingTunnel,omitempty"`
	EpsInterworkingInd                   EpsInterworkingIndication          `json:"epsInterworkingInd,omitempty"`
	ForwardingFTeid                      string                             `json:"forwardingFTeid,omitempty"`
	SmPolicyNotifyInd                    *SmPolicyNotifyInd                 `json:"smPolicyNotifyInd,omitempty"`
	AddUeLocation                        *UserLocation                      `json:"addUeLocation,omitempty"`
	UpCnxState                           UpCnxState                         `json:"upCnxState,omitempty"`
	HoState                              HoState                            `json:"hoState,omitempty"`
	MaNwUpgradeInd                       *bool                              `json:"maNwUpgradeInd,omitempty"`
	MaReleaseInd                         MaReleaseIndication                `json:"maReleaseInd,omitempty"`
	MaRequestInd                         *bool                              `json:"maRequestInd,omitempty"`
	SecondaryRatUsageDataReportContainer []string                           `json:"secondaryRatUsageDataReportContainer,omitempty"`
	N9UlForwardingTnlList                []IndirectDataForwardingTunnelInfo `json:"n9UlForwardingTnlList,omitempty"`
	NgApCause                            *NgApCause                         `json:"ngApCause,omitempty"`
	TraceData                            *TraceData                         `json:"traceData,omitempty"`
	UeTimeZone                           string                             `json:"ueTimeZone,omitempty"`
	N2SmInfoType                         N2SmInfoType                       `json:"n2SmInfoType,omitempty"`
	N9DlForwardingTnlList                []IndirectDataForwardingTunnelInfo `json:"n9DlForwardingTnlList,omitempty"`
	EpsBearerSetup                       []string                           `json:"epsBearerSetup,omitempty"`
	ServingNetwork                       *PlmnIdNid                         `json:"servingNetwork,omitempty"`
	N9ForwardingTunnel                   *TunnelInfo                        `json:"n9ForwardingTunnel,omitempty"`
	TargetId                             *NgRanTargetId                     `json:"targetId,omitempty"`
	FiveGMmCauseValue                    *int                               `json:"5gMmCauseValue,omitempty"`
	MoExpDataCounter                     *MoExpDataCounter                  `json:"moExpDataCounter,omitempty"`
	SkipN2PduSessionResRelInd            *bool                              `json:"skipN2PduSessionResRelInd,omitempty"`
	BackupAmfInfo                        []BackupAmfInfo                    `json:"backupAmfInfo,omitempty"`
	ToBeSwitched                         *bool                              `json:"toBeSwitched,omitempty"`
	UeLocation                           *UserLocation                      `json:"ueLocation,omitempty"`
	N1SmMsg                              *RefToBinaryData                   `json:"n1SmMsg,omitempty"`
	SNssai                               *Snssai                            `json:"sNssai,omitempty"`
	DdnFailureSubs                       *DdnFailureSubs                    `json:"ddnFailureSubs,omitempty"`
	ServingNfId                          string                             `json:"servingNfId,omitempty"`
	RatType                              RatType                            `json:"ratType,omitempty"`
}
