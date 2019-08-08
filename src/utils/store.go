package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"os/user"
)

// QuickLink A single QuickLink
type QuickLink struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// QuickLinks The struct that represents the store
type QuickLinks struct {
	Links []QuickLink `json:"links"`
}

func userHomePath() string {
	usr, _ := user.Current()
	return usr.HomeDir
}

func configDirPath() string {
	return userHomePath() + "/.qcdconfig"
}

func storePath() string {
	return userHomePath() + "/.qcdconfig/store.json"
}

// AddQL Add quicklink to the store
func AddQL(ql QuickLink) error {
	// Create store if it does not exist
	if _, err := os.Stat(storePath()); os.IsNotExist(err) {
		err = createStoreIfNotExist()
		if err != nil {
			return err
		}
	}

	storeData, err := readStore()
	if err != nil {
		return err
	}

	var alreadyInStore QuickLinks
	var newStore QuickLinks

	err = json.Unmarshal(storeData, &alreadyInStore)
	if err != nil {
		return err
	}

	// check if link name or path already exists
	for _, link := range alreadyInStore.Links {
		if link.Name == ql.Name {
			return errors.New("A quicklink with the name '" + link.Name + "' already exists")
		}
		if link.Path == ql.Path {
			return errors.New("A quicklink with the path " + link.Path + " already exists (it's called '" + link.Name + "')")
		}
	}

	// append the new ql to the existing array of qls
	newStore.Links = append(alreadyInStore.Links, ql)

	err = writeStore(newStore.Links)
	if err != nil {
		return err
	}

	return nil
}

// RemoveQL Remove a quicklink from the store
func RemoveQL(name string) error {
	storeData, err := readStore()
	if err != nil {
		return err
	}

	var store QuickLinks

	var newStoreLinks []QuickLink
	found := false

	err = json.Unmarshal(storeData, &store)
	if err != nil {
		return err
	}

	for _, ql := range store.Links {
		// only append links store do not match the name => in order to remove the item which has the name
		if ql.Name == name {
			found = true
		} else {
			newStoreLinks = append(newStoreLinks, ql)
		}
	}

	if found == false {
		return errors.New("Quicklink with the name '" + name + "' was not found")
	}

	err = writeStore(newStoreLinks)
	if err != nil {
		return err
	}
	return nil
}

// GetQLs get all quicklinks from the store file
func GetQLs() ([]QuickLink, error) {
	storeData, err := readStore()
	if err != nil {
		return nil, err
	}

	var store QuickLinks

	err = json.Unmarshal(storeData, &store)
	if err != nil {
		return nil, err
	}

	return store.Links, nil
}

// functions that creates the store and writes to it
func writeStore(qls []QuickLink) error {
	// Create store if it does not exist
	err := createStoreIfNotExist()
	if err != nil {
		return err
	}

	var newStore QuickLinks

	newStore.Links = qls

	json, err := json.Marshal(newStore)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(storePath(), json, 077)
	if err != nil {
		return err
	}

	return nil
}

func createStoreIfNotExist() error {

	if _, err := os.Stat(storePath()); os.IsNotExist(err) {
		err := os.MkdirAll(configDirPath(), 0777)
		if err != nil {
			return err
		}
		defaultStoreEntry := []byte(`{"links":[]}`)

		err = ioutil.WriteFile(storePath(), defaultStoreEntry, 0777)
		if err != nil {
			return err
		}
	}

	return nil
}

func readStore() ([]byte, error) {
	return ioutil.ReadFile(storePath())
}
