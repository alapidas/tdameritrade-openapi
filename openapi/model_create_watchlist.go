/*
 * TD Ameritrade API
 *
 * TD Ameritrade API
 *
 * API version: 3.0.1
 * Contact: austin.millan@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CreateWatchlist struct for CreateWatchlist
type CreateWatchlist struct {
	Name string `json:"name,omitempty"`
	WatchlistItems []CreateWatchlistWatchlistItems `json:"watchlistItems,omitempty"`
}