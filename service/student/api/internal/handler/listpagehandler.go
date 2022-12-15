package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/service/student/api/internal/logic"
	"go-zero-demo/service/student/api/internal/svc"
)

func listPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewListPageLogic(r.Context(), svcCtx)
		resp, err := l.ListPage()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
