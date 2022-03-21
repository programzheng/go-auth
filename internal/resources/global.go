package resources

import (
	"fmt"

	"github.com/programzheng/go-auth/config"
)

var env = config.New()

type Status string

// H is a shortcut for map[string]interface{}
type H map[string]interface{}

func checkIsError(value interface{}) bool {
	_, ok := value.(error)
	return ok
}

func ResponseDebug(response map[string]interface{}) {
	if env.GetString("API_DEBUG") == "true" {
		fmt.Printf("ResponseDebug:%v\n", response)
	}
}

func GlobalResponse(status Status, value interface{}) map[string]interface{} {
	response := make(map[string]interface{})
	response["status"] = status

	if status == "error" && checkIsError(value) {
		response["error"] = value.(error).Error()
		ResponseDebug(response)
		return response
	}

	response["value"] = value
	ResponseDebug(response)
	return response
}
