package schema

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"sort"
	"strconv"
	"time"
)

//false则表示有新数据
func DataInclude(data interface{}, entry Entry) bool {
	if entry.ValueType != Map {
		return true
	}
	mval, ok := data.(map[string]interface{})
	if !ok {
		return true
	}
	if len(mval) > len(entry.Schema) {
		return false
	}
	for k, v := range mval {
		notfind := true
		for _, sv := range entry.Schema {
			if sv.Key == k {
				notfind = false
				if sv.ValueType == Map && !DataInclude(v, sv) {
					return false
				}
			}
		}
		if notfind {
			return false
		}
	}
	return true
}

func DataConvert(data interface{}, schema Entry) (converted interface{}, err error) {
	return dataConvert(data, schema)
}

func dataConvert(data interface{}, schema Entry) (converted interface{}, err error) {
	switch schema.ValueType {
	case Long:
		value := reflect.ValueOf(data)
		switch value.Kind() {
		case reflect.Int64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return data, nil
		case reflect.Float32, reflect.Float64:
			return int64(value.Float()), nil
		case reflect.String:
			if converted, err = strconv.ParseInt(value.String(), 10, 64); err == nil {
				return
			}
			var floatc float64
			if floatc, err = strconv.ParseFloat(value.String(), 10); err == nil {
				converted = int64(floatc)
				return
			}
		}
	case Float:
		value := reflect.ValueOf(data)
		switch value.Kind() {
		case reflect.Int64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return data, nil
		case reflect.Float32, reflect.Float64:
			return data, nil
		case reflect.String:
			if converted, err = strconv.ParseFloat(value.String(), 10); err == nil {
				return
			}
		}
	case String:
		value := reflect.ValueOf(data)
		switch value.Kind() {
		case reflect.Int64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
			return strconv.FormatInt(value.Int(), 10), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return strconv.FormatUint(value.Uint(), 10), nil
		case reflect.Float32, reflect.Float64:
			return strconv.FormatFloat(value.Float(), 'f', -1, 64), nil
		default:
			return data, nil
		}
	case Date:
		return data, nil
	case Bool:
		return data, nil
	case Array:
		ret := make([]interface{}, 0)
		switch value := data.(type) {
		case []interface{}:
			for _, j := range value {
				vi, err := dataConvert(j, Entry{ValueType: schema.ElemType})
				if err != nil {
					log.Println(err)
					continue
				}
				ret = append(ret, vi)
			}
		case []string:
			for _, j := range value {
				vi, err := dataConvert(j, Entry{ValueType: schema.ElemType})
				if err != nil {
					log.Println(err)
					continue
				}
				ret = append(ret, vi)
			}
		case []int:
			for _, j := range value {
				vi, err := dataConvert(j, Entry{ValueType: schema.ElemType})
				if err != nil {
					log.Println(err)
					continue
				}
				ret = append(ret, vi)
			}
		case []int64:
			for _, j := range value {
				vi, err := dataConvert(j, Entry{ValueType: schema.ElemType})
				if err != nil {
					log.Println(err)
					continue
				}
				ret = append(ret, vi)
			}
		case []json.Number:
			for _, j := range value {
				vi, err := dataConvert(j, Entry{ValueType: schema.ElemType})
				if err != nil {
					log.Println(err)
					continue
				}
				ret = append(ret, vi)
			}
		case []float64:
			for _, j := range value {
				vi, err := dataConvert(j, Entry{ValueType: schema.ElemType})
				if err != nil {
					log.Println(err)
					continue
				}
				ret = append(ret, vi)
			}
		case []bool:
			for _, j := range value {
				vi, err := dataConvert(j, Entry{ValueType: schema.ElemType})
				if err != nil {
					log.Println(err)
					continue
				}
				ret = append(ret, vi)
			}
		case []float32:
			for _, j := range value {
				vi, err := dataConvert(j, Entry{ValueType: schema.ElemType})
				if err != nil {
					log.Println(err)
					continue
				}
				ret = append(ret, vi)
			}
		case []int8:
			for _, j := range value {
				vi, err := dataConvert(j, Entry{ValueType: schema.ElemType})
				if err != nil {
					log.Println(err)
					continue
				}
				ret = append(ret, vi)
			}
		case []int16:
			for _, j := range value {
				vi, err := dataConvert(j, Entry{ValueType: schema.ElemType})
				if err != nil {
					log.Println(err)
					continue
				}
				ret = append(ret, vi)
			}
		case []int32:
			for _, j := range value {
				vi, err := dataConvert(j, Entry{ValueType: schema.ElemType})
				if err != nil {
					log.Println(err)
					continue
				}
				ret = append(ret, vi)
			}
		case []uint:
			for _, j := range value {
				vi, err := dataConvert(j, Entry{ValueType: schema.ElemType})
				if err != nil {
					log.Println(err)
					continue
				}
				ret = append(ret, vi)
			}
		case []uint8:
			for _, j := range value {
				vi, err := dataConvert(j, Entry{ValueType: schema.ElemType})
				if err != nil {
					log.Println(err)
					continue
				}
				ret = append(ret, vi)
			}
		case []uint16:
			for _, j := range value {
				vi, err := dataConvert(j, Entry{ValueType: schema.ElemType})
				if err != nil {
					log.Println(err)
					continue
				}
				ret = append(ret, vi)
			}
		case []uint32:
			for _, j := range value {
				vi, err := dataConvert(j, Entry{ValueType: schema.ElemType})
				if err != nil {
					log.Println(err)
					continue
				}
				ret = append(ret, vi)
			}
		case []uint64:
			for _, j := range value {
				vi, err := dataConvert(j, Entry{ValueType: schema.ElemType})
				if err != nil {
					log.Println(err)
					continue
				}
				ret = append(ret, vi)
			}
		case string:
			newdata := make([]interface{}, 0)
			err = json.Unmarshal([]byte(value), &newdata)
			if err != nil {
				return
			}
			for _, j := range newdata {
				vi, err := dataConvert(j, Entry{ValueType: schema.ElemType})
				if err != nil {
					log.Println(err)
					continue
				}
				ret = append(ret, vi)
			}
		}
		return ret, nil
	case Map:
		switch value := data.(type) {
		case map[string]interface{}:
			return mapDataConvert(value, schema.Schema), nil
		case string:
			newdata := make(map[string]interface{})
			err = json.Unmarshal([]byte(value), &newdata)
			if err == nil {
				return mapDataConvert(newdata, schema.Schema), nil
			}
		}
	}
	return data, fmt.Errorf("can not convert data[%v] to type(%v), err %v", data, reflect.TypeOf(data), err)
}

