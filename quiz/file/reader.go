package file

import "errors"

func ReadFile(questions int) ([]string, error) {

	if(questions <= 0){
		return make([]string,0), errors.New("Pick at least 1 question")
	}
	return make([]string, 10), nil
}