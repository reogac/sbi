/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SdmSubscription struct {
	SubscriptionId             string                       `json:"subscriptionId,omitempty"`
	ImmediateReport            *bool                        `json:"immediateReport,omitempty"`
	UdrRestartInd              *bool                        `json:"udrRestartInd,omitempty"`
	CallbackReference          string                       `json:"callbackReference"`
	AmfServiceName             ServiceName                  `json:"amfServiceName,omitempty"`
	SingleNssai                *Snssai                      `json:"singleNssai,omitempty"`
	DataRestorationCallbackUri string                       `json:"dataRestorationCallbackUri,omitempty"`
	PlmnId                     *PlmnId                      `json:"plmnId,omitempty"`
	Report                     *ImmediateReport             `json:"report,omitempty"`
	UniqueSubscription         *bool                        `json:"uniqueSubscription,omitempty"`
	ImplicitUnsubscribe        *bool                        `json:"implicitUnsubscribe,omitempty"`
	Dnn                        string                       `json:"dnn,omitempty"`
	SupportedFeatures          string                       `json:"supportedFeatures,omitempty"`
	ContextInfo                *ContextInfo                 `json:"contextInfo,omitempty"`
	NfChangeFilter             *bool                        `json:"nfChangeFilter,omitempty"`
	ResetIds                   []string                     `json:"resetIds,omitempty"`
	UeConSmfDataSubFilter      *UeContextInSmfDataSubFilter `json:"ueConSmfDataSubFilter,omitempty"`
	NfInstanceId               string                       `json:"nfInstanceId"`
	Expires                    string                       `json:"expires,omitempty"`
	MonitoredResourceUris      []string                     `json:"monitoredResourceUris"`
}
