package kv

import (
	"encoding/gob"
	"errors"
	"fmt"
	"io/fs"
	"os"
)

type store struct {
	path string
	data map[string]string
}

func OpenStore(path string) (*store, error) {
	s := &store{
		path: path,
		data: map[string]string{},
	}
	f, err := os.Open(path)
	if errors.Is(err, fs.ErrNotExist) {
		return s, nil
	}
	if err != nil {
		return nil, err
	}
	defer f.Close()
	err = gob.NewDecoder(f).Decode(&s.data)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *store) Set(k, v string) {
	s.data[k] = v
}

func (s store) Get(k string) (string, bool) {
	v, ok := s.data[k]
	return v, ok
}

func (s store) Save() error {
	f, err := os.Create(s.path)
	if err != nil {
		return err
	}
	defer f.Close()
	return gob.NewEncoder(f).Encode(s.data)
}

func (s store) Dump() {
	for k, v := range s.data {
		fmt.Printf("%s: %s\n", k, v)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s <command> [key] [value]\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "  set [key] [value]    Set a key-value pair")
	fmt.Fprintln(os.Stderr, "  get [key]		Get the value for a key")
	fmt.Fprintln(os.Stderr, "  dump			Dump all key-value pairs")
}

func Main() int {
	if len(os.Args) < 2 {
		usage()
		return 1
	}
	store, err := OpenStore("kv.store")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening store: %v\n", err)
		return 1
	}
	switch os.Args[1] {
	case "set":
		if len(os.Args) != 4 {
			usage()
			return 1
		}
		store.Set(os.Args[2], os.Args[3])
		err := store.Save()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error saving store: %v\n", err)
			return 1
		}
	case "get":
		if len(os.Args) != 3 {
			usage()
			return 1
		}
		val, ok := store.Get(os.Args[2])
		if !ok {
			usage()
			return 1
		}
		fmt.Println(val)
	case "dump":
		store.Dump()
	default:
		usage()
		return 1
	}
	return 0
}
