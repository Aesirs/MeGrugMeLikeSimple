package handlers

import (
	"net/http"
	"strconv"
	"sync/atomic"

	"MeGrugMeLikeSimple/internal/views"
)

type HomeHandler struct {
	counter atomic.Int64
}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	views.HomePage().Render(r.Context(), w)
}

func (h *HomeHandler) Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<p class='font-semibold'>Hello from Chi! 🎉</p>"))
}

func (h *HomeHandler) Counter(w http.ResponseWriter, r *http.Request) {
	val := h.counter.Add(1)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(strconv.FormatInt(val, 10)))
}
