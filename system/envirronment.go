package System

import (
	"net/http"
)

// GetField from uri
func GetField(r *http.Request, index int) string{
	field := r.Context().Value("fields").([]string)
	return field[index]
}