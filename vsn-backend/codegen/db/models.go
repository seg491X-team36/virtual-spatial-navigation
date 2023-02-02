// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type UserAccountRole string

const (
	UserAccountRoleREGISTERED UserAccountRole = "REGISTERED"
	UserAccountRoleREJECTED   UserAccountRole = "REJECTED"
	UserAccountRoleACCEPTED   UserAccountRole = "ACCEPTED"
)

func (e *UserAccountRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserAccountRole(s)
	case string:
		*e = UserAccountRole(s)
	default:
		return fmt.Errorf("unsupported scan type for UserAccountRole: %T", src)
	}
	return nil
}

type NullUserAccountRole struct {
	UserAccountRole UserAccountRole
	Valid           bool // Valid is true if UserAccountRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserAccountRole) Scan(value interface{}) error {
	if value == nil {
		ns.UserAccountRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserAccountRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserAccountRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.UserAccountRole, nil
}

type Experiment struct {
	ID          uuid.UUID
	Name        string
	Description string
}

type ExperimentResult struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	ExperimentID uuid.UUID
	Completed    time.Time
}

type Invite struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	ExperimentID uuid.UUID
	Supervised   bool
}

type User struct {
	ID     uuid.UUID
	Email  string
	State  UserAccountRole
	Source string
}
