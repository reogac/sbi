/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SdmSubscription struct {
	ImplicitUnsubscribe        *bool                        `json:"implicitUnsubscribe,omitempty"`
	Expires                    string                       `json:"expires,omitempty"`
	CallbackReference          string                       `json:"callbackReference"`
	AmfServiceName             ServiceName                  `json:"amfServiceName,omitempty"`
	ContextInfo                *ContextInfo                 `json:"contextInfo,omitempty"`
	PlmnId                     *PlmnId                      `json:"plmnId,omitempty"`
	ResetIds                   []string                     `json:"resetIds,omitempty"`
	UdrRestartInd              *bool                        `json:"udrRestartInd,omitempty"`
	SingleNssai                *Snssai                      `json:"singleNssai,omitempty"`
	Dnn                        string                       `json:"dnn,omitempty"`
	SubscriptionId             string                       `json:"subscriptionId,omitempty"`
	SupportedFeatures          string                       `json:"supportedFeatures,omitempty"`
	DataRestorationCallbackUri string                       `json:"dataRestorationCallbackUri,omitempty"`
	UniqueSubscription         *bool                        `json:"uniqueSubscription,omitempty"`
	UeConSmfDataSubFilter      *UeContextInSmfDataSubFilter `json:"ueConSmfDataSubFilter,omitempty"`
	NfInstanceId               string                       `json:"nfInstanceId"`
	MonitoredResourceUris      []string                     `json:"monitoredResourceUris"`
	ImmediateReport            *bool                        `json:"immediateReport,omitempty"`
	Report                     *ImmediateReport             `json:"report,omitempty"`
	NfChangeFilter             *bool                        `json:"nfChangeFilter,omitempty"`
}
