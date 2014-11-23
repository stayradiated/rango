package rangolib

import "os/exec"

func RunHugo() ([]byte, error) {
	hugo := exec.Command("hugo")

	output, err := hugo.Output()
	if err != nil {
		return nil, err
	}

	return output, nil
}
