package main

// Constantes de error
var (
    ErrNotFound = DictionaryError("The word doesn't exists in the dictionary")
    ErrWordExists = DictionaryError("The word already exists")
    ErrWordNotExists = DictionaryError("The word doesn't exists")
)

// Definimos el tipo de error personalizado
type DictionaryError string

// Implementamos la interfaz de error con el metodo Error() y devolvemos la cadena
func (e DictionaryError) Error() string {
    return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(searching string) (string, error) {
    value, founded := d[searching]

    if !founded {
        return "", ErrNotFound
    }
    
    return value, nil
}


func (d Dictionary) Add(word, definition string) error {
    _, err := d.Search(word)

    switch err{
    case ErrNotFound:
        d[word] = definition
    case nil:
        return ErrWordExists
    default:
        return err
    }

	return nil
}

func (d Dictionary) Update(word, newDef string) error {
    _, err := d.Search(word)

    switch err{
   case ErrNotFound:
        return ErrWordNotExists
    case nil:
        d[word] = newDef
    default:
        return err
    }

    return nil
}

func (d Dictionary) Delete(word string) error {
    _, err := d.Search(word)

    switch err {
    case ErrNotFound:
        return ErrWordNotExists
    case nil:
        delete(d, word)
    default:
        return err
    }
    return nil
} 
