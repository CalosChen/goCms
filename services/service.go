package services

import "fmt"

type Service struct {
}

func (Service) GetConfig() {

	fmt.Println("getting config...")

}
