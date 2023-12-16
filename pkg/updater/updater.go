package updater

import "fmt"

func Update(application string) error {
	fmt.Printf("Updating %s...\n", application)
	return nil
}
