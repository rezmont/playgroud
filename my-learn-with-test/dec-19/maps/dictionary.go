package maps

import "errors"

var notFoundErr = errors.New("element not found")
var exsitsErr = errors.New("element exsits")

type Dictionary map[string]string

func (d Dictionary) Search(s string) (string, error) {
	def, ok := d[s]
	if !ok {
		return "", notFoundErr
	}
	return def, nil
}

func (d Dictionary) Add(s, def string) error {
	_, ok := d[s]
	if ok {
		return exsitsErr
	}
	d[s] = def
	return nil
}
