/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:27 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContext struct {
	PcfAmpBindingInfo          string                          `json:"pcfAmpBindingInfo,omitempty"`
	DisasterPlmn               *PlmnId                         `json:"disasterPlmn,omitempty"`
	PcfId                      string                          `json:"pcfId,omitempty"`
	SupportedCodecList         []string                        `json:"supportedCodecList,omitempty"`
	RestrictedPrimaryRatList   []string                        `json:"restrictedPrimaryRatList,omitempty"`
	AstiDistributionIndication *bool                           `json:"astiDistributionIndication,omitempty"`
	SnpnOnboardInd             *bool                           `json:"snpnOnboardInd,omitempty"`
	PcfUeAmbr                  *Ambr                           `json:"pcfUeAmbr,omitempty"`
	PcfBinding                 SbiBindingLevel                 `json:"pcfBinding,omitempty"`
	PcfUePolicyUri             string                          `json:"pcfUePolicyUri,omitempty"`
	RestrictedCoreNwTypeList   []string                        `json:"restrictedCoreNwTypeList,omitempty"`
	PraInAmPolicy              map[string]PresenceInfo         `json:"praInAmPolicy,omitempty"`
	PcfSetId                   string                          `json:"pcfSetId,omitempty"`
	CagData                    *CagData                        `json:"cagData,omitempty"`
	PcfUepBindingInfo          string                          `json:"pcfUepBindingInfo,omitempty"`
	UpdpSubscriptionData       *UpdpSubscriptionData           `json:"updpSubscriptionData,omitempty"`
	AmPolicyReqTriggerList     []string                        `json:"amPolicyReqTriggerList,omitempty"`
	RestrictedRatList          []string                        `json:"restrictedRatList,omitempty"`
	ManagementMdtInd           *bool                           `json:"managementMdtInd,omitempty"`
	ImmediateMdtConf           *ImmediateMdtConf               `json:"immediateMdtConf,omitempty"`
	EcRestrictionDataNb        *bool                           `json:"ecRestrictionDataNb,omitempty"`
	IabOperationAllowed        *bool                           `json:"iabOperationAllowed,omitempty"`
	ProseContext               *ProseContext                   `json:"proseContext,omitempty"`
	PcfUeCallbackInfo          *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	SmfSelInfo                 *SmfSelectionData               `json:"smfSelInfo,omitempty"`
	HNwPubKeyId                *int                            `json:"hNwPubKeyId,omitempty"`
	SubUeSliceMbrList          map[string]SliceMbr             `json:"subUeSliceMbrList,omitempty"`
	SmsfId                     string                          `json:"smsfId,omitempty"`
	SessionContextList         []PduSessionContext             `json:"sessionContextList,omitempty"`
	AusfGroupId                string                          `json:"ausfGroupId,omitempty"`
	UsedRfsp                   *int                            `json:"usedRfsp,omitempty"`
	HpcfSetId                  string                          `json:"hpcfSetId,omitempty"`
	EventSubscriptionList      []ExtAmfEventSubscription       `json:"eventSubscriptionList,omitempty"`
	CMsisdn                    string                          `json:"cMsisdn,omitempty"`
	PraInUePolicy              map[string]PresenceInfo         `json:"praInUePolicy,omitempty"`
	RoutingIndicator           string                          `json:"routingIndicator,omitempty"`
	DrxParameter               string                          `json:"drxParameter,omitempty"`
	PcfRfsp                    *int                            `json:"pcfRfsp,omitempty"`
	EpsInterworkingInfo        *EpsInterworkingInfo            `json:"epsInterworkingInfo,omitempty"`
	SmsfServiceSetId           string                          `json:"smsfServiceSetId,omitempty"`
	WlServAreaRes              *WirelineServiceAreaRestriction `json:"wlServAreaRes,omitempty"`
	PcfGroupId                 string                          `json:"pcfGroupId,omitempty"`
	SubUeAmbr                  *Ambr                           `json:"subUeAmbr,omitempty"`
	PcfAmPolicyUri             string                          `json:"pcfAmPolicyUri,omitempty"`
	HpcfId                     string                          `json:"hpcfId,omitempty"`
	SmallDataRateStatusInfos   []SmallDataRateStatusInfo       `json:"smallDataRateStatusInfos,omitempty"`
	UsedServiceAreaRestriction *ServiceAreaRestriction         `json:"usedServiceAreaRestriction,omitempty"`
	PcfUepServiceSetId         string                          `json:"pcfUepServiceSetId,omitempty"`
	RestrictedSecondaryRatList []string                        `json:"restrictedSecondaryRatList,omitempty"`
	RedCapInd                  *bool                           `json:"redCapInd,omitempty"`
	TsErrorBudget              *int                            `json:"tsErrorBudget,omitempty"`
	PcfUeSliceMbrList          map[string]SliceMbr             `json:"pcfUeSliceMbrList,omitempty"`
	DisasterRoamingInd         *bool                           `json:"disasterRoamingInd,omitempty"`
	SubRfsp                    *int                            `json:"subRfsp,omitempty"`
	FiveGMmCapability          string                          `json:"5gMmCapability,omitempty"`
	MmContextList              []MmContext                     `json:"mmContextList,omitempty"`
	SmsfSetId                  string                          `json:"smsfSetId,omitempty"`
	AsTimeDisParam             *AsTimeDistributionParam        `json:"asTimeDisParam,omitempty"`
	GpsiList                   []string                        `json:"gpsiList,omitempty"`
	MsClassmark2               string                          `json:"msClassmark2,omitempty"`
	LteCatMInd                 *bool                           `json:"lteCatMInd,omitempty"`
	UePositioningCap           string                          `json:"uePositioningCap,omitempty"`
	GroupList                  []string                        `json:"groupList,omitempty"`
	PcfAmpServiceSetId         string                          `json:"pcfAmpServiceSetId,omitempty"`
	ServiceAreaRestriction     *ServiceAreaRestriction         `json:"serviceAreaRestriction,omitempty"`
	StnSr                      string                          `json:"stnSr,omitempty"`
	SupiUnauthInd              *bool                           `json:"supiUnauthInd,omitempty"`
	ForbiddenAreaList          []Area                          `json:"forbiddenAreaList,omitempty"`
	TraceData                  *TraceData                      `json:"traceData,omitempty"`
	MoExpDataCounter           *MoExpDataCounter               `json:"moExpDataCounter,omitempty"`
	EcRestrictionDataWb        *EcRestrictionDataWb            `json:"ecRestrictionDataWb,omitempty"`
	AnalyticsSubscriptionList  []AnalyticsSubscription         `json:"analyticsSubscriptionList,omitempty"`
	SmsfBindingInfo            string                          `json:"smsfBindingInfo,omitempty"`
	Supi                       string                          `json:"supi,omitempty"`
	Pei                        string                          `json:"pei,omitempty"`
	UdmGroupId                 string                          `json:"udmGroupId,omitempty"`
	SeafData                   *SeafData                       `json:"seafData,omitempty"`
	UePolicyReqTriggerList     []string                        `json:"uePolicyReqTriggerList,omitempty"`
	ServiceGapExpiryTime       string                          `json:"serviceGapExpiryTime,omitempty"`
	V2xContext                 *V2xContext                     `json:"v2xContext,omitempty"`
	SmPolicyNotifyPduList      []PduSessionInfo                `json:"smPolicyNotifyPduList,omitempty"`
}
