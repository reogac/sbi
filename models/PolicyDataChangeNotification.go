/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:03 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyDataChangeNotification struct {
	UePolicySet             *UePolicySet                             `json:"uePolicySet,omitempty"`
	OpSpecDataMap           map[string]OperatorSpecificDataContainer `json:"opSpecDataMap,omitempty"`
	UsageMonId              string                                   `json:"usageMonId,omitempty"`
	OpSpecData              *OperatorSpecificDataContainer           `json:"opSpecData,omitempty"`
	UeId                    string                                   `json:"ueId,omitempty"`
	NotifId                 string                                   `json:"notifId,omitempty"`
	SlicePolicyData         *SlicePolicyData                         `json:"slicePolicyData,omitempty"`
	DelResources            []string                                 `json:"delResources,omitempty"`
	ReportedFragments       []NotificationItem                       `json:"reportedFragments,omitempty"`
	AmPolicyData            *AmPolicyData                            `json:"amPolicyData,omitempty"`
	SmPolicyData            *SmPolicyData                            `json:"smPolicyData,omitempty"`
	UsageMonData            *UsageMonData                            `json:"usageMonData,omitempty"`
	SponsorConnectivityData *SponsorConnectivityData                 `json:"SponsorConnectivityData,omitempty"`
	BdtData                 *BdtData                                 `json:"bdtData,omitempty"`
	SponsorId               string                                   `json:"sponsorId,omitempty"`
	Snssai                  *Snssai                                  `json:"snssai,omitempty"`
	PlmnUePolicySet         *UePolicySet                             `json:"plmnUePolicySet,omitempty"`
	BdtRefId                string                                   `json:"bdtRefId,omitempty"`
	PlmnId                  *PlmnId                                  `json:"plmnId,omitempty"`
}
