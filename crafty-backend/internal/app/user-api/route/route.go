package route

import (
	authMiddleware "github.com/Admin-OR-1-1/2110336-SE2-Crafty/crafty-backend/internal/app/user-api/middleware/auth"
	userRoute "github.com/Admin-OR-1-1/2110336-SE2-Crafty/crafty-backend/internal/app/user-api/pkg/user"
	"github.com/Admin-OR-1-1/2110336-SE2-Crafty/crafty-backend/internal/repository"
	"github.com/gofiber/fiber/v2"
)

type UserAPI struct {
	r *repository.Repositories
}

func SetupUserAPI(repos *repository.Repositories) *UserAPI {
	// userAPI := userapi.NewUserAPI(repos.UserRepository)
	// userAPI.SetupRoute()
	return &UserAPI{r: repos}
}

func (api *UserAPI) SetupRoute(ro fiber.Router) error {

	authMW := authMiddleware.NewAuthMiddleware()
	userRouter := userRoute.NewUserRouter(api.r)
	userRouter.SetupRoute(ro.Group("/user", authMW.VerifyToken))
	return nil
}
