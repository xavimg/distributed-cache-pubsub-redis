package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

type Store struct {
	cache Cacher
}

func NewStore(c Cacher) *Store {
	return &Store{
		cache: c,
	}
}

func (s *Store) GetUser(c echo.Context) error {
	timeStart := time.Now()

	id := c.Param("id")
	idint, _ := strconv.Atoi(id)

	var user User

	val, ok := s.cache.Get(idint)
	if ok {
		if err := s.cache.Remove(idint); err != nil {
			log.Fatal(err)
		}

		if err := json.Unmarshal(val, &user); err != nil {
			return err
		}
		timeSince := time.Since(timeStart)
		fmt.Println(fmt.Sprintf("fetching from cache...time needed %v ", timeSince))
		return c.JSON(http.StatusOK, &User{
			ID:       user.ID,
			Username: user.Username,
		})
	}

	if err := db.Where("id = ?", id).First(&user).Debug().Error; err != nil {
		log.Fatal(err)
		return err
	}

	userBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}

	if err := s.cache.Set(idint, userBytes); err != nil {
		return err
	}

	time.Sleep(time.Second * 5)
	timeSince := time.Since(timeStart)
	fmt.Println(fmt.Sprintf("fetching from db...time needed %s ", timeSince))

	return c.JSON(http.StatusOK, &User{
		ID:       user.ID,
		Username: user.Username,
	})
}
