package service

import (
	"context"
	"errors"

	"github.com/cxptek/vdex/config"
	"github.com/cxptek/vdex/models"
	"github.com/cxptek/vdex/models/pg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AddUser(ctx context.Context, user *models.User) error {
	return pg.SharedStore().AddUser(ctx, user)
}

func GetUserByID(ctx context.Context, userID int64) (*models.User, error) {
	return pg.SharedStore().GetUserByID(ctx, userID)
}
func GetUserByPublicID(ctx context.Context, publicID uuid.UUID) (*models.User, error) {
	return pg.SharedStore().GetUserByPublicID(ctx, publicID)
}

func GetUserByAddress(ctx context.Context, address string) (*models.User, error) {
	return pg.SharedStore().GetUserByAddress(ctx, address)
}

func CheckJWT(tokenStr string) (*models.User, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("cannot convert claim to MapClaims")
	}
	if !token.Valid {
		return nil, errors.New("token is invalid")
	}

	idVal, found := claim["publicId"]
	if !found {
		return nil, errors.New("bad token")
	}
	id := idVal.(string)
	addressVal, found := claim["address"]
	if !found {
		return nil, errors.New("bad token")
	}
	address := addressVal.(string)
	user := &models.User{
		PublicID: uuid.MustParse(id),
		Address:  address,
	}

	return user, nil
}

func GetUserFromContext(c *fiber.Ctx) (*models.User, error) {
	u := c.Locals("user").(*models.User)
	u, err := GetUserByPublicID(c.Context(), u.PublicID)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetUsersByIDs(ctx context.Context, userIDs []int64) ([]models.User, error) {
	return pg.SharedStore().GetUsersByIDs(ctx, userIDs)
}
