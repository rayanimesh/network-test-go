/*
 * DNS Service
 *
 * Sevice for managing DNS with route53
 *
 * API version: 0.0.1
 * Contact: iitr.animesh@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package route53dns

import (
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
