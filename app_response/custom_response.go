package custom_response

import (
	"github.com/1827mk/app-commons/convutil"
	"github.com/1827mk/app-commons/stringutil"
)

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type Error struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Fields  map[string]any `json:"fields,omitempty"`
}

func (err *Error) String() string {
	return stringutil.Json(*err)
}

func (err *Error) Error() string {
	return stringutil.Json(*err)
}

func (err *Error) ToMap() map[string]any {
	return convutil.Obj2Map(*err)
}

type Errors []Error

func (errs *Errors) String() string {
	return stringutil.Json(*errs)
}

func (errs *Errors) ToMap() map[string]any {
	errors := make(map[string]any)
	errors["errors"] = errs
	return convutil.Obj2Map(errors)
}

func (errs *Errors) Error() string {
	return stringutil.Json(*errs)
}
