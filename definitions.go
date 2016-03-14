package main

type ContentCode struct {
	number   string
	name     string
	dmapType int16
}

type Version struct {
	major uint16
	minor uint8
	patch uint8
}

type Song struct {
	Title  string
	Album  string
	Artist string
}

type Database struct {
	name  string
	songs []Song
}

const DmapChar int16 = 1
const DmapShort int16 = 3
const DmapLong int16 = 5
const DmapLongLong int16 = 7
const DmapString int16 = 9
const DmapVersion int16 = 11
const DmapContainer int16 = 12
