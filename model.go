package lorm

type Model struct {
	*Builder
	TableName string
	primaryKey string
	connection string
}

func (m *Model) SetTable(tableName string)  {
	m.TableName = tableName
}

func (m *Model) SetPrimaryKey(primaryKey string)  {
	m.primaryKey = primaryKey
}

func (m *Model) GetTable()  {
	return
}