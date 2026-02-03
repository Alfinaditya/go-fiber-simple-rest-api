package dto

import "github.com/Alfinaditya/go-fiber-simple-rest-api/pkg/utils"

// BaseResponse represents a standard API response structure
// Used for simple success/error responses
type BaseResponse struct {
	Error bool   `json:"error" example:"false"`
	Msg   string `json:"msg" example:"Success"`
}

// ValidationErrorResponse represents a validation error response with field-specific errors
type ValidationErrorResponse struct {
	Error  bool                    `json:"error" example:"true"`
	Msg    string                  `json:"msg" example:"Validation failed"`
	Errors []utils.ValidatorErrors `json:"errors"`
}

// ErrorDetailResponse represents a detailed error response
type ErrorDetailResponse struct {
	Error   bool   `json:"error" example:"true"`
	Msg     string `json:"msg" example:"Invalid request"`
	Details string `json:"details,omitempty" example:"Additional error information"`
}

// DataResponse represents a response with a single data object
type DataResponse[T any] struct {
	Error bool   `json:"error" example:"false"`
	Msg   string `json:"msg" example:"Success"`
	Data  T      `json:"data"`
}

// ListResponse represents a paginated list response with generic data type
// Used when returning lists/arrays with metadata
type ListResponse[T any] struct {
	BaseResponse
	Count int `json:"count" example:"5"`
	Data  T   `json:"data"`
}

// SuccessResponse creates a successful response
//
// Usage in handlers:
//
//	return c.JSON(dto.SuccessResponse("Author deleted successfully"))
//
// Response JSON:
//
//	{"error": false, "msg": "Author deleted successfully"}
func SuccessResponse(msg string) BaseResponse {
	return BaseResponse{
		Error: false,
		Msg:   msg,
	}
}

// ErrorResponse creates an error response
//
// Usage in handlers:
//
//	return c.Status(400).JSON(dto.ErrorResponse("Invalid author ID format"))
//
// Response JSON:
//
//	{"error": true, "msg": "Invalid author ID format"}
func ErrorResponse(msg string) BaseResponse {
	return BaseResponse{
		Error: true,
		Msg:   msg,
	}
}

// NewListResponse creates a list response with data
//
// Usage in handlers:
//
//	authors := []dto.AuthorResponse{ ... }
//	return c.JSON(dto.NewListResponse(authors, len(authors), ""))
//
// Response JSON:
//
//	{"error": false, "msg": "", "count": 5, "data": [...]}
//
// Or with custom message:
//
//	return c.JSON(dto.NewListResponse(authors, len(authors), "Authors retrieved successfully"))
func NewListResponse[T any](data T, count int, msg string) ListResponse[T] {
	return ListResponse[T]{
		BaseResponse: BaseResponse{
			Error: false,
			Msg:   msg,
		},
		Count: count,
		Data:  data,
	}
}

// ValidationErrorResponseFunc creates a validation error response
//
// Usage in handlers:
//
//	if errors := utils.ValidateStruct(req); errors != nil {
//	    return c.Status(400).JSON(dto.ValidationErrorResponseFunc(errors))
//	}
//
// Response JSON:
//
//	{"error": true, "msg": "Validation failed", "errors": {...}}
func ValidationErrorResponseFunc(errors []utils.ValidatorErrors) ValidationErrorResponse {
	return ValidationErrorResponse{
		Error:  true,
		Msg:    "Validation failed",
		Errors: errors,
	}
}

// ErrorDetailResponseFunc creates a detailed error response
//
// Usage in handlers:
//
//	return c.Status(500).JSON(dto.ErrorDetailResponseFunc("Database error", err.Error()))
//
// Response JSON:
//
//	{"error": true, "msg": "Database error", "details": "connection refused"}
func ErrorDetailResponseFunc(msg string, details string) ErrorDetailResponse {
	return ErrorDetailResponse{
		Error:   true,
		Msg:     msg,
		Details: details,
	}
}

// NewDataResponse creates a response with a single data object
//
// Usage in handlers:
//
//	author := dto.AuthorResponse{...}
//	return c.JSON(dto.NewDataResponse(author, "Author created successfully"))
//
// Response JSON:
//
//	{"error": false, "msg": "Author created successfully", "data": {...}}
func NewDataResponse[T any](data T, msg string) DataResponse[T] {
	return DataResponse[T]{
		Error: false,
		Msg:   msg,
		Data:  data,
	}
}
