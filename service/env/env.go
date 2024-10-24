package env

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

var envpath string = ".env"

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d) + "/.."
}

func Init(p string) {
	if p != "" {
		envpath = p
		log.Println("Using env file:", envpath)
	}
}

func GetEnvVariable(key string) string {
	err := godotenv.Load(rootDir() + "/" + envpath)
	if err != nil {
		log.Fatalf("Error loading env file:%v", err)
	}
	return os.Getenv(key)
}
