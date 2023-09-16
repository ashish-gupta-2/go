package models

// Job represents fields used in Job model.
type Person struct {
	Base
	Name string `json:"name,omitempty"`
	Age  string `json:"age,omitempty"`
	//Output Output `json:"output,omitempty" gorm:"type:bytea"`
}

// // Value implements driver.Valuer interface and returns json value.
// func (e Output) Value() (driver.Value, error) {
// 	if e == nil {
// 		return nil, nil
// 	}
// 	return json.Marshal(e)
// }

// // Scan implements sql.Scanner interface and scan value into json.
// func (e *Output) Scan(value interface{}) error {
// 	bytes, ok := value.([]byte)
// 	if !ok {
// 		return errors.New("field must be byte array, check the table schema")
// 	}
// 	return json.Unmarshal(bytes, &e)
// }

// // Value implements driver.Valuer interface and returns json value.
// func (e Input) Value() (driver.Value, error) {
// 	return json.Marshal(e)
// }

// // Scan implements sql.Scanner interface and scan value into json.
// func (e *Input) Scan(value interface{}) error {
// 	bytes, ok := value.([]byte)
// 	if !ok {
// 		return errors.New("field must be byte array, check the table schema")
// 	}
// 	return json.Unmarshal(bytes, &e)
// }

// // IsEmpty method checks for empty struct.
// func (e Input) IsEmpty() bool {
// 	return reflect.DeepEqual(e, Input{})
// }
