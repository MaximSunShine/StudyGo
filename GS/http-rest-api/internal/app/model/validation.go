package model

import validation "github.com/go-ozzo/ozzo-validation"

// если true возвращаем validation.Required(...), если false, то nil
func requiredIf(cond bool) validation.RuleFunc {

	return func(value interface{}) error {
		if cond {
			//проверка value на не пустое
			return validation.Validate(value, validation.Required)
		}
		return nil
	}
}
