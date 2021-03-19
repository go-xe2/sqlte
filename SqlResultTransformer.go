package sqlte

import "reflect"

type SqlTransformer func(results []reflect.Value, next SqlTransformer) (interface{}, error)
