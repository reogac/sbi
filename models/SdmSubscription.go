/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SdmSubscription struct {
	Expires                    string                       `json:"expires,omitempty"`
	ImmediateReport            *bool                        `json:"immediateReport,omitempty"`
	UniqueSubscription         *bool                        `json:"uniqueSubscription,omitempty"`
	ResetIds                   []string                     `json:"resetIds,omitempty"`
	UeConSmfDataSubFilter      *UeContextInSmfDataSubFilter `json:"ueConSmfDataSubFilter,omitempty"`
	NfInstanceId               string                       `json:"nfInstanceId"`
	ImplicitUnsubscribe        *bool                        `json:"implicitUnsubscribe,omitempty"`
	SingleNssai                *Snssai                      `json:"singleNssai,omitempty"`
	Dnn                        string                       `json:"dnn,omitempty"`
	SubscriptionId             string                       `json:"subscriptionId,omitempty"`
	DataRestorationCallbackUri string                       `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd              *bool                        `json:"udrRestartInd,omitempty"`
	AmfServiceName             ServiceName                  `json:"amfServiceName,omitempty"`
	MonitoredResourceUris      []string                     `json:"monitoredResourceUris"`
	ContextInfo                *ContextInfo                 `json:"contextInfo,omitempty"`
	CallbackReference          string                       `json:"callbackReference"`
	PlmnId                     *PlmnId                      `json:"plmnId,omitempty"`
	NfChangeFilter             *bool                        `json:"nfChangeFilter,omitempty"`
	Report                     *ImmediateReport             `json:"report,omitempty"`
	SupportedFeatures          string                       `json:"supportedFeatures,omitempty"`
}
