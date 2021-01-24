package modelloader

import "reflect"

const ptr = "ptr"

//ModelResult 解析模型结果
type ModelResult struct {
	modelName string

	fields   []ModelFeild
	numField uint
}

//ModelFeild 模型中的模板
type ModelFeild struct {
	feildType reflect.Type

	rawValue      reflect.Value
	formatedValue string

	feildIndex    uint
	feildName     string
	feildShowName string

	maxSize       uint
	htmlInputType uint

	autoGenerate bool
	generateWay  string

	sha256Hash bool
	editAble   bool
}

//NewModel 通过给定的
func NewModel(target interface{}) ModelResult {
	var models ModelResult
	targetType := typeElemOutPtr(reflect.TypeOf(target))
	targetValue := valueElemOutPtr(reflect.ValueOf(target))

	models.modelName = targetType.Name()
	models.numField = uint(targetType.NumField())

	for i := 0; i < int(models.numField); i++ {
		feildType := targetType.Field(i)
		feildValue := targetValue.Field(i)
		model := newModelFeild(feildType, feildValue, uint(i))
		models.fields = append(models.fields, model)
	}

	return models
}

func typeElemOutPtr(t reflect.Type) reflect.Type {
	if t.Kind().String() != ptr {
		return t
	}
	return typeElemOutPtr(t.Elem())
}

func valueElemOutPtr(t reflect.Value) reflect.Value {
	if t.Kind().String() != ptr {
		return t
	}
	return valueElemOutPtr(t.Elem())
}

func newModelFeild(t reflect.StructField, v reflect.Value, index uint) ModelFeild {
	var model ModelFeild
	model.feildType = t.Type
	model.rawValue = v

	model.feildName = t.Type.Name()
	model.feildIndex = index

	gormTagLoad(t, &model)
	adminTagLoad(t, &model)

	return model
}
