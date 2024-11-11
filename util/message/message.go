package message

import "fmt"

func NotValidError(field string) string {
	return fmt.Sprintf(`Field { %s } is not valid`, field)
}

func InternalServerError() string {
	return fmt.Sprintf("Internal server error")
}

func DuplicateAccountError() string {
	return fmt.Sprintf("Duplicated account")
}