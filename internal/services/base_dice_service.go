package services

type BaseDiceService struct{}

func NewBaseDiceService() *BaseDiceService {
	return &BaseDiceService{}
}

// IsWinningValue checks if the dice value is a winning one
// In this implementation, value 6 (for standard dice) is considered winning
func (receiver *BaseDiceService) IsWinningValue(value int) bool {
	return value == 6
}
