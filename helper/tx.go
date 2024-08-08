package helper

import "gorm.io/gorm"

func CommitOrRollback(err error, tx *gorm.DB){
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	tx.Commit()
}