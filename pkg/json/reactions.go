package json

import (
	"regexp"
)

// Reaction represents a valid emoji reaction
type Reaction string

// All valid reactions
const (
	Empty          Reaction = "" // Empty reaction for clearing
	ThumbsUp       Reaction = "👍"
	ThumbsDown     Reaction = "👎"
	Heart          Reaction = "❤️"
	Fire           Reaction = "🔥"
	HeartEyes      Reaction = "🥰"
	Clap           Reaction = "👏"
	Grin           Reaction = "😁"
	Thinking       Reaction = "🤔"
	MindBlown      Reaction = "🤯"
	Scream         Reaction = "😱"
	Rage           Reaction = "🤬"
	Cry            Reaction = "😢"
	Party          Reaction = "🎉"
	StarStruck     Reaction = "🤩"
	Vomit          Reaction = "🤮"
	Poop           Reaction = "💩"
	PrayingHands   Reaction = "🙏"
	Ok             Reaction = "👌"
	Dove           Reaction = "🕊"
	Clown          Reaction = "🤡"
	Yawn           Reaction = "🥱"
	WoozyFace      Reaction = "🥴"
	HeartEyesSmile Reaction = "😍"
	Whale          Reaction = "🐳"
	HeartOnFire    Reaction = "❤️‍🔥"
	NewMoon        Reaction = "🌚"
	HotDog         Reaction = "🌭"
	Hundred        Reaction = "💯"
	ROFL           Reaction = "🤣"
	Lightning      Reaction = "⚡️"
	Banana         Reaction = "🍌"
	Trophy         Reaction = "🏆"
	BrokenHeart    Reaction = "💔"
	Unamused       Reaction = "🤨"
	Neutral        Reaction = "😐"
	Strawberry     Reaction = "🍓"
	Champagne      Reaction = "🍾"
	Kiss           Reaction = "💋"
	MiddleFinger   Reaction = "🖕"
	Devil          Reaction = "😈"
	Sleep          Reaction = "😴"
	CryLoud        Reaction = "😭"
	Nerd           Reaction = "🤓"
	Ghost          Reaction = "👻"
	Programmer     Reaction = "👨‍💻"
	Eyes           Reaction = "👀"
	Pumpkin        Reaction = "🎃"
	SeeNoEvil      Reaction = "🙈"
	Innocent       Reaction = "😇"
	Fearful        Reaction = "😨"
	Handshake      Reaction = "🤝"
	Writing        Reaction = "✍️"
	Hugging        Reaction = "🤗"
	Salute         Reaction = "🫡"
	Santa          Reaction = "🎅"
	ChristmasTree  Reaction = "🎄"
	Snowman        Reaction = "☃️"
	NailPolish     Reaction = "💅"
	Zany           Reaction = "🤪"
	Moai           Reaction = "🗿"
	Cool           Reaction = "🆒"
	Cupid          Reaction = "💘"
	HearNoEvil     Reaction = "🙉"
	Unicorn        Reaction = "🦄"
	BlowingKiss    Reaction = "😘"
	Pill           Reaction = "💊"
	SpeakNoEvil    Reaction = "🙊"
	Sunglasses     Reaction = "😎"
	Alien          Reaction = "👾"
	ShrugMan       Reaction = "🤷‍♂️"
	Shrug          Reaction = "🤷"
	ShrugWoman     Reaction = "🤷‍♀️"
	Angry          Reaction = "😡"
)

// ValidReactions contains all valid reactions
var ValidReactions = map[Reaction]bool{
	ThumbsUp: true, ThumbsDown: true, Heart: true, Fire: true, HeartEyes: true,
	Clap: true, Grin: true, Thinking: true, MindBlown: true, Scream: true,
	Rage: true, Cry: true, Party: true, StarStruck: true, Vomit: true,
	Poop: true, PrayingHands: true, Ok: true, Dove: true, Clown: true,
	Yawn: true, WoozyFace: true, HeartEyesSmile: true, Whale: true, HeartOnFire: true,
	NewMoon: true, HotDog: true, Hundred: true, ROFL: true, Lightning: true,
	Banana: true, Trophy: true, BrokenHeart: true, Unamused: true, Neutral: true,
	Strawberry: true, Champagne: true, Kiss: true, MiddleFinger: true, Devil: true,
	Sleep: true, CryLoud: true, Nerd: true, Ghost: true, Programmer: true,
	Eyes: true, Pumpkin: true, SeeNoEvil: true, Innocent: true, Fearful: true,
	Handshake: true, Writing: true, Hugging: true, Salute: true, Santa: true,
	ChristmasTree: true, Snowman: true, NailPolish: true, Zany: true, Moai: true,
	Cool: true, Cupid: true, HearNoEvil: true, Unicorn: true, BlowingKiss: true,
	Pill: true, SpeakNoEvil: true, Sunglasses: true, Alien: true, ShrugMan: true,
	Shrug: true, ShrugWoman: true, Angry: true,
}

// IsValid checks if the reaction is valid
func (r Reaction) IsValid() bool {
	return ValidReactions[r]
}

// Pattern represents a pattern for matching and its corresponding reactions
type Pattern struct {
	Pattern   *regexp.Regexp
	Reactions []Reaction
}

// Holiday represents a holiday period with special reactions
type Holiday struct {
	StartDay   int
	StartMonth int
	EndDay     int
	EndMonth   int
	Reactions  []Reaction
	Patterns   []Pattern
}

// Config represents the complete reactions configuration
type Config struct {
	Everyday []Pattern
	Holidays map[string]Holiday
}
