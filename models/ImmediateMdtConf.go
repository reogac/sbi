/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ImmediateMdtConf struct {
	MeasurementNrList        []string                  `json:"measurementNrList,omitempty"`
	ReportInterval           ReportIntervalMdt         `json:"reportInterval,omitempty"`
	EventThresholdRsrqNr     *int                      `json:"eventThresholdRsrqNr,omitempty"`
	AreaScope                *AreaScope                `json:"areaScope,omitempty"`
	PositioningMethod        PositioningMethodMdt      `json:"positioningMethod,omitempty"`
	AddPositioningMethodList []string                  `json:"addPositioningMethodList,omitempty"`
	ReportingTriggerList     []string                  `json:"reportingTriggerList,omitempty"`
	CollectionPeriodRmmLte   CollectionPeriodRmmLteMdt `json:"collectionPeriodRmmLte,omitempty"`
	MdtAllowedPlmnIdList     []PlmnId                  `json:"mdtAllowedPlmnIdList,omitempty"`
	SensorMeasurementList    []string                  `json:"sensorMeasurementList,omitempty"`
	MeasurementLteList       []string                  `json:"measurementLteList,omitempty"`
	ReportAmount             ReportAmountMdt           `json:"reportAmount,omitempty"`
	EventThresholdRsrp       *int                      `json:"eventThresholdRsrp,omitempty"`
	JobType                  JobType                   `json:"jobType"`
	ReportIntervalNr         ReportIntervalNrMdt       `json:"reportIntervalNr,omitempty"`
	EventThresholdRsrq       *int                      `json:"eventThresholdRsrq,omitempty"`
	EventThresholdRsrpNr     *int                      `json:"eventThresholdRsrpNr,omitempty"`
	CollectionPeriodRmmNr    CollectionPeriodRmmNrMdt  `json:"collectionPeriodRmmNr,omitempty"`
	MeasurementPeriodLte     MeasurementPeriodLteMdt   `json:"measurementPeriodLte,omitempty"`
}
