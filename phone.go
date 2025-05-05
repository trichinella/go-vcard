package vcard

type PhoneType string

const (
	HOME  PhoneType = "HOME"  //indicate a telephone number associated with a residence +
	WORK  PhoneType = "WORK"  //indicate a telephone number associated with a place of work +
	PREF  PhoneType = "PREF"  //indicate a preferred-use telephone number
	VOICE PhoneType = "VOICE" //indicate a voice telephone number
	FAX   PhoneType = "FAX"   //indicate a facsimile telephone number
	MSG   PhoneType = "MSG"   //indicate the telephone number has voice messaging support
	CELL  PhoneType = "CELL"  //indicate a cellular telephone number +
	PAGER PhoneType = "PAGER" //indicate a paging device telephone number +
	BBS   PhoneType = "BBS"   //indicate a bulletin board system telephone number
	MODEM PhoneType = "MODEM" //indicate a MODEM connected telephone number
	CAR   PhoneType = "CAR"   //indicate a car-phone telephone number
	ISDN  PhoneType = "ISDN"  //indicate an ISDN service telephone number
	VIDEO PhoneType = "VIDEO" //indicate a video conferencing telephone number
	PCS   PhoneType = "PCS"   //indicate a personal communication services telephone number
)

type PhoneField struct {
	number     string
	phoneTypes []PhoneType
}

func NewPhone(number string) *PhoneField {
	return &PhoneField{
		number: number,
	}
}

func (f *PhoneField) Home() *PhoneField {
	f.phoneTypes = append(f.phoneTypes, HOME)

	return f
}

func (f *PhoneField) Work() *PhoneField {
	f.phoneTypes = append(f.phoneTypes, WORK)

	return f
}

func (f *PhoneField) Pref() *PhoneField {
	f.phoneTypes = append(f.phoneTypes, PREF)

	return f
}

func (f *PhoneField) Voice() *PhoneField {
	f.phoneTypes = append(f.phoneTypes, VOICE)

	return f
}

func (f *PhoneField) Fax() *PhoneField {
	f.phoneTypes = append(f.phoneTypes, FAX)

	return f
}

func (f *PhoneField) Msg() *PhoneField {
	f.phoneTypes = append(f.phoneTypes, MSG)

	return f
}

func (f *PhoneField) Cell() *PhoneField {
	f.phoneTypes = append(f.phoneTypes, CELL)

	return f
}

func (f *PhoneField) Pager() *PhoneField {
	f.phoneTypes = append(f.phoneTypes, PAGER)

	return f
}

func (f *PhoneField) Bbs() *PhoneField {
	f.phoneTypes = append(f.phoneTypes, BBS)

	return f
}

func (f *PhoneField) Modem() *PhoneField {
	f.phoneTypes = append(f.phoneTypes, MODEM)

	return f
}

func (f *PhoneField) Car() *PhoneField {
	f.phoneTypes = append(f.phoneTypes, CAR)

	return f
}

func (f *PhoneField) ISDN() *PhoneField {
	f.phoneTypes = append(f.phoneTypes, ISDN)

	return f
}

func (f *PhoneField) Video() *PhoneField {
	f.phoneTypes = append(f.phoneTypes, VIDEO)

	return f
}

func (f *PhoneField) PCS() *PhoneField {
	f.phoneTypes = append(f.phoneTypes, PCS)

	return f
}
