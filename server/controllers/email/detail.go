package email

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"pmail/dto"
	"pmail/dto/response"
	"pmail/services/auth"
	"pmail/services/detail"
)

type emailDetailRequest struct {
	ID int `json:"id"`
}

func EmailDetail(ctx *dto.Context, w http.ResponseWriter, req *http.Request) {
	reqBytes, err := io.ReadAll(req.Body)
	if err != nil {
		log.WithContext(ctx).Errorf("%+v", err)
	}
	var retData emailDetailRequest
	err = json.Unmarshal(reqBytes, &retData)
	if err != nil {
		log.WithContext(ctx).Errorf("%+v", err)
	}

	if retData.ID <= 0 {
		response.NewErrorResponse(response.ParamsError, "ID错误", "").FPrint(w)
		return
	}

	email, err := detail.GetEmailDetail(ctx, retData.ID, true)
	if err != nil {
		response.NewErrorResponse(response.ServerError, err.Error(), "").FPrint(w)
		return
	}

	// 检查是否有权限
	hasAuth := auth.HasAuth(ctx, email)
	if !hasAuth {
		response.NewErrorResponse(response.ParamsError, "", "").FPrint(w)
		return
	}

	response.NewSuccessResponse(email).FPrint(w)

}
