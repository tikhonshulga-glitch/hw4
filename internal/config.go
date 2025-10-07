package internal

import "flag"

// 1 Создание структуры конфиг для хранения инфы для запуска
type Config struct {
	Host string
	Port int
	//TODO: DB connection string
	//TODO: Debug
}

func ReadConfig() Config {

	var cfg Config

	// упаковал данные о хосте в переменную дата, результат запишется в переменную cfg
	flag.StringVar(&cfg.Host, "host", "localhost", "указание адреса для запуска сервера")
	flag.IntVar(&cfg.Port, "port", 8080, "указание порта для запуска сервера")

	flag.Parse()
	return cfg
}
