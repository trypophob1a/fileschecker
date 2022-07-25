package interfaces

type CheckFinder interface {
	SetResources(first, second string)
	Find(percent uint8, callback func(uniqueFilename string))
}

type SelfCheckFinder interface {
	Find(percent uint8, callback func(duplicateFilename string))
	SetResource(fileTxt string)
}
