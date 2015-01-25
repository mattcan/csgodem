package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"os"
	unsafe "unsafe"
)

const (
	DEMO_HEADER_ID string = "HL2DEMO"
	DEMO_PROTOCOL  int32  = 4
)

type DemoHeader struct {
	demoFileStamp   [8]byte
	demoProtocol    int32
	networkProtocol int32
	serverName      [260]byte
	clientName      [260]byte
	mapName         [260]byte
	gameDirectory   [260]byte
	playbackTime    float32
	playbackTicks   int32
	playbackFrames  int32
	signonLength    int32
}

// FileStampString removes NUL values first
func (dh *DemoHeader) FileStampString() string {
	return string(bytes.Trim(dh.demoFileStamp[:], "\x00"))
}

type DemoFile struct {
	FileBuffer    string
	fileBufferPos int

	FileName   string
	DemoHeader DemoHeader
}

func (d *DemoFile) Open(fileName string) bool {
	d.Close() // reset the structure

	//open file
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get length of file
	fi, err := f.Stat()
	if err != nil {
		log.Fatal(err)
		return false
	}

	length := fi.Size()

	// check size
	// http://stackoverflow.com/questions/23202864/assigning-a-type-uintptr-to-uint64-in-golang
	hdrSize := (int64)(unsafe.Pointer(unsafe.Sizeof(d.DemoHeader)))
	if length < hdrSize {
		log.Fatal("File is too small")
		return false
	}

	// fread?
	f.Seek(0, 0) // go back to the beginning of the file
	reader := bufio.NewReader(f)
	hdrBytes, err := reader.Peek(int(hdrSize))
	if err != nil {
		log.Fatal(err)
		return false
	}

	d.fillDemoHeader(hdrBytes)
	//d.debugHeader()

	// reduce the length var based on size of demoheader
	length -= hdrSize

	// check demofilestamp matches demo headerid
	if d.DemoHeader.FileStampString() != DEMO_HEADER_ID {
		log.Fatal("File stamp doesn't match")
	}

	// check demoprotocol is valid
	if d.DemoHeader.demoProtocol != DEMO_PROTOCOL {
		log.Fatal("Demo protocol is invalid")
		return false
	}

	d.fileBufferPos = 0
	d.FileName = fileName

	return true
}

func (d *DemoFile) debugHeader() {
	fmt.Println(string(d.DemoHeader.demoFileStamp[:]))
	fmt.Println(d.DemoHeader.demoProtocol)
	fmt.Println(d.DemoHeader.networkProtocol)
	fmt.Println(string(d.DemoHeader.serverName[:]))
	fmt.Println(string(d.DemoHeader.clientName[:]))
	fmt.Println(string(d.DemoHeader.mapName[:]))
	fmt.Println(string(d.DemoHeader.gameDirectory[:]))
	fmt.Println(d.DemoHeader.playbackTime)
	fmt.Println(d.DemoHeader.playbackTicks)
	fmt.Println(d.DemoHeader.playbackFrames)
	fmt.Println(d.DemoHeader.signonLength)
}

func (d *DemoFile) fillDemoHeader(header []byte) {
	var newHeader []byte = header

	// get the demo files stamp
	copy(d.DemoHeader.demoFileStamp[:], newHeader[0:7])

	// get demo protocol
	d.DemoHeader.demoProtocol = byteSliceToInt32(newHeader[8:12])

	// get protocol version
	d.DemoHeader.networkProtocol = byteSliceToInt32(newHeader[13:17])

	// servername, clientname, mapname, directory
	copy(d.DemoHeader.serverName[:], newHeader[16:275])
	copy(d.DemoHeader.clientName[:], newHeader[276:535])
	copy(d.DemoHeader.mapName[:], newHeader[536:795])
	copy(d.DemoHeader.gameDirectory[:], newHeader[796:1055])

	// playback
	d.DemoHeader.playbackTime = byteSliceToFloat32(newHeader[1056:1060])
	d.DemoHeader.playbackTicks = byteSliceToInt32(newHeader[1061:1065])
	d.DemoHeader.playbackTicks = byteSliceToInt32(newHeader[1066:1070])
	d.DemoHeader.signonLength = byteSliceToInt32(newHeader[1071:1075])
}

func byteSliceToInt32(data []byte) int32 {
	var result int32
	buf := bytes.NewBuffer(data)
	err := binary.Read(buf, binary.LittleEndian, &result)
	if err != nil {
		panic(err)
	}
	return result
}

func byteSliceToFloat32(data []byte) float32 {
	bits := binary.LittleEndian.Uint32(data)
	result := math.Float32frombits(bits)
	return result
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
