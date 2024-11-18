package env

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
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
	log.Println("Env file is :" + p)
	godotenv.Load(p)
}

func GetEnvVariable(key string) string {
	envPattern := regexp.MustCompile(`env\(["']([^"']+)["']\)`)
	match := envPattern.FindStringSubmatch(key)

	if len(match) == 2 {
		envVar := match[1]
		return os.Getenv(envVar)
	} else {
		return os.Getenv(key)
	}
}
