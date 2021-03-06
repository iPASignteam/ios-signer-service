package storage

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Profile interface {
	GetId() string
	GetCert() (ReadonlyFile, error)
	GetProv() (ReadonlyFile, error)
	GetPassword() (ReadonlyFile, error)
	GetName() (string, error)
}

func newProfile(id string) *profile {
	return &profile{id}
}

type profile struct {
	id string
}

type ProfileError struct {
	Message string
	Id      string
	Err     error
}

func (e *ProfileError) Error() string {
	return fmt.Sprintf("%s %s: %s", e.Message, e.Id, e.Err)
}

func (p *profile) GetId() string {
	return p.id
}

func (p *profile) GetCert() (ReadonlyFile, error) {
	file, err := os.Open(profileCertPath(p.id))
	if err != nil {
		return nil, &ProfileError{"open ProfilesCertPath", p.id, err}
	}
	return file, nil
}

func (p *profile) GetProv() (ReadonlyFile, error) {
	file, err := os.Open(profileProvPath(p.id))
	if err != nil {
		return nil, &ProfileError{"open ProfilesProvPath", p.id, err}
	}
	return file, nil
}

func (p *profile) GetPassword() (ReadonlyFile, error) {
	file, err := os.Open(profilePassPath(p.id))
	if err != nil {
		return nil, &ProfileError{"open ProfilesPassPath", p.id, err}
	}
	return file, nil
}

func (p *profile) GetName() (string, error) {
	bytes, err := ioutil.ReadFile(profileNamePath(p.id))
	if err != nil {
		return "", &ProfileError{"read file ProfilesNamePath", p.id, err}
	}
	return string(bytes), nil
}
