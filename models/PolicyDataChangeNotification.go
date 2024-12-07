/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:34 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyDataChangeNotification struct {
	UsageMonId              string                                   `json:"usageMonId,omitempty"`
	DelResources            []string                                 `json:"delResources,omitempty"`
	SlicePolicyData         *SlicePolicyData                         `json:"slicePolicyData,omitempty"`
	UePolicySet             *UePolicySet                             `json:"uePolicySet,omitempty"`
	BdtRefId                string                                   `json:"bdtRefId,omitempty"`
	OpSpecData              *OperatorSpecificDataContainer           `json:"opSpecData,omitempty"`
	OpSpecDataMap           map[string]OperatorSpecificDataContainer `json:"opSpecDataMap,omitempty"`
	SponsorConnectivityData *SponsorConnectivityData                 `json:"SponsorConnectivityData,omitempty"`
	BdtData                 *BdtData                                 `json:"bdtData,omitempty"`
	SmPolicyData            *SmPolicyData                            `json:"smPolicyData,omitempty"`
	Snssai                  *Snssai                                  `json:"snssai,omitempty"`
	UsageMonData            *UsageMonData                            `json:"usageMonData,omitempty"`
	UeId                    string                                   `json:"ueId,omitempty"`
	SponsorId               string                                   `json:"sponsorId,omitempty"`
	PlmnId                  *PlmnId                                  `json:"plmnId,omitempty"`
	NotifId                 string                                   `json:"notifId,omitempty"`
	ReportedFragments       []NotificationItem                       `json:"reportedFragments,omitempty"`
	AmPolicyData            *AmPolicyData                            `json:"amPolicyData,omitempty"`
	PlmnUePolicySet         *UePolicySet                             `json:"plmnUePolicySet,omitempty"`
}
