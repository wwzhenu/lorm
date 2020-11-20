package lorm

type ConnectionInterface interface {
	table(table string) Builder
	get(query string,bindings []interface{})map[string]interface{}
}
