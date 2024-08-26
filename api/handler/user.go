package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "api_gateway/genproto/user"
)

// @Summary      Get user by email
// @Description  This endpoint retrieves user details by email.
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        email  path      string  true  "User Email"
// @Success      200    {object}  user.GetUserResponse
// @Failure      500    {object}  string
// @Router       /user/getbyuser/{email} [get]
func (h Handler) GetUSerByEmail(c *gin.Context) {
	req := pb.GetUSerByEmailReq{
		Email: c.Param("email"),
	}

	res, err := h.UserService.GetUSerByEmail(c, &req)
	if err != nil {
		h.Log.Error("GetUSerByEmail funksiyasida xatolik.", "error", err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}

// @Summary      Update user details
// @Description  This endpoint updates the user's details based on the provided information.
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user  body      user.UpdateUserRequest  true  "User Update Data"
// @Success      200    {object}  user.UpdateUserRespose
// @Failure      400    {object}  string
// @Failure      500    {object}  string
// @Router       /user/update_user [put]
func (h Handler) UpdateUser(c *gin.Context) {
	req := pb.UpdateUserRequest{}

	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	res, err := h.UserService.UpdateUser(c, &req)
	if err != nil {
		h.Log.Error("UpdateUser funksiyasoga xabar yuborishda xatolik", "error", err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}

// @Summary      Delete user
// @Description  This endpoint deletes a user based on the provided user ID.
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id  path      string  true  "User ID"
// @Success      200    {object}  user.DeleteUserr
// @Failure      400    {object}  string
// @Failure      500    {object}  string
// @Router       /user/delete_user/{id} [delete]
func (h Handler) DeleteUser(c *gin.Context) {
	req := pb.UserId{
		Id: c.Param("id"),
	}

	res, err := h.UserService.DeleteUser(c, &req)
	if err != nil {
		h.Log.Error("DeleteUserga malumot yuborishda xatolik", "error", err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}
