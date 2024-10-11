package response

import "net/http"

func ServiceUnavailableMessage(msg any) (int, any) {
	return http.StatusServiceUnavailable, map[string]any{
		"status":  http.StatusText(http.StatusServiceUnavailable),
		"code":    http.StatusServiceUnavailable,
		"message": msg,
	}
}

func Created(data map[string]any) (int, any) {
	result := map[string]any{
		"status":  http.StatusText(http.StatusCreated),
		"code":    http.StatusCreated,
		"message": "successfully",
	}
	for key, value := range data {
		result[key] = value
	}
	return http.StatusCreated, result
}
