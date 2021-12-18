package dockervol_test

import (
	"testing"

	"github.com/atjhoendz/dockervol/dockervol"
)

var (
	mockListVolume = []string{
		"d1690bb8c291024afa32c2df9945015346f426068185a84b5248445c34a24f19",
		"some_db_data",
		"d1690bb8c291024afa32c2df9945015346f426068185a84b5248445c34a24f191110",
		"ca39795842781ff69a7e49f1131ce4530937c483824ec04509b0236e4e24bbba",
	}
	mockDuplicateListVolume = []string{
		"d1690bb8c291024afa32c2df9945015346f426068185a84b5248445c34a24f19",
		"some_db_data",
		"d1690bb8c291024afa32c2df9945015346f426068185a84b5248445c34a24f19",
		"ca39795842781ff69a7e49f1131ce4530937c483824ec04509b0236e4e24bbba",
	}
)

func Test_ParseValidVolume(t *testing.T) {
	var parsedVolume dockervol.DockerVolume
	err := dockervol.ParseVolume("fffb099b33e6aefee6ad039c068a031318169e003dd7c02d035f43e62fc20282", &parsedVolume)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(parsedVolume)

	if parsedVolume.Name == "" {
		t.Fatal("Parsed volume name cannot be empty")
	}
}

func Test_ParseInvalidVolumeName(t *testing.T) {
	var parsedVolume dockervol.DockerVolume
	err := dockervol.ParseVolume("wrongvolumename_thatdoesn'texist", &parsedVolume)

	if err == nil {
		t.Fatal("Must return error because volume doesn't exist")
	}

	if parsedVolume.Name != "" {
		t.Fatal("Parsed invalid volume name must be empty")
	}
}

func Test_GetAnonymousVolume(t *testing.T) {
	anonymousVolumes := dockervol.GetAnonymousVolume(mockListVolume)

	if anonymousVolumes == nil {
		t.Fatal("The result is invalid")
	}

	if len(*anonymousVolumes) != 2 {
		t.Fatal("Must return data with length 2")
	}

	if (*anonymousVolumes)[0] != mockListVolume[0] {
		t.Fatal("The result is wrong")
	}
}

func Test_GetAnonymousVolume_DuplicateData(t *testing.T) {
	anonymousVolume := dockervol.GetAnonymousVolume(mockDuplicateListVolume)

	if anonymousVolume == nil {
		t.Fatal("The result is invalid")
	}

	if len(*anonymousVolume) != 2 {
		t.Fatal("Should be return with length 2")
	}

	if (*anonymousVolume)[1] != mockDuplicateListVolume[3] {
		t.Fatal("The result is wrong")
	}
}
