package installer

import "fmt"

func Install(application string) error {
	fmt.Printf("Installing %s...\n", application)
	return nil
}
