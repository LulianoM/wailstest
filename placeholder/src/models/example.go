package models

type Example struct {
	Metadata
	SomeStringField string  `json:"some_string_field" validate:"required"`
	SomeIntField    int     `json:"some_int_field"`
	SomeFloatField  float32 `json:"some_float_field" validate:"required"`
}
