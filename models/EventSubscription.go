/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EventSubscription struct {
	NotificationMethod NotificationMethod            `json:"notificationMethod,omitempty"`
	NfTypes            []string                      `json:"nfTypes,omitempty"`
	NetworkArea        *NetworkAreaInfo              `json:"networkArea,omitempty"`
	MaxTopAppUlNbr     *int                          `json:"maxTopAppUlNbr,omitempty"`
	RedTransReqs       []RedundantTransmissionExpReq `json:"redTransReqs,omitempty"`
	DnPerfReqs         []DnPerformanceReq            `json:"dnPerfReqs,omitempty"`
	NfInstanceIds      []string                      `json:"nfInstanceIds,omitempty"`
	CongThresholds     []ThresholdLevel              `json:"congThresholds,omitempty"`
	DisperReqs         []DispersionRequirement       `json:"disperReqs,omitempty"`
	AnySlice           *bool                         `json:"anySlice,omitempty"`
	TgtUe              *TargetUeInformation          `json:"tgtUe,omitempty"`
	ListOfAnaSubsets   []string                      `json:"listOfAnaSubsets,omitempty"`
	NfSetIds           []string                      `json:"nfSetIds,omitempty"`
	AppIds             []string                      `json:"appIds,omitempty"`
	Dnais              []string                      `json:"dnais,omitempty"`
	Event              NwdafEvent                    `json:"event"`
	NsiIdInfos         []NsiIdInfo                   `json:"nsiIdInfos,omitempty"`
	RanUeThrouThds     []string                      `json:"ranUeThrouThds,omitempty"`
	ExptUeBehav        *ExpectedUeBehaviourData      `json:"exptUeBehav,omitempty"`
	WlanReqs           []WlanPerformanceReq          `json:"wlanReqs,omitempty"`
	MaxTopAppDlNbr     *int                          `json:"maxTopAppDlNbr,omitempty"`
	QosFlowRetThds     []RetainabilityThreshold      `json:"qosFlowRetThds,omitempty"`
	RepetitionPeriod   *int                          `json:"repetitionPeriod,omitempty"`
	AppServerAddrs     []AddrFqdn                    `json:"appServerAddrs,omitempty"`
	Dnns               []string                      `json:"dnns,omitempty"`
	LadnDnns           []string                      `json:"ladnDnns,omitempty"`
	LoadLevelThreshold *int                          `json:"loadLevelThreshold,omitempty"`
	QosRequ            *QosRequirement               `json:"qosRequ,omitempty"`
	Snssaia            []Snssai                      `json:"snssaia,omitempty"`
	BwRequs            []BwRequirement               `json:"bwRequs,omitempty"`
	ExptAnaType        ExpectedAnalyticsType         `json:"exptAnaType,omitempty"`
	MatchingDir        MatchingDirection             `json:"matchingDir,omitempty"`
	VisitedAreas       []NetworkAreaInfo             `json:"visitedAreas,omitempty"`
	NsiLevelThrds      []int                         `json:"nsiLevelThrds,omitempty"`
	NwPerfRequs        []NetworkPerfRequirement      `json:"nwPerfRequs,omitempty"`
	ExtraReportReq     *EventReportingRequirement    `json:"extraReportReq,omitempty"`
	NfLoadLvlThds      []ThresholdLevel              `json:"nfLoadLvlThds,omitempty"`
	ExcepRequs         []Exception                   `json:"excepRequs,omitempty"`
	RatFreqs           []RatFreqInformation          `json:"ratFreqs,omitempty"`
	UpfInfo            *UpfInformation               `json:"upfInfo,omitempty"`
}
