package request

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

func HandleQueryParams(c *gin.Context) map[string]interface{} {
	queryParams := c.Request.URL.Query()
	paramMap := make(map[string]interface{})
	filters := make(map[string]interface{})

	// Check and handle "page" parameter
	if pageValue, exists := queryParams["page"]; exists && len(pageValue) > 0 {
		if page, err := strconv.Atoi(pageValue[0]); err == nil {
			paramMap["page"] = page
		} else {
			paramMap["page"] = pageValue[0]
		}
		delete(queryParams, "page")
	} else {
		paramMap["page"] = 1
	}

	// Check and handle "itemsPerPage" parameter
	if itemsPerPageValue, exists := queryParams["itemsPerPage"]; exists && len(itemsPerPageValue) > 0 {
		if itemsPerPage, err := strconv.Atoi(itemsPerPageValue[0]); err == nil {
			paramMap["itemsPerPage"] = itemsPerPage
		} else {
			paramMap["itemsPerPage"] = itemsPerPageValue[0]
		}
		delete(queryParams, "itemsPerPage")
	} else {
		paramMap["itemsPerPage"] = 20
	}

	// Group remaining parameters into filters
	for key, value := range queryParams {
		if len(value) == 1 {
			if intValue, err := strconv.Atoi(value[0]); err == nil {
				filters[key] = intValue
			} else {
				filters[key] = value[0]
			}
		} else {
			filters[key] = value
		}
	}

	// Add filters to paramMap if there are any
	if len(filters) > 0 {
		paramMap["filters"] = filters
	}

	// Group path parameters
	paramMap["path"] = handleDynamicPath(c)

	return paramMap
}

func handleDynamicPath(c *gin.Context) map[string]interface{} {
	params := make(map[string]interface{})

	// Iterate over all path parameters and add them to the map
	for _, param := range c.Params {
		params[param.Key] = param.Value
	}

	return params
}
