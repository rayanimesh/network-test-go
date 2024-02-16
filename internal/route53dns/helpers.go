/*
 * DNS Service
 *
 * Sevice for managing DNS with route53
 *
 * API version: 0.0.1
 * Contact: iitr.animesh@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package route53dns

import (
	"reflect"
)

// Response return a ImplResponse struct filled
func Response(code int, body interface{}) ImplResponse {
	return ImplResponse {
		Code: code,
		Body: body,
	}
}

// IsZeroValue checks if the val is the zero-ed value.
func IsZeroValue(val interface{}) bool {
	return val == nil || reflect.DeepEqual(val, reflect.Zero(reflect.TypeOf(val)).Interface())
}

// AssertRecurseInterfaceRequired recursively checks each struct in a slice against the callback.
// This method traverse nested slices in a preorder fashion.
func AssertRecurseInterfaceRequired[T any](obj interface{}, callback func(T) error) error {
	return AssertRecurseValueRequired(reflect.ValueOf(obj), callback)
}

// AssertRecurseValueRequired checks each struct in the nested slice against the callback.
// This method traverse nested slices in a preorder fashion. ErrTypeAssertionError is thrown if
// the underlying struct does not match type T.
func AssertRecurseValueRequired[T any](value reflect.Value, callback func(T) error) error {
	switch value.Kind() {
	// If it is a struct we check using callback
	case reflect.Struct:
		obj, ok := value.Interface().(T)
		if !ok {
			return ErrTypeAssertionError
		}

		if err := callback(obj); err != nil {
			return err
		}

	// If it is a slice we continue recursion
	case reflect.Slice:
		for i := 0; i < value.Len(); i += 1 {
			if err := AssertRecurseValueRequired(value.Index(i), callback); err != nil {
				return err
			}
		}
	}
	return nil
}
