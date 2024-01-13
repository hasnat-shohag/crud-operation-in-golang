package main

// import (
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

type User struct {
	ID             uint `gorm:"primary_key"`
	Name           string
	Email          string
	Age            int
	// gorm.DeletedAt `gorm:"index"`
}

func main() {

	// create connection
	// dsn := "root:root@tcp(localhost:3306)/Test?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Info),
	// }) //Logger: logger.Default.LogMode(logger.Info),
	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// create the table
	//err = db.AutoMigrate(&User{})
	//if err != nil {
	//	panic("failed to connect database")
	//}

	// insert data
	//singleUser := User{Name: "Anna", Email: "Anna@gmail.com", Age: 23}
	//result := db.Create(singleUser)
	//multipleUser := []User{
	//	{
	//		Name:  "Mahfuz",
	//		Email: "fklasjdf",
	//		Age:   25,
	//	},
	//	{
	//		Name:  "Mahfuz2",
	//		Email: "fklasjdf2",
	//		Age:   25,
	//	},
	//}
	//result := db.Create(&multipleUser)
	//result := db.Create(multipleUser)
	//if result.Error != nil {
	//	panic("failed to connect database")
	//}
	//fmt.Println("Rows affected : ", result.RowsAffected)
	//fmt.Println("Auto generated user ID : ", singleUser.ID)

	// read data
	//var user User
	//result := db.Where("name = ?", "mahfuz").First(&user)
	//if result.Error != nil {
	//	panic("failed to connect database")
	//}
	//fmt.Println("User : ", user)
	//var multipleUser []User
	//result := db.Find(&multipleUser)
	//if result.Error != nil {
	//	panic("failed to connect database")
	//}
	//for idx, user := range multipleUser {
	//	fmt.Println("User ", idx+1, " : ", user.Name, user.Email, user.Age)
	//}
	//fmt.Println("Rows affected : ", result.RowsAffected)

	// update data
	//updateUser := User{ID: 10}
	//result := db.Save(&updateUser)
	//result := db.Model(&User{ID: 10}).Updates(User{Name: "zyz", Email: "xyz@gmail.com", Age: 25})
	//if result.Error != nil {
	//	panic("failed to connect database")
	//}
	//fmt.Println("Rows affected : ", result.RowsAffected)

	// Hard delete data
	//var user User
	//result := db.Delete(user, 1) // 1st parameter is the model, 2nd parameter is ID
	//if result.Error != nil {
	//	panic("failed to connect database")
	//}
	//fmt.Println("Rows affected : ", result.RowsAffected)

	// Soft delete data
	//result := db.Delete(&User{}, 1) // 1st parameter is the model, 2nd parameter is ID})
	//if result.Error != nil {
	//	panic("failed to connect database")
	//}
	//fmt.Println("Rows affected : ", result.RowsAffected)
}
