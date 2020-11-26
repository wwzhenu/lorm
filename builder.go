package lorm

import (
	"fmt"
	"reflect"
)

type Where struct {
	column string
	operator string
	value string
}

type Aggregate struct {
	method string
	column string
}

type Builder struct {
	*Model
	*Grammar
	connection string
	Wheres []Where
	From string
	Columns string
	distinct bool
	Aggregate Aggregate
	Bindings map[string][]interface{}
}

func (builder *Builder) Where(column string, operator string,value string) *Builder{
	where := Where{column:column,operator:operator,value:"?"}
	builder.Wheres = append(builder.Wheres,where)
	builder.Bindings["where"] = append(builder.Bindings["where"],value)
	return builder
}

func (builder *Builder) Get(columns string,dest interface{}) {
	builder.Columns = columns
	value := reflect.ValueOf(dest)
	direct := reflect.Indirect(value)
	//base := value.Type().Elem()
	base := reflect.ValueOf(builder.RealModel).Type()
	sql := builder.ToSql()
	fmt.Println(sql)
	smt,err := GetConnection(builder.Model.Connection).Prepare(sql)
	if err != nil {
		panic(err)
	}
	rs,err := smt.Query(20)
	if err != nil {
		panic(err)
	}
	rsColumns,_ := rs.Columns()
	tagField := TagFiledMap(builder.RealModel)
	//s := reflect.Value{}
	//if reflect.TypeOf(dest).Kind() == reflect.Slice{
	//	s = reflect.ValueOf(dest)
	//	fmt.Println("aaaaaa")
	//}
	k := 0
	for rs.Next() {
		data := make([]interface{},len(rsColumns))
		fmt.Println(reflect.ValueOf(builder.RealModel).Type())
		fmt.Println(base)
		vp := reflect.New(base)
		fmt.Println(vp.Type())
		vv := reflect.Indirect(vp)
		for i,v := range rsColumns{
			fieldName := tagField[v]
			fmt.Println(fieldName)
			//value := builder.RealModelValue
			//fmt.Println(value.FieldByName(fieldName).Addr().Interface())
			//fmt.Println(&builder.Model.id)
			//data[i] = value.FieldByName(fieldName).Addr().Interface()
			data[i] = vv.FieldByName(fieldName).Addr().Interface()
		}

		rs.Scan(data...)
		fmt.Println(vv.FieldByName("Name"))
		direct.Set(reflect.Append(direct, vp))
		k++
	}


}

func (builder *Builder) Distinct() *Builder{
	builder.distinct = true
	return builder
}

func (builder *Builder)ToSql()string  {
	return builder.compileSelect(builder)
}

func (builder *Builder) Count(column string) *Builder{
	builder.Aggregate.method = "COUNT"
	builder.Aggregate.column = column
	return builder
}

func (builder *Builder) Sum(column string) *Builder{
	builder.Aggregate.method = "SUM"
	builder.Aggregate.column = column
	return builder
}

type U struct {
	Id int
}