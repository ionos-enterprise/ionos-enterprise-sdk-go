package profitbricks

import (
	"errors"
	"fmt"
	ionossdk "github.com/ionos-cloud/ionos-cloud-sdk-go/v5"
	"reflect"
	"time"
)

type ConversionTarget string
const (
	ConversionTargetCore ConversionTarget = "core"
	ConversionTargetCompat = "compat"
)

var primitiveTypes = []reflect.Kind{
	reflect.String, reflect.Bool, reflect.Chan, reflect.Complex64, reflect.Complex128,
	reflect.Float32, reflect.Float64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
	reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
}

/* compat field -> core field */
var CompatCoreFieldMap = map[string]string{
	"ID": "Id",
	"PBType": "Type",
	"RAM": "Ram",
	"VMState": "VmState",
	"CPUFamily": "CpuFamily",
	"CPUHotPlug": "CpuHotPlug",
	"CPUHotUnplug": "CpuHotUnplug",
	"RAMHotPlug": "RamHotPlug",
	"RAMHotUnplug": "RamHotUnplug",
	"IPs": "Ips",
	"IPConsumers": "IpConsumers",
	"FirewallRules": "Firewallrules",
	"ReserveIP": "ReserveIp",
	"URL": "Url",
	"NodePools": "Nodepools",
}

/* custom caster, such as from *Type to string or *time.Time to string */
type CustomCaster struct {
	From string
	To string

	/* performs the actual cast */
	Cast func(value reflect.Value) reflect.Value
}

const (
	dateLayout = "2006-01-02T15:04:05.000Z"
)

