package student

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/service/student/api/internal/logic/student"
	"go-zero-demo/service/student/api/internal/svc"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := student.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
