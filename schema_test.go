package schema

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAssertSchema(t *testing.T) {
	var data map[string]interface{}
	dc := json.NewDecoder(strings.NewReader("{\"a\":123,\"b\":123.1,\"c\":\"123\",\"d\":true,\"e\":[1,2,3],\"f\":[1.2,2.1,3.1],\"g\":{\"g1\":\"1\"}}"))
	dc.UseNumber()
	err := dc.Decode(&data)
	emp := formValueType("e", Array)
	emp.ElemType = Long
	fmp := formValueType("f", Array)
	fmp.ElemType = Float
	gmp := formValueType("g", Map)
	gmp.Schema = Schemas{
		{
			Key:       "g1",
			ValueType: String,
		},
	}

	exp := map[string]Entry{
		"a": formValueType("a", Long),
		"b": formValueType("b", Float),
		"c": formValueType("c", String),
		"d": formValueType("d", Bool),
		"e": emp,
		"f": fmp,
		"g": gmp,
	}
	assert.NoError(t, err)
	vt := AssertSchema(data)
	assert.Equal(t, exp, vt)
	data = map[string]interface{}{
		"a": 1,
		"b": time.Now().Format(time.RFC3339Nano),
		"c": time.Now().Format(time.RFC3339),
		"d": 1.0,
		"e": int64(32),
		"f": "123",
		"g": true,
		"m": nil,
		"h": map[string]interface{}{
			"h5": map[string]interface{}{
				"h51": 1,
			},
		},
		"h1": map[string]interface{}{
			"h1": 123,
		},
		"h2": map[string]interface{}{
			"h2": "123",
		},
		"h3": map[string]interface{}{
			"h3": 123.1,
		},
		"h4": map[string]interface{}{
			"h4": map[string]interface{}{},
		},
		"i": false,
	}
	hmp := formValueType("h", Map)
	hmp.Schema = Schemas{
		{
			Key:       "h5",
			ValueType: Map,
			Schema: Schemas{
				{
					Key:       "h51",
					ValueType: Long,
				},
			},
		},
	}
	hmp1 := formValueType("h1", Map)
	hmp1.Schema = Schemas{
		{
			Key:       "h1",
			ValueType: Long,
		},
	}
	hmp2 := formValueType("h2", Map)
	hmp2.Schema = Schemas{
		Entry{
			Key:       "h2",
			ValueType: String,
		},
	}
	hmp3 := formValueType("h3", Map)
	hmp3.Schema = Schemas{
		Entry{
			Key:       "h3",
			ValueType: Float,
		},
	}
	hmp4 := formValueType("h4", Map)
	hmp4.Schema = Schemas{
		Entry{
			Key:       "h4",
			ValueType: Map,
		},
	}

	exp = map[string]Entry{
		"a":  formValueType("a", Long),
		"b":  formValueType("b", Date),
		"c":  formValueType("c", Date),
		"d":  formValueType("d", Float),
		"e":  formValueType("e", Long),
		"f":  formValueType("f", String),
		"g":  formValueType("g", Bool),
		"h":  hmp,
		"h1": hmp1,
		"h2": hmp2,
		"h3": hmp3,
		"h4": hmp4,

		"i": formValueType("i", Bool),
	}
	vt = AssertSchema(data)
	assert.EqualValues(t, exp, vt)
}

