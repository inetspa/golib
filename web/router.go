package web

import (
	"fmt"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

// Query Uid
func GetQueryUid(c echo.Context, queryName string, uid *uuid.UUID) error {
	return extractUid(c.Param(queryName), uid)
}

func extractUid(uidString string, uid *uuid.UUID) error {
	var err error
	if *uid, err = uuid.FromString(uidString); err != nil {
		return err
	}
	return nil
}

// Middleware router skipper
type SkipperPath struct {
	Prefix string
	Paths  map[string]bool
}

func (s *SkipperPath) Add(path string, method string) {
	s.Paths[s.key(path, method)] = true
}

func (s *SkipperPath) Delete(path string, method string) {
	delete(s.Paths, s.key(path, method))
}

func (s *SkipperPath) key(path string, method string) string {
	return fmt.Sprintf("%s%s", method, path)
}

func (s *SkipperPath) TestContext(c echo.Context) bool {
	if active, ok := s.Paths[s.key(c.Path(), c.Request().Method)]; ok && active {
		return true
	}
	return false
}

func NewSkipper(prefix string) SkipperPath {
	return SkipperPath{
		Prefix: prefix,
		Paths:  map[string]bool{},
	}
}
