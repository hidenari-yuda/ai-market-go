package repository_test

import (
	"os"
	"testing"
	"time"

	// "github.com/hidenari-yuda/ai-market-go/domain/config"
	"github.com/hidenari-yuda/ai-market-go/domain/utils"
	"github.com/hidenari-yuda/ai-market-go/infra/database"
	"github.com/hidenari-yuda/ai-market-go/tests"
)

var (
	dbm    *database.Db
	// cfg    config.Config
	helper *tests.Helper
)

func TestMain(m *testing.M) {
	time.Local = utils.Tokyo

	// cfg, err := config.New()
	// if err != nil {
	// 	panic(err)
	// }

	dbm = database.NewDb()
	helper = tests.NewHelper()

	// Do tests
	code := m.Run()

	os.Exit(code)
}
