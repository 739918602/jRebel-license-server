package handler

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
)

func UUID(w http.ResponseWriter, req *http.Request) {
	uuid := uuid.NewV4().String()
	fmt.Printf("UUIDv4: %s\n", uuid)
	w.Write([]byte(uuid))
}
