package helper

import(
	"errors"
	"github.com/gin-gonic/gin"
)

func CheckUserType(ctx *gin.Context, userTypeRole string)(err error){
	userType :=  ctx.GetString("user_type")
	err = nil

	if userType != userTypeRole{
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	return err
}

func MatchUserTypeToUserId(ctx *gin.Context, userId string) (err error){
	userType :=  ctx.GetString("user_type")
	uid := ctx.GetString("id")

	err=nil

	if userType =="USER" && uid != userId {
		err = errors.New("Unauthorized to access this resource")
		return err
	}

	err = CheckUserType(ctx, userType)
	return err
}