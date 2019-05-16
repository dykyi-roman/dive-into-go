package Domain

import (
	"net/http"

	Domain "github.com/dykyi-roman/dive-into-go/server/Domain/Model"
)

// Interface
type DataBaseInterface interface {
	Get() []Domain.User
	Insert(r *http.Request) bool
	Update(r *http.Request) bool
	Delete(r *http.Request) bool
}
