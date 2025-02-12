package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Port        int
	Directory   string
	StoragePath string
}

func ConfigLoad() error {
	port := flag.Int("port", 8080, "port of srever")
	directory := flag.String("dir", "data", "data directory")
	help := flag.Bool("help", false, "help")

	flag.Parse()

	if help != nil && *help {
		printHelp()
		os.Exit(0)
	}

	if err := validatePath(*directory); err != nil {
		return err
	}

	storagePath, _ := filepath.Abs(*directory)

	if *port < 1024 {
		return errors.New("port couldn't be equal less than 1024")
	}

	cfg = Config{*port, *directory, storagePath}
	return cfg.CreateStorage()
}

func GetStoragePath() string {
	return cfg.StoragePath
}

func GetConfigPort() int {
	return cfg.Port
}

var cfg Config

func validatePath(path string) error {
	c, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	return isNotInProgramFiles(c)
}

func isNotInProgramFiles(path string) error {
	cmdPath, _ := filepath.Abs("cmd")
	internalPath, _ := filepath.Abs("internal")
	modelsPath, _ := filepath.Abs("models")
	if strings.HasPrefix(path, cmdPath) || strings.HasPrefix(path, internalPath) || strings.HasPrefix(path, modelsPath) {
		return errors.New("path to storage directory is inside a program file")
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("could not parse path for data directory")
	}

	if strings.Contains(path, "..") {
		return fmt.Errorf("breaking code traversal")
	}

	if filepath.IsAbs(path) && strings.Contains(path, "..") {
		return fmt.Errorf("relative paths or directory traversal is not allowed")
	}

	projectRootDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("could not start server on %s", projectRootDir)
	}
	if !strings.HasPrefix(absPath, projectRootDir) {
		return fmt.Errorf("path is outside of the project directory")
	}

	if absPath == projectRootDir {
		return fmt.Errorf("cannot create file directly inside the project root directory")
	}

	return nil
}

func printHelp() {
	fmt.Println(`Coffee Shop Management System

Usage:
  hot-coffee [--port <N>] [--dir <S>] 
  hot-coffee --help`)
	fmt.Println("\nOptions:")
	flag.PrintDefaults()
}

func (cfg Config) CreateStorage() error {
	if _, err := os.Stat(cfg.StoragePath); os.IsNotExist(err) {
		os.Mkdir(cfg.StoragePath, 0o777)
	}
	if _, err := os.Stat(filepath.Join(cfg.StoragePath, "orders.json")); os.IsNotExist(err) {
		if file, err := os.Create(filepath.Join(cfg.StoragePath, "orders.json")); err != nil {
			return err
		} else {
			defer file.Close()
		}
	}
	if _, err := os.Stat(filepath.Join(cfg.StoragePath, "inventory.json")); os.IsNotExist(err) {
		if file, err := os.Create(filepath.Join(cfg.StoragePath, "inventory.json")); err != nil {
			return err
		} else {
			defer file.Close()
		}
	}
	if _, err := os.Stat(filepath.Join(cfg.StoragePath, "menu_items.json")); os.IsNotExist(err) {
		if file, err := os.Create(filepath.Join(cfg.StoragePath, "menu_items.json")); err != nil {
			return err
		} else {
			defer file.Close()
		}
	}
	return nil
}
