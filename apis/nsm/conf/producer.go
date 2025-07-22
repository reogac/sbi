/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:16 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package conf

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnGetUdrConfiguration(ctx sbi.RequestContext, prod Producer) {

	// call application handler
	rsp, prob := prod.HandleGetUdrConfiguration()

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnGetUdmConfiguration(ctx sbi.RequestContext, prod Producer) {

	// call application handler
	rsp, prob := prod.HandleGetUdmConfiguration()

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnGetNssfConfiguration(ctx sbi.RequestContext, prod Producer) {

	// call application handler
	rsp, prob := prod.HandleGetNssfConfiguration()

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnGetSessionManagementConfiguration(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetSessionManagementConfigurationParams

	// read 'uuid'
	params.Uuid = ctx.Param("uuid")
	if len(params.Uuid) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "uuid is required"), nil)
		return
	}

	// read 'slice'
	sliceStr := ctx.Param("slice")
	if len(sliceStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "slice is required"), nil)
		return
	}

	if params.Slice, err = models.SnssaiFromString(sliceStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse slice failed: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetSessionManagementConfiguration(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnGetUserPlaneConfiguration(ctx sbi.RequestContext, prod Producer) {
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.UserPlaneConfigurationRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetUserPlaneConfiguration(body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

type Producer interface {
	HandleGetUdrConfiguration() (*models.UdrConfiguration, *models.ProblemDetails)

	HandleGetUdmConfiguration() (*models.UdmConfiguration, *models.ProblemDetails)

	HandleGetNssfConfiguration() (*models.NssfConfiguration, *models.ProblemDetails)

	HandleGetSessionManagementConfiguration(*GetSessionManagementConfigurationParams) (*models.SessionManagementConfiguration, *models.ProblemDetails)

	HandleGetUserPlaneConfiguration(*models.UserPlaneConfigurationRequest) (*models.UserPlaneConfigurationResponse, *models.ProblemDetails)
}