func mapDataConvert(mpvalue map[string]interface{}, schemas Schemas) (converted interface{}) {
	for _, v := range schemas {
		if subv, ok := mpvalue[v.Key]; ok {
			subconverted, err := dataConvert(subv, v)
			if err != nil {
				log.Println(err)
				continue
			}
			mpvalue[v.Key] = subconverted
		}
	}
	return mpvalue
}

func IsInvalid(value interface{}, schemeType string) bool {
	if value == nil {
		return true
	}
	rv := reflect.ValueOf(value)
	if (rv.Kind() == reflect.Map || rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface || rv.Kind() == reflect.Slice) && rv.IsNil() {
		return true
	}
	str, ok := value.(string)
	if !ok || str != "" {
		return false
	}
	switch schemeType {
	case Array, Map, Long, Float:
		if str == "" {
			return true
		}
	}
	return false
}

func MergeSchemas(sa, sb Schemas) (ret Schemas, err error) {
	ret = make(Schemas, 0)
	if sa == nil && sb == nil {
		return
	}
	if sa == nil {
		ret = sb
		return
	}
	if sb == nil {
		ret = sa
		return
	}
	sort.Sort(sa)
	sort.Sort(sb)
	i, j := 0, 0
	for {
		if i >= len(sa) {
			break
		}
		if j >= len(sb) {
			break
		}
		if sa[i].Key < sb[j].Key {
			ret = append(ret, sa[i])
			i++
			continue
		}
		if sa[i].Key > sb[j].Key {
			ret = append(ret, sb[j])
			j++
			continue
		}
		if sa[i].ValueType != sb[j].ValueType {
			err = fmt.Errorf("type conflict: key %v old type is <%v> want change to type <%v>", sa[i].Key, sa[i].ValueType, sb[j].ValueType)
			return
		}
		if sa[i].ValueType == Map {
			if sa[i].Schema, err = MergeSchemas(sa[i].Schema, sb[j].Schema); err != nil {
				return
			}
		}
		ret = append(ret, sa[i])
		i++
		j++
	}
	for ; i < len(sa); i++ {
		ret = append(ret, sa[i])
	}
	for ; j < len(sb); j++ {
		ret = append(ret, sb[j])
	}
	return
}

