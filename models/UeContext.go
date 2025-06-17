/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:41 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContext struct {
	PraInUePolicy              map[string]PresenceInfo         `json:"praInUePolicy,omitempty"`
	DrxParameter               string                          `json:"drxParameter,omitempty"`
	MsClassmark2               string                          `json:"msClassmark2,omitempty"`
	GroupList                  []string                        `json:"groupList,omitempty"`
	AmPolicyReqTriggerList     []string                        `json:"amPolicyReqTriggerList,omitempty"`
	AnalyticsSubscriptionList  []AnalyticsSubscription         `json:"analyticsSubscriptionList,omitempty"`
	SnpnOnboardInd             *bool                           `json:"snpnOnboardInd,omitempty"`
	UdmGroupId                 string                          `json:"udmGroupId,omitempty"`
	RoutingIndicator           string                          `json:"routingIndicator,omitempty"`
	PcfBinding                 SbiBindingLevel                 `json:"pcfBinding,omitempty"`
	RestrictedPrimaryRatList   []string                        `json:"restrictedPrimaryRatList,omitempty"`
	LteCatMInd                 *bool                           `json:"lteCatMInd,omitempty"`
	PcfRfsp                    *int                            `json:"pcfRfsp,omitempty"`
	SubUeSliceMbrList          map[string]SliceMbr             `json:"subUeSliceMbrList,omitempty"`
	PcfGroupId                 string                          `json:"pcfGroupId,omitempty"`
	SubUeAmbr                  *Ambr                           `json:"subUeAmbr,omitempty"`
	ForbiddenAreaList          []Area                          `json:"forbiddenAreaList,omitempty"`
	IabOperationAllowed        *bool                           `json:"iabOperationAllowed,omitempty"`
	WlServAreaRes              *WirelineServiceAreaRestriction `json:"wlServAreaRes,omitempty"`
	GpsiList                   []string                        `json:"gpsiList,omitempty"`
	Pei                        string                          `json:"pei,omitempty"`
	ServiceAreaRestriction     *ServiceAreaRestriction         `json:"serviceAreaRestriction,omitempty"`
	AstiDistributionIndication *bool                           `json:"astiDistributionIndication,omitempty"`
	DisasterRoamingInd         *bool                           `json:"disasterRoamingInd,omitempty"`
	SupiUnauthInd              *bool                           `json:"supiUnauthInd,omitempty"`
	PcfUeAmbr                  *Ambr                           `json:"pcfUeAmbr,omitempty"`
	FiveGMmCapability          string                          `json:"5gMmCapability,omitempty"`
	ServiceGapExpiryTime       string                          `json:"serviceGapExpiryTime,omitempty"`
	EcRestrictionDataWb        *EcRestrictionDataWb            `json:"ecRestrictionDataWb,omitempty"`
	PcfAmpBindingInfo          string                          `json:"pcfAmpBindingInfo,omitempty"`
	UpdpSubscriptionData       *UpdpSubscriptionData           `json:"updpSubscriptionData,omitempty"`
	SmsfSetId                  string                          `json:"smsfSetId,omitempty"`
	HNwPubKeyId                *int                            `json:"hNwPubKeyId,omitempty"`
	SmsfId                     string                          `json:"smsfId,omitempty"`
	PcfUepServiceSetId         string                          `json:"pcfUepServiceSetId,omitempty"`
	HpcfId                     string                          `json:"hpcfId,omitempty"`
	PcfUepBindingInfo          string                          `json:"pcfUepBindingInfo,omitempty"`
	SmsfBindingInfo            string                          `json:"smsfBindingInfo,omitempty"`
	UsedRfsp                   *int                            `json:"usedRfsp,omitempty"`
	PcfId                      string                          `json:"pcfId,omitempty"`
	SessionContextList         []PduSessionContext             `json:"sessionContextList,omitempty"`
	EpsInterworkingInfo        *EpsInterworkingInfo            `json:"epsInterworkingInfo,omitempty"`
	CMsisdn                    string                          `json:"cMsisdn,omitempty"`
	CagData                    *CagData                        `json:"cagData,omitempty"`
	TsErrorBudget              *int                            `json:"tsErrorBudget,omitempty"`
	SeafData                   *SeafData                       `json:"seafData,omitempty"`
	PcfUePolicyUri             string                          `json:"pcfUePolicyUri,omitempty"`
	ImmediateMdtConf           *ImmediateMdtConf               `json:"immediateMdtConf,omitempty"`
	EcRestrictionDataNb        *bool                           `json:"ecRestrictionDataNb,omitempty"`
	UsedServiceAreaRestriction *ServiceAreaRestriction         `json:"usedServiceAreaRestriction,omitempty"`
	DisasterPlmn               *PlmnId                         `json:"disasterPlmn,omitempty"`
	SubRfsp                    *int                            `json:"subRfsp,omitempty"`
	EventSubscriptionList      []ExtAmfEventSubscription       `json:"eventSubscriptionList,omitempty"`
	MmContextList              []MmContext                     `json:"mmContextList,omitempty"`
	ManagementMdtInd           *bool                           `json:"managementMdtInd,omitempty"`
	PcfAmpServiceSetId         string                          `json:"pcfAmpServiceSetId,omitempty"`
	HpcfSetId                  string                          `json:"hpcfSetId,omitempty"`
	SmPolicyNotifyPduList      []PduSessionInfo                `json:"smPolicyNotifyPduList,omitempty"`
	SmsfServiceSetId           string                          `json:"smsfServiceSetId,omitempty"`
	RestrictedRatList          []string                        `json:"restrictedRatList,omitempty"`
	SmallDataRateStatusInfos   []SmallDataRateStatusInfo       `json:"smallDataRateStatusInfos,omitempty"`
	MoExpDataCounter           *MoExpDataCounter               `json:"moExpDataCounter,omitempty"`
	SmfSelInfo                 *SmfSelectionData               `json:"smfSelInfo,omitempty"`
	PcfUeSliceMbrList          map[string]SliceMbr             `json:"pcfUeSliceMbrList,omitempty"`
	Supi                       string                          `json:"supi,omitempty"`
	RedCapInd                  *bool                           `json:"redCapInd,omitempty"`
	TraceData                  *TraceData                      `json:"traceData,omitempty"`
	StnSr                      string                          `json:"stnSr,omitempty"`
	SupportedCodecList         []string                        `json:"supportedCodecList,omitempty"`
	V2xContext                 *V2xContext                     `json:"v2xContext,omitempty"`
	PraInAmPolicy              map[string]PresenceInfo         `json:"praInAmPolicy,omitempty"`
	UePositioningCap           string                          `json:"uePositioningCap,omitempty"`
	AusfGroupId                string                          `json:"ausfGroupId,omitempty"`
	PcfAmPolicyUri             string                          `json:"pcfAmPolicyUri,omitempty"`
	AsTimeDisParam             *AsTimeDistributionParam        `json:"asTimeDisParam,omitempty"`
	RestrictedCoreNwTypeList   []string                        `json:"restrictedCoreNwTypeList,omitempty"`
	PcfUeCallbackInfo          *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	RestrictedSecondaryRatList []string                        `json:"restrictedSecondaryRatList,omitempty"`
	ProseContext               *ProseContext                   `json:"proseContext,omitempty"`
	PcfSetId                   string                          `json:"pcfSetId,omitempty"`
	UePolicyReqTriggerList     []string                        `json:"uePolicyReqTriggerList,omitempty"`
}
