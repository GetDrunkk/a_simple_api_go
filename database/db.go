package database

import (
	"fmt"
	"strconv"

	"github.com/tidwall/buntdb"
)

type Logsingle struct {
	Postal_code   int `json:"postal_code"`
	Request_count int `json:"request_count"`
}

type Logs struct {
	Access_logs []Logsingle `json:"access_logs"`
}

func Insert(postal_code int) {
	db, err := buntdb.Open("data.db")
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	err_find := db.View(func(tx *buntdb.Tx) error {
		_, err := tx.Get(strconv.Itoa(postal_code))
		if err != nil {
			return err
		}
		return nil
	})
	if err_find == buntdb.ErrNotFound {
		err_insert := db.Update(func(tx *buntdb.Tx) error {
			_, _, err := tx.Set(strconv.Itoa(postal_code), "1", nil)
			return err
		})
		_ = err_insert
	} else {
		var val int
		err := db.View(func(tx *buntdb.Tx) error {
			num, err := tx.Get(strconv.Itoa(postal_code))
			val, _ = strconv.Atoi(num)
			if err != nil {
				return err
			}
			return nil
		})
		_ = err
		val = val + 1
		err_insert := db.Update(func(tx *buntdb.Tx) error {
			_, _, err := tx.Set(strconv.Itoa(postal_code), strconv.Itoa(val), nil)
			return err
		})
		_ = err_insert
	}
}

func Read_log() (res Logs) {
	db, err := buntdb.Open("data.db")
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()

	var arr_log []Logsingle

	err_view := db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("", func(key, value string) bool {
			key_int, _ := strconv.Atoi(key)
			value_int, _ := strconv.Atoi(value)
			arr_log = append(arr_log, Logsingle{Postal_code: key_int, Request_count: value_int})
			return true // continue iteration
		})
		return err
	})
	_ = err_view
	res = Logs{Access_logs: arr_log}
	return res
}
