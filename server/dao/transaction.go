package dao

import "database/sql"

func dbTransaction(tx *sql.Tx, s func(*sql.Tx) error) error {
	var err error

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = s(tx)
	return err
}