func AssertSchema(data Data) (valueType map[string]Entry) {
	valueType = make(map[string]Entry)
	for k, v := range data {
		switch nv := v.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			valueType[k] = formValueType(k, Long)
		case float32, float64:
			valueType[k] = formValueType(k, Float)
		case bool:
			valueType[k] = formValueType(k, Bool)
		case json.Number:
			_, err := nv.Int64()
			if err == nil {
				valueType[k] = formValueType(k, Long)
			} else {
				valueType[k] = formValueType(k, Float)
			}
		case map[string]interface{}:
			sc := formValueType(k, Map)
			follows := AssertSchema(Data(nv))
			for _, m := range follows {
				sc.Schema = append(sc.Schema, m)
			}
			valueType[k] = sc
		case []interface{}:
			sc := formValueType(k, Array)
			if len(nv) > 0 {
				switch nnv := nv[0].(type) {
				case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
					sc.ElemType = Long
				case float32, float64:
					sc.ElemType = Float
				case bool:
					sc.ElemType = Bool
				case json.Number:
					_, err := nnv.Int64()
					if err == nil {
						sc.ElemType = Long
					} else {
						sc.ElemType = Float
					}
				case nil: // 不处理，不加入
				case string:
					sc.ElemType = String
				default:
					sc.ElemType = String
				}
				valueType[k] = sc
			}
			//对于里面没有元素的interface，不添加进去，因为无法判断类型
		case []int, []int8, []int16, []int32, []int64, []uint, []uint8, []uint16, []uint32, []uint64:
			sc := formValueType(k, Array)
			sc.ElemType = Long
			valueType[k] = sc
		case []float32, []float64:
			sc := formValueType(k, Array)
			sc.ElemType = Float
			valueType[k] = sc
		case []bool:
			sc := formValueType(k, Array)
			sc.ElemType = Bool
			valueType[k] = sc
		case []string:
			sc := formValueType(k, Array)
			sc.ElemType = Bool
			valueType[k] = sc
		case []json.Number:
			sc := formValueType(k, Array)
			sc.ElemType = Float
			valueType[k] = sc
		case nil: // 不处理，不加入
		case string:
			_, err := time.Parse(time.RFC3339, nv)
			if err == nil {
				valueType[k] = formValueType(k, Date)
			} else {
				valueType[k] = formValueType(k, String)
			}
		case time.Time, *time.Time:
			valueType[k] = formValueType(k, Date)
		default:
			valueType[k] = formValueType(k, String)
			log.Printf("find undetected key(%v)-type(%v), read it as string\n", k, reflect.TypeOf(v))
		}
	}
	return
}

func formValueType(key, vtype string) Entry {
	return Entry{
		Key:       key,
		ValueType: vtype,
	}
}

func GetDefault(t Entry) (result interface{}) {
	switch t.ValueType {
	case Long:
		result = 0
	case Float:
		result = 0.0
	case String:
		result = ""
	case Date:
		result = time.Now().Format(time.RFC3339Nano)
	case Bool:
		result = false
	case Map:
		result = make(map[string]interface{})
	case Array:
		switch t.ElemType {
		case String:
			result = make([]string, 0)
		case Float:
			result = make([]float64, 0)
		case Long:
			result = make([]int64, 0)
		case Bool:
			result = make([]bool, 0)
		}
	}
	return
}
