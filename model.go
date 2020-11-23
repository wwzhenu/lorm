package lorm

type Model struct {
	TableName string
	PrimaryKey string
	Connection string
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
	}
}