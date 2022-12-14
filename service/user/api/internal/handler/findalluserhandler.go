package handler

import (
	"net/http"

	"book/service/user/api/internal/logic"
	"book/service/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
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
