package types

import telebot "gopkg.in/telebot.v3"

//go:generate mockgen -destination=../../../mocks/telebot/types/mock_route_context.go -package=types_mocks github.com/fromsi/tg_reaction/internal/telebot/types RouteContext
type RouteContext interface {
	telebot.Context
}