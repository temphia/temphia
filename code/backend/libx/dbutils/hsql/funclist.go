package hsql

var funcWhiteList = map[string]bool{
	"TOLOWER":           true,
	"ASCII":             true,
	"CHAR_LENGTH":       true,
	"CHARACTER_LENGTH":  true,
	"CONCAT":            true,
	"CONCAT_WS":         true,
	"FIELD":             true,
	"FIND_IN_SET":       true,
	"FORMAT":            true,
	"INSERT":            true,
	"INSTR":             true,
	"LCASE":             true,
	"LEFT":              true,
	"LENGTH":            true,
	"LOCATE":            true,
	"LOWER":             true,
	"LPAD":              true,
	"LTRIM":             true,
	"MID":               true,
	"POSITION":          true,
	"REPEAT":            true,
	"REPLACE":           true,
	"REVERSE":           true,
	"RIGHT":             true,
	"RPAD":              true,
	"RTRIM":             true,
	"SPACE":             true,
	"STRCMP":            true,
	"SUBSTR":            true,
	"SUBSTRING":         true,
	"SUBSTRING_INDEX":   true,
	"TRIM":              true,
	"UCASE":             true,
	"UPPER":             true,
	"ABS":               true,
	"ACOS":              true,
	"ASIN":              true,
	"ATAN":              true,
	"ATAN2":             true,
	"AVG":               true,
	"CEIL":              true,
	"CEILING":           true,
	"COS":               true,
	"COT":               true,
	"COUNT":             true,
	"DEGREES":           true,
	"DIV":               true,
	"EXP":               true,
	"FLOOR":             true,
	"GREATEST":          true,
	"LEAST":             true,
	"LN":                true,
	"LOG":               true,
	"LOG10":             true,
	"LOG2":              true,
	"MAX":               true,
	"MIN":               true,
	"MOD":               true,
	"PI":                true,
	"POW":               true,
	"POWER":             true,
	"RADIANS":           true,
	"RAND":              true,
	"ROUND":             true,
	"SIGN":              true,
	"SIN":               true,
	"SQRT":              true,
	"SUM":               true,
	"TAN":               true,
	"TRUNCATE":          true,
	"ADDDATE":           true,
	"ADDTIME":           true,
	"CURDATE":           true,
	"CURRENT_DATE":      true,
	"CURRENT_TIME":      true,
	"CURRENT_TIMESTAMP": true,
	"CURTIME":           true,
	"DATE":              true,
	"DATEDIFF":          true,
	"DATE_ADD":          true,
	"DATE_FORMAT":       true,
	"DATE_SUB":          true,
	"DAY":               true,
	"DAYNAME":           true,
	"DAYOFMONTH":        true,
	"DAYOFWEEK":         true,
	"DAYOFYEAR":         true,
	"EXTRACT":           true,
	"FROM_DAYS":         true,
	"HOUR":              true,
	"LAST_DAY":          true,
	"LOCALTIME":         true,
	"LOCALTIMESTAMP":    true,
	"MAKEDATE":          true,
	"MAKETIME":          true,
	"MICROSECOND":       true,
	"MINUTE":            true,
	"MONTH":             true,
	"MONTHNAME":         true,
	"NOW":               true,
	"PERIOD_ADD":        true,
	"PERIOD_DIFF":       true,
	"QUARTER":           true,
	"SECOND":            true,
	"SEC_TO_TIME":       true,
	"STR_TO_DATE":       true,
	"SUBDATE":           true,
	"SUBTIME":           true,
	"SYSDATE":           true,
	"TIME":              true,
	"TIME_FORMAT":       true,
	"TIME_TO_SEC":       true,
	"TIMEDIFF":          true,
	"TIMESTAMP":         true,
	"TO_DAYS":           true,
	"WEEK":              true,
	"WEEKDAY":           true,
	"WEEKOFYEAR":        true,
	"YEAR":              true,
	"YEARWEEK":          true,
	"BIN":               true,
	"BINARY":            true,
	"CAST":              true,
	"COALESCE":          true,
	"CONV":              true,
	"CONVERT":           true,
	"IF":                true,
	"IFNULL":            true,
	"ISNULL":            true,
	"NULLIF":            true,
}