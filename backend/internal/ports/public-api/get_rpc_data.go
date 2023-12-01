package public_api

import (
	"github.com/go-chi/render"
	"go.uber.org/zap"
	"net/http"
	"public-rpc/internal/app/public-api/query"
	"public-rpc/models"
)

func GetRPCDataHandler(handler query.GetRPCDataHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		ctx := r.Context()

		handler.Logger.Info("", zap.Any("query", r.URL.Query()))

		network := r.URL.Query().Get("network")
		q := query.GetRPCDataQuery{Network: network}

		rpcData, err := handler.Handle(ctx, q)

		if err != nil {
			handler.Logger.Error("GetRPCDataHandler error", zap.Error(err))
			httpErr := models.HttpError{
				Code:    http.StatusInternalServerError,
				Message: "Internal server error",
			}
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, httpErr)
		} else {
			render.JSON(w, r, rpcData)
		}
	}
}
