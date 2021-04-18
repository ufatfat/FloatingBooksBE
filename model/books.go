package model

import "time"

type (
	BorrowBook struct {
		Name			string		`json:"name"`
		StudentID		string		`json:"student_id"`
		Phone			string		`json:"phone"`
		BookID			int16		`json:"book_id"`
		BorrowTimestamp	time.Time	`json:"borrow_timestamp"`
		ReturnTimestamp	time.Time	`json:"return_timestamp"`
	}

	ReturnBook struct {
		LocationID		uint8		`json:"location_id"`
		BookID			int16		`json:"book_id"`
		Thoughts		string		`json:"thoughts"`
	}
)