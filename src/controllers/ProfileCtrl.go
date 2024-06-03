package controllers

import "github.com/gin-gonic/gin"

// #region Profile Controller

func GetProfile(context *gin.Context) {

}

func CreateProfile(context *gin.Context) {

}

func UpdateProfile(context *gin.Context) {

}

func DeleteProfile(context *gin.Context) {

}

// #endregion

// #region Photos Controller

func GetPhotos(context *gin.Context) {

}

func CreatePhoto(context *gin.Context) {

}

func UpdatePhoto(context *gin.Context) {

}

func DeletePhoto(context *gin.Context) {

}

// #endregion

func ProfileController(r *gin.Engine) {
	profile := r.Group("/profile")
	{
		profile.GET("/me", GetProfile)
		profile.GET("/:id", GetProfile)
		profile.POST("/", CreateProfile)
		profile.PUT("/:id", UpdateProfile)
		profile.DELETE("/:id", DeleteProfile)

		photos := profile.Group("/photos")
		{
			photos.GET("/", GetPhotos)
			photos.POST("/", CreatePhoto)
			photos.PUT("/:id", UpdatePhoto)
			photos.DELETE("/:id", DeletePhoto)
		}
	}
}
