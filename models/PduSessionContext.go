/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:17 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionContext struct {
	MaPduSession               *bool              `json:"maPduSession,omitempty"`
	NrfDiscoveryUri            string             `json:"nrfDiscoveryUri,omitempty"`
	PduSessionId               int                `json:"pduSessionId"`
	AccessType                 AccessType         `json:"accessType"`
	IsmfId                     string             `json:"ismfId,omitempty"`
	IsmfSetId                  string             `json:"ismfSetId,omitempty"`
	SmContextRef               string             `json:"smContextRef"`
	PgwFqdn                    string             `json:"pgwFqdn,omitempty"`
	PlmnId                     *PlmnId            `json:"plmnId,omitempty"`
	SelectedDnn                string             `json:"selectedDnn,omitempty"`
	HsmfServiceSetId           string             `json:"hsmfServiceSetId,omitempty"`
	NsInstance                 string             `json:"nsInstance,omitempty"`
	Dnn                        string             `json:"dnn"`
	NrfManagementUri           string             `json:"nrfManagementUri,omitempty"`
	AdditionalAccessType       AccessType         `json:"additionalAccessType,omitempty"`
	SmfBinding                 SbiBindingLevel    `json:"smfBinding,omitempty"`
	VsmfBinding                SbiBindingLevel    `json:"vsmfBinding,omitempty"`
	InterPlmnApiRoot           string             `json:"interPlmnApiRoot,omitempty"`
	VsmfId                     string             `json:"vsmfId,omitempty"`
	IsmfBinding                SbiBindingLevel    `json:"ismfBinding,omitempty"`
	AdditionalSnssai           *Snssai            `json:"additionalSnssai,omitempty"`
	PgwIpAddr                  *IpAddress         `json:"pgwIpAddr,omitempty"`
	VsmfServiceSetId           string             `json:"vsmfServiceSetId,omitempty"`
	SmfServiceInstanceId       string             `json:"smfServiceInstanceId,omitempty"`
	CnAssistedRanPara          *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	SNssai                     Snssai             `json:"sNssai"`
	AllocatedEbiList           []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	HsmfId                     string             `json:"hsmfId,omitempty"`
	HsmfSetId                  string             `json:"hsmfSetId,omitempty"`
	VsmfBindingInfo            string             `json:"vsmfBindingInfo,omitempty"`
	IsmfBindingInfo            string             `json:"ismfBindingInfo,omitempty"`
	AnchorSmfSupportedFeatures string             `json:"anchorSmfSupportedFeatures,omitempty"`
	AnchorSmfOauth2Required    *bool              `json:"anchorSmfOauth2Required,omitempty"`
	VsmfSetId                  string             `json:"vsmfSetId,omitempty"`
	IsmfServiceSetId           string             `json:"ismfServiceSetId,omitempty"`
	NrfAccessTokenUri          string             `json:"nrfAccessTokenUri,omitempty"`
	SmfBindingInfo             string             `json:"smfBindingInfo,omitempty"`
}
