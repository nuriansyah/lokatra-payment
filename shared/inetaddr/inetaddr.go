package inetaddr

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"net/netip"
	"strings"
)

// NullIP stores an optional IP address and supports SQL scan/value conversion.
type NullIP struct {
	Addr  netip.Addr
	Valid bool
}

func FromAddr(addr netip.Addr) NullIP {
	return NullIP{Addr: addr, Valid: true}
}

func FromString(value string) (NullIP, error) {
	if strings.TrimSpace(value) == "" {
		return NullIP{}, nil
	}
	addr, err := netip.ParseAddr(strings.TrimSpace(value))
	if err != nil {
		return NullIP{}, err
	}
	return NullIP{Addr: addr, Valid: true}, nil
}

func (n *NullIP) Scan(value interface{}) error {
	if n == nil {
		return fmt.Errorf("inetaddr: NullIP receiver is nil")
	}

	switch v := value.(type) {
	case nil:
		n.Valid = false
		n.Addr = netip.Addr{}
		return nil
	case []byte:
		return n.scanString(string(v))
	case string:
		return n.scanString(v)
	default:
		n.Valid = false
		n.Addr = netip.Addr{}
		return fmt.Errorf("inetaddr: cannot scan type %T into NullIP", value)
	}
}

func (n *NullIP) scanString(value string) error {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		n.Valid = false
		n.Addr = netip.Addr{}
		return nil
	}
	addr, err := netip.ParseAddr(trimmed)
	if err != nil {
		n.Valid = false
		n.Addr = netip.Addr{}
		return err
	}
	n.Addr = addr
	n.Valid = true
	return nil
}

func (n NullIP) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Addr.String(), nil
}

func (n NullIP) String() string {
	if !n.Valid {
		return ""
	}
	return n.Addr.String()
}

func (n NullIP) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.Addr.String())
}
