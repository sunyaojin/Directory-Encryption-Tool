package main

import (
	"flag"
	"log"
	"os"
	"strings"
)

func main() {
	return
	logger, unlocker := readFromUnlockInputs()
	//
	// get pubkeypath
	//
	pubKeyPath := unlocker.publicKeyFilePath
	//
	// validate pubkey subject
	//
	unlocker.validatePubKeyFileSubject(pubKeyPath)
	//
	// validate signature
	//
	var isValidSig bool
	isValidSig = unlocker.readSigAndValidate(unlocker.directoryPath,pubKeyPath)
	if !isValidSig {
		return
	}
	//
	// fetch aes key, import from keyfile, and decrypt
	//
	var aes []byte
	aes = unlocker.importEncryptedAES()
	//
	// filewalk, getting list of files in directory
	//
	fileList := unlocker.fileWalkRecursive(unlocker.directoryPath)
	//
	// decrypt the files in this list.
	//
	for _,file := range fileList {
		unlocker.decryptFileAndReplace(file,aes)
	}
	//
	// close the logger
	//
	closelockLogger(logger)
}
//
// if keyfile exists, do nothing, else create keyfile
//
func createKeyfile(filename string) {

	var dirPath string
	//
	// - - - - - - append keyfile to path
	//
	tempDirPath := strings.Split(filename,"/")
	tempDirPath = append(tempDirPath,"keyfile")
	dirPath = strings.Join(tempDirPath,"/")

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// path/to/whatever does not exist
		//
		//
		outFile, _ := os.Create(dirPath)
		defer outFile.Close()
		_ = outFile.Sync()
		outFile.Close()
	}
}

func readFromUnlockInputs() ( *log.Logger, *Unlocker) {
	logger := createlockLogger()
	u := &Unlocker{}
	u.setFlagParameters()
	flag.Parse()
	return logger, u
}

func closelockLogger(logger *log.Logger) {
	logger.Println("\n...........END...............")
}


func createlockLogger() *log.Logger {
	loggerFile, err := os.OpenFile("text.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	logger := log.New(loggerFile, "Keygen|", log.LstdFlags)
	logger.Println("\n\n\n------------------------Log File Created----------------------")
	return logger
}