var CustomCasters = []CustomCaster{
	{
		From: "*Type",
		To: "string",
		Cast: func (value reflect.Value) reflect.Value{
			v := value.Interface().(*ionossdk.Type)
			if v == nil {
				return reflect.ValueOf("")
			}
			return reflect.ValueOf(string(*(v)))
		},
	},
	{
		From: "string",
		To: "*Type",
		Cast: func (value reflect.Value) reflect.Value {
			v := value.Interface().(string)
			if v == "" {
				/* we need to avoid sending the type otherwise the api returns 422 because
				 * type is read only */
				return reflect.Zero(reflect.TypeOf((*ionossdk.Type)(nil)))
			}
			ret := ionossdk.Type(value.Interface().(string))
			return reflect.ValueOf(&ret)
		},
	},
	{
		From: "string",
		To: "*Time",
		Cast: func (value reflect.Value) reflect.Value {
			str := value.Interface().(string)
			if str == "" {
				return reflect.Zero(reflect.TypeOf((*time.Time)(nil)))
			}
			createdDate, err := time.Parse(dateLayout, str)
			if err != nil {
				panic("Error parsing date " + str + "; expecting format " + dateLayout)
			}
			return reflect.ValueOf(&createdDate)
		},
	},
	{
		From: "*Time",
		To: "string",
		Cast: func (value reflect.Value) reflect.Value {
			t := value.Interface().(*time.Time)
			return reflect.ValueOf(t.Format(dateLayout))
		},
	},
	{
		From: "string",
		To: "*IonosTime",
		Cast: func (value reflect.Value) reflect.Value {
			str := value.Interface().(string)
			if str == "" {
				return reflect.Zero(reflect.TypeOf((*ionossdk.IonosTime)(nil)))
			}
			createdDate, err := time.Parse(dateLayout, str)
			if err != nil {
				panic("Error parsing date " + str + "; expecting format " + dateLayout)
			}
			return reflect.ValueOf(&ionossdk.IonosTime{Time: createdDate})
		},
	},
	{
		From: "*IonosTime",
		To: "string",
		Cast: func (value reflect.Value) reflect.Value {
			t := value.Interface().(*ionossdk.IonosTime)
			return reflect.ValueOf(t.Format(dateLayout))
		},
	},
	{
		From: "*int32",
		To: "int",
		Cast: func (value reflect.Value) reflect.Value {
			v := value.Interface().(*int32)
			if v == nil {
				return reflect.ValueOf(0)
			}
			return reflect.ValueOf(int(*v))
		},
	},
	{
		From: "int",
		To: "*int32",
		Cast: func (value reflect.Value) reflect.Value {
			v := value.Interface().(int)
			if v == 0 {
				return reflect.Zero(reflect.TypeOf((*int32)(nil)))
			}
			vInt32 := int32(v)
			return reflect.ValueOf(&vInt32)
		},
	},
	{
		From: "*float32",
		To: "int",
		Cast: func (value reflect.Value) reflect.Value {
			v := value.Interface().(*float32)
			if v == nil {
				return reflect.ValueOf(0)
			}
			return reflect.ValueOf(int(*v))
		},
	},
	{
		From: "int",
		To: "*float32",
		Cast: func (value reflect.Value) reflect.Value {
			v := value.Interface().(int)
			if v == 0 {
				return reflect.Zero(reflect.TypeOf((*float32)(nil)))
			}
			vFloat32 := float32(v)
			return reflect.ValueOf(&vFloat32)
		},
	},
	{
		From: "string",
		To: "*string",
		Cast: func(value reflect.Value) reflect.Value {
			v := value.Interface().(string)
			if v == "" {
				return reflect.Zero(reflect.TypeOf((*string)(nil)))
			}
			newV := v
			return reflect.ValueOf(&newV)
		},
	},
	{
		From: "int32",
		To: "*int32",
		Cast: func(value reflect.Value) reflect.Value {
			v := value.Interface().(int32)
			if v == 0 {
				return reflect.Zero(reflect.TypeOf((*int32)(nil)))
			}
			return reflect.ValueOf(&v)
		},
	},
	{
		From: "int64",
		To: "*int64",
		Cast: func(value reflect.Value) reflect.Value {
			v := value.Interface().(int64)
			if v == 0 {
				return reflect.Zero(reflect.TypeOf((*int64)(nil)))
			}
			return reflect.ValueOf(&v)
		},
	},
	{
		From: "int",
		To: "*int",
		Cast: func(value reflect.Value) reflect.Value {
			v := value.Interface().(int)
			if v == 0 {
				return reflect.Zero(reflect.TypeOf((*int)(nil)))
			}
			return reflect.ValueOf(&v)
		},
	},
	{
		From: "*int32",
		To: "*int",
		Cast: func (value reflect.Value) reflect.Value {
			v := value.Interface().(*int32)
			if v == nil {
				return reflect.Zero(reflect.TypeOf((*int)(nil)))
			}
			vInt := int(*v)
			return reflect.ValueOf(&vInt)
		},
	},
	{
		From: "*int",
		To: "*int32",
		Cast: func (value reflect.Value) reflect.Value {
			v := value.Interface().(*int)
			if v == nil {
				return reflect.Zero(reflect.TypeOf((*int32)(nil)))
			}
			vInt32 := int32(*v)
			return reflect.ValueOf(&vInt32)
		},
	},
	{
		From: "*float32",
		To: "float64",
		Cast: func (value reflect.Value) reflect.Value {
			v := value.Interface().(*float32)
			if v == nil {
				return reflect.ValueOf(0)
			}
			vFloat64 := float64(*v)
			return reflect.ValueOf(vFloat64)
		},
	},
	{
		From: "float64",
		To: "*float32",
		Cast: func (value reflect.Value) reflect.Value {
			v := value.Interface().(float64)
			if v == 0 {
				return reflect.Zero(reflect.TypeOf((*float32)(nil)))
			}
			vFloat32 := float32(v)
			return reflect.ValueOf(&vFloat32)
		},
	},
	{
		From: "*Time",
		To: "Time",
		Cast: func (value reflect.Value) reflect.Value {
			v := value.Interface().(*time.Time)
			return reflect.ValueOf(*v)
		},
	},
	{
		From: "Time",
		To: "*Time",
		Cast: func (value reflect.Value) reflect.Value {
			v := value.Interface().(time.Time)
			return reflect.ValueOf(&v)
		},
	},

	{
		From: "*IonosTime",
		To: "Time",
		Cast: func (value reflect.Value) reflect.Value {
			v := value.Interface().(*ionossdk.IonosTime)
			return reflect.ValueOf(v.Time)
		},
	},
	{
		From: "Time",
		To: "*IonosTime",
		Cast: func (value reflect.Value) reflect.Value {
			v := value.Interface().(time.Time)
			return reflect.ValueOf(&ionossdk.IonosTime{Time: v})
		},
	},


	{
		From: "Time",
		To: "Time",
		Cast: func (value reflect.Value) reflect.Value {
			v := value.Interface().(time.Time)
			return reflect.ValueOf(v)
		},
	},
	{
		From: "*map[string]string",
		To: "interface {}",
		Cast: func (value reflect.Value) reflect.Value {
			ret := map[string]string{}
			mapValue := value.Elem()
			for _, k := range mapValue.MapKeys() {
				v := mapValue.MapIndex(k)
				key := k.Interface().(string)
				switch t := v.Interface().(type) {
				case string:
					ret[key] = t
				default:
					/* todo: warn about non string values */
				}
			}
			return reflect.ValueOf(ret)
		},
	},
	{
		From: "*int32",
		To: "uint32",
		Cast: func(value reflect.Value) reflect.Value {
			v := value.Interface().(*int32)
			if *v == 0 {
				return reflect.ValueOf(0)
			}
			return reflect.ValueOf(uint32(*v))
		},
	},
}

