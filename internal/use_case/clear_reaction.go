package use_case

//go:generate mockgen -destination=../../mocks/use_case/mock_clear_reaction_use_case.go -package=use_case_mocks github.com/fromsi/tg_reaction/internal/use_case ClearReactionUseCase
type ClearReactionUseCase interface {
	Execute(chatId int64, messageId int) error
}
