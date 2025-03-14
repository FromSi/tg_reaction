package json

import (
	"testing"
)

func TestReaction_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		reaction Reaction
		want     bool
	}{
		{
			name:     "valid reaction ThumbsUp",
			reaction: ThumbsUp,
			want:     true,
		},
		{
			name:     "valid reaction Heart",
			reaction: Heart,
			want:     true,
		},
		{
			name:     "valid reaction Fire",
			reaction: Fire,
			want:     true,
		},
		{
			name:     "invalid reaction empty",
			reaction: "",
			want:     false,
		},
		{
			name:     "invalid reaction random string",
			reaction: "invalid",
			want:     false,
		},
		{
			name:     "invalid reaction random emoji",
			reaction: "🌈",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.reaction.IsValid(); got != tt.want {
				t.Errorf("Reaction.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidReactionsContainsAllConstants(t *testing.T) {
	// Check that all reaction constants are in ValidReactions
	reactions := []Reaction{
		ThumbsUp, ThumbsDown, Heart, Fire, HeartEyes,
		Clap, Grin, Thinking, MindBlown, Scream,
		Rage, Cry, Party, StarStruck, Vomit,
		Poop, PrayingHands, Ok, Dove, Clown,
		Yawn, WoozyFace, HeartEyesSmile, Whale, HeartOnFire,
		NewMoon, HotDog, Hundred, ROFL, Lightning,
		Banana, Trophy, BrokenHeart, Unamused, Neutral,
		Strawberry, Champagne, Kiss, MiddleFinger, Devil,
		Sleep, CryLoud, Nerd, Ghost, Programmer,
		Eyes, Pumpkin, SeeNoEvil, Innocent, Fearful,
		Handshake, Writing, Hugging, Salute, Santa,
		ChristmasTree, Snowman, NailPolish, Zany, Moai,
		Cool, Cupid, HearNoEvil, Unicorn, BlowingKiss,
		Pill, SpeakNoEvil, Sunglasses, Alien, ShrugMan,
		Shrug, ShrugWoman, Angry,
	}

	for _, reaction := range reactions {
		if !ValidReactions[reaction] {
			t.Errorf("Reaction %q is not in ValidReactions", reaction)
		}
	}

	// Check that the number of reactions matches
	if len(reactions) != len(ValidReactions) {
		t.Errorf("Number of reactions mismatch: constants=%d, ValidReactions=%d", len(reactions), len(ValidReactions))
	}
}
