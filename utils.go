package main

import (
	"log"
	"os"
	"os/user"
	"fmt"
	"path/filepath"
)

//logger function 
func logToFile(input string) error {
	fileName := getWorkingDirectory()+"/log4j.log"
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	log.SetOutput(file)
	log.Println(input)
	return nil
}

func getWorkingDirectory() string{
	usr,err:=user.Current()
	dirPath := filepath.Join(usr.HomeDir, ".xnotepro")

	_, err = os.Stat(dirPath)
	if os.IsNotExist(err) {
		err = os.Mkdir(dirPath, 0755) // 0755 is the permission mode
		if err != nil {
			fmt.Println("Error creating directory:", err)
			os.Exit(1)
		}
		return dirPath
	} else if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	} 
return dirPath}

