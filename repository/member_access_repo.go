package repository

import (
	"context"
	"errors"
	"github.com/Budi721/dating_app/entity"
	"github.com/Budi721/dating_app/utils/logger"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type MemberAccessRepo interface {
	UpdateVerification(id string) error
	Create(user *entity.MemberUserAccess) error
	FindByUsernamePasswordVerified(username string, password string) (*entity.MemberUserAccess, error)
}

type memberAccessRepo struct {
	conn *pgx.Conn
}

func (m *memberAccessRepo) FindByUsernamePasswordVerified(username string, password string) (*entity.MemberUserAccess, error) {
	logger.Log.Debug().Msgf("Login user and password %s: %s", username, password)
	var memberAccess entity.MemberUserAccess
	err := m.conn.QueryRow(
		context.Background(),
		"SELECT member_id FROM member_access WHERE verification_status = 'Y' AND user_name = $1 AND user_password = $2",
		username, password,
	).Scan(&memberAccess.MemberId)
	if err != nil {
		return nil, err
	}

	return &memberAccess, nil
}

func (m *memberAccessRepo) UpdateVerification(id string) error {
	logger.Log.Debug().Msgf("User id to activation %s", id)
	tx, err := m.conn.BeginTx(context.Background(), pgx.TxOptions{})
	defer func(tx pgx.Tx) {
		if err != nil {
			logger.Log.Error().Err(err).Msg("failed update verification")
			_ = tx.Rollback(context.Background())
		} else {
			_ = tx.Commit(context.Background())
		}
	}(tx)

	res, err := tx.Exec(context.Background(), "UPDATE member_access SET verification_status = 'Y' WHERE member_id = $1", id)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return errors.New("update verification failed")
	}

	var email string
	err = tx.QueryRow(context.Background(), "SELECT user_name FROM member_access WHERE member_id = $1", id).Scan(&email)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		context.Background(),
		"INSERT INTO member_contact_information(contact_information_id, member_id, email) VALUES($1, $2, $3)",
		uuid.New().String(), id, email,
	)

	return nil
}

func (m *memberAccessRepo) Create(user *entity.MemberUserAccess) error {
	logger.Log.Debug().Msgf("Create %s", user.Username)
	res, err := m.conn.Exec(
		context.Background(),
		"INSERT INTO member_access VALUES($1, $2, $3, $4, $5)",
		user.MemberId,
		user.Username,
		user.Password,
		user.JoinDate,
		user.VerificationStatus,
	)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed insert member access")
		return err
	}

	if res.RowsAffected() == 0 {
		logger.Log.Error().Msg("Failed insert member access")
		return errors.New("insert member failed")
	}

	return nil
}

func NewMemberAccessRepo(conn *pgx.Conn) MemberAccessRepo {
	return &memberAccessRepo{
		conn: conn,
	}
}
