package main

import "errors"

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	// 第一引数に値、第二引数に数値が格納されたかどうかのboolean値がはいるらしい
	definition, ok := d[word]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}
	return definition, nil
}
