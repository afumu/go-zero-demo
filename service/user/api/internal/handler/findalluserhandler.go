package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/service/user/api/internal/logic"
	"go-zero-demo/service/user/api/internal/svc"
)

func findAllUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewFindAllUserLogic(r.Context(), svcCtx)
		resp, err := l.FindAllUser()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
