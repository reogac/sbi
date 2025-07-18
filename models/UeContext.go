/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContext struct {
	PcfUeAmbr                  *Ambr                           `json:"pcfUeAmbr,omitempty"`
	PcfUepBindingInfo          string                          `json:"pcfUepBindingInfo,omitempty"`
	ServiceGapExpiryTime       string                          `json:"serviceGapExpiryTime,omitempty"`
	SupportedCodecList         []string                        `json:"supportedCodecList,omitempty"`
	EcRestrictionDataNb        *bool                           `json:"ecRestrictionDataNb,omitempty"`
	PraInAmPolicy              map[string]PresenceInfo         `json:"praInAmPolicy,omitempty"`
	AusfGroupId                string                          `json:"ausfGroupId,omitempty"`
	FiveGMmCapability          string                          `json:"5gMmCapability,omitempty"`
	UsedRfsp                   *int                            `json:"usedRfsp,omitempty"`
	PcfUepServiceSetId         string                          `json:"pcfUepServiceSetId,omitempty"`
	PcfUePolicyUri             string                          `json:"pcfUePolicyUri,omitempty"`
	MoExpDataCounter           *MoExpDataCounter               `json:"moExpDataCounter,omitempty"`
	SmsfSetId                  string                          `json:"smsfSetId,omitempty"`
	Pei                        string                          `json:"pei,omitempty"`
	UdmGroupId                 string                          `json:"udmGroupId,omitempty"`
	TraceData                  *TraceData                      `json:"traceData,omitempty"`
	UsedServiceAreaRestriction *ServiceAreaRestriction         `json:"usedServiceAreaRestriction,omitempty"`
	PcfAmpBindingInfo          string                          `json:"pcfAmpBindingInfo,omitempty"`
	DisasterRoamingInd         *bool                           `json:"disasterRoamingInd,omitempty"`
	GroupList                  []string                        `json:"groupList,omitempty"`
	CMsisdn                    string                          `json:"cMsisdn,omitempty"`
	HpcfSetId                  string                          `json:"hpcfSetId,omitempty"`
	ImmediateMdtConf           *ImmediateMdtConf               `json:"immediateMdtConf,omitempty"`
	SmsfServiceSetId           string                          `json:"smsfServiceSetId,omitempty"`
	DrxParameter               string                          `json:"drxParameter,omitempty"`
	PcfBinding                 SbiBindingLevel                 `json:"pcfBinding,omitempty"`
	PcfRfsp                    *int                            `json:"pcfRfsp,omitempty"`
	PcfId                      string                          `json:"pcfId,omitempty"`
	ServiceAreaRestriction     *ServiceAreaRestriction         `json:"serviceAreaRestriction,omitempty"`
	SessionContextList         []PduSessionContext             `json:"sessionContextList,omitempty"`
	LteCatMInd                 *bool                           `json:"lteCatMInd,omitempty"`
	IabOperationAllowed        *bool                           `json:"iabOperationAllowed,omitempty"`
	GpsiList                   []string                        `json:"gpsiList,omitempty"`
	HNwPubKeyId                *int                            `json:"hNwPubKeyId,omitempty"`
	UePositioningCap           string                          `json:"uePositioningCap,omitempty"`
	SmsfBindingInfo            string                          `json:"smsfBindingInfo,omitempty"`
	RoutingIndicator           string                          `json:"routingIndicator,omitempty"`
	EventSubscriptionList      []ExtAmfEventSubscription       `json:"eventSubscriptionList,omitempty"`
	ForbiddenAreaList          []Area                          `json:"forbiddenAreaList,omitempty"`
	RestrictedCoreNwTypeList   []string                        `json:"restrictedCoreNwTypeList,omitempty"`
	ManagementMdtInd           *bool                           `json:"managementMdtInd,omitempty"`
	AstiDistributionIndication *bool                           `json:"astiDistributionIndication,omitempty"`
	WlServAreaRes              *WirelineServiceAreaRestriction `json:"wlServAreaRes,omitempty"`
	Supi                       string                          `json:"supi,omitempty"`
	SmsfId                     string                          `json:"smsfId,omitempty"`
	UpdpSubscriptionData       *UpdpSubscriptionData           `json:"updpSubscriptionData,omitempty"`
	PcfUeSliceMbrList          map[string]SliceMbr             `json:"pcfUeSliceMbrList,omitempty"`
	SubUeAmbr                  *Ambr                           `json:"subUeAmbr,omitempty"`
	SmallDataRateStatusInfos   []SmallDataRateStatusInfo       `json:"smallDataRateStatusInfos,omitempty"`
	CagData                    *CagData                        `json:"cagData,omitempty"`
	ProseContext               *ProseContext                   `json:"proseContext,omitempty"`
	PcfUeCallbackInfo          *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	SeafData                   *SeafData                       `json:"seafData,omitempty"`
	MmContextList              []MmContext                     `json:"mmContextList,omitempty"`
	RestrictedPrimaryRatList   []string                        `json:"restrictedPrimaryRatList,omitempty"`
	HpcfId                     string                          `json:"hpcfId,omitempty"`
	RestrictedRatList          []string                        `json:"restrictedRatList,omitempty"`
	UePolicyReqTriggerList     []string                        `json:"uePolicyReqTriggerList,omitempty"`
	SmPolicyNotifyPduList      []PduSessionInfo                `json:"smPolicyNotifyPduList,omitempty"`
	SubUeSliceMbrList          map[string]SliceMbr             `json:"subUeSliceMbrList,omitempty"`
	PcfAmpServiceSetId         string                          `json:"pcfAmpServiceSetId,omitempty"`
	V2xContext                 *V2xContext                     `json:"v2xContext,omitempty"`
	RedCapInd                  *bool                           `json:"redCapInd,omitempty"`
	AnalyticsSubscriptionList  []AnalyticsSubscription         `json:"analyticsSubscriptionList,omitempty"`
	TsErrorBudget              *int                            `json:"tsErrorBudget,omitempty"`
	SmfSelInfo                 *SmfSelectionData               `json:"smfSelInfo,omitempty"`
	SubRfsp                    *int                            `json:"subRfsp,omitempty"`
	EpsInterworkingInfo        *EpsInterworkingInfo            `json:"epsInterworkingInfo,omitempty"`
	EcRestrictionDataWb        *EcRestrictionDataWb            `json:"ecRestrictionDataWb,omitempty"`
	SnpnOnboardInd             *bool                           `json:"snpnOnboardInd,omitempty"`
	AsTimeDisParam             *AsTimeDistributionParam        `json:"asTimeDisParam,omitempty"`
	PcfGroupId                 string                          `json:"pcfGroupId,omitempty"`
	PcfAmPolicyUri             string                          `json:"pcfAmPolicyUri,omitempty"`
	AmPolicyReqTriggerList     []string                        `json:"amPolicyReqTriggerList,omitempty"`
	StnSr                      string                          `json:"stnSr,omitempty"`
	MsClassmark2               string                          `json:"msClassmark2,omitempty"`
	RestrictedSecondaryRatList []string                        `json:"restrictedSecondaryRatList,omitempty"`
	PraInUePolicy              map[string]PresenceInfo         `json:"praInUePolicy,omitempty"`
	DisasterPlmn               *PlmnId                         `json:"disasterPlmn,omitempty"`
	SupiUnauthInd              *bool                           `json:"supiUnauthInd,omitempty"`
	PcfSetId                   string                          `json:"pcfSetId,omitempty"`
}
