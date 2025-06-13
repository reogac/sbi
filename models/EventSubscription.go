/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EventSubscription struct {
	RedTransReqs       []RedundantTransmissionExpReq `json:"redTransReqs,omitempty"`
	LadnDnns           []string                      `json:"ladnDnns,omitempty"`
	NetworkArea        *NetworkAreaInfo              `json:"networkArea,omitempty"`
	ExptAnaType        ExpectedAnalyticsType         `json:"exptAnaType,omitempty"`
	ExcepRequs         []Exception                   `json:"excepRequs,omitempty"`
	Dnais              []string                      `json:"dnais,omitempty"`
	VisitedAreas       []NetworkAreaInfo             `json:"visitedAreas,omitempty"`
	Snssaia            []Snssai                      `json:"snssaia,omitempty"`
	NfInstanceIds      []string                      `json:"nfInstanceIds,omitempty"`
	CongThresholds     []ThresholdLevel              `json:"congThresholds,omitempty"`
	NfSetIds           []string                      `json:"nfSetIds,omitempty"`
	AppServerAddrs     []AddrFqdn                    `json:"appServerAddrs,omitempty"`
	AnySlice           *bool                         `json:"anySlice,omitempty"`
	Dnns               []string                      `json:"dnns,omitempty"`
	Event              NwdafEvent                    `json:"event"`
	QosRequ            *QosRequirement               `json:"qosRequ,omitempty"`
	NwPerfRequs        []NetworkPerfRequirement      `json:"nwPerfRequs,omitempty"`
	WlanReqs           []WlanPerformanceReq          `json:"wlanReqs,omitempty"`
	AppIds             []string                      `json:"appIds,omitempty"`
	MaxTopAppDlNbr     *int                          `json:"maxTopAppDlNbr,omitempty"`
	NsiLevelThrds      []int                         `json:"nsiLevelThrds,omitempty"`
	TgtUe              *TargetUeInformation          `json:"tgtUe,omitempty"`
	ExptUeBehav        *ExpectedUeBehaviourData      `json:"exptUeBehav,omitempty"`
	RatFreqs           []RatFreqInformation          `json:"ratFreqs,omitempty"`
	ExtraReportReq     *EventReportingRequirement    `json:"extraReportReq,omitempty"`
	MatchingDir        MatchingDirection             `json:"matchingDir,omitempty"`
	NfLoadLvlThds      []ThresholdLevel              `json:"nfLoadLvlThds,omitempty"`
	BwRequs            []BwRequirement               `json:"bwRequs,omitempty"`
	ListOfAnaSubsets   []string                      `json:"listOfAnaSubsets,omitempty"`
	DnPerfReqs         []DnPerformanceReq            `json:"dnPerfReqs,omitempty"`
	NotificationMethod NotificationMethod            `json:"notificationMethod,omitempty"`
	RanUeThrouThds     []string                      `json:"ranUeThrouThds,omitempty"`
	RepetitionPeriod   *int                          `json:"repetitionPeriod,omitempty"`
	NsiIdInfos         []NsiIdInfo                   `json:"nsiIdInfos,omitempty"`
	QosFlowRetThds     []RetainabilityThreshold      `json:"qosFlowRetThds,omitempty"`
	DisperReqs         []DispersionRequirement       `json:"disperReqs,omitempty"`
	UpfInfo            *UpfInformation               `json:"upfInfo,omitempty"`
	LoadLevelThreshold *int                          `json:"loadLevelThreshold,omitempty"`
	NfTypes            []string                      `json:"nfTypes,omitempty"`
	MaxTopAppUlNbr     *int                          `json:"maxTopAppUlNbr,omitempty"`
}
