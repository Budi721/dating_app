package repository

import (
	"context"
	"github.com/Budi721/dating_app/entity"
	"github.com/Budi721/dating_app/utils/logger"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type MemberPreferenceRepo interface {
	Create(newMemberPref *entity.MemberPreferences, newMemberInterest []*entity.MemberInterest) error
	FindById(id string) (*entity.MemberPreferences, []*entity.MemberInterest, error)
}

type memberPreferenceRepo struct {
	conn *pgx.Conn
}

func (m *memberPreferenceRepo) FindById(id string) (*entity.MemberPreferences, []*entity.MemberInterest, error) {
	var pref *entity.MemberPreferences
	err := m.conn.QueryRow(context.Background(),
		"SELECT looking_gender, looking_domicile, looking_start_age, looking_end_age FROM member_preference WHERE member_id = $1",
		id,
	).Scan(
		&pref.LookingForGender,
		&pref.LookingForDomicile,
		&pref.LookingForStartAge,
		&pref.LookingForEndAge,
	)
	if err != nil {
		logger.Log.Error().Err(err).Msg("find by id preference fail")
		return nil, nil, err
	}

	rows, err := m.conn.Query(context.Background(), "SELECT interest_id FROM member_interest WHERE member_id = $1", id)
	if err != nil {
		logger.Log.Error().Err(err).Msg("find by id interest fail")
		return nil, nil, err
	}

	var interest []*entity.MemberInterest
	for rows.Next() {
		var itr *entity.MemberInterest
		err := rows.Scan(&itr.InterestId)
		if err != nil {
			return nil, nil, err
		}
		interest = append(interest, itr)
	}

	return pref, interest, nil
}

func (m *memberPreferenceRepo) Create(newMemberPref *entity.MemberPreferences, newMemberInterest []*entity.MemberInterest) error {
	logger.Log.Debug().Msgf("create preference %s", newMemberPref.MemberId)
	tx, err := m.conn.BeginTx(context.Background(), pgx.TxOptions{})
	defer func(tx pgx.Tx) {
		if err != nil {
			logger.Log.Error().Err(err).Msg("failed create preference")
			_ = tx.Rollback(context.Background())
		} else {
			_ = tx.Commit(context.Background())
		}
	}(tx)

	_, err = tx.Exec(
		context.Background(),
		"INSERT INTO member_preference VALUES($1, $2, $3, $4, $5, $6)",
		uuid.New().String(),
		newMemberPref.MemberId,
		newMemberPref.LookingForGender,
		newMemberPref.LookingForDomicile,
		newMemberPref.LookingForStartAge,
		newMemberPref.LookingForEndAge,
	)

	for _, interest := range newMemberInterest {
		_, err = tx.Exec(context.Background(),
			"INSERT INTO member_interest VALUES($1, $2)",
			interest.InterestId,
			interest.MemberId,
		)
	}

	if err != nil {
		return err
	}

	return nil
}

func NewMemberPreferenceRepo(conn *pgx.Conn) MemberPreferenceRepo {
	return &memberPreferenceRepo{
		conn: conn,
	}
}
