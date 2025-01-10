package main

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(k string) (error, string) {
	w, ok := d[k]

	if !ok {
		return ErrNotFound, ""
	}

	return nil, w
}

func (d Dictionary) Add(k, v string) error {
	err, _ := d.Search(k)
	switch err {
	case ErrNotFound:
		d[k] = v
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(k, v string) error {
	err, _ := d.Search(k)

	switch err {
	case nil:
		d[k] = v
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(k string) error {
	err, _ := d.Search(k)

	switch err {
	case nil:
		delete(d, k)
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}

	return nil
}
