/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionContext struct {
	HsmfId                     string             `json:"hsmfId,omitempty"`
	IsmfServiceSetId           string             `json:"ismfServiceSetId,omitempty"`
	NrfManagementUri           string             `json:"nrfManagementUri,omitempty"`
	AnchorSmfSupportedFeatures string             `json:"anchorSmfSupportedFeatures,omitempty"`
	SelectedDnn                string             `json:"selectedDnn,omitempty"`
	AllocatedEbiList           []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	PgwIpAddr                  *IpAddress         `json:"pgwIpAddr,omitempty"`
	AnchorSmfOauth2Required    *bool              `json:"anchorSmfOauth2Required,omitempty"`
	CnAssistedRanPara          *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	NrfDiscoveryUri            string             `json:"nrfDiscoveryUri,omitempty"`
	VsmfBindingInfo            string             `json:"vsmfBindingInfo,omitempty"`
	AdditionalSnssai           *Snssai            `json:"additionalSnssai,omitempty"`
	IsmfBindingInfo            string             `json:"ismfBindingInfo,omitempty"`
	SNssai                     Snssai             `json:"sNssai"`
	VsmfServiceSetId           string             `json:"vsmfServiceSetId,omitempty"`
	NsInstance                 string             `json:"nsInstance,omitempty"`
	SmfBindingInfo             string             `json:"smfBindingInfo,omitempty"`
	NrfAccessTokenUri          string             `json:"nrfAccessTokenUri,omitempty"`
	AccessType                 AccessType         `json:"accessType"`
	VsmfId                     string             `json:"vsmfId,omitempty"`
	IsmfSetId                  string             `json:"ismfSetId,omitempty"`
	MaPduSession               *bool              `json:"maPduSession,omitempty"`
	SmfServiceInstanceId       string             `json:"smfServiceInstanceId,omitempty"`
	Dnn                        string             `json:"dnn"`
	HsmfServiceSetId           string             `json:"hsmfServiceSetId,omitempty"`
	VsmfSetId                  string             `json:"vsmfSetId,omitempty"`
	IsmfBinding                SbiBindingLevel    `json:"ismfBinding,omitempty"`
	PduSessionId               int                `json:"pduSessionId"`
	PgwFqdn                    string             `json:"pgwFqdn,omitempty"`
	VsmfBinding                SbiBindingLevel    `json:"vsmfBinding,omitempty"`
	IsmfId                     string             `json:"ismfId,omitempty"`
	InterPlmnApiRoot           string             `json:"interPlmnApiRoot,omitempty"`
	PlmnId                     *PlmnId            `json:"plmnId,omitempty"`
	SmContextRef               string             `json:"smContextRef"`
	AdditionalAccessType       AccessType         `json:"additionalAccessType,omitempty"`
	HsmfSetId                  string             `json:"hsmfSetId,omitempty"`
	SmfBinding                 SbiBindingLevel    `json:"smfBinding,omitempty"`
}
