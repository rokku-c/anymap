package anymap

type AnyMap struct {
	o interface{}
}

func New(o interface{}) *AnyMap {
	am := &AnyMap{}
	am.o = o
	return am
}

func (am *AnyMap) ObjectOf(key string) *AnyMap {
	if am == nil || am.o == nil {
		return am
	}
	o := am.o.(map[string]interface{})[key]
	return New(o)
}

func (am *AnyMap) Value() interface{} {
	if am == nil || am.o == nil {
		return nil
	}
	return am.o
}

func (am *AnyMap) Bool(default_value_if_null bool) bool {
	if am == nil || am.o == nil {
		return default_value_if_null
	}
	return am.o.(bool)
}

func (am *AnyMap) Int(default_value_if_null int) int {
	if am == nil || am.o == nil {
		return default_value_if_null
	}
	return am.o.(int)
}

func (am *AnyMap) String(default_value_if_null string) string {
	if am == nil || am.o == nil {
		return default_value_if_null
	}
	return am.o.(string)
}

func (am *AnyMap) Keys() []string {
	if am == nil || am.o == nil {
		return []string{}
	}
	var keys []string
	for k := range am.o.(map[string]interface{}) {
		keys = append(keys, k)
	}
	return keys
}
