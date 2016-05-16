package actions

import (
	"log"
	"os/exec"
)

func Call(c string, repo string, tag string) error {
	log.Println("reload in progress...")
	out, err := exec.Command(c, repo, tag).Output()
	if err != nil {
		if out != nil {
			log.Println(string(out))
		}
		log.Println("reload error!")
		log.Println(err)
		return err
	}
	log.Println(string(out))

	log.Printf("reload done.")

	return nil
}
