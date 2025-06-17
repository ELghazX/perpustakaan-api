package connection

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/elghazx/perpustakaan/internal/config"
	_ "github.com/lib/pq"
)

func GetDatabase(conf config.Database) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable Timezone=%s",
		conf.Host,
		conf.Port,
		conf.Pass,
		conf.User,
		conf.Name,
		conf.Tz,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Gagal konek database`", err.Error())
	}
	if err := db.Ping(); err != nil {
		log.Fatal("gagal ping koneksi database", err.Error())
	}
	return db
}
