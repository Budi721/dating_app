package entity

type MemberPreferences struct {
	PreferenceId       string
	MemberId           string
	LookingForGender   string
	LookingForDomicile string
	LookingForStartAge int
	LookingForEndAge   int
}
