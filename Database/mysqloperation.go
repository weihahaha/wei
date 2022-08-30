package Database

// 创建表
func (m *Mysql)CreateTable(tables ...interface{})  {
	migrator := m.db.Migrator()
	for _, v := range tables{
		migrator.AutoMigrate(&v)
	}
}

// 删除表
func (m *Mysql)DropTable(tables ...interface{})  {
	migrator := m.db.Migrator()
	for _, v := range tables{
		migrator.DropTable(&v)
	}
}

