package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

const rows, columns = 9, 9

type Grid [rows][columns]int8

type safeWriter struct {
	w   io.Writer
	err error
}

// エラーインターフェース
type error interface {
	Error() string
}
type SudokuError []error

// エラーの定義
var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit  = errors.New("invalid digit")
)

func main() {
	list_28_13()
}

func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}

func list_28_13() {
	var g Grid
	err := g.Set(10, 0, 15)
	if err != nil {
		if errs, ok := err.(SudokuError); ok {
			fmt.Printf("%d error(s) occurred:\n", len(errs))
			for _, e := range errs {
				fmt.Printf("- %v\n", e)
			}
		}

		os.Exit(1)
	}
}

func (g *Grid) Set(row, column int, digit int8) error {
	var errs SudokuError
	if !inBounds(row, column) {
		errs = append(errs, ErrBounds)
	}

	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}

	if len(errs) > 0 {
		return errs
	}

	g[row][column] = digit
	return nil
}

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}

	if column < 0 || column >= columns {
		return false
	}
	return true
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return
	}
	_, sw.err = fmt.Fprintln(sw.w, s)
}

func list_28_3() {
	err := proverds("proverds.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func proverds(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	// deferはreturnの直前に実行される
	defer f.Close()

	sw := safeWriter{w: f}
	sw.writeln("Errors are values.")
	sw.writeln("Don't just check errors. handle them gracefully.")
	sw.writeln("Don't panic.")
	sw.writeln("Make the zero value useful.")
	sw.writeln("The bigger the interface. the weaker the abstraction.")
	sw.writeln("interface{} save nothing.")
	sw.writeln("Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.")
	sw.writeln("Documentation is for users.")
	sw.writeln("A little copying is  better than a little dependency.")
	sw.writeln("Clear is better then clever.")
	sw.writeln("Concurrency is not parallelism.")
	sw.writeln("Don't communicate by sharing memory, share memory by communicating.")
	sw.writeln("Channels orchestrate; mutexes serialize.")
	return sw.err
}

func list_28_1() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		// 実行時例外をスローさせるには0以外を渡す
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
