package main
import(
  "encoding/json"
)


/**
 *
 */
type Actions struct {
    Type string `json:"type"`
    To  []string `json:"to"`
    Body string `json:"body"`
}
/**
 *
 */
type TrackerResponse struct {
    Data   map[string]interface{}   `json:"data"`
    Actions  []Actions `json:"actions"`
}

/**
 * Convert a string to a map of interface
 */
func getJson (message []byte) TrackerResponse{
  var trackerResponse TrackerResponse
  err := json.Unmarshal(message, &trackerResponse)
  if err != nil {
      panic(err)
  }
  return trackerResponse
}
