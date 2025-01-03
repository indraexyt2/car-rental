package cmd

import (
	"car-rental-ums/helpers"
	"car-rental-ums/internal/api"
	"car-rental-ums/internal/interfaces"
	"car-rental-ums/internal/repository"
	"car-rental-ums/internal/services"
	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	d := DependencyInject()
	r := gin.Default()

	userV1 := r.Group("/api/user/v1")
	userV1.POST("/register", d.RegisterAPI.Register)
	userV1.PUT("/email-verify/:token", d.EmailVerifyAPI.EmailVerify)
	userV1.GET("/resend-email-verify", d.ResendEmailVerifyAPI.ResendEmailVerify)
	userV1.POST("/login", d.LoginApi.Login)

	userV1Auth := userV1.Use(d.MiddlewareValidateAuthToken)
	userV1Auth.DELETE("/logout", d.LogoutAPI.Logout)

	err := r.Run(":" + helpers.GetEnv("APP_PORT"))
	if err != nil {
		helpers.Logger.Fatal("failed to run server: ", err)
	}
}

type Dependency struct {
	UserRepository       interfaces.IUserRepository
	RegisterAPI          interfaces.IRegisterAPI
	EmailVerifyAPI       interfaces.IEmailVerifyAPI
	LoginApi             interfaces.ILoginAPI
	ResendEmailVerifyAPI interfaces.IResendEmailVerifyAPI
	LogoutAPI            interfaces.ILogoutAPI
}

func DependencyInject() *Dependency {
	userRepo := &repository.UserRepo{DB: helpers.DB}

	registerSvc := &services.RegisterService{UserRepository: userRepo}
	registerApi := &api.RegisterAPI{RegisterSVC: registerSvc}

	verifyEmailSvc := &services.EmailVerifyService{UserRepo: userRepo}
	verifyEmailApi := &api.EmailVerifyAPI{EmailVerifySVC: verifyEmailSvc}

	loginSvc := &services.LoginService{UserRepo: userRepo}
	loginApi := &api.LoginAPI{LoginSVC: loginSvc}

	resendVerifyEmailSvc := &services.ResendEmailVerifyService{UserRepo: userRepo}
	resendVerifyEmailApi := &api.ResendEmailVerifyAPI{ResendEmailVerifySVC: resendVerifyEmailSvc}

	logoutSvc := &services.LogoutService{UserRepository: userRepo}
	logoutApi := &api.LogoutAPI{LogoutSVC: logoutSvc}

	return &Dependency{
		UserRepository:       userRepo,
		RegisterAPI:          registerApi,
		EmailVerifyAPI:       verifyEmailApi,
		LoginApi:             loginApi,
		ResendEmailVerifyAPI: resendVerifyEmailApi,
		LogoutAPI:            logoutApi,
	}
}
