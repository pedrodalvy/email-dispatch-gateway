package campaign

type RepositoryInterface interface {
	Create(campaign *Campaign) error
	GetByID(id string) (campaign *Campaign, err error)
	Update(campaign *Campaign) error
	Delete(campaign *Campaign) error
}
