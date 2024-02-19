package main

import (
	"net/http"
	"strings"
	"context"
)

type Config struct {
	Cookies []string `json:"cookies,omitempty"`
}

type FilterPlugin struct {
	next    http.Handler
	cookies []string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &FilterPlugin{
		next:    next,
		cookies: config.Cookies,
	}, nil
}

func (p *FilterPlugin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// Copia las cookies a una nueva variable
	cookies := make([]*http.Cookie, len(req.Cookies()))
	copy(cookies, req.Cookies())

	// Filtra las cookies que no estén en la lista de cookies a mantener
	filteredCookies := make([]*http.Cookie, 0)
	for _, cookie := range cookies {
		if contains(p.cookies, cookie.Name) {
			filteredCookies = append(filteredCookies, cookie)
		}
	}

	// Reemplaza las cookies en la solicitud con las cookies filtradas
	req.Header.Set("Cookie", strings.Join(cookieStrings(filteredCookies), "; "))

	// Llama al siguiente middleware o handler en la cadena
	p.next.ServeHTTP(rw, req)
}

// Función auxiliar que devuelve true si un slice de strings contiene un valor determinado
func contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

// Función auxiliar que convierte un slice de cookies en un slice de strings
func cookieStrings(cookies []*http.Cookie) []string {
	cookieStrings := make([]string, len(cookies))
	for i, cookie := range cookies {
		cookieStrings[i] = cookie.String()
	}
	return cookieStrings
}


