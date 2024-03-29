package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// Path path of a file
type Path = string

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "clear files",
	Long:  "clear unused files",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("require a dir path")
		}

		dir := args[0]
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			return err
		}

		for _, file := range files {
			path := getFilePath(dir, file)

			if !isDeleteable(path) {
				continue
			}

			removeFile(path)
		}

		return nil
	},
}

// flag vars
var deleteTorrent bool;
var deleteCompression bool;

// constants
var compressionExts = []string{".7z", ".gz", ".rar", ".zip"}

func init() {
	rootCmd.AddCommand(clearCmd)

	clearCmd.Flags().BoolVar(&deleteTorrent, "torrent", false, "clear *.torrent files")
	clearCmd.Flags().BoolVar(&deleteCompression, "comp", false, "clear *.(gz|zip|rar|7z) files")
}

func isDeleteable(path Path) bool {
	fileExt := filepath.Ext(path)

	if deleteTorrent && fileExt == ".torrent" {
		return true
	}

	if deleteCompression {
		for _, ext := range compressionExts {
			if ext == fileExt {
				return true
			}
		}
	}

	return false
}

func getFilePath(dir string, file os.FileInfo) Path {
	path := fmt.Sprintf("%s/%s", dir, file.Name())

	return path
}

func removeFile(path Path) {
	err := os.RemoveAll(path)
	if err != nil {
		log.Fatalf("failed to delete file: %s", path)
	}

	log.Printf("deleted %s\n", path)
}
