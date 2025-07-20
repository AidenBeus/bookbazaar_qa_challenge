package main 

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type book struct{
	ID     string  `json:"id"`
	Title  string
	Author string

}