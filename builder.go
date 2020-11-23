package lorm

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

func (builder *Builder) Get(columns string) *Builder{
	builder.Columns = columns
	return builder
}

func (builder *Builder) Distinct() *Builder{
	builder.distinct = true
	return builder
}

func (builder Builder)ToSql()string  {
	return builder.compileSelect(builder)
}

func (builder *Builder) Count(column string) *Builder{
	builder.Aggregate.method = "COUNT"
	builder.Aggregate.column = column
	return builder
}

func (builder *Builder) Sum(column string){
	builder.Aggregate.method = "SUM"
	builder.Aggregate.column = column
}