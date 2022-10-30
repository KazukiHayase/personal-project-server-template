package gqlmodel

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

// ID
// NOTE: Goではint64、GraphQLではstringとして扱う
func MarshalID(i int64) graphql.Marshaler {
	return graphql.MarshalString(strconv.FormatInt(i, 10))
}

func UnmarshalID(v interface{}) (int64, error) {
	switch v := v.(type) {
	case string:
		return strconv.ParseInt(v, 10, 64)
	case int:
		return int64(v), nil
	case int64:
		return v, nil
	case json.Number:
		return strconv.ParseInt(string(v), 10, 64)
	default:
		return 0, fmt.Errorf("%T is not an int64", v)
	}
}
