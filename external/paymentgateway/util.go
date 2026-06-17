package paymentgateway

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func BasicAuthValue(username, password string) string {
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))
}

func RawJSON(v any) json.RawMessage {
	b, _ := json.Marshal(v)
	return b
}

func AmountToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(strings.TrimSpace(s), 64)
}

func AmountToInt64(s string) (int64, error) {
	f, err := AmountToFloat64(s)
	if err != nil {
		return 0, err
	}
	return int64(f), nil
}

func StringAmount(v any) string {
	switch t := v.(type) {
	case string:
		return t
	case float64:
		return fmt.Sprintf("%.2f", t)
	case int:
		return fmt.Sprintf("%d.00", t)
	case int64:
		return fmt.Sprintf("%d.00", t)
	default:
		return ""
	}
}

func PtrTime(t time.Time) *time.Time { return &t }
