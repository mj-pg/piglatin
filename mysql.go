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
func (m *MySQL) Save(text, translated string) error {
	query := "INSERT INTO pig_latin(text, translation) VALUES (?, ?)"
	_, err := m.Exec(query, text, translated)
	if err != nil {
		return fmt.Errorf("db insert: %w", err)
	}
	return nil
}

// Get returns all the translations from DB.
func (m *MySQL) Get() ([][2]string, error) {
	query := "SELECT text, translation FROM pig_latin"
	rows, err := m.Query(query)
	if err != nil {
		return nil, fmt.Errorf("db query: %w", err)
	}
	defer rows.Close()

	var ret [][2]string
	for rows.Next() {
		var text, translated string
		err := rows.Scan(&text, &translated)
		if err != nil {
			return nil, fmt.Errorf("db scan: %w", err)
		}
		ret = append(ret, [2]string{text, translated})
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("db rows: %w", err)
	}

	return ret, nil
}

/*
// GetByText returns the saved translation of the text.
func (mysql *MySQL) GetByText(text string) (string, error) {
	query := "SELECT translation FROM pig_latin WHERE text = ?"
	row := mysql.QueryRow(query, text)
	var translated string
	if err := row.Scan(&translated); err != nil {
		return "", fmt.Errorf("db query: %w", err)
	}
	return translated, nil
}
*/
