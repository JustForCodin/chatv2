package user

func FromDto(dto UserDto) User {
	return User{ID: dto.ID, Name: dto.Name, Email: dto.Email, Status: dto.Status,
		RegisteredAt: dto.RegisteredAt}
}

func ToDto(user User) UserDto {
	return UserDto{ID: user.ID, Name: user.Name, Email: user.Email, Status: user.Status,
		RegisteredAt: user.RegisteredAt}
}
