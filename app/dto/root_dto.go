package dto

type GetUsersResponse struct {
	ListResponse[[]UserResponse]
}
