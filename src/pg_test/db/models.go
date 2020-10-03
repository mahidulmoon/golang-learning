package db



type User struct {
    Id     int64  `json:"id"`
    Name   string `json:"name"`
    Email  string `json:"email"`
}

// type About struct {
//     about string `json: about`
// }