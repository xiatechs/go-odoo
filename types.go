package odoo

import (
	"encoding/json"
)

// OMany2one is a struct that hold Many2one value from Odoo
type OMany2one struct {
	ID   int
	Name string
}

// UnmarshalJSON parses Many2one struct from JSON
func (m *OMany2one) UnmarshalJSON(b []byte) error {
	var i interface{}
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}
	switch i.(type) {
	case []interface{}:
		slice := i.([]interface{})
		*m = OMany2one{
			ID:   slice[0].(int),
			Name: slice[1].(string),
		}
	case bool:
		*m = OMany2one{}
	}
	return nil
}
