package compaign

type Repository interface {
	Save(campaign *Campaign) error
}
