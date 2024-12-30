package model

import "testing"

func TestPostion_String(t *testing.T) {
	tests := []struct {
		name string
		p    Position
		want string
	}{
		{
			name: "Unspecified",
			p:    Unspecified,
			want: "unspecified",
		},
		{
			name: "Pitcher",
			p:    Pitcher,
			want: "pitcher",
		},
		{
			name: "Catcher",
			p:    Catcher,
			want: "catcher",
		},
		{
			name: "FirstBase",
			p:    FirstBase,
			want: "firstbase",
		},
		{
			name: "SecondBase",
			p:    SecondBase,
			want: "secondbase",
		},
		{
			name: "ThirdBase",
			p:    ThirdBase,
			want: "thirdbase",
		},
		{
			name: "ShortStop",
			p:    ShortStop,
			want: "shortstop",
		},
		{
			name: "LeftField",
			p:    LeftField,
			want: "leftfield",
		},
		{
			name: "CenterField",
			p:    CenterField,
			want: "centerfield",
		},
		{
			name: "RightField",
			p:    RightField,
			want: "rightfield",
		},
		{
			name: "DesignatedHeitter",
			p:    DesignatedHeitter,
			want: "designatedheitter",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("Position.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPostionFromString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want Position
	}{
		{
			name: "Unspecified",
			s:    "unspecified",
			want: Unspecified,
		},
		{
			name: "Pitcher",
			s:    "pitcher",
			want: Pitcher,
		},
		{
			name: "Catcher",
			s:    "catcher",
			want: Catcher,
		},
		{
			name: "FirstBase",
			s:    "firstbase",
			want: FirstBase,
		},
		{
			name: "SecondBase",
			s:    "secondbase",
			want: SecondBase,
		},
		{
			name: "ThirdBase",
			s:    "thirdbase",
			want: ThirdBase,
		},
		{
			name: "ShortStop",
			s:    "shortstop",
			want: ShortStop,
		},
		{
			name: "LeftField",
			s:    "leftfield",
			want: LeftField,
		},
		{
			name: "CenterField",
			s:    "centerfield",
			want: CenterField,
		},
		{
			name: "RightField",
			s:    "rightfield",
			want: RightField,
		},
		{
			name: "DesignatedHeitter",
			s:    "designatedheitter",
			want: DesignatedHeitter,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPositionFromString(tt.s); got != tt.want {
				t.Errorf("NewPositionFromString = %v, want %v", got, tt.want)
			}
		})
	}
}
