package main

import (
	"errors"
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

type DB_ERR struct {
	code  int
	level string
	msg   string
}

func (de DB_ERR) Error() string {
	return fmt.Sprintf("level:%s	code:%d	msg:%s", de.level, de.code, de.msg)
}

func main() {
	formatted_error := fmt.Errorf("this is simple formatted error")
	log.Println(formatted_error.Error())

	err := getDB_ERR(1, "tolerable", "this is error msg")
	log.Println(err.Error())

	// errors.Is() // to compare errors
	// errors.As() // to typecast errors to specific type

	// log.Fatalln("prints this msg and exits the process (no defer statements are run)")
	eg_panic()

	if err := print_log_at_source(); err != nil {
		return
	}
}

func getDB_ERR(code int, level string, msg string) DB_ERR {
	return DB_ERR{code, level, msg}
}

func eg_panic() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recovered from: ", err)
		}
	}()

	/*
		recovery is possible only if the panic is from same go-routine
		if recovered:
			then the caller execution works just fine
		else:
			all the defer statements from the caller hirarchy are executed and prg exits
	*/
	panic("im panic")
}

func print_log_at_source() error {

	err := call1()
	if err != nil {
		return err
	}
	return nil
}

func call1() error {

	err := call2()
	if err != nil {
		return err
	}
	return nil
}

func call2() error {

	log.Println("im error in call2")
	return errors.New("call2 error")
}
