package router

import (
	"go-tools/controller"
	"go-tools/durable"
	"go-tools/repository"
	"go-tools/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(nrm gin.HandlerFunc) *gin.Engine {

	var (
		userRepository   repository.UserRepository   = repository.NewUserRepository(durable.GormDB)
		userService      service.UserService         = service.NewUserService(userRepository)
		userController   controller.UserController   = controller.NewUserController(userService)
		healthController controller.HealthController = controller.NewHealthController()
		googleService    service.OAuthService = service.NewGoogleService("470277026302-tok0dbh86nt9jrntqbs472abgmvbin7u.apps.googleusercontent.com", "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", "https://in.linkedin.com/in/mohamed-sameem-17778416b")
   		facebookService  service.OAuthService = service.NewFacebookService("facebookClientID", "facebookClientSecret", "facebookRedirectURL")
    	githubService    service.OAuthService = service.NewGithubService("githubClientID", "githubClientSecret", "githubRedirectURL")
    	loginService     *service.LoginService = service.NewLoginService(googleService, facebookService, githubService)
    	loginController  *controller.LoginController   = controller.NewLoginController(loginService)
    )

	r := gin.Default()
	grp1 := r.Group("/user/v1")
	{
		grp1.Use(nrm)
		grp1.GET("/fetch_all", userController.GetAllUsers)
		grp1.GET("/user_id/:userId", userController.GetUserByUserId)
		grp1.POST("/create_new", userController.CreateNewUser)
		grp1.PUT("/update", userController.UpdateUser)
		grp1.POST("/deletebyUserId/:userId", userController.DeleteUserById)
	}

	grp2 := r.Group("/")
	{
		grp2.GET("/health", healthController.GetHealth)
	}

	grp3 := r.Group("/")
	{
		grp3.GET("/google-login", loginController.GoogleLogin)
		grp3.GET("/google-signup", loginController.GoogleSignup)
	}

	return r
}
