/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:26 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package sdm

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnGetAmData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params GetAmDataParams

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdNidFromString(plmnIdStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)))
			return
		}
	}

	// read 'adjacent-plmns'
	adjacentPlmnsStr := ctx.Param("adjacent-plmns")
	if len(adjacentPlmnsStr) > 0 {
		if params.AdjacentPlmns, err = models.ArrayOfPlmnIdFromString(adjacentPlmnsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse adjacent-plmns failed: %+v", err)))
			return
		}
	}

	// read 'disaster-roaming-ind'
	disasterRoamingIndStr := ctx.Param("disaster-roaming-ind")
	if len(disasterRoamingIndStr) > 0 {
		var disasterRoamingIndTmp bool
		if disasterRoamingIndTmp, err = models.BoolFromString(disasterRoamingIndStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse disaster-roaming-ind failed: %+v", err)))
			return
		}
		params.DisasterRoamingInd = &disasterRoamingIndTmp
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	rsp, prob := prod.HandleGetAmData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnGetUeCtxInSmsfData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params GetUeCtxInSmsfDataParams

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetUeCtxInSmsfData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnUnsubscribeForSharedData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'subscriptionId'
	var subscriptionId string
	subscriptionId = ctx.Param("subscriptionId")
	if len(subscriptionId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subscriptionId is required"))
		return
	}

	// call application handler
	prob := prod.HandleUnsubscribeForSharedData(subscriptionId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnGetDataSets(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params GetDataSetsParams

	// read 'disaster-roaming-ind'
	disasterRoamingIndStr := ctx.Param("disaster-roaming-ind")
	if len(disasterRoamingIndStr) > 0 {
		var disasterRoamingIndTmp bool
		if disasterRoamingIndTmp, err = models.BoolFromString(disasterRoamingIndStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse disaster-roaming-ind failed: %+v", err)))
			return
		}
		params.DisasterRoamingInd = &disasterRoamingIndTmp
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'dataset-names'
	datasetNamesStr := ctx.Param("dataset-names")
	if len(datasetNamesStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "dataset-names is required"))
		return
	}

	if params.DatasetNames, err = models.ArrayOfStringFromString(datasetNamesStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse dataset-names failed: %+v", err)))
		return
	}

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdNidFromString(plmnIdStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)))
			return
		}
	}

	// call application handler
	rsp, prob := prod.HandleGetDataSets(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnGetSmsData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params GetSmsDataParams

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdFromString(plmnIdStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)))
			return
		}
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	rsp, prob := prod.HandleGetSmsData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnSorAckInfo(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'supi'
	var supi string
	supi = ctx.Param("supi")
	if len(supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// decode request body
	var body *models.AcknowledgeInfo
	if ctx.HasRequestBody() {
		body = new(models.AcknowledgeInfo)
		if err = ctx.DecodeRequest(body); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
			return
		}
	}

	// call application handler
	prob := prod.HandleSorAckInfo(supi, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnUpuAck(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'supi'
	var supi string
	supi = ctx.Param("supi")
	if len(supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// decode request body
	var body *models.AcknowledgeInfo
	if ctx.HasRequestBody() {
		body = new(models.AcknowledgeInfo)
		if err = ctx.DecodeRequest(body); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
			return
		}
	}

	// call application handler
	prob := prod.HandleUpuAck(supi, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnSNSSAIsAck(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'supi'
	var supi string
	supi = ctx.Param("supi")
	if len(supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// decode request body
	var body *models.AcknowledgeInfo
	if ctx.HasRequestBody() {
		body = new(models.AcknowledgeInfo)
		if err = ctx.DecodeRequest(body); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
			return
		}
	}

	// call application handler
	prob := prod.HandleSNSSAIsAck(supi, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnUpdateSORInfo(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'supi'
	var supi string
	supi = ctx.Param("supi")
	if len(supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// decode request body
	var body *models.SorUpdateInfo
	if ctx.HasRequestBody() {
		body = new(models.SorUpdateInfo)
		if err = ctx.DecodeRequest(body); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
			return
		}
	}

	// call application handler
	rsp, prob := prod.HandleUpdateSORInfo(supi, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnModifySharedDataSubs(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params ModifySharedDataSubsParams

	// read 'subscriptionId'
	params.SubscriptionId = ctx.Param("subscriptionId")
	if len(params.SubscriptionId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subscriptionId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	body := new(models.SdmSubsModification)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifySharedDataSubs(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnGetIndividualSharedData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params GetIndividualSharedDataParams

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'sharedDataId'
	sharedDataIdStr := ctx.Param("sharedDataId")
	if len(sharedDataIdStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "sharedDataId is required"))
		return
	}

	if params.SharedDataId, err = models.ArrayOfStringFromString(sharedDataIdStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse sharedDataId failed: %+v", err)))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// call application handler
	rsp, prob := prod.HandleGetIndividualSharedData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnGetUeCtxInSmfData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params GetUeCtxInSmfDataParams

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetUeCtxInSmfData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnGetMultipleIdentifiers(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params GetMultipleIdentifiersParams

	// read 'gpsi-list'
	gpsiListStr := ctx.Param("gpsi-list")
	if len(gpsiListStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "gpsi-list is required"))
		return
	}

	if params.GpsiList, err = models.ArrayOfStringFromString(gpsiListStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse gpsi-list failed: %+v", err)))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetMultipleIdentifiers(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnGetLcsPrivacyData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params GetLcsPrivacyDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	rsp, prob := prod.HandleGetLcsPrivacyData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnGetProseData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params GetProseDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetProseData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnGetUeCtxInAmfData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params GetUeCtxInAmfDataParams

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetUeCtxInAmfData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnGetV2xData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params GetV2xDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetV2xData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnGetMbsData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params GetMbsDataParams

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	rsp, prob := prod.HandleGetMbsData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnGetSupiOrGpsi(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params GetSupiOrGpsiParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'mtc-provider-info'
	params.MtcProviderInfo = ctx.Param("mtc-provider-info")

	// read 'app-port-id'
	appPortIdStr := ctx.Param("app-port-id")
	if len(appPortIdStr) > 0 {
		if params.AppPortId, err = models.AppPortIdFromString(appPortIdStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse app-port-id failed: %+v", err)))
			return
		}
	}

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'requested-gpsi-type'
	params.RequestedGpsiType = ctx.Param("requested-gpsi-type")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'af-id'
	params.AfId = ctx.Param("af-id")

	// read 'af-service-id'
	params.AfServiceId = ctx.Param("af-service-id")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// call application handler
	rsp, prob := prod.HandleGetSupiOrGpsi(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnGetNSSAI(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params GetNSSAIParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdFromString(plmnIdStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)))
			return
		}
	}

	// read 'disaster-roaming-ind'
	disasterRoamingIndStr := ctx.Param("disaster-roaming-ind")
	if len(disasterRoamingIndStr) > 0 {
		var disasterRoamingIndTmp bool
		if disasterRoamingIndTmp, err = models.BoolFromString(disasterRoamingIndStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse disaster-roaming-ind failed: %+v", err)))
			return
		}
		params.DisasterRoamingInd = &disasterRoamingIndTmp
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetNSSAI(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnGetLcsMoData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params GetLcsMoDataParams

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	rsp, prob := prod.HandleGetLcsMoData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnModify(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params ModifyParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'subscriptionId'
	params.SubscriptionId = ctx.Param("subscriptionId")
	if len(params.SubscriptionId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subscriptionId is required"))
		return
	}

	// decode request body
	body := new(models.SdmSubsModification)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleModify(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnGetTraceConfigData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params GetTraceConfigDataParams

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdFromString(plmnIdStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)))
			return
		}
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetTraceConfigData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnGetSmfSelData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params GetSmfSelDataParams

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdFromString(plmnIdStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)))
			return
		}
	}

	// read 'disaster-roaming-ind'
	disasterRoamingIndStr := ctx.Param("disaster-roaming-ind")
	if len(disasterRoamingIndStr) > 0 {
		var disasterRoamingIndTmp bool
		if disasterRoamingIndTmp, err = models.BoolFromString(disasterRoamingIndStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse disaster-roaming-ind failed: %+v", err)))
			return
		}
		params.DisasterRoamingInd = &disasterRoamingIndTmp
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetSmfSelData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnUnsubscribe(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params UnsubscribeParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'subscriptionId'
	params.SubscriptionId = ctx.Param("subscriptionId")
	if len(params.SubscriptionId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subscriptionId is required"))
		return
	}

	// call application handler
	prob := prod.HandleUnsubscribe(&params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnGetSharedData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params GetSharedDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'shared-data-ids'
	sharedDataIdsStr := ctx.Param("shared-data-ids")
	if len(sharedDataIdsStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "shared-data-ids is required"))
		return
	}

	if params.SharedDataIds, err = models.ArrayOfStringFromString(sharedDataIdsStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse shared-data-ids failed: %+v", err)))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetSharedData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnGetEcrData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params GetEcrDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetEcrData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnGetLcsBcaData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params GetLcsBcaDataParams

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdFromString(plmnIdStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)))
			return
		}
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetLcsBcaData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnGetUcData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var params GetUcDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'uc-purpose'
	params.UcPurpose = ctx.Param("uc-purpose")

	// call application handler
	rsp, prob := prod.HandleGetUcData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnSubscribe(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// decode request body
	body := new(models.SdmSubscription)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleSubscribe(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnCAGAck(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'supi'
	var supi string
	supi = ctx.Param("supi")
	if len(supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// decode request body
	var body *models.AcknowledgeInfo
	if ctx.HasRequestBody() {
		body = new(models.AcknowledgeInfo)
		if err = ctx.DecodeRequest(body); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
			return
		}
	}

	// call application handler
	prob := prod.HandleCAGAck(supi, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnGetGroupIdentifiers(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params GetGroupIdentifiersParams

	// read 'int-group-id'
	params.IntGroupId = ctx.Param("int-group-id")

	// read 'ue-id-ind'
	ueIdIndStr := ctx.Param("ue-id-ind")
	if len(ueIdIndStr) > 0 {
		var ueIdIndTmp bool
		if ueIdIndTmp, err = models.BoolFromString(ueIdIndStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse ue-id-ind failed: %+v", err)))
			return
		}
		params.UeIdInd = &ueIdIndTmp
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'af-id'
	params.AfId = ctx.Param("af-id")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ext-group-id'
	params.ExtGroupId = ctx.Param("ext-group-id")

	// call application handler
	rsp, prob := prod.HandleGetGroupIdentifiers(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnGetSmsMngtData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params GetSmsMngtDataParams

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdFromString(plmnIdStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)))
			return
		}
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	rsp, prob := prod.HandleGetSmsMngtData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnSubscribeToSharedData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// decode request body
	body := new(models.SdmSubscription)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleSubscribeToSharedData(body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnGetSmData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params GetSmDataParams

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'single-nssai'
	singleNssaiStr := ctx.Param("single-nssai")
	if len(singleNssaiStr) > 0 {
		if params.SingleNssai, err = models.SnssaiFromString(singleNssaiStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse single-nssai failed: %+v", err)))
			return
		}
	}

	// read 'dnn'
	params.Dnn = ctx.Param("dnn")

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdFromString(plmnIdStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)))
			return
		}
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// call application handler
	rsp, prob := prod.HandleGetSmData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

type Producer interface {
	HandleGetAmData(*GetAmDataParams) (*models.AccessAndMobilitySubscriptionData, *models.ProblemDetails)

	HandleGetUeCtxInSmsfData(*GetUeCtxInSmsfDataParams) (*models.UeContextInSmsfData, *models.ProblemDetails)

	HandleUnsubscribeForSharedData(string) *models.ProblemDetails

	HandleGetDataSets(*GetDataSetsParams) (*models.SubscriptionDataSets, *models.ProblemDetails)

	HandleGetSmsData(*GetSmsDataParams) (*models.SmsSubscriptionData, *models.ProblemDetails)

	HandleSorAckInfo(string, *models.AcknowledgeInfo) *models.ProblemDetails

	HandleUpuAck(string, *models.AcknowledgeInfo) *models.ProblemDetails

	HandleSNSSAIsAck(string, *models.AcknowledgeInfo) *models.ProblemDetails

	HandleUpdateSORInfo(string, *models.SorUpdateInfo) (*models.SorInfo, *models.ProblemDetails)

	HandleModifySharedDataSubs(*ModifySharedDataSubsParams, *models.SdmSubsModification) (*models.Schema, *models.ProblemDetails)

	HandleGetIndividualSharedData(*GetIndividualSharedDataParams) (*models.SharedData, *models.ProblemDetails)

	HandleGetUeCtxInSmfData(*GetUeCtxInSmfDataParams) (*models.UeContextInSmfData, *models.ProblemDetails)

	HandleGetMultipleIdentifiers(*GetMultipleIdentifiersParams) (*map[string]models.SupiInfo, *models.ProblemDetails)

	HandleGetLcsPrivacyData(*GetLcsPrivacyDataParams) (*models.LcsPrivacyData, *models.ProblemDetails)

	HandleGetProseData(*GetProseDataParams) (*models.ProseSubscriptionData, *models.ProblemDetails)

	HandleGetUeCtxInAmfData(*GetUeCtxInAmfDataParams) (*models.UeContextInAmfData, *models.ProblemDetails)

	HandleGetV2xData(*GetV2xDataParams) (*models.V2xSubscriptionData, *models.ProblemDetails)

	HandleGetMbsData(*GetMbsDataParams) (*models.MbsSubscriptionData, *models.ProblemDetails)

	HandleGetSupiOrGpsi(*GetSupiOrGpsiParams) (*models.IdTranslationResult, *models.ProblemDetails)

	HandleGetNSSAI(*GetNSSAIParams) (*models.Nssai, *models.ProblemDetails)

	HandleGetLcsMoData(*GetLcsMoDataParams) (*models.LcsMoData, *models.ProblemDetails)

	HandleModify(*ModifyParams, *models.SdmSubsModification) (*models.Schema, *models.ProblemDetails)

	HandleGetTraceConfigData(*GetTraceConfigDataParams) (*models.TraceDataResponse, *models.ProblemDetails)

	HandleGetSmfSelData(*GetSmfSelDataParams) (*models.SmfSelectionSubscriptionData, *models.ProblemDetails)

	HandleUnsubscribe(*UnsubscribeParams) *models.ProblemDetails

	HandleGetSharedData(*GetSharedDataParams) (*[]models.SharedData, *models.ProblemDetails)

	HandleGetEcrData(*GetEcrDataParams) (*models.EnhancedCoverageRestrictionData, *models.ProblemDetails)

	HandleGetLcsBcaData(*GetLcsBcaDataParams) (*models.LcsBroadcastAssistanceTypesData, *models.ProblemDetails)

	HandleGetUcData(*GetUcDataParams) (*models.UcSubscriptionData, *models.ProblemDetails)

	HandleSubscribe(string, *models.SdmSubscription) (*models.SdmSubscription, *models.ProblemDetails)

	HandleCAGAck(string, *models.AcknowledgeInfo) *models.ProblemDetails

	HandleGetGroupIdentifiers(*GetGroupIdentifiersParams) (*models.GroupIdentifiers, *models.ProblemDetails)

	HandleGetSmsMngtData(*GetSmsMngtDataParams) (*models.SmsManagementSubscriptionData, *models.ProblemDetails)

	HandleSubscribeToSharedData(*models.SdmSubscription) (*models.SdmSubscription, *models.ProblemDetails)

	HandleGetSmData(*GetSmDataParams) (*models.SmSubsData, *models.ProblemDetails)
}
