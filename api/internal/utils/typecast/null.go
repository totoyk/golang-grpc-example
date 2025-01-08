package typecast

import (
	"database/sql"
	"time"
)

func NullString(s *string) sql.NullString {
	if s == nil || *s == "" {
		return sql.NullString{}
	}
	return sql.NullString{
		String: *s,
		Valid:  true,
	}
}

func NullInt64(i *int64) sql.NullInt64 {
	if i == nil {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: *i,
		Valid: true,
	}
}

func NullTime(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{}
	}
	return sql.NullTime{
		Time:  *t,
		Valid: true,
	}
}

func NullStringToString(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return ""
}

func NullInt64ToInt64(i sql.NullInt64) int64 {
	if i.Valid {
		return i.Int64
	}
	return 0
}

func NullTimeToTime(t sql.NullTime) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}

func NullBoolToBool(b sql.NullBool) bool {
	if b.Valid {
		return b.Bool
	}
	return false
}
