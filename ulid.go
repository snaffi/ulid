package ulid

import (
	"crypto/rand"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
)

type ULID ulid.ULID

func (u ULID) MarshalText() ([]byte, error) {

	return ulid.ULID(u).MarshalText()
}

func (u *ULID) UnmarshalText(v []byte) error {
	parsed, err := ulid.ParseStrict(string(v))
	if err != nil {
		return err
	}
	*u = ULID(parsed)
	return nil
}

func (u ULID) String() string {
	if u.IsZero() {
		return ""
	}
	return ulid.ULID(u).String()
}

func (u ULID) Time() time.Time {
	return ulid.Time(ulid.ULID(u).Time())
}

var zeroValue = [16]byte{}

func (u ULID) IsZero() bool {
	return [16]byte(u) == zeroValue
}

type ULIDS []ULID

func (us ULIDS) Join(sep string) string {
	switch len(us) {
	case 0:
		return ""
	case 1:
		return us[0].String()
	}
	n := len(sep) * (len(us) - 1)
	for i := 0; i < len(us); i++ {
		n += len(us[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(us[0].String())
	for _, s := range us[1:] {
		b.WriteString(sep)
		b.WriteString(s.String())
	}
	return b.String()
}

func (us ULIDS) Strings() []string {
	res := make([]string, 0, len(us))
	for _, u := range us {
		res = append(res, u.String())
	}
	return res
}

func Parse(val string) (ULID, error) {
	id, err := ulid.ParseStrict(val)
	if err != nil {
		return [16]byte{}, err
	}
	return ULID(id), nil
}

func MustParse(val string) (ULID) {
	id := ulid.MustParse(val)
	return ULID(id)
}

func ParseSlice(val []string) ([]ULID, error) {
	result := make([]ULID, len(val))
	for i, v := range val {
		id, err := Parse(v)
		if err != nil {
			return nil, err
		}
		result[i] = id
	}
	return result, nil
}

func New() ULID {
	now := time.Now()

	entropy := ulid.Monotonic(rand.Reader, 0)
	id := ulid.MustNew(ulid.Timestamp(now), entropy)
	return ULID(id)
}

func Sequence() func() ULID {
	now := time.Now()
	entropy := ulid.Monotonic(rand.Reader, 0)
	return func() ULID {
		id := ulid.MustNew(ulid.Timestamp(now), entropy)
		return ULID(id)
	}
}
