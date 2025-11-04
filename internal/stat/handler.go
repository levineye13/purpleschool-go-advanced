package stat

import (
	"net/http"
	"purpleschool-go/advanced/pkg/res"
	"time"
)

type StatHandlerDeps struct {
	StatService    *StatService
	StatRepository *StatRepository
}

type StatHandler struct {
	baseUrl        string
	StatService    *StatService
	StatRepository *StatRepository
}

const (
	GroupByDay   = "day"
	GroupByMonth = "month"
	TimeFormat   = "2006-05-02"
)

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	statHandler := &StatHandler{
		baseUrl:        "/stat",
		StatService:    deps.StatService,
		StatRepository: deps.StatRepository,
	}

	router.Handle("GET "+statHandler.baseUrl, statHandler.GetStat())
}

func (handler *StatHandler) GetStat() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		query := req.URL.Query()

		from, fromErr := time.Parse(TimeFormat, query.Get("from"))
		to, toErr := time.Parse(TimeFormat, query.Get("to"))
		by := query.Get("by")

		if (by != GroupByDay && by != GroupByMonth) || fromErr != nil || toErr != nil {
			http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		stats := handler.StatRepository.GetAll(from, to, by)

		res.Json(rw, http.StatusOK, stats)
	}
}
