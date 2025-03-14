package use_case

//go:generate mockgen -destination=../../mocks/use_case/mock_poop_reaction.go -package=use_case_mocks github.com/fromsi/tg_reaction/internal/use_case PoopReactionUseCase
type PoopReactionUseCase interface {
	Execute(chatId int64, messageId int) error
}
