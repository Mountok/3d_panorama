package cfg

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Cfg struct {
	Port       string
	DBname     string
	DBuser     string
	DBpass string
	DBhost     string
	DBport     string
}

func LoadConfig() Cfg {
	v := viper.New()
	v.SetEnvPrefix("3D_PARONAMA_BACK") // превикс
	v.Set("PORT", "8080") // порт сервера
	v.Set("DBNAME", "h_2024_pano_images") // название базы данных
	v.Set("DBUSER", "postgres") // пользователь бд
	v.Set("DBPASS", "root") // пароль пользователя
	v.Set("DBHOST", "") // локальный хостинг по дефолту
	v.Set("DBPORT", "5433") // порт базы данных
	v.AutomaticEnv()

	var cfg Cfg

	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Panic(err)
	}
	return cfg

}

func (cfg *Cfg) GetDBConnetcUrl() string { //маленький метод для сборки строки соединения с БД
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBuser,
		cfg.DBpass,
		cfg.DBhost,
		cfg.DBport,
		cfg.DBname,
	)
}