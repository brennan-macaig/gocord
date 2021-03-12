package gocord

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

type Database struct {
	// Mutex to lock/unlock as needed
	mu *sync.Mutex

	// All the messages that are stored by the bot
	Msgs Messages

	// The config for the bot
	Conf Config
}

type Messages map[string][]string

func MakeDatabase(c Config) Database {
	db := Database{
		Conf: c,
		Msgs: make(map[string][]string),
		mu:   &sync.Mutex{},
	}
	return db
}

// Get the database messages from the database file
// If the file does not exist, return error
func (d *Database) GetDatabaseMessages(path string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("could not open db file - %w", err)
	}
	err = json.Unmarshal(dat, &d.Msgs)
	if err != nil {
		return fmt.Errorf("could not unmarshal db - %w", err)
	}
	return nil
}

// Write the database messages to the database file
// If the file does not exist, create it
func (d *Database) WriteDatabaseMessages(path string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	byt, err := json.MarshalIndent(d.Msgs, "", " ")
	if err != nil {
		return fmt.Errorf("could not marshal db - %w", err)
	}

	err = ioutil.WriteFile(path, byt, 0664)
	if err != nil {
		return fmt.Errorf("could not write db file - %w", err)
	}
	return nil
}
