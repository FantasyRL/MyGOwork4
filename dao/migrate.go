package dao

func migration() {
	//DB.Set("gorm:table_options", "charset=utf8mb4").
	//	AutoMigrate(&db.User{}).
	//	AutoMigrate(&db2.Video{})
	//DB.Model(&db2.Video{}).AddForeignKey("uid", "user(id)", "CASCADE", "CASCADE")
	//CASCADE:在父表上update/delete记录时，同步update/delete子表的匹配记录
}
