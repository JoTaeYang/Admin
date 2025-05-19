package factory

import "github.com/JoTaeYang/Admin/gpkg/repo"

type RepoFactory struct {
	identityRepo repo.IdentityRepository
	currencyRepo repo.CurrencyRepository
	authRepo     repo.AuthRepository
	profileRepo  repo.ProfileRepository
}

func NewFactory() RepoFactory {
	return RepoFactory{
		identityRepo: repo.IdentityRepository{},
		currencyRepo: repo.CurrencyRepository{},
		authRepo:     repo.AuthRepository{},
		profileRepo:  repo.ProfileRepository{},
	}
}

func (f *RepoFactory) Identity() repo.IdentityRepository {
	return f.identityRepo
}

func (f *RepoFactory) Currency() repo.CurrencyRepository {
	return f.currencyRepo
}

func (f *RepoFactory) Auth() repo.AuthRepository {
	return f.authRepo
}

func (f *RepoFactory) Profile() repo.ProfileRepository {
	return f.profileRepo
}
