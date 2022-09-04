package dynhub

import (
	"net/mail"
	"regexp"
	"strings"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"

	"github.com/thoas/go-funk"
)

type Validator struct {
	columns map[string]*entities.Column
}

func (v *Validator) ValidateRows(rows []map[string]any) error {
	for _, row := range rows {

		err := v.ValidateRow(row)
		if err != nil {
			return err
		}
	}

	return nil
}

func (v *Validator) ValidateRow(row map[string]any) error {

	for key, cell := range row {
		if store.IsMeta(key) {
			continue
		}
		col := v.columns[key]
		if col == nil {
			pp.Println("emty col")
			continue
		}

		switch col.Ctype {
		case store.CtypeShortText:
			if !col.StrictPattern || col.Pattern == "" {
				continue
			}

			str, ok := cell.(string)
			if !ok {
				return easyerr.Error("bad type")
			}
			err := textValidator(col, str)
			if err != nil {
				return err
			}
		case store.CtypePhone:
		case store.CtypeSelect:
			str, ok := cell.(string)
			if !ok {
				return easyerr.Error("bad type")
			}

			if !funk.ContainsString(col.Options, str) {
				return easyerr.Error("invalid select option")
			}

		case store.CtypeRFormula:
		case store.CtypeFile:
		case store.CtypeMultiFile:
		case store.CtypeCheckBox:
		case store.CtypeCurrency:
		case store.CtypeNumber:
		case store.CtypeLocation:
		case store.CtypeDateTime:

		default:
			panic("not implemented")
		}

	}

	return nil

}

func RegexValidateBuilder(regex string) func(value string) error {
	reg := regexp.MustCompile(regex)
	return func(value string) error {
		if reg.MatchString(value) {
			return easyerr.Error("regex err")
		}
		return nil
	}
}

var Exprs = map[string]func(col *entities.Column, value string) error{
	"email": func(col *entities.Column, value string) error {
		_, err := mail.ParseAddress(value)
		return err
	},
}

func textValidator(col *entities.Column, value string) error {
	if strings.HasPrefix(col.Pattern, "regex ") {
		expr := regexp.MustCompile(strings.TrimPrefix(col.Pattern, "regex "))
		if expr.MatchString(value) {
			return easyerr.Error("regex expr not match")
		}
	}

	expr, ok := Exprs[col.Pattern]
	if !ok {
		return easyerr.Error("pattern not found")
	}

	return expr(col, value)
}
