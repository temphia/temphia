package dyndb

const (

	// column types ctype
	CtypeShortText = "shorttext"
	CtypePhone     = "phonenumber"
	CtypeSelect    = "select"
	CtypeRFormula  = "rowformula"
	CtypeFile      = "file"
	CtypeMultiFile = "multifile"
	CtypeCheckBox  = "checkbox"
	CtypeCurrency  = "currency"
	CtypeNumber    = "number"
	CtypeLocation  = "location"
	CtypeDateTime  = "datetime"

	CtypeMultSelect  = "multiselect"
	CtypeLongText    = "longtext"
	CtypeSingleUser  = "singleuser"
	CtypeMultiUser   = "multiuser"
	CtypeEmail       = "email"
	CtypeJSON        = "json"
	CtypeRangeNumber = "rangenumber"
	CtypeColor       = "color"
)

var AllColumns = map[string]struct{}{
	CtypeShortText:   {},
	CtypePhone:       {},
	CtypeSelect:      {},
	CtypeRFormula:    {},
	CtypeFile:        {},
	CtypeMultiFile:   {},
	CtypeCheckBox:    {},
	CtypeCurrency:    {},
	CtypeNumber:      {},
	CtypeLocation:    {},
	CtypeDateTime:    {},
	CtypeMultSelect:  {},
	CtypeLongText:    {},
	CtypeSingleUser:  {},
	CtypeMultiUser:   {},
	CtypeEmail:       {},
	CtypeJSON:        {},
	CtypeRangeNumber: {},
	CtypeColor:       {},
}

const (
	RefHardPriId = "hard_pri"
	RefSoftPriId = "soft_pri"
	RefHardText  = "hard_text"
	RefSoftText  = "soft_text"
	RefHardMulti = "hard_multi"
)

const (
	// index types
	IndexUnique = "unique"
	IndexNormal = "normal"
	IndexFTS    = "fts"
)
