/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyDataChangeNotification struct {
	UsageMonId              string                                   `json:"usageMonId,omitempty"`
	ReportedFragments       []NotificationItem                       `json:"reportedFragments,omitempty"`
	Snssai                  *Snssai                                  `json:"snssai,omitempty"`
	SmPolicyData            *SmPolicyData                            `json:"smPolicyData,omitempty"`
	SponsorConnectivityData *SponsorConnectivityData                 `json:"SponsorConnectivityData,omitempty"`
	UeId                    string                                   `json:"ueId,omitempty"`
	DelResources            []string                                 `json:"delResources,omitempty"`
	SlicePolicyData         *SlicePolicyData                         `json:"slicePolicyData,omitempty"`
	UsageMonData            *UsageMonData                            `json:"usageMonData,omitempty"`
	PlmnUePolicySet         *UePolicySet                             `json:"plmnUePolicySet,omitempty"`
	OpSpecData              *OperatorSpecificDataContainer           `json:"opSpecData,omitempty"`
	BdtRefId                string                                   `json:"bdtRefId,omitempty"`
	PlmnId                  *PlmnId                                  `json:"plmnId,omitempty"`
	NotifId                 string                                   `json:"notifId,omitempty"`
	UePolicySet             *UePolicySet                             `json:"uePolicySet,omitempty"`
	BdtData                 *BdtData                                 `json:"bdtData,omitempty"`
	OpSpecDataMap           map[string]OperatorSpecificDataContainer `json:"opSpecDataMap,omitempty"`
	SponsorId               string                                   `json:"sponsorId,omitempty"`
	AmPolicyData            *AmPolicyData                            `json:"amPolicyData,omitempty"`
}
