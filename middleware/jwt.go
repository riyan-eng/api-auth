package middleware

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtMiddleware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/riyan-eng/api-auth/config"
)

// for middleware
func JWTProtected() func(*fiber.Ctx) error {
	config := jwtMiddleware.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET_KEY")),
		ContextKey:   "jwt",
		ErrorHandler: jwtError,
	}

	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "bad",
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"data":    err.Error(),
		"message": "bad",
	})
}

// generate
type AccessToken struct {
	Token   string
	Expired int64
}

func GenerateNewAccessToken(userID string, role []string) (*AccessToken, error) {
	accessToken := new(AccessToken)
	secret := os.Getenv("JWT_SECRET_KEY")
	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTE_COUNT"))
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()
	claims["user_id"] = userID
	claims["role"] = role
	accessToken.Expired = claims["exp"].(int64)

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}
	accessToken.Token = t
	return accessToken, nil
}

// parser
type TokenMetadata struct {
	Expires int64
	UserID  string
	Role    []interface{}
}

func ExtractTokenMetadata(bearToken string) (*TokenMetadata, error) {
	token, err := verifyToken(bearToken)
	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		expires := int64(claims["exp"].(float64))
		userID := claims["user_id"].(string)
		role := claims["role"].([]interface{})
		// fmt.Println(role)
		// fmt.Println(reflect.TypeOf(role))
		return &TokenMetadata{
			Expires: expires,
			UserID:  userID,
			Role:    role,
		}, nil
	}
	return nil, err
}

func extractToken(bearToken string) string {
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}
	return ""
}

func verifyToken(bearToken string) (*jwt.Token, error) {
	tokenString := extractToken(bearToken)
	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET_KEY")), nil
}

type RefreshToken struct {
	Token   string
	UserID  string
	Expired int64
}

func GenerateNewRefreshToken(userID string) (*RefreshToken, error) {
	var err error
	refreshToken := new(RefreshToken)
	refreshToken.Token = uuid.NewString()
	refreshToken.UserID = userID

	validUntil := time.Now().Add(time.Minute * 30).Format(time.RFC3339)

	// query
	queryDeleteToken := fmt.Sprintf(`
		delete from management.refresh_token where user_id='%s'
	`, userID)
	queryInsertToken := fmt.Sprintf(`
		insert into management.refresh_token(id, user_id, valid_until) values('%s', '%s', '%v')
	`, refreshToken.Token, userID, validUntil)

	// transaction
	ctx := context.Background()
	tx, _ := config.DB.BeginTx(ctx, nil)
	defer tx.Rollback()

	_, err = tx.Exec(queryDeleteToken)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(queryInsertToken)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return refreshToken, nil
}

type RefreshTokenMetaData struct {
	Valid  bool
	UserID string
}

func ValidRefreshToken(bearToken string) (*RefreshTokenMetaData, error) {
	refreshTokenMetaData := new(RefreshTokenMetaData)

	token := extractToken(bearToken)
	var expiry time.Time
	var userID string
	query := fmt.Sprintf(`
		select rt.user_id, rt.valid_until from management.refresh_token rt where rt.id='%s'
	`, token)
	err := config.DB.QueryRow(query).Scan(&userID, &expiry)

	if err == sql.ErrNoRows {
		return refreshTokenMetaData, errors.New("token can't use")
	} else if err != nil {
		return refreshTokenMetaData, err
	}

	if time.Now().Before(expiry) {
		refreshTokenMetaData.Valid = true
		refreshTokenMetaData.UserID = userID
		return refreshTokenMetaData, nil
	}

	return refreshTokenMetaData, errors.New("token expired")
}
