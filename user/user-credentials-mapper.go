package user

func UserFromDto(dto UserCredentialsDTO) UserCredentials {
	return UserCredentials{Email: dto.Email, Password: dto.Password}
}

func UserToDto(userCredentials UserCredentials) UserCredentialsDTO {
	return UserCredentialsDTO{Email: userCredentials.Email, Password: userCredentials.Password}
}
