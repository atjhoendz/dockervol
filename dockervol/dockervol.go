package dockervol

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

type DockerVolume struct {
	Name      string
	Driver    string
	CreatedAt time.Time
}

// TODO: add more accurate identification
func getAnonymousVolume(listVolume []string) *[]string {
	keys := make(map[string]bool)
	var anonVolume []string

	for _, volume := range listVolume {
		// check duplicate volume name
		if _, val := keys[volume]; !val {
			keys[volume] = true

			if len(volume) == 64 {
				anonVolume = append(anonVolume, volume)
			}
		}
	}

	return &anonVolume
}

func parseVolume(volumeName string, parsedVolume *DockerVolume) error {
	byteVolumeData, err := exec.Command("docker", "volume", "inspect", volumeName).Output()

	if err != nil {
		return err
	}

	var arrVolumeData []DockerVolume
	err = json.Unmarshal(byteVolumeData, &arrVolumeData)

	if err != nil {
		return err
	}

	*parsedVolume = arrVolumeData[0]

	return nil
}

func parseVolumes(volumes []string, parsedVolumes *[]DockerVolume) {
	for i, volume := range volumes {
		var parsedVolume DockerVolume
		err := parseVolume(volume, &parsedVolume)
		if err != nil {
			log.Println(err)
		} else {
			*parsedVolumes = append(*parsedVolumes, parsedVolume)
		}

		fmt.Printf("Parsed volume %d of %d\n", i, len(volumes))
	}
}

func execRemoveListVolume(volumes []string) {
	for i, volume := range volumes {
		_, err := exec.Command("docker", "volume", "rm", volume).Output()

		if err != nil {
			log.Printf("Volume with name %s cannot be removed, caused by %v\n", volume, err)
		} else {
			log.Printf("Volume %s removed successfully | Progress: %d of %d volume removed\n", volume, i+1, len(volumes))
		}
	}
}

func RemoveAnonymousVolume() error {
	byteVolumes, err := exec.Command("docker", "volume", "ls", "--format", `{{.Name}}`).Output()

	if err != nil {
		return err
	}

	volumes := string(byteVolumes)

	listVolume := strings.Split(volumes, "\n")

	anonymousVolumes := getAnonymousVolume(listVolume)

	execRemoveListVolume(*anonymousVolumes)

	return nil
}
