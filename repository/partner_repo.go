package repository

import (
	"context"
	"errors"
	"github.com/Budi721/dating_app/entity"
	"github.com/Budi721/dating_app/utils/logger"
	"github.com/jackc/pgx/v4"
)

type PartnerRepo interface {
	Find(
		id string,
		byGender string,
		byDomicile string,
		byStartAge int,
		byEndAge int,
		limit int,
		offset int,
	) ([]*entity.MemberPersonalInformation, error)
	Create(id string, id2 string) error
	FindAll(id string) ([]*entity.MemberPersonalInformation, error)
}

type partnerRepo struct {
	conn *pgx.Conn
}

func (p *partnerRepo) FindAll(id string) ([]*entity.MemberPersonalInformation, error) {
	logger.Log.Debug().Msgf("Find all member id %s", id)
	rows, err := p.conn.Query(context.Background(),
		"SELECT pi.member_id, pi.name, pi.photo_path, pi.self_description FROM member_personal_information pi JOIN ( SELECT partner_id FROM member_partner WHERE member_id = $1 ) p ON pi.member_id = p.partner_id", id)
	if err != nil {
		return nil, err
	}
	var members []*entity.MemberPersonalInformation
	for rows.Next() {
		var member *entity.MemberPersonalInformation
		err := rows.Scan(&member.MemberId, &member.Name, &member.RecentPhotoPath, &member.SelfDescription)
		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}

	return members, nil
}

func (p *partnerRepo) Create(id string, id2 string) error {
	logger.Log.Debug().Msgf("Create new patrner for %s and %s", id, id2)
	res, err := p.conn.Exec(context.Background(), "INSERT INTO member_partner VALUES($1, $2)", id, id2)
	if err != nil {
		logger.Log.Error().Err(err).Msg("failed insert member partner")
		return err
	}
	if res.RowsAffected() == 0 {
		logger.Log.Error().Msgf("no row affected in member partner")
		return errors.New("failed member partner")
	}

	return nil
}

func (p *partnerRepo) Find(id string, byGender string, byDomicile string, byStartAge int, byEndAge int, limit int, offset int) ([]*entity.MemberPersonalInformation, error) {
	logger.Log.Debug().Msgf("Finding partner %v", id)
	rows, err := p.conn.Query(
		context.Background(),
		"SELECT pi.member_id, pi.name, pi.gender, pi.bod, pi.photo_path, pi.self_description FROM member_personal_information pi JOIN member_preference p ON pi.member_id = p.member_id AND date_part('year', age(now(), pi.bod)) BETWEEN $4 AND $5 AND p.member_id != $1 AND p.looking_gender <> $2 AND p.looking_domicile = $3 ORDER BY pi.member_id LIMIT $6 OFFSET $7",
		id, byGender, byDomicile, byStartAge, byEndAge, limit, offset,
	)
	if err != nil {
		return nil, err
	}

	var member []*entity.MemberPersonalInformation
	for rows.Next() {
		var m *entity.MemberPersonalInformation
		if err := rows.Scan(&m.MemberId, &m.Name, &m.Gender, &m.Bod, &m.RecentPhotoPath, &m.SelfDescription); err != nil {
			return nil, err
		}
		member = append(member, m)
	}

	return member, nil
}

func NewPartnerRepo(conn *pgx.Conn) PartnerRepo {
	return &partnerRepo{
		conn: conn,
	}
}
