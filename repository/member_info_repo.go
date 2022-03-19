package repository

import (
	"context"
	"errors"
	"github.com/Budi721/dating_app/entity"
	"github.com/Budi721/dating_app/utils/logger"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type MemberInfoRepo interface {
	Create(member *entity.Member) (*entity.Member, error)
	FindById(memberId string) (*entity.Member, error)
	Update(path string, id string) error
}

type memberInfoRepo struct {
	conn *pgx.Conn
}

func (m *memberInfoRepo) FindById(memberId string) (*entity.Member, error) {
	logger.Log.Debug().Msgf("find by id %s", memberId)
	var member *entity.Member

	err := m.conn.QueryRow(context.Background(), `
		SELECT pi.bod, pi.gender, pi.member_id, pi.self_description,
			pi.name, pi.photo_path, ci.instagram_id, ci.twitter_id,
			ai.city
		FROM member_personal_information pi 
		JOIN member_address_information ai 
			ON pi.member_id = ai.member_id 
		JOIN member_contact_information ci 
			ON pi.member_id = ci.member_id
		WHERE pi.member_id = $1
	`, memberId).Scan(
		&member.PersonalInfo.Bod,
		&member.PersonalInfo.Gender,
		&member.PersonalInfo.MemberId,
		&member.PersonalInfo.SelfDescription,
		&member.PersonalInfo.Name,
		&member.PersonalInfo.RecentPhotoPath,
		&member.ContactInfo.InstagramId,
		&member.ContactInfo.TwitterId,
		&member.AddressInfo.City,
	)

	if err != nil {
		logger.Log.Error().Err(err).Msg("failed get profile")
		return nil, err
	}

	return member, nil
}

func (m *memberInfoRepo) Update(path string, id string) error {
	logger.Log.Debug().Msgf("update profile %s %s", path, id)
	res, err := m.conn.Exec(context.Background(),
		"UPDATE member_personal_information SET photo_path = $2 WHERE member_id = $1", id, path)
	if err != nil {
		logger.Log.Error().Err(err).Msg("failed upload profile")
		return err
	}

	if res.RowsAffected() == 0 {
		logger.Log.Error().Msgf("no row affected %s", path)
		return errors.New("failed upload photo")
	}

	return nil
}

func (m *memberInfoRepo) Create(member *entity.Member) (*entity.Member, error) {
	logger.Log.Debug().Msgf("Create %s", member.PersonalInfo.MemberId)
	tx, err := m.conn.BeginTx(context.Background(), pgx.TxOptions{})
	defer func(tx pgx.Tx) {
		if err != nil {
			logger.Log.Error().Err(err).Msg("failed update profile")
			_ = tx.Rollback(context.Background())
		} else {
			_ = tx.Commit(context.Background())
		}
	}(tx)

	_, err = tx.Exec(context.Background(),
		"INSERT INTO member_personal_information(personal_information_id, member_id, bod, gender, photo_path, self_description) VALUES($1, $2, $3, $4, $5, $6)",
		uuid.New().String(), member.PersonalInfo.MemberId, member.PersonalInfo.Bod, member.PersonalInfo.Gender, member.PersonalInfo.RecentPhotoPath, member.PersonalInfo.SelfDescription,
	)

	_, err = tx.Exec(context.Background(),
		"INSERT INTO member_address_information(address_information_id, member_id, address, city, postal_code) VALUES($1, $2, $3, $4, $5)",
		uuid.New().String(), member.PersonalInfo.MemberId, member.AddressInfo.Address, member.AddressInfo.City, member.AddressInfo.PostalCode,
	)

	_, err = tx.Exec(context.Background(),
		"UPDATE member_contact_information SET mobile_phone = $2, instagram_id = $3, twitter_id = $4 WHERE member_id = $1",
		member.PersonalInfo.MemberId, member.ContactInfo.MobilePhoneNumber, member.ContactInfo.InstagramId, member.ContactInfo.TwitterId,
	)

	if err != nil {
		logger.Log.Error().Err(err).Msg("failed update profile")
		return nil, err
	}

	return member, nil
}

func NewMemberInfoRepo(conn *pgx.Conn) MemberInfoRepo {
	return &memberInfoRepo{
		conn: conn,
	}
}
