package lorm

import (
	"fmt"
	"reflect"
	"strings"
)

type Grammar struct {

}

type selectComponent struct {
	varName string
	method  reflect.Value
}
//var selectComponents =  map[string]{ "aggregate":reflect.ValueOf(compileFrom),"columns","from","joins","wheres","groups","havings","orders","limit","offset","unions","lock",}
//var selectComponents =  map[string]reflect.Value{ "columns":reflect.ValueOf(compileColumns),"from":reflect.ValueOf(compileFrom),"wheres":reflect.ValueOf(compileWheres)}

var selectComponents =  []selectComponent{
	{
		varName: "Aggregate",
		method:  reflect.ValueOf(compileAggregate),
	},
	{
	varName: "Columns",
	method:  reflect.ValueOf(compileColumns),
	},
	{
	varName: "From",
	method:  reflect.ValueOf(compileFrom),
	},
	{
	varName: "Wheres",
	method:  reflect.ValueOf(compileWheres),
	},
}



func (*Grammar)compileSelect(builder *Builder)string  {
	return strings.Join(compileComponents(builder)," ")
}

func compileComponents(builder *Builder)[]string{
	var sql []string
	for _,v := range selectComponents{
		k := v.varName
		param := reflect.ValueOf(builder).Elem().FieldByName(k).String()
		if param != "" {
			method := v.method
			params := []reflect.Value{reflect.ValueOf(builder),reflect.ValueOf(param)}
			sql = append(sql,method.Call(params)[0].String())
		}

	}
	return sql
}

func compileAggregate(builder *Builder, table string)string  {
	if builder.Columns != "" {
		return ""
	}
	sel := "SELECT "
	column := builder.Aggregate.column
	if builder.distinct && column != "*"{
		column =  "DISTINCT " + column
	}
	return fmt.Sprintf(sel+"%s(%s) AS aggregate",builder.Aggregate.method,column)
}

func compileFrom(builder *Builder, table string)string  {
	return "FROM "+table
}

func compileWheres(builder *Builder, table string)string  {
	rs := "WHERE "
	if len(builder.Wheres)>0 {
		for _,v := range builder.Wheres{
			tmp := v.column +" "+ v.operator +" "+ v.value+""+" AND "
			rs = rs+tmp
		}
		rs = strings.Trim(rs," AND ")
	}
	return rs
}

func compileColumns(builder *Builder, table string)string  {
	if builder.distinct && builder.Columns != "*" {
		return "SELECT DISTINCT "+builder.Columns
	}
	return "SELECT "+builder.Columns
}

func UcFirst(str string)string  {
	for k,v := range str{
		return strings.ToUpper(string(v)) + str[k+1:]
	}
	return ""
}