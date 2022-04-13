package handler

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Handler struct {
	ctx context.Context
}

func NewHandler(ctx context.Context) *Handler {
	return &Handler{
		ctx: ctx,
	}
}

func (h *Handler) InitAppHandler() {
	runtime.EventsOn(h.ctx, "login", loginHandler)
}

func loginHandler(loginData ...interface{}) {
	fmt.Println(loginData...)
}
