package entity

type Message struct {
    Role    string // "system", "user", "assistant"
    Content string
}


type ChatResponse struct {
    Content string
    Raw     string // optional: raw JSON if you want debugging
}
