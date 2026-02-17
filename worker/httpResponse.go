package worker

type ErrResponse struct {
	HTTPStatusCode int    // Use int for status codes
	Message        string // Exported field (Capitalized)
}
