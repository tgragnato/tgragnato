---
title: mdadm Recovery
---

```go
package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	input := []string{
		"/dev/sda",
		"/dev/sdb1",
		"/dev/sdd",
		"/dev/sde",
		"/dev/sdf",
		"missing",
	}
	permutations := generatePermutations(input)

	for _, perm := range permutations {
		if perm[4] != "missing" {
			continue
		}

		cmd := exec.Command(
			"mdadm", "--create", "/dev/md0",
			"--level=6", "--raid-devices=6",
			perm[0], perm[1], perm[2], perm[3], perm[4], perm[5],
			"--assume-clean", "--readonly",
		)
		log.Println(cmd)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				log.Printf("Command mdadm create finished with error (%d): \n\t%v\n", exitError.ExitCode(), exitError)
			} else {
				log.Println(err.Error())
			}
		}

		cmd = exec.Command(
			"mount", "/dev/md0", "/root/data",
		)
		log.Println(cmd)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				log.Printf("Command mount finished with error (%d): \n\t%v\n", exitError.ExitCode(), exitError)
			} else {
				log.Println(err.Error())
			}
		} else {
			log.Println("WOW")
			log.Println(perm)
			return
		}

		cmd = exec.Command(
			"mdadm", "--stop", "/dev/md0",
		)
		log.Println(cmd)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				log.Printf("Command mdadm stop finished with error (%d): \n\t%v\n", exitError.ExitCode(), exitError)
			} else {
				log.Println(err.Error())
			}
			return
		}
	}
}

func generatePermutations(arr []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
```