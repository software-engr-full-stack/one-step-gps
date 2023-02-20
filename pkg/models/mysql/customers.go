package mysql

import (
    "fmt"

    "database/sql"

    "github.com/software-engr-full-stack/one-step-gps/pkg/models"
)

const tableName = "customers"

type CustomerModel struct {
    DB *sql.DB
}

func (m *CustomerModel) Validate() error {
    if m.DB == nil {
        return fmt.Errorf("DB must not be nil")
    }

    return nil
}

func (m *CustomerModel) Get(id int) (*models.Customer, error) {
    stmt := fmt.Sprintf(
        `SELECT id, name, business_category, payload, img_url, created, updated FROM %s
        WHERE id = ?`,
        tableName,
    )

    row := m.DB.QueryRow(stmt, id)

    cm := &models.Customer{}
    err := row.Scan(&cm.ID, &cm.Name, &cm.BusinessCategory, &cm.Payload, &cm.ImgURL, &cm.Created, &cm.Updated)
    if err == sql.ErrNoRows {
        return nil, models.ErrNoRecord
    } else if err != nil {
        return nil, err
    }

    return cm, nil
}

func (m *CustomerModel) Latest() ([]*models.Customer, error) {
    stmt := fmt.Sprintf(
        `SELECT id, name, business_category, payload, img_url, created, updated FROM %s
        ORDER BY created DESC LIMIT 10`,
        tableName,
    )
    rows, err := m.DB.Query(stmt)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    customers := []*models.Customer{}
    for rows.Next() {
        cm := &models.Customer{}
        err = rows.Scan(&cm.ID, &cm.Name, &cm.BusinessCategory, &cm.Payload, &cm.ImgURL, &cm.Created, &cm.Updated)
        if err != nil {
            return nil, err
        }
        customers = append(customers, cm)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return customers, nil
}
