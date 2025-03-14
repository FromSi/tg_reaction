package routes

import (
	"github.com/fromsi/tg_reaction/internal/telebot/types"
)

//go:generate mockgen -destination=../../../mocks/telebot/routes/mock_route.go -package=routes_mocks github.com/fromsi/tg_reaction/internal/telebot/routes Route
type Route interface {
	Endpoint() string
	Handle(types.RouteContext) error
}