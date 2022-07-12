package interfaces

type Finder interface {
	SetResources(first, second string)
	Find(percent uint8, callback func(filename string))
}
