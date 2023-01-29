package model

import validation "github.com/go-ozzo/ozzo-validation"

// если true возвращаем validation.Required(...), если false, то nil
func requiredIf(cond bool) validation.RuleFunc {

	return func(value interface{}) error {
		if cond {
			return validation.Validate(value, validation.Required)
		} //проверка value на не пустое
		return nil
	}
}
