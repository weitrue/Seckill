/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2022/12/21 10:22 PM
 * Description:
 **/

package meta

// Field key-value
type Field interface {
	Key() string
	Value() interface{}
	meta()
}

type field struct {
	key   string
	value interface{}
}

func (m *field) Key() string {
	return m.key
}

func (m *field) Value() interface{} {
	return m.value
}

func (m *field) meta() {}

// NewField create meat
func NewField(key string, value any) Field {
	return &field{key: key, value: value}
}
