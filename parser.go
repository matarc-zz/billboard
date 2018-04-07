package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Parser struct {
	scanner *bufio.Scanner
}

func NewParser(reader io.Reader) *Parser {
	r := new(Parser)
	r.scanner = bufio.NewScanner(reader)
	r.scanner.Split(bufio.ScanLines)
	return r
}

func (p *Parser) ParseLine() (listWords []string, width, height uint64, err error) {
	p.scanner.Scan()
	if err = p.scanner.Err(); err != nil {
		return
	}
	return parseLine(strings.NewReader(p.scanner.Text()))
}

func parseLine(reader io.Reader) (listWords []string, width, height uint64, err error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	if err = scanner.Err(); err != nil {
		return
	}
	if width, err = strconv.ParseUint(scanner.Text(), 10, 32); err != nil {
		return
	}
	scanner.Scan()
	if err = scanner.Err(); err != nil {
		return
	}
	if height, err = strconv.ParseUint(scanner.Text(), 10, 32); err != nil {
		return
	}

	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			return
		}
		listWords = append(listWords, scanner.Text())
	}
	return
}

func (p *Parser) ParseNumberOfLines() (numberOfLines int, err error) {
	p.scanner.Scan()
	if err = p.scanner.Err(); err != nil {
		return
	}
	return strconv.Atoi(p.scanner.Text())
}
