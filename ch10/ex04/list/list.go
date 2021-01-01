package list

import (
	"bytes"
	"encoding/json"
	"os/exec"
)

type Package struct {
	ImportPath string   `json:"ImportPath"`
	Deps       []string `json:"Deps"`
}

func Get(pkgstr string) ([]Package, error) {
	out, err := exec.Command("go", "list", "-json", pkgstr).Output()
	if err != nil {
		return nil, err
	}

	pkgs := make([]Package, 0)
	decoder := json.NewDecoder(bytes.NewBuffer(out))
	for {
		var pkg Package
		if err := decoder.Decode(&pkg); err != nil {
			break
		}

		pkgs = append(pkgs, pkg)
	}

	return pkgs, nil
}
