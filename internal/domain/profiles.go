package domain

type (
    Profile struct {
        Username  string `json:"username"`
        Email     string `json:"email"`
        Image     string `json:"image"`
        Bio       string `json:"bio"`
        Following bool   `json:"following"`
    }

    ProfileResponse struct {
        Profile *Profile `json:"profile"`
    }
)
