/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContext struct {
	MsClassmark2               string                          `json:"msClassmark2,omitempty"`
	SupportedCodecList         []string                        `json:"supportedCodecList,omitempty"`
	ProseContext               *ProseContext                   `json:"proseContext,omitempty"`
	PcfAmpBindingInfo          string                          `json:"pcfAmpBindingInfo,omitempty"`
	HNwPubKeyId                *int                            `json:"hNwPubKeyId,omitempty"`
	HpcfId                     string                          `json:"hpcfId,omitempty"`
	StnSr                      string                          `json:"stnSr,omitempty"`
	CMsisdn                    string                          `json:"cMsisdn,omitempty"`
	PraInUePolicy              map[string]PresenceInfo         `json:"praInUePolicy,omitempty"`
	UePositioningCap           string                          `json:"uePositioningCap,omitempty"`
	GpsiList                   []string                        `json:"gpsiList,omitempty"`
	PcfUepServiceSetId         string                          `json:"pcfUepServiceSetId,omitempty"`
	UePolicyReqTriggerList     []string                        `json:"uePolicyReqTriggerList,omitempty"`
	ManagementMdtInd           *bool                           `json:"managementMdtInd,omitempty"`
	ImmediateMdtConf           *ImmediateMdtConf               `json:"immediateMdtConf,omitempty"`
	UpdpSubscriptionData       *UpdpSubscriptionData           `json:"updpSubscriptionData,omitempty"`
	AusfGroupId                string                          `json:"ausfGroupId,omitempty"`
	GroupList                  []string                        `json:"groupList,omitempty"`
	DrxParameter               string                          `json:"drxParameter,omitempty"`
	PcfBinding                 SbiBindingLevel                 `json:"pcfBinding,omitempty"`
	EcRestrictionDataNb        *bool                           `json:"ecRestrictionDataNb,omitempty"`
	PcfGroupId                 string                          `json:"pcfGroupId,omitempty"`
	PcfRfsp                    *int                            `json:"pcfRfsp,omitempty"`
	SmsfId                     string                          `json:"smsfId,omitempty"`
	SmfSelInfo                 *SmfSelectionData               `json:"smfSelInfo,omitempty"`
	PcfAmPolicyUri             string                          `json:"pcfAmPolicyUri,omitempty"`
	EcRestrictionDataWb        *EcRestrictionDataWb            `json:"ecRestrictionDataWb,omitempty"`
	AnalyticsSubscriptionList  []AnalyticsSubscription         `json:"analyticsSubscriptionList,omitempty"`
	PcfUeSliceMbrList          map[string]SliceMbr             `json:"pcfUeSliceMbrList,omitempty"`
	PraInAmPolicy              map[string]PresenceInfo         `json:"praInAmPolicy,omitempty"`
	SmPolicyNotifyPduList      []PduSessionInfo                `json:"smPolicyNotifyPduList,omitempty"`
	PcfUeCallbackInfo          *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	RestrictedSecondaryRatList []string                        `json:"restrictedSecondaryRatList,omitempty"`
	SmsfBindingInfo            string                          `json:"smsfBindingInfo,omitempty"`
	SubUeSliceMbrList          map[string]SliceMbr             `json:"subUeSliceMbrList,omitempty"`
	ServiceAreaRestriction     *ServiceAreaRestriction         `json:"serviceAreaRestriction,omitempty"`
	RestrictedPrimaryRatList   []string                        `json:"restrictedPrimaryRatList,omitempty"`
	MoExpDataCounter           *MoExpDataCounter               `json:"moExpDataCounter,omitempty"`
	SmsfServiceSetId           string                          `json:"smsfServiceSetId,omitempty"`
	DisasterRoamingInd         *bool                           `json:"disasterRoamingInd,omitempty"`
	PcfSetId                   string                          `json:"pcfSetId,omitempty"`
	MmContextList              []MmContext                     `json:"mmContextList,omitempty"`
	ServiceGapExpiryTime       string                          `json:"serviceGapExpiryTime,omitempty"`
	RestrictedCoreNwTypeList   []string                        `json:"restrictedCoreNwTypeList,omitempty"`
	CagData                    *CagData                        `json:"cagData,omitempty"`
	TraceData                  *TraceData                      `json:"traceData,omitempty"`
	V2xContext                 *V2xContext                     `json:"v2xContext,omitempty"`
	SupiUnauthInd              *bool                           `json:"supiUnauthInd,omitempty"`
	UdmGroupId                 string                          `json:"udmGroupId,omitempty"`
	AmPolicyReqTriggerList     []string                        `json:"amPolicyReqTriggerList,omitempty"`
	RoutingIndicator           string                          `json:"routingIndicator,omitempty"`
	PcfUeAmbr                  *Ambr                           `json:"pcfUeAmbr,omitempty"`
	PcfId                      string                          `json:"pcfId,omitempty"`
	SmsfSetId                  string                          `json:"smsfSetId,omitempty"`
	WlServAreaRes              *WirelineServiceAreaRestriction `json:"wlServAreaRes,omitempty"`
	SubRfsp                    *int                            `json:"subRfsp,omitempty"`
	RestrictedRatList          []string                        `json:"restrictedRatList,omitempty"`
	AstiDistributionIndication *bool                           `json:"astiDistributionIndication,omitempty"`
	SmallDataRateStatusInfos   []SmallDataRateStatusInfo       `json:"smallDataRateStatusInfos,omitempty"`
	PcfUepBindingInfo          string                          `json:"pcfUepBindingInfo,omitempty"`
	DisasterPlmn               *PlmnId                         `json:"disasterPlmn,omitempty"`
	FiveGMmCapability          string                          `json:"5gMmCapability,omitempty"`
	ForbiddenAreaList          []Area                          `json:"forbiddenAreaList,omitempty"`
	SessionContextList         []PduSessionContext             `json:"sessionContextList,omitempty"`
	EventSubscriptionList      []ExtAmfEventSubscription       `json:"eventSubscriptionList,omitempty"`
	LteCatMInd                 *bool                           `json:"lteCatMInd,omitempty"`
	TsErrorBudget              *int                            `json:"tsErrorBudget,omitempty"`
	SnpnOnboardInd             *bool                           `json:"snpnOnboardInd,omitempty"`
	Supi                       string                          `json:"supi,omitempty"`
	UsedRfsp                   *int                            `json:"usedRfsp,omitempty"`
	SeafData                   *SeafData                       `json:"seafData,omitempty"`
	UsedServiceAreaRestriction *ServiceAreaRestriction         `json:"usedServiceAreaRestriction,omitempty"`
	SubUeAmbr                  *Ambr                           `json:"subUeAmbr,omitempty"`
	HpcfSetId                  string                          `json:"hpcfSetId,omitempty"`
	EpsInterworkingInfo        *EpsInterworkingInfo            `json:"epsInterworkingInfo,omitempty"`
	RedCapInd                  *bool                           `json:"redCapInd,omitempty"`
	IabOperationAllowed        *bool                           `json:"iabOperationAllowed,omitempty"`
	AsTimeDisParam             *AsTimeDistributionParam        `json:"asTimeDisParam,omitempty"`
	Pei                        string                          `json:"pei,omitempty"`
	PcfAmpServiceSetId         string                          `json:"pcfAmpServiceSetId,omitempty"`
	PcfUePolicyUri             string                          `json:"pcfUePolicyUri,omitempty"`
}