func TestDataInclude(t *testing.T) {
	tests := []struct {
		value  interface{}
		left   interface{}
		schema Entry
		exp    bool
	}{
		{
			value: map[string]interface{}{},
			left:  map[string]interface{}{},
			schema: Entry{
				Key:       "hello",
				ValueType: Map,
			},
			exp: true,
		},
		{
			value: 123,
			left:  123,
			schema: Entry{
				Key:       "hello",
				ValueType: Map,
			},
			exp: true,
		},
		{
			value: map[string]interface{}{},
			left:  map[string]interface{}{},
			schema: Entry{
				Key:       "hello",
				ValueType: Long,
			},
			exp: true,
		},
		{
			value: map[string]interface{}{
				"x": 123,
			},
			left: map[string]interface{}{
				"x": 123,
			},
			schema: Entry{
				Key:       "hello",
				ValueType: Map,
				Schema:    Schemas{},
			},
			exp: false,
		},
		{
			value: map[string]interface{}{
				"x": 123,
			},
			left: map[string]interface{}{
				"x": 123,
			},
			schema: Entry{
				Key:       "hello",
				ValueType: Map,
				Schema: Schemas{
					Entry{
						Key:       "x",
						ValueType: Long,
					},
				},
			},
			exp: true,
		},
		{
			value: map[string]interface{}{
				"x": 123,
			},
			left: map[string]interface{}{
				"x": 123,
			},
			schema: Entry{
				Key:       "hello",
				ValueType: Map,
				Schema: Schemas{
					Entry{
						Key:       "x",
						ValueType: Map,
					},
				},
			},
			exp: true,
		},
		{
			value: map[string]interface{}{
				"x": map[string]interface{}{},
			},
			left: map[string]interface{}{
				"x": map[string]interface{}{},
			},
			schema: Entry{
				Key:       "hello",
				ValueType: Map,
				Schema: Schemas{
					Entry{
						Key:       "x",
						ValueType: Long,
					},
				},
			},
			exp: true,
		},
		{
			value: map[string]interface{}{
				"x": map[string]interface{}{
					"y": 123,
				},
			},
			left: map[string]interface{}{
				"x": map[string]interface{}{
					"y": 123,
				},
			},
			schema: Entry{
				Key:       "hello",
				ValueType: Map,
				Schema: Schemas{
					Entry{
						Key:       "x",
						ValueType: Map,
					},
				},
			},
			exp: false,
		},
		{
			value: map[string]interface{}{
				"x": map[string]interface{}{
					"y": 123,
				},
			},
			left: map[string]interface{}{
				"x": map[string]interface{}{
					"y": 123,
				},
			},
			schema: Entry{
				Key:       "hello",
				ValueType: Map,
				Schema: Schemas{
					Entry{
						Key:       "x",
						ValueType: Map,
						Schema: Schemas{
							Entry{
								Key: "y",
							},
						},
					},
				},
			},
			exp: true,
		},
		{
			value: map[string]interface{}{
				"x": map[string]interface{}{
					"y": 123,
				},
				"z": 123,
			},
			left: map[string]interface{}{
				"x": map[string]interface{}{
					"y": 123,
				},
				"z": 123,
			},
			schema: Entry{
				Key:       "hello",
				ValueType: Map,
				Schema: Schemas{
					Entry{
						Key:       "x",
						ValueType: Map,
						Schema: Schemas{
							Entry{
								Key: "y",
							},
						},
					},
				},
			},
			exp: false,
		},
		{
			value: map[string]interface{}{
				"x": map[string]interface{}{
					"y": 123,
					"z": 123,
				},
			},
			left: map[string]interface{}{
				"x": map[string]interface{}{
					"y": 123,
					"z": 123,
				},
			},
			schema: Entry{
				Key:       "hello",
				ValueType: Map,
				Schema: Schemas{
					Entry{
						Key:       "x",
						ValueType: Map,
						Schema: Schemas{
							Entry{
								Key: "y",
							},
							Entry{
								Key: "z",
							},
						},
					},
				},
			},
			exp: true,
		},
		{
			value: map[string]interface{}{
				"x": map[string]interface{}{
					"y": 123,
					"z": 123,
					"a": true,
				},
			},
			left: map[string]interface{}{
				"x": map[string]interface{}{
					"y": 123,
					"z": 123,
					"a": true,
				},
			},
			schema: Entry{
				Key:       "hello",
				ValueType: Map,
				Schema: Schemas{
					Entry{
						Key:       "x",
						ValueType: Map,
						Schema: Schemas{
							Entry{
								Key: "y",
							},
							Entry{
								Key: "z",
							},
						},
					},
				},
			},
			exp: false,
		},
	}
	for _, ti := range tests {
		got := DataInclude(ti.value, ti.schema)
		assert.Equal(t, ti.exp, got)
		assert.Equal(t, ti.left, ti.value)
	}
}

