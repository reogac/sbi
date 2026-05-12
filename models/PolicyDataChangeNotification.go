/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:46 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyDataChangeNotification struct {
	UeId                    string                                   `json:"ueId,omitempty"`
	NotifId                 string                                   `json:"notifId,omitempty"`
	SponsorConnectivityData *SponsorConnectivityData                 `json:"SponsorConnectivityData,omitempty"`
	BdtData                 *BdtData                                 `json:"bdtData,omitempty"`
	SponsorId               string                                   `json:"sponsorId,omitempty"`
	PlmnId                  *PlmnId                                  `json:"plmnId,omitempty"`
	ReportedFragments       []NotificationItem                       `json:"reportedFragments,omitempty"`
	SlicePolicyData         *SlicePolicyData                         `json:"slicePolicyData,omitempty"`
	AmPolicyData            *AmPolicyData                            `json:"amPolicyData,omitempty"`
	UePolicySet             *UePolicySet                             `json:"uePolicySet,omitempty"`
	SmPolicyData            *SmPolicyData                            `json:"smPolicyData,omitempty"`
	OpSpecData              *OperatorSpecificDataContainer           `json:"opSpecData,omitempty"`
	BdtRefId                string                                   `json:"bdtRefId,omitempty"`
	DelResources            []string                                 `json:"delResources,omitempty"`
	Snssai                  *Snssai                                  `json:"snssai,omitempty"`
	PlmnUePolicySet         *UePolicySet                             `json:"plmnUePolicySet,omitempty"`
	UsageMonData            *UsageMonData                            `json:"usageMonData,omitempty"`
	UsageMonId              string                                   `json:"usageMonId,omitempty"`
	OpSpecDataMap           map[string]OperatorSpecificDataContainer `json:"opSpecDataMap,omitempty"`
}
