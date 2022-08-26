package oracle

import (
	"github.com/doug-martin/goqu/v9"
)

func DialectOptions() *goqu.SQLDialectOptions {
	do := goqu.DefaultDialectOptions()

	do.QuoteIdentifiers = false
	do.PlaceHolderFragment = []byte(":")
	do.IncludePlaceholderNum = true
	do.LimitFragment = []byte(" FETCH NEXT ")
	do.LimitEndFragment = []byte(" ROWS ONLY ")

	return do
}

func init() {
	goqu.RegisterDialect("oracle", DialectOptions())
	goqu.RegisterDialect("oci8", DialectOptions())
}
