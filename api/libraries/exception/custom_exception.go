package exception

import "errors"

func NotExistException() error {
	return errors.New("data not exist")
}

func SomeProductNotAvailableException() error  {
	return errors.New("some product not available")
}

func ProductNotAvailableException() error  {
	return errors.New("all product not available")
}

func ProductMustInSameStore() error {
	return errors.New("product must in same store")
}