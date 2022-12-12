package wizard

import "github.com/temphia/temphia/code/executors/backend/wizard/wmodels"

const ErrStringFieldEmpty = "Field cannot be empty"
const ErrStringWrongDataType = "Wrong data type"

func (sw *SimpleWizard) validateField(field *wmodels.Field, data interface{}) string {

	if data == nil {
		if field.Optional {
			return ""
		} else {
			return ErrStringFieldEmpty
		}
	}

	switch field.Type {
	case wmodels.BASIC_SHORTTEXT, wmodels.BASIC_LONGTEXT, wmodels.BASIC_SELECT:

		dstr, ok := data.(string)
		if !ok {
			return ErrStringWrongDataType
		}

		if len(dstr) == 0 {
			return ErrStringFieldEmpty
		}
	case wmodels.BASIC_EMAIL:

	case wmodels.BASIC_RANGE:
	case wmodels.BASIC_MULTI_SELECT:
	case wmodels.BASIC_PHONE:
	case wmodels.BASIC_CHECKBOX:
	case wmodels.BASIC_COLOR:
	case wmodels.BASIC_DATE:
	case wmodels.BASIC_DATETIME:

	case wmodels.BASIC_NUMBER:

	}

	return ""

}
