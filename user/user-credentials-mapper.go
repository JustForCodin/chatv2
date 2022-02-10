package user

func UserFromDto(dto UserCredentialsDTO) UserCredentials {
	return UserCredentials{}
}

func UserToDto(userCredentials UserCredentials) UserCredentialsDTO {
	return UserCredentialsDTO{}
}
