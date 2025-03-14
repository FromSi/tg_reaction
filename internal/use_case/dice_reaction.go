package use_case

//go:generate mockgen -destination=../../mocks/use_case/mock_dice_reaction_use_case.go -package=use_case_mocks github.com/fromsi/tg_reaction/internal/use_case DiceReactionUseCase
type DiceReactionUseCase interface {
	Execute(chatId int64, messageId int, diceValue int) error
}
