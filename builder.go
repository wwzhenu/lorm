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
	limit int
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
	base := value.Type().Elem()
	//fmt.Println(value.Elem().Type())
	//fmt.Println(value.Type())// 值的数据类型 User等
	//fmt.Println(value.Elem().Type().Elem().Kind())// slice中元素kind
	//fmt.Println(value.Kind()) // ptr struct等等
	//fmt.Println(base)
	base = value.Elem().Type().Elem()//若dest为指针使用此
	//fmt.Println(base)

	//base := value.Type().Elem()//若dest为非指针使用此
	//base := reflect.ValueOf(builder.RealModel).Type()
	sql := builder.ToSql()
	fmt.Println(sql)
	smt,err := GetConnection(builder.Model.Connection).Prepare(sql)
	if err != nil {
		panic(err)
	}
	fmt.Println(builder.Bindings)
	rs,err := smt.Query(builder.Bindings["where"]...)
	if err != nil {
		panic(err)
	}
	rsColumns,_ := rs.Columns()
	tagField := TagFiledMap(builder.RealModel)
	for rs.Next() {
		data := make([]interface{},len(rsColumns))
		vp := reflect.New(base)
		vv := reflect.Indirect(vp)
		for i,v := range rsColumns{
			fieldName := tagField[v]
			data[i] = vv.FieldByName(fieldName).Addr().Interface()
		}

		rs.Scan(data...)
		direct.Set(reflect.Append(direct, vv))
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

func (builder *Builder) Limit(offset int) *Builder{
	builder.limit = offset
	return builder
}

func (builder *Builder) First(columns string,dest interface{}){
	builder.Columns = columns
	builder.Limit(1)
	value := reflect.ValueOf(dest)
	direct := reflect.Indirect(value)
	base := value.Type().Elem()
	//fmt.Println(value.Elem().Type())
	//fmt.Println(value.Type())// 值的数据类型 User等
	//fmt.Println(value.Elem().Type().Elem().Kind())// slice中元素kind
	//fmt.Println(value.Kind()) // ptr struct等等
	//fmt.Println(base)
	fmt.Println(base)

	//base := value.Type().Elem()//若dest为非指针使用此
	//base := reflect.ValueOf(builder.RealModel).Type()
	sql := builder.ToSql()
	fmt.Println(sql)
	smt,err := GetConnection(builder.Model.Connection).Prepare(sql)
	if err != nil {
		panic(err)
	}
	fmt.Println(builder.Bindings)
	rs,err := smt.Query(builder.Bindings["where"]...)
	if err != nil {
		panic(err)
	}
	rsColumns,_ := rs.Columns()
	tagField := TagFiledMap(builder.RealModel)
	for rs.Next() {
		data := make([]interface{},len(rsColumns))
		for i,v := range rsColumns{
			fieldName := tagField[v]
			data[i] = direct.FieldByName(fieldName).Addr().Interface()
		}
		rs.Scan(data...)
	}

}

func (builder *Builder) Value(column string,dest interface{}){
	builder.Columns = column
	builder.Limit(1)

	//base := value.Type().Elem()//若dest为非指针使用此
	//base := reflect.ValueOf(builder.RealModel).Type()
	sql := builder.ToSql()
	fmt.Println(sql)
	smt,err := GetConnection(builder.Model.Connection).Prepare(sql)
	if err != nil {
		panic(err)
	}
	fmt.Println(builder.Bindings)
	smt.QueryRow(builder.Bindings["where"]...).Scan(dest)
}
