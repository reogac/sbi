/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:51 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SdmSubscription struct {
	NfInstanceId               string                       `json:"nfInstanceId"`
	Expires                    string                       `json:"expires,omitempty"`
	UeConSmfDataSubFilter      *UeContextInSmfDataSubFilter `json:"ueConSmfDataSubFilter,omitempty"`
	SubscriptionId             string                       `json:"subscriptionId,omitempty"`
	ImmediateReport            *bool                        `json:"immediateReport,omitempty"`
	SupportedFeatures          string                       `json:"supportedFeatures,omitempty"`
	NfChangeFilter             *bool                        `json:"nfChangeFilter,omitempty"`
	UniqueSubscription         *bool                        `json:"uniqueSubscription,omitempty"`
	CallbackReference          string                       `json:"callbackReference"`
	SingleNssai                *Snssai                      `json:"singleNssai,omitempty"`
	Dnn                        string                       `json:"dnn,omitempty"`
	DataRestorationCallbackUri string                       `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd              *bool                        `json:"udrRestartInd,omitempty"`
	Report                     *ImmediateReport             `json:"report,omitempty"`
	ContextInfo                *ContextInfo                 `json:"contextInfo,omitempty"`
	AmfServiceName             ServiceName                  `json:"amfServiceName,omitempty"`
	MonitoredResourceUris      []string                     `json:"monitoredResourceUris"`
	PlmnId                     *PlmnId                      `json:"plmnId,omitempty"`
	ImplicitUnsubscribe        *bool                        `json:"implicitUnsubscribe,omitempty"`
	ResetIds                   []string                     `json:"resetIds,omitempty"`
}
