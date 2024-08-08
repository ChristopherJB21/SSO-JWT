package user

type UserResponse struct {
	IDUser   uint   `json:"IDUser"`
	UserName string `json:"UserName"`
}

func ToUserResponse(item User) UserResponse {
	return UserResponse{
		IDUser:   item.ID,
		UserName: item.UserName,
	}
}

func ToUsersResponses(items []User) []UserResponse {
	var userResponses []UserResponse
	for _, item := range items {
		userResponses = append(userResponses, ToUserResponse(item))
	}
	return userResponses
}
