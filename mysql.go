package main

import (
	"database/sql"
	"fmt"
)

type MySQLConfig struct {
	User string `json:"user"`
	Pass string `json:"pass"`
	Host string `json:"host"`
	Port int    `json:"port"`
	Name string `json:"name"`
}

// MySQL is a connection to a mysql DB.
type MySQL struct {
	*sql.DB
}

func NewMySQL(cfg MySQLConfig) (*MySQL, error) {

	// connect to db
	//
	dbInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.Name)
	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	mysql := MySQL{
		DB: db,
	}
	return &mysql, nil
}

// Save saves the translation in DB.
func (m *MySQL) Save(text, pig_latin string) error {
	_, err := m.Exec("INSERT INTO pig_latin VALUES (?, ?)",
		text, pig_latin)
	if err != nil {
		return fmt.Errorf("db save: %w", err)
	}
	return nil
}

// Get returns all the translations from DB.
func (m *MySQL) Get() ([][2]string, error) {
	rows, err := m.Query("select text, translation from pig_latin")
	if err != nil {
		return nil, fmt.Errorf("db query: %w", err)
	}
	defer rows.Close()

	var ret [][2]string
	for rows.Next() {
		var text, transl string
		err := rows.Scan(&text, &transl)
		if err != nil {
			return nil, fmt.Errorf("db scan: %w", err)
		}
		ret = append(ret, [2]string{text, transl})
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("db rows: %w", err)
	}

	return ret, nil
}
