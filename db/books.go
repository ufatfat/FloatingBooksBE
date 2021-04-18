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

func ReturnBook (returnInfo *model.ReturnBook) (ok bool) {
	if err := Mysql.Table("records").Where("book_id=?", returnInfo.BookID).Updates(map[string]interface{}{"is_returned": 1, "return_timestamp": time.Now(), "thoughts": returnInfo.Thoughts}).Error; err != nil {
		return false
	}
	if ok, _ = changeBookStatus(false, returnInfo.BookID); !ok {
		return ok
	}
	if ok = changeBookLocation(returnInfo.LocationID, returnInfo.BookID); !ok {
		return ok
	}
	return true
}
func BorrowABook (borrowInfo *model.BorrowBook) (ok bool, msg string) {
	if ok, msg = createUser(borrowInfo); !ok {
		return
	}
	if ok, msg = checkBorrow(borrowInfo); !ok {
		return
	}

	borrowInfo.BorrowTimestamp = time.Now().Local()

	Mysql.Table("records").Select("book_id", "student_id", "borrow_timestamp").Create(&borrowInfo)
	if ok, msg = changeBookStatus(true, borrowInfo.BookID); !ok {
		return
	}
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

func changeBookLocation (locationID uint8, bookID string) (ok bool) {
	if err := Mysql.Table("books").Where("book_id=?", bookID).Update("book_location", locationID).Error; err != nil {
		return
	}
	return true
}

func changeBookStatus (isLend bool, bookID string) (ok bool, msg string) {
	if err := Mysql.Table("books").Where("book_id=?", bookID).Update("is_lend", isLend).Error; err != nil {
		return
	}
	return true, ""
}