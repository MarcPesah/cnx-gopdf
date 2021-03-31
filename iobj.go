package gopdf

import (
	"io"
)

//IObj inteface for all pdf object
type IObj interface {
	init(func() *GoPdf2)
	getType() string
	write(w io.Writer, objID int) error
}
