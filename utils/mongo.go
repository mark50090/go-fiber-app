package utils

import (
	"reflect"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	mapFilterCollection = map[string]func(reflect.Value) bson.E{
		"start_date": func(field reflect.Value) bson.E {
			if field.Kind() != reflect.String {
				return bson.E{Key: "date", Value: nil} // Handle non-string cases
			}
			dateStr := field.Interface().(string)
			parsedTime, err := ParseDate(dateStr)
			if err != nil {
				return bson.E{Key: "date", Value: nil} // Handle parsing errors
			}
			return bson.E{Key: "date", Value: bson.M{"$gte": time.Date(parsedTime.Year(), parsedTime.Month(), parsedTime.Day(), 00, 00, 00, 00, parsedTime.Location())}}
		},

		"end_date": func(field reflect.Value) bson.E {
			if field.Kind() != reflect.String {
				return bson.E{Key: "date", Value: nil} // Handle non-string cases
			}
			dateStr := field.Interface().(string)
			parsedTime, err := ParseDate(dateStr)
			if err != nil {
				return bson.E{Key: "date", Value: nil} // Handle parsing errors
			}
			return bson.E{Key: "date", Value: bson.M{"$lte": time.Date(parsedTime.Year(), parsedTime.Month(), parsedTime.Day(), 23, 59, 59, 99, parsedTime.Location())}}
		},

		"service_type": func(field reflect.Value) bson.E {
			serviceType := field.Interface().(string)
			return bson.E{Key: "service_type", Value: strings.ToLower(serviceType)}
		},
		"code": func(field reflect.Value) bson.E {
			code := field.Interface().(string)
			return bson.E{Key: "reject_detail.code", Value: code}
		},
	}
)

func GetProjectionBson(query interface{}) bson.D {
	val := reflect.ValueOf(query)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()

	projection := bson.D{}

	for i := 0; i < val.NumField(); i++ {
		tag := typ.Field(i).Tag.Get("bson")
		if tag == "" || tag == "-" {
			continue
		}

		projection = append(projection, bson.E{Key: tag, Value: 1})
	}

	return projection
}

func ConvertStructToBsonFilter(query interface{}) (bson.D, error) {
	val := reflect.ValueOf(query)
	typ := val.Type()
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	filter := bson.D{}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		tag := typ.Field(i).Tag.Get("query")
		if tag == "" {
			continue
		}

		if field.Kind() == reflect.Ptr && field.IsNil() {
			continue
		}

		if field.Kind() == reflect.Ptr {
			field = field.Elem()
		}

		if (field.Kind() == reflect.String && field.String() == "") ||
			(field.Kind() == reflect.Int && field.Int() == 0) ||
			(field.Kind() == reflect.Slice && field.Len() == 0) {
			continue
		}

		if val, ok := mapFilterCollection[tag]; ok {
			filter = append(filter, val(field))
			continue
		}

		if field.Kind() == reflect.Slice {
			filter = append(filter, bson.E{
				Key:   tag,
				Value: bson.M{"$in": field.Interface()},
			})
		} else {
			filter = append(filter, bson.E{Key: tag, Value: field.Interface()})
		}

	}
	return filter, nil
}
