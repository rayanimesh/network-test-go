/*
 * DNS Service
 *
 * Sevice for managing DNS with route53
 *
 * API version: 0.0.1
 * Contact: iitr.animesh@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	_ "embed"
	"log"
	"net/http"

	route53dns "github.com/GIT_USER_ID/GIT_REPO_ID/internal/route53dns"

	"github.com/flowchartsman/swaggerui"
)

//go:embed api/openapi.yaml
var spec []byte

func main() {

	DefaultAPIService := route53dns.NewDefaultAPIService()
	DefaultAPIController := route53dns.NewDefaultAPIController(DefaultAPIService)

	router := route53dns.NewRouter(DefaultAPIController)

	router.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger", swaggerui.Handler(spec)))
	log.Println("serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
