package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func IsValidWeeks(weeks int) bool {
	validWeeks := []int{3, 5, 10, 15, 20}
	for _, w := range validWeeks {
		if w == weeks {
			return true
		}
	}
	return false
}

func IsValidStadium(stadium int) bool {
	valid := []int{1, 2}
	for _, w := range valid {
		if w == stadium {
			return true
		}
	}
	return false
}

func ParseJsonString(i interface{}) (string, error) {
	json, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	return string(json), nil
}

// helper para obtener y validar un booleano de la consulta
func GetQueryBool(c *gin.Context, key string, defaultValue bool) (bool, error) {
	v, exists := c.GetQuery(key)
	if !exists {
		return defaultValue, nil
	}
	return strconv.ParseBool(v)
}

// helper para obtener y validar una matriz de cadenas de la consulta
func GetQueryStringArray(c *gin.Context, key string, defaultValue []string) ([]string, error) {
	v, exists := c.GetQuery(key)
	if !exists {
		return defaultValue, nil
	}
	arr := strings.Split(v, ",")
	for _, item := range arr {
		if _, err := strconv.Atoi(item); err != nil {
			return nil, fmt.Errorf("invalid value for %s parameter: %v", key, err)
		}
	}
	return arr, nil
}

// helper para obtener y validar un entero de la consulta
func GetQueryInt(c *gin.Context, key string, defaultValue int, validator func(int) bool) (int, error) {
	v, exists := c.GetQuery(key)
	if !exists {
		return defaultValue, nil
	}
	intValue, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("invalid %s parameter: %v", key, err)
	}
	if validator != nil && !validator(intValue) {
		return 0, fmt.Errorf("invalid value for %s parameter: %d", key, intValue)
	}
	return intValue, nil
}
