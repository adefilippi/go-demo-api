package env

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d) + "/.."
}

func Init(p string) {
	if p == "" {
		p = ".env"
	}
	log.Println("Before : ", os.Getenv("ENV"))
	godotenv.Load(rootDir() + "/" + p)
	log.Println("Before : ", os.Getenv("ENV"))
}

func GetEnvVariable(key string) string {
	return os.Getenv(key)
}
