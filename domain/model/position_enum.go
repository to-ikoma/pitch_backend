package model

type Position int

const (
	Unspecified Position = iota
	Pitcher
	Catcher
	FirstBase
	SecondBase
	ThirdBase
	ShortStop
	LeftField
	CenterField
	RightField
	DesignatedHeitter
)

func (p Position) String() string {
	switch p {
	case Unspecified:
		return "unspecified"
	case Pitcher:
		return "pitcher"
	case Catcher:
		return "catcher"
	case FirstBase:
		return "firstbase"
	case SecondBase:
		return "secondbase"
	case ThirdBase:
		return "thirdbase"
	case ShortStop:
		return "shortstop"
	case LeftField:
		return "leftfield"
	case CenterField:
		return "centerfield"
	case RightField:
		return "rightfield"
	case DesignatedHeitter:
		return "designatedheitter"
	default:
		return "unspecified"
	}
}

func NewPositionFromString(s string) Position {
	switch s {
	case Unspecified.String():
		return Unspecified
	case Pitcher.String():
		return Pitcher
	case Catcher.String():
		return Catcher
	case FirstBase.String():
		return FirstBase
	case SecondBase.String():
		return SecondBase
	case ThirdBase.String():
		return ThirdBase
	case ShortStop.String():
		return ShortStop
	case LeftField.String():
		return LeftField
	case CenterField.String():
		return CenterField
	case RightField.String():
		return RightField
	case DesignatedHeitter.String():
		return DesignatedHeitter
	default:
		return Unspecified
	}
}
