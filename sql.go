package ulid

import (
	"database/sql/driver"
	"fmt"

	"github.com/oklog/ulid/v2"
	"github.com/snaffi/errors"
)

func (u ULID) Value() (driver.Value, error) {
	return u.String(), nil
}

func (u *ULID) Scan(src any) error {
	if src == nil {
		return nil
	}
	data, ok := src.(string)
	if !ok {
		return fmt.Errorf("ulid expect string data from database, got: %T", src)
	}
	id, err := ulid.ParseStrict(data)
	if err != nil {
		return errors.Wrap(fmt.Sprintf("could not parse ulid, got: %s", data), err)
	}
	*u = ULID(id)
	return nil
}

func (us ULIDS) Value() (driver.Value, error) {
	return fmt.Sprintf("{%s}", us.Join(",")), nil
}
