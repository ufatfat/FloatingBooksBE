package db

import (
	"FloatingBooks/model"
	"gorm.io/gorm"
	"time"
)

func GetBookName (bookID int16) (bookName string, err error) {
	err = Mysql.Table("books").Select("book_name").Where("book_id=?", bookID).Take(&bookName).Error
	return
}

func BorrowABook (borrowInfo *model.BorrowBook) (ok bool, msg string) {
	if ok, msg = createUser(borrowInfo); !ok {
		return
	}
	if ok, msg = checkBorrow(borrowInfo); !ok {
		return
	}

	borrowInfo.BorrowTimestamp = time.Now()

	Mysql.Table("records").Select("book_id", "student_id", "borrow_timestamp").Create(borrowInfo)
	return true, "借书成功！"
}

func createUser (borrowInfo *model.BorrowBook) (ok bool, msg string) {
	var id int32
	if err := Mysql.Table("students").Select("id").Where("student_id=? and phone=?", borrowInfo.StudentID, borrowInfo.Phone).Take(&id).Error; err == gorm.ErrRecordNotFound {
		Mysql.Table("students").Select("student_id", "phone", "name").Create(borrowInfo)
		return true, ""
	} else if err != nil {
		return false, err.Error()
	}
	return true, ""
}

func checkBorrow (borrowInfo *model.BorrowBook) (ok bool, msg string) {
	return true, ""
}