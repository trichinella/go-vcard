# go-vcard
vCard with golang

```
file, err := os.CreateTemp("/tmp", "vcard")
defer func() {
    _ = file.Close()
}()

card := &vcard.Card{}
card.SetName("Carl Carlsen")
card.SetNickname("Like a pro")
card.AddPhone(*vcard.NewPhone("+3 123 456 78 99").Work().Cell())
card.AddAddress(*vcard.NewAddress("My sweet house").Home().Pref())
	
photo, err := vcard.NewFilePhoto("/path/to/photo")
if err != nil {
    return err
}
card.SetPhoto(photo)

err = card.Write(file)
if err != nil {
    return err
}
```