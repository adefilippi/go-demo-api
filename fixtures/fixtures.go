package fixtures

import (
	"log"
	"path"
	"path/filepath"
	"runtime"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/lunmy/go-api-core/database"
)

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func SetupFixtures() {
	log.Println("Loading fixtures...")
	database.Setup("config/database.yml")
	db, _ := database.GetDB("default").DB()
	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory(rootDir()+"/fixtures/ORM"),
		testfixtures.DangerousSkipTestDatabaseCheck(),
	)

	if err != nil {
		log.Println("Error:", err)
		panic(err)
	}

	if err := fixtures.Load(); err != nil {
		log.Println("Error:", err)
	}

	log.Println("Fixtures loaded !")
}

func DumpFixtures() {
	log.Println("Dumping fixtures...")
	database.Setup("config/database.yml")
	db, _ := database.GetDB("default").DB()
	dumper, err := testfixtures.NewDumper(
		testfixtures.DumpDatabase(db),
		testfixtures.DumpDialect("postgres"), // or your database of choice
		testfixtures.DumpDirectory(rootDir()+"/fixtures/ORM"),
		testfixtures.DumpTables( // optional, will dump all table if not given
			"model",
			"media_object",
		),
	)
	if err != nil {
		log.Println("Error:", err)
		panic(err)
	}
	if err := dumper.Dump(); err != nil {
		log.Println("Error:", err)
	}
	log.Println("Fixtures dumped !")
}
