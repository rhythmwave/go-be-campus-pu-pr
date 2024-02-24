package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/sccicitb/pupr-backend/constants"
)

func checkOrderKeyRegex(order string) bool {
	snakeCaseRegex := regexp.MustCompile(`^[a-z]+\.?[a-z]+(\_{1}([a-z]+))*$`)
	match := snakeCaseRegex.MatchString(order)

	return match
}

// QueryOperation QueryOperation
func QueryOperation(query *string, order map[string]string, group string, limit uint32, page uint32) *constants.ErrorResponse {
	orderVal := []string{}
	limitVal := constants.DefaultLimit
	pageVal := constants.DefaultPage

	if group != "" {
		groupQueryResult := fmt.Sprintf("%s GROUP BY %s", *query, group)
		*query = groupQueryResult
	}
	if order != nil {
		var invalidKeys []string
		var invalidValues []string
		for key, val := range order {
			validOrder := checkOrderKeyRegex(key)
			if !validOrder {
				invalidKeys = append(invalidKeys, key)
			}
			// for i, v := range constants.ValidOrderKey() {
			// 	if v == key {
			// 		break
			// 	}
			// 	if i == len(constants.ValidOrderKey())-1 {
			// 		invalidKeys = append(invalidKeys, key)
			// 	}
			// }
			for i, v := range constants.ValidOrderValue() {
				if v == val {
					break
				}
				if i == len(constants.ValidOrderValue())-1 {
					invalidValues = append(invalidValues, key)
				}
			}
		}
		if len(invalidKeys) != 0 {
			return ErrInvalidOrderKey(invalidKeys)
		}

		if len(invalidValues) != 0 {
			return ErrInvalidOrder(invalidValues)
		}

		for key, val := range order {
			orderVal = append(orderVal, fmt.Sprintf("%s %s", key, val))
		}

		orderQueryResult := fmt.Sprintf("%s ORDER BY %s", *query, strings.Join(orderVal, ","))
		*query = orderQueryResult
	}

	if limit != 0 || page != 0 {
		if limit != 0 {
			limitVal = limit
		}
		if page != 0 {
			pageVal = page
		}

		offsetVal := (pageVal - 1) * limit
		limitQueryResult := fmt.Sprintf("%s LIMIT %d OFFSET %d", *query, limitVal, offsetVal)
		*query = limitQueryResult
	}

	return nil
}
