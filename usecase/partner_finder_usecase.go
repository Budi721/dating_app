package usecase

import (
	"github.com/Budi721/dating_app/entity"
	"github.com/Budi721/dating_app/repository"
	"github.com/Budi721/dating_app/utils/base64"
	"github.com/Budi721/dating_app/utils/read"
)

type PartnerFinderUseCase interface {
	ViewPartner(id string, pageNo int, pageSize int) ([]*entity.MemberPersonalInformation, error)
	MatchPartner(memberId string, partnerId string) error
	ListPartner(memberId string) ([]*entity.MemberPersonalInformation, error)
}

type partnerFinderUseCase struct {
	partnerRepo repository.PartnerRepo
	prefRepo    repository.MemberPreferenceRepo
}

func (p *partnerFinderUseCase) ListPartner(memberId string) ([]*entity.MemberPersonalInformation, error) {
	return p.partnerRepo.FindAll(memberId)
}

func (p *partnerFinderUseCase) MatchPartner(memberId string, partnerId string) error {
	return p.partnerRepo.Create(memberId, partnerId)
}

func (p *partnerFinderUseCase) ViewPartner(id string, pageNo int, pageSize int) ([]*entity.MemberPersonalInformation, error) {
	preferences, intrs, err := p.prefRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	var interests []string
	for _, intr := range intrs {
		interests = append(interests, intr.InterestId)
	}
	partners, err := p.partnerRepo.Find(id, preferences.LookingForGender, preferences.LookingForDomicile, preferences.LookingForStartAge, preferences.LookingForEndAge, pageSize, pageNo)
	for _, partner := range partners {
		buf := read.ReadFile(partner.RecentPhotoPath)
		imgBase64Str := base64.ToBase64(buf)
		partner.RecentPhotoPath = imgBase64Str
	}

	return partners, nil
}

func NewPartnerFinderUseCase(partnerRepo repository.PartnerRepo, prefRepo repository.MemberPreferenceRepo) PartnerFinderUseCase {
	return &partnerFinderUseCase{
		partnerRepo: partnerRepo,
		prefRepo:    prefRepo,
	}
}
