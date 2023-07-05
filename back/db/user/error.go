package user

import (
	"errors"
	"fmt"
)

var (
	ERR_FORBIDEN_OPPERATION = errors.New("Opperration forbiden")
	ERR_NO_FIELD_NAME       = errors.New("No Field find")
	ERR_BAD_VALUE_TYPE      = errors.New("Bad value kind")
)

func NoFieldName(f_name string) error {
	return fmt.Errorf("%w (%s)", ERR_NO_FIELD_NAME, f_name)
}
