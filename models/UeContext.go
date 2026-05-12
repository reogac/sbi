/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContext struct {
	LteCatMInd                 *bool                           `json:"lteCatMInd,omitempty"`
	MoExpDataCounter           *MoExpDataCounter               `json:"moExpDataCounter,omitempty"`
	CagData                    *CagData                        `json:"cagData,omitempty"`
	ProseContext               *ProseContext                   `json:"proseContext,omitempty"`
	PcfUeCallbackInfo          *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	DisasterRoamingInd         *bool                           `json:"disasterRoamingInd,omitempty"`
	UsedRfsp                   *int                            `json:"usedRfsp,omitempty"`
	PcfSetId                   string                          `json:"pcfSetId,omitempty"`
	HpcfId                     string                          `json:"hpcfId,omitempty"`
	ServiceGapExpiryTime       string                          `json:"serviceGapExpiryTime,omitempty"`
	MsClassmark2               string                          `json:"msClassmark2,omitempty"`
	SupportedCodecList         []string                        `json:"supportedCodecList,omitempty"`
	ManagementMdtInd           *bool                           `json:"managementMdtInd,omitempty"`
	PraInAmPolicy              map[string]PresenceInfo         `json:"praInAmPolicy,omitempty"`
	CMsisdn                    string                          `json:"cMsisdn,omitempty"`
	PcfId                      string                          `json:"pcfId,omitempty"`
	PcfAmPolicyUri             string                          `json:"pcfAmPolicyUri,omitempty"`
	MmContextList              []MmContext                     `json:"mmContextList,omitempty"`
	EpsInterworkingInfo        *EpsInterworkingInfo            `json:"epsInterworkingInfo,omitempty"`
	IabOperationAllowed        *bool                           `json:"iabOperationAllowed,omitempty"`
	RestrictedSecondaryRatList []string                        `json:"restrictedSecondaryRatList,omitempty"`
	AsTimeDisParam             *AsTimeDistributionParam        `json:"asTimeDisParam,omitempty"`
	AmPolicyReqTriggerList     []string                        `json:"amPolicyReqTriggerList,omitempty"`
	StnSr                      string                          `json:"stnSr,omitempty"`
	EcRestrictionDataWb        *EcRestrictionDataWb            `json:"ecRestrictionDataWb,omitempty"`
	UdmGroupId                 string                          `json:"udmGroupId,omitempty"`
	RestrictedRatList          []string                        `json:"restrictedRatList,omitempty"`
	TsErrorBudget              *int                            `json:"tsErrorBudget,omitempty"`
	FiveGMmCapability          string                          `json:"5gMmCapability,omitempty"`
	PcfUepServiceSetId         string                          `json:"pcfUepServiceSetId,omitempty"`
	EcRestrictionDataNb        *bool                           `json:"ecRestrictionDataNb,omitempty"`
	Supi                       string                          `json:"supi,omitempty"`
	SubRfsp                    *int                            `json:"subRfsp,omitempty"`
	SubUeAmbr                  *Ambr                           `json:"subUeAmbr,omitempty"`
	UsedServiceAreaRestriction *ServiceAreaRestriction         `json:"usedServiceAreaRestriction,omitempty"`
	ImmediateMdtConf           *ImmediateMdtConf               `json:"immediateMdtConf,omitempty"`
	PcfUeAmbr                  *Ambr                           `json:"pcfUeAmbr,omitempty"`
	SubUeSliceMbrList          map[string]SliceMbr             `json:"subUeSliceMbrList,omitempty"`
	SessionContextList         []PduSessionContext             `json:"sessionContextList,omitempty"`
	AstiDistributionIndication *bool                           `json:"astiDistributionIndication,omitempty"`
	DisasterPlmn               *PlmnId                         `json:"disasterPlmn,omitempty"`
	Pei                        string                          `json:"pei,omitempty"`
	SnpnOnboardInd             *bool                           `json:"snpnOnboardInd,omitempty"`
	PcfUeSliceMbrList          map[string]SliceMbr             `json:"pcfUeSliceMbrList,omitempty"`
	AusfGroupId                string                          `json:"ausfGroupId,omitempty"`
	PcfRfsp                    *int                            `json:"pcfRfsp,omitempty"`
	SeafData                   *SeafData                       `json:"seafData,omitempty"`
	ServiceAreaRestriction     *ServiceAreaRestriction         `json:"serviceAreaRestriction,omitempty"`
	RoutingIndicator           string                          `json:"routingIndicator,omitempty"`
	HNwPubKeyId                *int                            `json:"hNwPubKeyId,omitempty"`
	DrxParameter               string                          `json:"drxParameter,omitempty"`
	PcfAmpServiceSetId         string                          `json:"pcfAmpServiceSetId,omitempty"`
	PcfUePolicyUri             string                          `json:"pcfUePolicyUri,omitempty"`
	RestrictedPrimaryRatList   []string                        `json:"restrictedPrimaryRatList,omitempty"`
	PcfGroupId                 string                          `json:"pcfGroupId,omitempty"`
	ForbiddenAreaList          []Area                          `json:"forbiddenAreaList,omitempty"`
	RestrictedCoreNwTypeList   []string                        `json:"restrictedCoreNwTypeList,omitempty"`
	EventSubscriptionList      []ExtAmfEventSubscription       `json:"eventSubscriptionList,omitempty"`
	SmfSelInfo                 *SmfSelectionData               `json:"smfSelInfo,omitempty"`
	SmsfSetId                  string                          `json:"smsfSetId,omitempty"`
	SmsfBindingInfo            string                          `json:"smsfBindingInfo,omitempty"`
	UePolicyReqTriggerList     []string                        `json:"uePolicyReqTriggerList,omitempty"`
	RedCapInd                  *bool                           `json:"redCapInd,omitempty"`
	PcfUepBindingInfo          string                          `json:"pcfUepBindingInfo,omitempty"`
	UpdpSubscriptionData       *UpdpSubscriptionData           `json:"updpSubscriptionData,omitempty"`
	SmsfServiceSetId           string                          `json:"smsfServiceSetId,omitempty"`
	GroupList                  []string                        `json:"groupList,omitempty"`
	HpcfSetId                  string                          `json:"hpcfSetId,omitempty"`
	TraceData                  *TraceData                      `json:"traceData,omitempty"`
	AnalyticsSubscriptionList  []AnalyticsSubscription         `json:"analyticsSubscriptionList,omitempty"`
	PraInUePolicy              map[string]PresenceInfo         `json:"praInUePolicy,omitempty"`
	WlServAreaRes              *WirelineServiceAreaRestriction `json:"wlServAreaRes,omitempty"`
	SupiUnauthInd              *bool                           `json:"supiUnauthInd,omitempty"`
	GpsiList                   []string                        `json:"gpsiList,omitempty"`
	SmsfId                     string                          `json:"smsfId,omitempty"`
	V2xContext                 *V2xContext                     `json:"v2xContext,omitempty"`
	PcfAmpBindingInfo          string                          `json:"pcfAmpBindingInfo,omitempty"`
	SmPolicyNotifyPduList      []PduSessionInfo                `json:"smPolicyNotifyPduList,omitempty"`
	UePositioningCap           string                          `json:"uePositioningCap,omitempty"`
	PcfBinding                 SbiBindingLevel                 `json:"pcfBinding,omitempty"`
	SmallDataRateStatusInfos   []SmallDataRateStatusInfo       `json:"smallDataRateStatusInfos,omitempty"`
}
