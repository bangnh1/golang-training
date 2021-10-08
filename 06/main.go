package main

import (
	"log"
	"sync"

	"github.com/bangnh1/golang-training/06/cfg"
	"github.com/bangnh1/golang-training/06/utils"
	"github.com/bangnh1/golang-training/06/videoConverter"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func init() {
	cfg.SetupConfig()
}

func main() {

	watcher, err := fsnotify.NewWatcher()
	var wg sync.WaitGroup
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	resumeHashing(&wg)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("Create a file:", event.Name)
					filename, err := utils.GetFileName(event.Name)
					if err != nil {
						continue
					}
					log.Println("Filename:", filename)
					wg.Add(1)
					go videoConverter.HashVideoToHLS(filename, &wg)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(viper.GetString(cfg.ConfigKeyUploadDir))
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

// Resume hashing after crashed
func resumeHashing(wg *sync.WaitGroup) {

	existedFiles, _ := utils.ListAllFiles(viper.GetString(cfg.ConfigKeyUploadDir))
	hashedFile := utils.ReadLogFile()
	hashingFile := utils.GetHashingFiles(existedFiles, hashedFile)
	log.Printf("Hashing file: %v", hashingFile)

	for _, file := range hashingFile {
		wg.Add(1)
		go videoConverter.HashVideoToHLS(file, wg)
	}
}
