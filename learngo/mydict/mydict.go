package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotfound   = errors.New("Not found")
	errWordExists = errors.New("That word already exists")
	errCantUpdate = errors.New("Cant update non-existing word")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]

	if exists {
		return value, nil
	}
	return "", errNotfound
}

// Add a word in the dictinoary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotfound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

// Update dictinoary with word
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotfound:
		return errCantUpdate
	}
	return nil
}

// Delete aw ord in the dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
