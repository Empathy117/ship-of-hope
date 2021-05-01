package dto

import "github.com/empathy117/ship-of-hope/model"

type UserDto struct {
	Name string `json:"name"`
	Telephone string `json:"telephone"`
	Goal int `json:"goal"`
	Rock int `json:"rock"`
	Paper int `json:"paper"`
	Scissor int `json:"scissor"`
	IsPlaying bool `json:"isplaying"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name: user.Name,
		Telephone: user.Telephone,
		Goal: user.Goal,
		Rock: user.Rock,
		Paper: user.Paper,
		Scissor: user.Scissor,
		IsPlaying: user.IsPlaying,
	}
}