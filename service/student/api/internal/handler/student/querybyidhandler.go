package student

import (
	"net/http"

	"book/service/student/api/internal/logic/student"
	"book/service/student/api/internal/svc"
	"book/service/student/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StudentIdParam
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := student.NewQueryByIdLogic(r.Context(), svcCtx)
		resp, err := l.QueryById(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