func TestMergePandoraSchemas(t *testing.T) {
	tests := []struct {
		oldScs Schemas
		newScs Schemas
		exp    Schemas
		err    bool
	}{
		{
			oldScs: Schemas{},
			newScs: Schemas{},
			exp:    Schemas{},
		},
		{
			oldScs: Schemas{},
			newScs: Schemas{
				Entry{Key: "abc"},
			},
			exp: Schemas{Entry{Key: "abc"}},
		},
		{
			oldScs: Schemas{
				Entry{Key: "abc"},
			},
			newScs: Schemas{
				Entry{Key: "abc"},
			},
			exp: Schemas{Entry{Key: "abc"}},
		},
		{
			oldScs: Schemas{
				Entry{Key: "abc", ValueType: "string"},
			},
			newScs: Schemas{
				Entry{Key: "abc", ValueType: "float"},
			},
			exp: Schemas{Entry{Key: "abc"}},
			err: true,
		},
		{
			oldScs: Schemas{
				Entry{Key: "a"},
			},
			newScs: Schemas{
				Entry{Key: "b"},
			},
			exp: Schemas{Entry{Key: "a"}, Entry{Key: "b"}},
		},
		{
			oldScs: Schemas{
				Entry{Key: "a"},
			},
			newScs: Schemas{
				Entry{Key: "a"},
				Entry{Key: "b"},
			},
			exp: Schemas{Entry{Key: "a"}, Entry{Key: "b"}},
		},
		{
			oldScs: Schemas{
				Entry{Key: "b"},
				Entry{Key: "c"},
			},
			newScs: Schemas{
				Entry{Key: "a"},
				Entry{Key: "b"},
			},
			exp: Schemas{Entry{Key: "a"}, Entry{Key: "b"}, Entry{Key: "c"}},
		},
		{
			oldScs: Schemas{
				Entry{Key: "b", ValueType: Map, Schema: Schemas{
					Entry{Key: "a"},
				}},
				Entry{Key: "c"},
			},
			newScs: Schemas{
				Entry{Key: "a"},
				Entry{Key: "b", ValueType: Map, Schema: Schemas{
					Entry{Key: "b"},
				}},
			},
			exp: Schemas{Entry{Key: "a"}, Entry{Key: "b", ValueType: Map, Schema: Schemas{
				Entry{Key: "a"},
				Entry{Key: "b"},
			}}, Entry{Key: "c"}},
		},
		{
			oldScs: Schemas{
				Entry{Key: "b", ValueType: Map, Schema: Schemas{
					Entry{Key: "a"},
				}},
				Entry{Key: "c"},
			},
			newScs: Schemas{
				Entry{Key: "a"},
				Entry{Key: "b", ValueType: Map, Schema: Schemas{
					Entry{Key: "a"},
				}},
			},
			exp: Schemas{Entry{Key: "a"}, Entry{Key: "b", ValueType: Map, Schema: Schemas{
				Entry{Key: "a"},
			}}, Entry{Key: "c"}},
		},
		{
			oldScs: Schemas{
				Entry{Key: "b", ValueType: Map, Schema: Schemas{
					Entry{Key: "y"},
				}},
				Entry{Key: "c"},
			},
			newScs: Schemas{
				Entry{Key: "a"},
				Entry{Key: "b", ValueType: Map, Schema: Schemas{
					Entry{Key: "x"},
				}},
			},
			exp: Schemas{Entry{Key: "a"}, Entry{Key: "b", ValueType: Map, Schema: Schemas{
				Entry{Key: "x"},
				Entry{Key: "y"},
			}}, Entry{Key: "c"}},
		},
		{
			oldScs: Schemas{
				Entry{Key: "b", ValueType: Map, Schema: Schemas{
					Entry{Key: "y", ValueType: Map},
				}},
				Entry{Key: "c"},
			},
			newScs: Schemas{
				Entry{Key: "a"},
				Entry{Key: "b", ValueType: Map, Schema: Schemas{
					Entry{Key: "y", ValueType: String},
				}},
			},
			exp: Schemas{Entry{Key: "a"}, Entry{Key: "b", ValueType: Map, Schema: Schemas{
				Entry{Key: "x"},
				Entry{Key: "y"},
			}}, Entry{Key: "c"}},
			err: true,
		},
		{
			oldScs: Schemas{
				Entry{Key: "b", ValueType: Map, Schema: Schemas{
					Entry{Key: "y", ValueType: Map, Schema: Schemas{
						Entry{Key: "11"},
					}},
				}},
				Entry{Key: "c"},
			},
			newScs: Schemas{
				Entry{Key: "a"},
				Entry{Key: "b", ValueType: Map, Schema: Schemas{
					Entry{Key: "y", ValueType: Map, Schema: Schemas{
						Entry{Key: "11"},
					}},
				}},
			},
			exp: Schemas{Entry{Key: "a"}, Entry{Key: "b", ValueType: Map, Schema: Schemas{
				Entry{Key: "y", ValueType: Map, Schema: Schemas{
					Entry{Key: "11"},
				}},
			}}, Entry{Key: "c"}},
		},
		{
			oldScs: Schemas{
				Entry{Key: "b", ValueType: Map, Schema: Schemas{
					Entry{Key: "y", ValueType: Map, Schema: Schemas{
						Entry{Key: "11"},
					}},
				}},
				Entry{Key: "c"},
			},
			newScs: Schemas{
				Entry{Key: "a"},
				Entry{Key: "b", ValueType: Map, Schema: Schemas{
					Entry{Key: "y", ValueType: Map, Schema: Schemas{
						Entry{Key: "21"},
						Entry{Key: "11"},
					}},
				}},
			},
			exp: Schemas{Entry{Key: "a"}, Entry{Key: "b", ValueType: Map, Schema: Schemas{
				Entry{Key: "y", ValueType: Map, Schema: Schemas{
					Entry{Key: "11"},
					Entry{Key: "21"},
				}},
			}}, Entry{Key: "c"}},
		},
	}
	for idx, ti := range tests {
		got, err := MergeSchemas(ti.oldScs, ti.newScs)
		if ti.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, ti.exp, got, fmt.Sprintf("index %v", idx))
		}
	}
}

