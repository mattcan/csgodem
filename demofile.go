package main

import (
	//"os"
)

type DemoHeader struct {
	demoFileStamp   string
	demoProtocol    int32
	networkProtocol int32
	serverName      string
	clientName      string
	mapName         string
	gameDirectory   string
	playbackTime    float32
	playbackTicks   int32
	playbackFrames  int32
	signonLength    int32
}

type DemoFile struct {
	FileBuffer    string
	fileBufferPos int

	FileName   string
	DemoHeader DemoHeader
}

func (d *DemoFile) Open(fileName string) bool {
	d.Close() // reset the structure

	return false
}

func (d *DemoFile) Close() {
	d.FileName = ""

	d.fileBufferPos = 0
	d.FileBuffer = ""
}

func (d *DemoFile) ReadRawData(buffer []byte, length int32) int32 {
	return 0
}

func (d *DemoFile) ReadSequenceInfo(seqNrIn *int32, seqNrOutAck *int32) {
}

func (d *DemoFile) ReadCmdInfo(info *DemoCmdInfo) {
}

func (d *DemoFile) ReadCmdHeader(cmd *string, tick *int32, playerSlot *string) {
}

/*func (d *DemoFile) ReadDemoHeader() DemoHeader {
	return nil
}*/
