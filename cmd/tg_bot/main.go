package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fromsi/tg_reaction/internal/adapters"
	"github.com/fromsi/tg_reaction/internal/configs"
	"github.com/fromsi/tg_reaction/internal/services"
	"github.com/fromsi/tg_reaction/internal/telebot/routes"
	"github.com/fromsi/tg_reaction/internal/use_case"
	"github.com/fromsi/tg_reaction/pkg/json"
	"go.uber.org/fx"
	telebot "gopkg.in/telebot.v3"
)

type FxBeforeParams struct {
	fx.In

	Config configs.Config
}

type FxAfterParams struct {
	fx.In

	Config configs.Config
	Routes []routes.Route `group:"routes"`
}

func main() {
	fx.New(CreateApp()).Run()
}

func CreateApp() fx.Option {
	routeAnnotationGroup := fx.ResultTags(`group:"routes"`)

	configJson, err := json.Load("config.json")
	if err != nil {
		log.Fatalf("Configuration loading error: %v", err)
	}

	return fx.Options(
		fx.Provide(
			func() *json.Config {
				return configJson
			},

			func(fxBeforeParams FxBeforeParams) (adapters.TeleBot, error) {
				return NewTelegramBot(fxBeforeParams)
			},

			fx.Annotate(
				adapters.NewTelebotBotAdapter,
				fx.As(new(adapters.BotAdapter)),
			),

			fx.Annotate(
				configs.NewBaseConfig,
				fx.As(new(configs.Config)),
				fx.As(new(configs.AppConfig)),
			),

			fx.Annotate(
				services.NewBaseRegexService,
				fx.As(new(services.RegexService)),
			),
			fx.Annotate(
				services.NewBaseTelebotMethodService,
				fx.As(new(services.TelebotMethodService)),
			),
			fx.Annotate(
				services.NewBaseDiceService,
				fx.As(new(services.DiceService)),
			),

			fx.Annotate(
				use_case.NewBaseRegexReactionUseCase,
				fx.As(new(use_case.RegexReactionUseCase)),
			),
			fx.Annotate(
				use_case.NewBasePartyReactionUseCase,
				fx.As(new(use_case.PartyReactionUseCase)),
			),
			fx.Annotate(
				use_case.NewBaseCryLoudReactionUseCase,
				fx.As(new(use_case.CryLoudReactionUseCase)),
			),
			fx.Annotate(
				use_case.NewBaseEyesReactionUseCase,
				fx.As(new(use_case.EyesReactionUseCase)),
			),
			fx.Annotate(
				use_case.NewBaseDiceReactionUseCase,
				fx.As(new(use_case.DiceReactionUseCase)),
			),

			fx.Annotate(
				routes.NewTextRoute,
				fx.As(new(routes.Route)),
				routeAnnotationGroup,
			),
			fx.Annotate(
				routes.NewMediaRoute,
				fx.As(new(routes.Route)),
				routeAnnotationGroup,
			),
			fx.Annotate(
				routes.NewDiceRoute,
				fx.As(new(routes.Route)),
				routeAnnotationGroup,
			),
			fx.Annotate(
				routes.NewDocumentRoute,
				fx.As(new(routes.Route)),
				routeAnnotationGroup,
			),
			fx.Annotate(
				routes.NewStickerRoute,
				fx.As(new(routes.Route)),
				routeAnnotationGroup,
			),
			fx.Annotate(
				routes.NewNewChatPhotoRoute,
				fx.As(new(routes.Route)),
				routeAnnotationGroup,
			),
			fx.Annotate(
				routes.NewNewChatTitleRoute,
				fx.As(new(routes.Route)),
				routeAnnotationGroup,
			),
			fx.Annotate(
				routes.NewNewMemberRoute,
				fx.As(new(routes.Route)),
				routeAnnotationGroup,
			),
			fx.Annotate(
				routes.NewLeftChatMemberRoute,
				fx.As(new(routes.Route)),
				routeAnnotationGroup,
			),
			fx.Annotate(
				routes.NewChatPhotoDeletedRoute,
				fx.As(new(routes.Route)),
				routeAnnotationGroup,
			),
			fx.Annotate(
				routes.NewPinnedMessageRoute,
				fx.As(new(routes.Route)),
				routeAnnotationGroup,
			),
		),
		fx.Invoke(SetupBot),
	)
}

func NewTelegramBot(fxBeforeParams FxBeforeParams) (adapters.TeleBot, error) {
	p := telebot.Settings{
		Token:  fxBeforeParams.Config.GetTelegramToken(),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(p)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot: %w", err)
	}

	return bot, nil
}

func SetupBot(lifecycle fx.Lifecycle, bot adapters.TeleBot, fxAfterParams FxAfterParams) {
	for _, appRoute := range fxAfterParams.Routes {
		bot.Handle(appRoute.Endpoint(), func(context telebot.Context) error {
			return appRoute.Handle(context)
		})
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(context context.Context) error {
			go bot.Start()
			return nil
		},
		OnStop: func(context context.Context) error {
			bot.Stop()
			return nil
		},
	})
}
