package lorm

import (
	"reflect"
)

type Model struct {
	TableName string
	PrimaryKey string
	Connection string
	id int
	RealModel interface{}
	RealModelValue reflect.Value
}

func (m *Model) SetTable(tableName string)  {
	m.TableName = tableName
}

func (m *Model) SetPrimaryKey(primaryKey string)  {
	m.PrimaryKey = primaryKey
}

func (m *Model) GetTable()string  {
	return m.TableName
}

func (m *Model) Query() *Builder  {
	return &Builder{
		Model : m,
		From: m.GetTable(),
		Bindings: map[string][]interface{}{},
		Grammar: &Grammar{},
	}
}

func (m *Model)Table(tableName string) *Model  {
	m.SetTable(tableName)
	return  m
}

func TagFiledMap(m interface{})map[string]string  {
	value := reflect.TypeOf(m)
	filedNum := value.NumField()
	rs := make(map[string]string)
	for i:=0;i<filedNum;i++{
		filed := value.Field(i)
		filedName := filed.Name
		tagName := filed.Tag.Get("sql")
		if tagName == "" {
			continue
		}
		rs[tagName] = filedName
	}
	return rs
}