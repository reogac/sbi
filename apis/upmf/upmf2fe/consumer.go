package upmf2fe

import (
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"github.com/reogac/sbi/models/n43"
	"fmt"
	"net/http"
)

const SERVICE_PATH string = "upmf2fe"

var _routes = sbi.SbiRoutes{
	{
		Label:   "TopoQuery",
		Method:  http.MethodGet,
		Path:    "topo",
		Handler: OnGetTopo,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   "upmf2fe",
		Routes:  _routes,
		Handler: p,
	}
}

func GetTopo(client sbi.ConsumerClient) (topo *n43.TopoUpf, err error) {

	//create a request
	path := fmt.Sprintf("%s/path", SERVICE_PATH)
	req := sbi.NewRequest(path, http.MethodGet, nil)

	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if response.GetCode() == 200 {
		topo = &n43.TopoUpf{}
		err = response.DecodeBody(topo)
	} else {
		prob := &models.ProblemDetails{}
		if err = response.DecodeBody(prob); err == nil {
			err = fmt.Errorf("%s: Problem=%s", response.GetStatus(), prob.Detail)
		}
	}
	return
}
