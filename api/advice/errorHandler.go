package advice

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// --------------------------------------------------------------------
// APIError defines a custom error type with a status code and message.
// --------------------------------------------------------------------
type ErrorResponse struct {
    StatusCode int
    Message    string
}
 
// ------------------------------------
// NewAPIError creates a new APIError.
// ------------------------------------
func NewAPIError(statusCode int, message string) *ErrorResponse {
    return &ErrorResponse{
        StatusCode: statusCode,
        Message:    message,
    }
}

// ----------------------------------------
// NewError handles regular error messages.
// ----------------------------------------
func NewError (message string) *ErrorResponse {
    return &ErrorResponse {
        Message: message,
    }
}
 
func (e *ErrorResponse) Error() string {
    return e.Message
}
 
// ----------------------------------------------------------------------
// ErrorMiddleware adapts the existing middleware logic to work with Gin.
// ----------------------------------------------------------------------
func ErrorMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            // Recovering from panic (if any) and converting it to an ErrorResponse
            if rec := recover(); rec != nil {
                var errResp *ErrorResponse
 
                // Type assertion to check if it's an ErrorResponse
                if err, ok := rec.(*ErrorResponse); ok {
                    errResp = err
                } else {
                    // For non-ErrorResponse types, returning a generic server error
                    errResp = NewAPIError(http.StatusInternalServerError, "Internal Server Error")
                }
 
                // Responding with the error response in JSON format
                c.JSON(errResp.StatusCode, errResp)
                c.Abort() // Prevents calling any remaining handlers
            }
        }()
 
        c.Next() // Processing the request to next middleware
    }
}