/* fieldMap is always compat -> core */
func convert(from reflect.Value, to reflect.Value, fieldMap map[string]string) error {

	fromType := from.Type()
	toType := to.Type()

	/* jumping through pointers */
	fromDeepType := getTypeFromPtr(fromType)
	toDeepType := getTypeFromPtr(toType)

	if caster := findCustomCaster(from, to); caster != nil {
		/* custom caster found, apply it */

		/* intentionally avoiding to use setReflectValue here since Cast controls the returned pointer if any and
		 * might return nil */
		to.Set(caster.Cast(from))
		return nil
	}

	/* check nulls */
	if fromType.Kind() == reflect.Ptr && from.IsNil() && toType.Kind() == reflect.Ptr {
		to.Set(reflect.Zero(toType))
		return nil
	}

	switch {
	case isPrimitiveType(toDeepType.Kind()):
		if !isPrimitiveType(fromDeepType.Kind()) {
			return errors.New(fmt.Sprintf(
				"[convert] to is primitive %s but from is not primitive (%s)",
				getTypeName(toType), getTypeName(fromType)))
		}

		if err := setPrimitiveFields(from, to); err != nil {
			return err
		}

	case toDeepType.Kind() == reflect.Struct:
		if fromDeepType.Kind() != reflect.Struct {
			return errors.New(fmt.Sprintf("[convert] from is %s and to is %s",
				getTypeName(from.Type()), getTypeName(to.Type())))
		}

		for i := 0; i < toDeepType.NumField(); i++ {
			toField := toDeepType.Field(i)
			toFieldName := toField.Name
			toFieldValue := getFieldByName(to, toFieldName)

			fromFieldName := getFieldMapping(fieldMap, toFieldName)

			if _, ok := fromDeepType.FieldByName(fromFieldName); !ok {
				if fromFieldName != toFieldName {
					fromFieldName = toFieldName
					/* fall back to the original name */
					if _, ok := fromDeepType.FieldByName(fromFieldName); !ok {
						/* 'to' field not found in 'from'; might happen if core has new fields added that compat doesn't
						 * know about; todo: log it */
						continue
					}
				} else {
					continue
				}
			}

			fromFieldValue := getFieldByName(from, fromFieldName)
			if fromFieldValue.Type().Kind() == reflect.Ptr && fromFieldValue.IsNil() {
				/* skip nil */
				continue
			}

			if toField.Type.Kind() == reflect.Ptr {
				toFieldValue.Set(reflect.New(toField.Type.Elem()))
			}
			if err := convert(fromFieldValue, toFieldValue, fieldMap); err != nil {
				return err
			}

		}

	case toDeepType.Kind() == reflect.Slice:
		if fromDeepType.Kind() != reflect.Slice {
			return errors.New(fmt.Sprintf("[convert] cannot convert %s to %s",
				getTypeName(fromType), getTypeName(toType)))
		}

		/* make the slice */
		fromLen := getSliceLen(from)

		/* if from is empty and to is a pointer, leave to nil */
		if fromLen == 0 && toType.Kind() == reflect.Ptr {
			to.Set(reflect.Zero(toType))
			return nil
		}

		setReflectValue(to, reflect.MakeSlice(toDeepType, fromLen, fromLen * 2))

		/* convert every element */
		for i := 0; i < fromLen; i++ {
			fromElement := getSliceElement(from, i)
			toElement := getSliceElement(to, i)

			if fromElement.Type().Kind() == reflect.Ptr && fromElement.IsNil() {
				/* skip nils */
				continue
			}

			if toElement.Type().Kind() == reflect.Ptr {
				/* malloc here */
				toElement.Set(reflect.New(toElement.Type().Elem()))
			}
			if err := convert(fromElement, toElement, fieldMap); err != nil {
				return err
			}
		}

	default:
		fromTypeName := getTypeName(fromType)
		fmt.Println(fromTypeName)
		return errors.New(fmt.Sprintf("[convert] conversion to %s not implemented", getTypeName(toType)))
	}

	return nil
}

