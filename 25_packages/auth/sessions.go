package auth

func processSession() string{
	return "user is logged in"
}

func GetSession() string{
	return processSession()
}