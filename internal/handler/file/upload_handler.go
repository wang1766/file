package file

import (
	"github.com/suyuan32/simple-admin-file/internal/logic/file"
	"github.com/suyuan32/simple-admin-file/internal/svc"
	"github.com/suyuan32/simple-admin-file/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strconv"
)

// swagger:route post /upload file Upload
//
// Upload file | 上传文件
//
// Upload file | 上传文件
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UploadReq
//
// Responses:
//  200: UploadResp

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	defer func() {
		if err := recover(); err != nil {
			logx.Error("UploadHandler error")
		}
	}()
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(svcCtx.Config.UploadConf.MaxVideoSize)
		if err != nil {
			httpx.Error(w, err)
		}
		createID := r.FormValue("createID")
		departmentID := r.FormValue("departmentID")
		categoryID := r.FormValue("categoryID")

		i, err := strconv.ParseInt(categoryID, 10, 64)
		if err != nil {
			httpx.Error(w, err)
		}

		req := &types.UploadReq{
			CreateId:     &createID,
			DepartmentID: &departmentID,
			CategoryID:   &i,
		}
		l := file.NewUploadLogic(r, svcCtx)
		resp, err := l.Upload(req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