func convertToCore(from interface{}, to interface{}) error {
	fromValue := reflect.ValueOf(from)
	toValue := reflect.ValueOf(to)

	if fromValue.Type().Kind() != reflect.Ptr {
		return errors.New("[convertToCore] from is not a pointer")
	}

	if toValue.Type().Kind() != reflect.Ptr {
		return errors.New("[convertToCore] to is not a pointer")
	}
	return convert(fromValue, toValue, revertMap(CompatCoreFieldMap))
}

func convertToCompat(from interface{}, to interface{}) error {
	fromValue := reflect.ValueOf(from)
	toValue := reflect.ValueOf(to)

	if fromValue.Type().Kind() != reflect.Ptr {
		return errors.New("[convertToCompat] from is not a pointer")
	}

	if toValue.Type().Kind() != reflect.Ptr {
		return errors.New("[convertToCompat] to is not a pointer")
	}
	return convert(fromValue, toValue, CompatCoreFieldMap)
}

func setReflectValue(v reflect.Value, value reflect.Value) {
	if v.Type().Kind() == reflect.Ptr {
		v.Set(reflect.New(v.Elem().Type()))
		v.Elem().Set(value)
	} else {
		v.Set(value)
	}
}

func setPrimitiveFields(from reflect.Value, to reflect.Value) error {

	if getTypeFromPtr(from.Type()) != getTypeFromPtr(to.Type()) {
		return errors.New(
			fmt.Sprintf(
				"[convert] primitives from and to are not of the same type: %s and %s",
				getTypeName(from.Type()), getTypeName(to.Type())))
	}

	if from.Type().Kind() == reflect.Ptr {
		if to.Type().Kind() == reflect.Ptr {
			/* pointer to pointer */
			to.Set(from)
		} else {
			/* dereference from */
			to.Set(from.Elem())
		}
	} else {
		if to.Type().Kind() == reflect.Ptr {
			/* allocate and set */
			to.Set(reflect.New(to.Type().Elem()))
			to.Elem().Set(from)
		} else {
			to.Set(from)
		}
	}

	return nil
}

func isPrimitiveType(t reflect.Kind) bool {
	for _, e := range primitiveTypes {
		if t == e {
			return true
		}
	}

	return false
}

func getTypeFromPtr(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}

	return t
}

/* todo */
func getFieldMapping(fieldMap map[string]string, field string) string {
	for k, v := range fieldMap {
		if k == field {
			return v
		}
	}
	return field
}

func revertMap(fieldMap map[string]string) map[string]string {
	ret := map[string]string{}
	for k, v := range fieldMap {
		ret[v] = k
	}

	return ret
}

func getTypeName(t reflect.Type) string {

	switch t.Kind() {
	case reflect.Ptr:
		return "*" + getTypeName(t.Elem())
	case reflect.Slice, reflect.Array:
		return "[]" + getTypeName(t.Elem())
	default:
		if t.Name() == "" {
			return t.String()
		} else {
			return t.Name()
		}
	}

}

/* note: assuming t is slice or ptr to slice */
func getSliceLen(t reflect.Value) int {
	if t.Type().Kind() == reflect.Ptr {
		return t.Elem().Len()
	} else {
		return t.Len()
	}
}

func getSliceElement(slice reflect.Value, i int) reflect.Value {
	if slice.Type().Kind() == reflect.Ptr {
		return slice.Elem().Index(i)
	} else {
		return slice.Index(i)
	}
}

func getFieldByName(s reflect.Value, f string) reflect.Value {
	if s.Type().Kind() == reflect.Ptr {
		return s.Elem().FieldByName(f)
	} else {
		return s.FieldByName(f)
	}
}

func findCustomCaster(from reflect.Value, to reflect.Value) *CustomCaster {
	for _, caster := range CustomCasters {
		if caster.From == getTypeName(from.Type()) && caster.To == getTypeName(to.Type()) {
			return &caster
		}
	}

	return nil
}
