package city

type City struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Cities []City