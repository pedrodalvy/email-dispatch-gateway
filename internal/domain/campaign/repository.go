package campaign

type Repository interface {
	Save(campaign *Campaign) error
	GetByID(id string) (campaign *Campaign, err error)
	Update(campaign *Campaign) error
}
