package user

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type UserRequests struct {
}

func (r UserRequests) UserCreate(ctx *gin.Context, UserDB IUserDB) error {
	UserDB.CreateUser(ctx.MustGet("UserId").(string), 1)
	return nil
}
func (r UserRequests) UserGet(ctx *gin.Context, UserDB IUserDB) (*User, error) {
	User, err := UserDB.GetUser(ctx.MustGet("UserId").(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": utils.InternalErrorDB})
		return nil, err
	}
	return User, nil
}
func (r UserRequests) UserDelete(ctx *gin.Context, UserDB IUserDB) error {
	err := UserDB.DeleteUser(ctx.MustGet("UserId").(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": utils.InternalErrorDB})
		return err
	}
	return nil
}
func (r UserRequests) UserProfileUpdate(ctx *gin.Context, UserDB IUserDB, UserUtils IUserUtils) error {
	payload, err := utils.GetPayloadFromBody(ctx, &UserProfile{})
	if err != nil {
		return err
	}
	updatePayload, err := UserUtils.InspectProfileUpdatePayload(*payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.InternalErrorInvalidPayload})
		return err
	}
	err = UserDB.UpdateProfile(ctx.MustGet("UserId").(string), updatePayload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": utils.InternalErrorDB})
		return err
	}
	return nil
}
func (r UserRequests) UserAddressRegister(ctx *gin.Context, UserDB IUserDB) error {
	payload, err := utils.GetPayloadFromBody(ctx, &UserAddressRegisterPayload{})
	if err != nil {
		return err
	}
	err = UserDB.RegisterAddress(ctx.MustGet("UserId").(string), *payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": utils.InternalErrorDB})
		return err
	}
	return nil
}
func (r UserRequests) UserAddressUpdate(ctx *gin.Context, UserDB IUserDB, UserUtils IUserUtils) error {
	payload, err := utils.GetPayloadFromBody(ctx, &UserAddress{})
	if err != nil {
		return err
	}
	updatePayload, err := UserUtils.InspectAddressUpdatePayload(*payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.InternalErrorInvalidPayload})
		return err
	}
	err = UserDB.UpdateAddress(ctx.MustGet("UserId").(string), updatePayload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": utils.InternalErrorDB})
		return err
	}
	return nil
}
