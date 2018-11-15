package sql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"longan_data_center/config"
)

//人脸注册表结构
type FacesRegisterTable struct {
	Id 				int    `gorm: "primary_key;AUTO_INCREMENT:number"`
	UserId  		uint64
	Name 			string
	Photo   		string `gorm:"column:face_photo"`
	Feature 		[]byte `gorm:"column:face_feature"`
	Size			uint64 `gorm:"column:face_size"`
	AcqTime			uint64 `gorm:"column:acquisition_time"`
	Score           uint32 `gorm:"column:face_score"`
}

//陌生人表结构
type FacesStrangerTable struct {
	Id 				int    `gorm: "primary_key;AUTO_INCREMENT:number"`
	UserId  		uint64
	Name 			string
	Photo   		string `gorm: "column:face_photo"`
	Feature 		[]byte `gorm: "column:face_feature"`
	Size			uint64 `gorm: "column:face_size"`
	AcqTime			uint64 `gorm: "column:acquisition_time"`
	Score           uint32 `gorm: "column:face_score"`
}

//数据中心操作
type DC struct {
	db *gorm.DB
}

//连接数据库
func (dc *DC) Connect() error {
	dataSource := config.DCConfig.Mysql.DataSource
	db, err := gorm.Open("mysql", dataSource)
	dc.db = db;
	return err;
}

//关闭连接
func (dc *DC) Close()  error {
	return dc.db.Close()
}

//全局禁用表名复数，比如FacesRegisterTable表默认映射为faces_register_tables,
// 最后面那个s有时不是我们需要的，这时可以调用这个接口，来阻止这种行为
func (dc *DC) SingularTable(enable bool) {
	dc.db.SingularTable(enable)
}

func (dc *DC) AddRecord(value interface{}) bool {
	dc.db.Create(value)
	return !dc.db.NewRecord(value)
}

func (dc *DC) GetAllEmployeesRecord()  []FacesRegisterTable {
	var employees []FacesRegisterTable
	dc.db.Find(&employees)
	return employees
}