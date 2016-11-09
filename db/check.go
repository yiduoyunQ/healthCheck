package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func healthCheck(addr string, timeout, rtimeout time.Duration) chan error {
	errCh := make(chan error)

	go func(addr string, ch chan error, timeout, rtimeout time.Duration) {

		ch <- dbCheck(addr, timeout, rtimeout)
		close(ch)

	}(addr, errCh, timeout, rtimeout)

	return errCh
}

func dbCheck(addr string, timeout, rtimeout time.Duration) error {
	end := time.Now().Add(timeout)
	rend := time.Now().Add(rtimeout)

	// DB insert test
	query := "INSERT INTO " + fTableName + " VALUES(1,'a');"
	err := dbExec(addr, query, timeout, rtimeout)
	if err != nil {
		log.Println("db insert error:", err)

		return err
	}

	timeout = end.Sub(time.Now())
	rtimeout = rend.Sub(time.Now())

	if timeout > 0 && rtimeout > 0 {
		// DB query test
		err = dbQuery(addr, timeout, rtimeout)
		if err != nil {
			log.Println("db query error:", err)
		}

	} else {
		err = errors.New("timeout before execute query db")
	}

	timeout = end.Sub(time.Now())
	rtimeout = rend.Sub(time.Now())
	if timeout <= 0 || rtimeout <= 0 {
		timeout = time.Second
		rtimeout = time.Second
	}

	// DB Delete test
	query = "DELETE FROM " + fTableName + ";"
	_err := dbExec(addr, query, timeout, readTimeout)
	if _err != nil {
		log.Println("db delete error:", _err)

		if err == nil {
			err = _err
		}
	}

	return err
}

func dbQuery(addr string, timeout, rtimeout time.Duration) error {

	source := fmt.Sprintf("%s:%s@tcp(%s)/%s?timeout=%s&readTimeout=%s", fUser, fPassword, addr, fDBName, timeout.String(), readTimeout.String())

	db, err := sql.Open(mysqlDriver, source)
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM " + fTableName + ";")
	if err != nil {
		return err
	}

	for rows.Next() {
		var (
			id    int
			value string
		)

		err := rows.Scan(&id, &value)
		if err != nil {
			rows.Close()

			return err
		}
	}

	return rows.Err()
}

func dbExec(addr, query string, timeout, rtimeout time.Duration) error {

	source := fmt.Sprintf("%s:%s@tcp(%s)/%s?timeout=%s&readTimeout=%s", fUser, fPassword, addr, fDBName, timeout.String(), readTimeout.String())

	db, err := sql.Open(mysqlDriver, source)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(query)


	return err
}
