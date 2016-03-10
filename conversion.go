package main

func shortToByteArray(i int16) []byte {
	data := [2]byte{
		byte((i >> 8) & 0xFF),
		byte(i & 0xFF),
	}
	return data[:]
}

func intToByteArray(i int) []byte {
	data := [4]byte{
		byte((i >> 24) & 0xFF),
		byte((i >> 16) & 0xFF),
		byte((i >> 8) & 0xFF),
		byte(i & 0xFF),
	}
	return data[:]
}

func longToByteArray(l int64) []byte {
	data := [8]byte{
		byte((l >> 56) & 0xFF),
		byte((l >> 48) & 0xFF),
		byte((l >> 40) & 0xFF),
		byte((l >> 32) & 0xFF),
		byte((l >> 24) & 0xFF),
		byte((l >> 16) & 0xFF),
		byte((l >> 8) & 0xFF),
		byte(l & 0xFF),
	}
	return data[:]
}

func stringToData(s string) []byte {
	data := intToByteArray(len(s))
	data = append(data, s...)
	return data
}

func charToData(b byte) []byte {
	data := intToByteArray(1)
	data = append(data, b)
	return data
}

func shortToData(i int16) []byte {
	data := intToByteArray(2)
	data = append(data, shortToByteArray(i)...)
	return data
}

func intToData(i int) []byte {
	data := intToByteArray(4)
	data = append(data, intToByteArray(i)...)
	return data
}

func longToData(i int64) []byte {
	data := intToByteArray(8)
	data = append(data, longToByteArray(i)...)
	return data
}

func versionToData(version Version) []byte {
	data := intToByteArray(4)
	versionData := [4]byte{
		byte((version.major >> 8) & 0xFF),
		byte(version.major & 0xFF),
		byte(version.minor & 0xFF),
		byte(version.patch & 0xFF),
	}
	data = append(data, versionData[:]...)
	return data
}

func contentCodeToData(contentCode ContentCode) []byte {
	data := []byte{}

	data = append(data, "mdcl"...)
	data = append(data, intToByteArray(12+8+len(contentCode.name)+10)...)

	data = append(data, "mcnm"...)
	data = append(data, stringToData(contentCode.number)...)

	data = append(data, "mcna"...)
	data = append(data, stringToData(contentCode.name)...)

	data = append(data, "mcty"...)
	data = append(data, shortToData(contentCode.dmapType)...)

	return data
}

func listingItemToData(listingItem ListingItem) []byte {
	data := []byte{}

	data = append(data, "mlit"...)
	data = append(data, intToByteArray(12+16+8+len(listingItem.itemName)+12+12)...)

	data = append(data, "miid"...)
	data = append(data, intToData(listingItem.itemId)...)

	data = append(data, "mper"...)
	data = append(data, longToData(listingItem.persistentId)...)

	data = append(data, "minm"...)
	data = append(data, stringToData(listingItem.itemName)...)

	data = append(data, "mimc"...)
	data = append(data, intToData(listingItem.itemCount)...)

	data = append(data, "mctc"...)
	data = append(data, intToData(listingItem.containerCount)...)

	return data
}