func TestCheckIgnore(t *testing.T) {
	tests := []struct {
		v   interface{}
		tp  string
		exp bool
	}{
		{
			exp: true,
		},
		{
			v:   "",
			tp:  String,
			exp: false,
		},
		{
			v:   "xs",
			tp:  String,
			exp: false,
		},
		{
			v:   123,
			tp:  Float,
			exp: false,
		},
	}
	for _, ti := range tests {
		got := IsInvalid(ti.v, ti.tp)
		assert.Equal(t, ti.exp, got)
	}
}

func TestConvertData(t *testing.T) {
	type helloint int
	tests := []struct {
		v      interface{}
		schema Entry
		exp    interface{}
	}{
		{
			v: helloint(1),
			schema: Entry{
				ValueType: Long,
			},
			exp: helloint(1),
		},
		{
			v: helloint(1),
			schema: Entry{
				ValueType: String,
			},
			exp: "1",
		},
		{
			v: json.Number("1"),
			schema: Entry{
				ValueType: Long,
			},
			exp: int64(1),
		},
		{
			v: "1",
			schema: Entry{
				ValueType: Long,
			},
			exp: int64(1),
		},
		{
			v: []int{1, 2, 3},
			schema: Entry{
				ValueType: Array,
				ElemType:  Long,
			},
			exp: []interface{}{1, 2, 3},
		},
		{
			v: []int{1, 2, 3},
			schema: Entry{
				ValueType: Array,
				ElemType:  String,
			},
			exp: []interface{}{"1", "2", "3"},
		},
		{
			v: []interface{}{1, 2, 3},
			schema: Entry{
				ValueType: Array,
				ElemType:  String,
			},
			exp: []interface{}{"1", "2", "3"},
		},
		{
			v: `[1, 2, 3]`,
			schema: Entry{
				ValueType: Array,
				ElemType:  String,
			},
			exp: []interface{}{"1", "2", "3"},
		},
		{
			v: `["1", "2", "3"]`,
			schema: Entry{
				ValueType: Array,
				ElemType:  Float,
			},
			exp: []interface{}{float64(1), float64(2), float64(3)},
		},
		{
			v: "1.1",
			schema: Entry{
				ValueType: Float,
			},
			exp: float64(1.1),
		},
		{
			v: map[string]interface{}{
				"a": 123,
			},
			schema: Entry{
				ValueType: Map,
				Schema: Schemas{
					{ValueType: String, Key: "a"},
				},
			},
			exp: map[string]interface{}{
				"a": "123",
			},
		},
		{
			v: map[string]interface{}{
				"a": 123,
			},
			schema: Entry{
				ValueType: Map,
				Schema: Schemas{
					{ValueType: Float, Key: "a"},
				},
			},
			exp: map[string]interface{}{
				"a": 123,
			},
		},
		{
			v: map[string]interface{}{
				"a": "123",
				"b": "hello",
			},
			schema: Entry{
				ValueType: Map,
				Schema: Schemas{
					{ValueType: Float, Key: "a"},
					{ValueType: String, Key: "b"},
				},
			},
			exp: map[string]interface{}{
				"a": float64(123),
				"b": "hello",
			},
		},
		{
			v: `{
				"a": "123",
				"b": "hello"
			}`,
			schema: Entry{
				ValueType: Map,
				Schema: Schemas{
					{ValueType: Float, Key: "a"},
					{ValueType: String, Key: "b"},
				},
			},
			exp: map[string]interface{}{
				"a": float64(123),
				"b": "hello",
			},
		},
		{
			v: `{
				"a": "123.23",
				"b": "hello"
			}`,
			schema: Entry{
				ValueType: Map,
				Schema: Schemas{
					{ValueType: Long, Key: "a"},
					{ValueType: String, Key: "b"},
				},
			},
			exp: map[string]interface{}{
				"a": int64(123),
				"b": "hello",
			},
		},
		{
			v: `{
				"a": "123.23",
				"b": {
					"c":123
				}
			}`,
			schema: Entry{
				ValueType: Map,
				Schema: Schemas{
					{ValueType: Long, Key: "a"},
					{ValueType: Map, Key: "b", Schema: Schemas{
						{ValueType: Long, Key: "c"}},
					},
				},
			},
			exp: map[string]interface{}{
				"a": int64(123),
				"b": map[string]interface{}{
					"c": int64(123),
				},
			},
		},
	}
	for _, ti := range tests {
		got, err := dataConvert(ti.v, ti.schema)
		assert.NoError(t, err)
		assert.Equal(t, ti.exp, got)
	}
}
