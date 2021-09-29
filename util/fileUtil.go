package util

import (
	"os"
	"path"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateFile(filename string) error {
	if res, err := PathExists(filename); !res {
		if err != nil {
			return err
		}
		if resDir,dirErr:=PathExists(path.Dir(filename));!resDir{
			if dirErr != nil {
				panic(dirErr)
			}
			err:=os.MkdirAll(path.Dir(filename),0700)
			if err != nil {
				panic(err)
			}
		}
		f, err := os.Create(filename)
		if err != nil {
			return err
		}
		err = f.Close()
		return err
	}
	return nil
}
