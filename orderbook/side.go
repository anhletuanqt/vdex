package orderbook

import (
	"encoding/json"
	"reflect"
)

// Side of the order
type Side string

// Sell (asks) or Buy (bids)
const (
	Sell = Side("sell")
	Buy  = Side("buy")
)

// String implements fmt.Stringer interface
func (s Side) String() string {
	return string(s)
}

func (s Side) Opposite() Side {
	if s == Buy {
		return Sell
	}
	return Buy
}

func SideFromStr(s string) Side {
	if s == "buy" {
		return Buy
	}

	return Sell
}

// MarshalJSON implements json.Marshaler interface
func (s Side) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

// UnmarshalJSON implements json.Unmarshaler interface
func (s *Side) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"buy"`:
		*s = Buy
	case `"sell"`:
		*s = Sell
	default:
		return &json.UnsupportedValueError{
			Value: reflect.New(reflect.TypeOf(data)),
			Str:   string(data),
		}
	}

	return nil
}
