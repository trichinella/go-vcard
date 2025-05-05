package vcard

type AddressType string

const (
	AddressTypeDom    AddressType = "dom"    //domestic delivery address
	AddressTypeINTL   AddressType = "intl"   //indicate an international delivery address
	AddressTypePostal AddressType = "postal" //indicate a postal delivery address
	AddressTypeVoice  AddressType = "parcel" //indicate a parcel delivery address
	AddressTypeHome   AddressType = "home"   //indicate a delivery address for a residence
	AddressTypeWork   AddressType = "work"   //indicate delivery address for a place of work
	AddressTypePref   AddressType = "pref"   //indicate the preferred delivery address when more than one address is specified
)

type AddressField struct {
	addressRaw   string
	addressTypes []AddressType
}

func NewAddress(address string) *AddressField {
	return &AddressField{
		addressRaw: address,
	}
}

func (f *AddressField) Dom() *AddressField {
	f.addressTypes = append(f.addressTypes, AddressTypeDom)

	return f
}

func (f *AddressField) INTL() *AddressField {
	f.addressTypes = append(f.addressTypes, AddressTypeINTL)

	return f
}

func (f *AddressField) Postal() *AddressField {
	f.addressTypes = append(f.addressTypes, AddressTypePostal)

	return f
}

func (f *AddressField) Voice() *AddressField {
	f.addressTypes = append(f.addressTypes, AddressTypeVoice)

	return f
}

func (f *AddressField) Home() *AddressField {
	f.addressTypes = append(f.addressTypes, AddressTypeHome)

	return f
}

func (f *AddressField) Work() *AddressField {
	f.addressTypes = append(f.addressTypes, AddressTypeWork)

	return f
}

func (f *AddressField) Pref() *AddressField {
	f.addressTypes = append(f.addressTypes, AddressTypePref)

	return f
}
