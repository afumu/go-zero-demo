package student

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/service/student/api/internal/logic/student"
	"go-zero-demo/service/student/api/internal/svc"
	"go-zero-demo/service/student/api/internal/types"
)

func EditHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Student
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := student.NewEditLogic(r.Context(), svcCtx)
		resp, err := l.Edit(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
