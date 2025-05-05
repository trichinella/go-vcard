package vcard

import (
	"io"
	"strings"
	"time"
)

type Card struct {
	name         string
	nickname     string
	role         string
	title        string
	birthdate    time.Time
	phones       []PhoneField
	organization string
	photo        Photo
	addresses    []AddressField
}

func (c *Card) SetName(name string) {
	c.name = name
}

func (c *Card) SetNickname(nickname string) {
	c.nickname = nickname
}

func (c *Card) SetBirthdate(birthdate time.Time) {
	c.birthdate = birthdate
}

func (c *Card) SetRole(role string) {
	c.role = role
}

func (c *Card) AddAddress(address AddressField) {
	c.addresses = append(c.addresses, address)
}

func (c *Card) AddPhone(phone PhoneField) {
	c.phones = append(c.phones, phone)
}

func (c *Card) SetOrganization(organization string) {
	c.organization = organization
}

func (c *Card) SetTitle(title string) {
	c.title = title
}

func (c *Card) SetPhoto(photo Photo) {
	c.photo = photo
}

func (c *Card) Write(dst io.Writer) error {
	_, err := dst.Write([]byte(StartFile + "\n"))
	if err != nil {
		return err
	}

	if c.name != "" {
		_, err = dst.Write([]byte(`FN` + ":" + c.name + "\n"))
	}
	if c.nickname != "" {
		_, err = dst.Write([]byte(`NICKNAME` + ":" + c.nickname + "\n"))
	}
	if !c.birthdate.IsZero() {
		_, err = dst.Write([]byte(`BDAY` + ":" + c.birthdate.Format("2006-01-02") + "\n"))
	}
	if len(c.phones) != 0 {
		for _, phone := range c.phones {
			if len(phone.phoneTypes) == 0 {
				continue
			}
			phoneRow := `TEL;TYPE=`
			for _, phoneType := range phone.phoneTypes {
				phoneRow += string(phoneType) + ","
			}
			phoneRow = strings.TrimRight(phoneRow, ",")
			phoneRow += ":" + phone.number + "\n"
			_, err = dst.Write([]byte(phoneRow))
		}
	}
	if c.organization != "" {
		_, err = dst.Write([]byte(`ORG` + ":" + c.organization + "\n"))
	}
	if c.role != "" {
		_, err = dst.Write([]byte(`ROLE` + ":" + c.role + "\n"))
	}
	if c.title != "" {
		_, err = dst.Write([]byte(`TITLE` + ":" + c.title + "\n"))
	}
	if len(c.addresses) != 0 {
		for _, address := range c.addresses {
			if len(address.addressTypes) == 0 {
				continue
			}
			addressRow := `ADR;TYPE=`
			for _, phoneType := range address.addressTypes {
				addressRow += string(phoneType) + ","
			}
			addressRow = strings.TrimRight(addressRow, ",")
			addressRow += ":" + address.addressRaw + "\n"
			_, err = dst.Write([]byte(addressRow))
		}
	}
	if c.photo != nil {
		_, err = dst.Write([]byte(c.photo.String()))
	}

	_, err = dst.Write([]byte(EndFile + "\n"))

	return err
}
