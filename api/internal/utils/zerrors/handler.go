package zerrors

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

type UniqueError struct {
	Code uint16
	Msg  string
}

func (e *UniqueError) Error() string {
	return fmt.Errorf("Error %d: %s", e.Code, e.Msg).Error()
}

//	if err := Func(); err != nil {
//		 	return zerrors.New("code", "msg")
//	}
func New(code uint16, msg string) error {
	return &UniqueError{
		Code: code,
		Msg:  msg,
	}
}

//	if err := zerrors.As(err); err != nil {
//		return nil, err
//	}
func As(err error) *UniqueError {
	if errors.Is(err, &UniqueError{}) {
		return err.(*UniqueError)
	}
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		return New(mysqlErr.Number, mysqlErr.Error()).(*UniqueError)
	}
	// Only when using QueryRow
	if errors.Is(err, sql.ErrNoRows) {
		return New(404, sql.ErrNoRows.Error()).(*UniqueError)
	}
	if errors.Is(err, sql.ErrTxDone) {
		return New(500, sql.ErrTxDone.Error()).(*UniqueError)
	}
	if err != nil {
		return New(501, err.Error()).(*UniqueError)
	}
	return nil
}
