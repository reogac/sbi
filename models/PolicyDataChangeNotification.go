/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:52 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyDataChangeNotification struct {
	UsageMonId              string                                   `json:"usageMonId,omitempty"`
	UePolicySet             *UePolicySet                             `json:"uePolicySet,omitempty"`
	BdtRefId                string                                   `json:"bdtRefId,omitempty"`
	SmPolicyData            *SmPolicyData                            `json:"smPolicyData,omitempty"`
	SponsorConnectivityData *SponsorConnectivityData                 `json:"SponsorConnectivityData,omitempty"`
	OpSpecDataMap           map[string]OperatorSpecificDataContainer `json:"opSpecDataMap,omitempty"`
	SponsorId               string                                   `json:"sponsorId,omitempty"`
	ReportedFragments       []NotificationItem                       `json:"reportedFragments,omitempty"`
	Snssai                  *Snssai                                  `json:"snssai,omitempty"`
	AmPolicyData            *AmPolicyData                            `json:"amPolicyData,omitempty"`
	PlmnUePolicySet         *UePolicySet                             `json:"plmnUePolicySet,omitempty"`
	UeId                    string                                   `json:"ueId,omitempty"`
	PlmnId                  *PlmnId                                  `json:"plmnId,omitempty"`
	DelResources            []string                                 `json:"delResources,omitempty"`
	NotifId                 string                                   `json:"notifId,omitempty"`
	BdtData                 *BdtData                                 `json:"bdtData,omitempty"`
	OpSpecData              *OperatorSpecificDataContainer           `json:"opSpecData,omitempty"`
	UsageMonData            *UsageMonData                            `json:"usageMonData,omitempty"`
	SlicePolicyData         *SlicePolicyData                         `json:"slicePolicyData,omitempty"`
}
