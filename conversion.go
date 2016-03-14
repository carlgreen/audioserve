package main

import "log"

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
	headerData := []byte("mdcl")

	data := []byte{}

	data = append(data, "mcnm"...)
	data = append(data, stringToData(contentCode.number)...)

	data = append(data, "mcna"...)
	data = append(data, stringToData(contentCode.name)...)

	data = append(data, "mcty"...)
	data = append(data, shortToData(contentCode.dmapType)...)

	headerData = append(headerData, intToByteArray(len(data))...)
	data = append(headerData, data...)

	return data
}

func databaseToData(database Database) []byte {
	headerData := []byte("mlit")

	data := []byte{}

	// assume only one database
	data = append(data, "miid"...)
	data = append(data, intToData(1)...)

	// assume only one database
	data = append(data, "mper"...)
	data = append(data, longToData(1)...)

	data = append(data, "minm"...)
	data = append(data, stringToData(database.name)...)

	data = append(data, "mimc"...)
	data = append(data, intToData(len(database.songs))...)

	// no playlist support
	data = append(data, "mctc"...)
	data = append(data, intToData(0)...)

	headerData = append(headerData, intToByteArray(len(data))...)
	data = append(headerData, data...)

	return data
}

func songToData(fields []string, song Song) []byte {
	headerData := []byte("mlit")

	data := []byte{}

	kindInd := index(fields, "dmap.itemkind")
	if kindInd > -1 {
		data = append(data, "mikd"...)
		data = append(data, charToData(2)...)

		fields = append(fields[:kindInd], fields[kindInd+1:]...)
	}

	for _, field := range fields {
		switch field {
		case "dmap.itemid":
			data = append(data, "miid"...)
			data = append(data, intToData(1)...)
		case "dmap.itemname":
			data = append(data, "minm"...)
			data = append(data, stringToData(song.Title)...)
		case "dmap.persistentid":
			data = append(data, "mper"...)
			data = append(data, longToData(1)...)
		case "daap.songalbum":
			data = append(data, "asal"...)
			data = append(data, stringToData(song.Album)...)
		case "daap.songartist":
			data = append(data, "asar"...)
			data = append(data, stringToData(song.Artist)...)
		default:
			log.Printf("unexpected field: %s", field)
		}
	}

	headerData = append(headerData, intToByteArray(len(data))...)
	data = append(headerData, data...)

	return data
}

func index(s []string, e string) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}
