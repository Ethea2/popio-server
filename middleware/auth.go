package middleware

import (
	"os"

	"github.com/Ethea2/nat-server/utils"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Protected() func()
