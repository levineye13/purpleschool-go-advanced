package stat

import (
	"fmt"
	"net/http"
	"time"
)

type StatHandlerDeps struct {
	StatService *StatService
}

type StatHandler struct {
	baseUrl     string
	StatService *StatService
}

const (
	FilterDay   = "day"
	FilterMonth = "month"
	TimeFormat  = "2006-05-02"
)

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	statHandler := &StatHandler{
		baseUrl:     "/stat",
		StatService: deps.StatService,
	}

	router.Handle("GET "+statHandler.baseUrl, statHandler.GetStat())
}

func (handler *StatHandler) GetStat() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		query := req.URL.Query()

		from, fromErr := time.Parse(TimeFormat, query.Get("from"))
		to, toErr := time.Parse(TimeFormat, query.Get("to"))
		by := query.Get("by")

		if (by != FilterDay && by != FilterMonth) || fromErr != nil || toErr != nil {
			http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		fmt.Println(from, to, by)
	}
}
