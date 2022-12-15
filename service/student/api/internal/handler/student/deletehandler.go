package student

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/service/student/api/internal/logic/student"
	"go-zero-demo/service/student/api/internal/svc"
	"go-zero-demo/service/student/api/internal/types"
)

func DeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StudentIdParam
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := student.NewDeleteLogic(r.Context(), svcCtx)
		resp, err := l.Delete(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
