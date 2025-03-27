package geometry

import (
	"database/sql/driver"
	"fmt"

	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/wkt"
)

type Geometry struct {
	geom.T
}

func (g Geometry) Value() (driver.Value, error) {
	if g.T == nil {
		return nil, nil
	}
	wktStr, err := wkt.Marshal(g.T)
	if err != nil {
		return nil, err
	}
	return wktStr, nil
}

func (g *Geometry) Scan(src interface{}) error {
	if src == nil {
		g.T = nil
		return nil
	}
	var s string
	switch v := src.(type) {
	case []byte:
		s = string(v)
	case string:
		s = v
	default:
		return fmt.Errorf("unsupported type for Geometry: %T", src)
	}
	t, err := wkt.Unmarshal(s)
	if err != nil {
		return err
	}
	g.T = t
	return nil
}
