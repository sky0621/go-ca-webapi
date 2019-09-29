package controller

import (
	"github.com/google/uuid"
	"strings"
)

func generateID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
