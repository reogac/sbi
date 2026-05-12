/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SdmSubscription struct {
	CallbackReference          string                       `json:"callbackReference"`
	AmfServiceName             ServiceName                  `json:"amfServiceName,omitempty"`
	SingleNssai                *Snssai                      `json:"singleNssai,omitempty"`
	Dnn                        string                       `json:"dnn,omitempty"`
	PlmnId                     *PlmnId                      `json:"plmnId,omitempty"`
	SupportedFeatures          string                       `json:"supportedFeatures,omitempty"`
	UeConSmfDataSubFilter      *UeContextInSmfDataSubFilter `json:"ueConSmfDataSubFilter,omitempty"`
	MonitoredResourceUris      []string                     `json:"monitoredResourceUris"`
	Report                     *ImmediateReport             `json:"report,omitempty"`
	ContextInfo                *ContextInfo                 `json:"contextInfo,omitempty"`
	UniqueSubscription         *bool                        `json:"uniqueSubscription,omitempty"`
	ResetIds                   []string                     `json:"resetIds,omitempty"`
	DataRestorationCallbackUri string                       `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd              *bool                        `json:"udrRestartInd,omitempty"`
	Expires                    string                       `json:"expires,omitempty"`
	SubscriptionId             string                       `json:"subscriptionId,omitempty"`
	ImmediateReport            *bool                        `json:"immediateReport,omitempty"`
	NfInstanceId               string                       `json:"nfInstanceId"`
	NfChangeFilter             *bool                        `json:"nfChangeFilter,omitempty"`
	ImplicitUnsubscribe        *bool                        `json:"implicitUnsubscribe,omitempty"`
}
