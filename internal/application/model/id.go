package model

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Id uuid.UUID

func NewId() (id Id, err error) {
	uuid, err := uuid.NewV7()
	if err != nil {
		return id, err
	}
	id = Id(uuid)
	return id, nil
}

func ScanId(value any) (Id, error) {
	id := uuid.UUID{}
	err := id.Scan(value)
	if err != nil {
		return Id{}, err
	}
	return Id(id), nil
}

func (id Id) String() string {
	return uuid.UUID(id).String()
}

func (id Id) PgId() pgtype.UUID {
	return pgtype.UUID{
		Valid: true,
		Bytes: uuid.UUID(id),
	}
}
