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
		return "", ErrNotFound
	}
	return definition, nil
}

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return nil
	case nil:
		d[word] = newDefinition
	}
	return nil

}
