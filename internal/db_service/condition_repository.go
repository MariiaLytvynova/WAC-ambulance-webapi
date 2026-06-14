package db_service
//package ambulance_wl
import (
	"context"
	"database/sql"
	//"github.com/MariiaLytvynova/WAC-ambulance-webapi/internal/ambulance_wl"
)
type ConditionDB struct {
	Value string
	Code string
	Reference string
	TypicalDurationMinutes int32
}

func (s *PostgresService) GetCondition(
	ctx context.Context,
	code string,
) (*ConditionDB, error){
	var c ConditionDB //prazdna struktura, ktora sa naplni z databazy

	err := s.DB.QueryRowContext(
		ctx,
		`
		SELECT
			value,
			code,
			reference,
			typical_duration_minutes
		FROM conditions
		WHERE code = $1
		`,
		code, //ziskali sme odpoved od db
	).Scan( //naplname strukturu condition z databazy
		&c.Value,
		&c.Code,
		&c.Reference,
		&c.TypicalDurationMinutes,
	)

	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}

	if err != nil { //ak chyba - vrat error
		return nil, err
	}

	return &c, nil
}