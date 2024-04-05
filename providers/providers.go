package providers

import (
	"github.com/pg06pg/stay-sorted/repository/user_repository"
)

// Repositories
var UserRepository = user_repository.UserRepositoryImpl{}
