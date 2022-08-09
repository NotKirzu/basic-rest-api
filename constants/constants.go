package constants

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	DNI       string `json:"dni"`
	ID        int    `json:"id"`
}

func RemoveIndex(user *[]User, i int) {
	*user = append((*user)[:i], (*user)[i+1:]...)
}

func GetIndexFromID(users []User, id int) (int, bool) {
	var index int
	var found bool
	for j := range users {
		if users[j].ID == id {
			index = j
			found = true
			break
		}
	}

	return index, found
}
