package mysql

import (
    "fmt"
    "strings"

    "database/sql"

    "github.com/software-engr-full-stack/one-step-gps/pkg/models"
)

type CustomerModel struct {
    DB     *sql.DB
    DBName string
}

func (m *CustomerModel) Validate() error {
    if m.DB == nil {
        return fmt.Errorf("DB must not be nil")
    }

    m.DBName = strings.TrimSpace(m.DBName)
    if m.DBName == "" {
        return fmt.Errorf("DBName must not be blank")
    }

    return nil
}

func (m *CustomerModel) Insert(title, content, expires string) (int, error) {
    return 0, nil
}

// TODO: paginate
func (m *CustomerModel) All() ([]*models.Customer, error) {
    return nil, nil
}

func (m *CustomerModel) Get(id int) (*models.Customer, error) {
    stmt := `SELECT id, title, content, created, expires FROM snippets
    WHERE expires > UTC_TIMESTAMP() AND id = ?`

    row := m.DB.QueryRow(stmt, id)

    s := &models.Customer{}
    err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
    if err == sql.ErrNoRows {
        return nil, models.ErrNoRecord
    } else if err != nil {
        return nil, err
    }

    return s, nil
}

func (m *CustomerModel) Latest() ([]*models.Customer, error) {
    return nil, nil
}
