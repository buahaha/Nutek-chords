package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"testing"

	"gitlab.com/gomidi/midi/reader"
)

const TestDirectory string = "test"

var Chord []uint8

type noteTest struct{}

func (pr noteTest) noteOn(p *reader.Position, channel, key, vel uint8) {
	// fmt.Printf("Track: %v Pos: %v NoteOn (ch %v: key %v vel: %v)\n", p.Track, p.AbsoluteTicks, channel, key, vel)
	Chord = append(Chord, key)
}

var p noteTest

// TestMajorTriad
func TestMajorTriad(t *testing.T) {

	// to disable logging, pass mid.NoLogger() as option
	rd := reader.New(reader.NoLogger(),
		// set the functions for the messages you are interested in
		reader.NoteOn(p.noteOn),
	)

	MajorTriad(C2, strconv.Itoa(1)+". "+Notes[0], TestDirectory)
	MajorTriad(C2+7, strconv.Itoa(8)+". "+Notes[7], TestDirectory)
	files, err := ioutil.ReadDir(TestDirectory)
	if err != nil {
		t.Errorf("could not read %v directory\n", TestDirectory)
	}
	for _, s := range files {
		err := reader.ReadSMFFile(rd, TestDirectory+"/"+s.Name())

		if err != nil {
			t.Errorf("could not read SMF file %v\n", s)
		}
		stringSlice := s.Name()[len(s.Name())-7:]
		expectedSuffix := "maj.mid"
		if stringSlice != expectedSuffix {
			t.Errorf("file name is wrong, %v should have suffix %v", s.Name(), expectedSuffix)
		}
		if Chord[1]-Chord[0] != 4 || Chord[2]-Chord[0] != 7 {
			t.Errorf("failed to create Major chord from notes: %v %v, %v %v, %v %v\n",
				NoteToName(Chord[0]), Chord[0],
				NoteToName(Chord[1]), Chord[1],
				NoteToName(Chord[2]), Chord[2])
		}
		Chord = nil
	}

	os.RemoveAll(TestDirectory)
}

func TestMinorTriad(t *testing.T) {

	// to disable logging, pass mid.NoLogger() as option
	rd := reader.New(reader.NoLogger(),
		// set the functions for the messages you are interested in
		reader.NoteOn(p.noteOn),
	)

	MinorTriad(C2, strconv.Itoa(0+1)+". "+Notes[0], TestDirectory)
	MinorTriad(C2+7, strconv.Itoa(0+1+7)+". "+Notes[7], TestDirectory)

	files, err := ioutil.ReadDir(TestDirectory)
	if err != nil {
		t.Errorf("could not read %v directory\n", TestDirectory)
	}
	for _, s := range files {
		err := reader.ReadSMFFile(rd, TestDirectory+"/"+s.Name())

		if err != nil {
			t.Errorf("could not read SMF file %v\n", s)
		}
		stringSlice := s.Name()[len(s.Name())-7:]
		expectedSuffix := "min.mid"
		if stringSlice != expectedSuffix {
			t.Errorf("file name is wrong, %v should have suffix %v", s.Name(), expectedSuffix)
		}
		if Chord[1]-Chord[0] != 3 || Chord[2]-Chord[0] != 7 {
			t.Errorf("failed to create Minor chord from notes: %v %v, %v %v, %v %v",
				NoteToName(Chord[0]), Chord[0],
				NoteToName(Chord[1]), Chord[1],
				NoteToName(Chord[2]), Chord[2])
		}
		Chord = nil
	}

	os.RemoveAll(TestDirectory)
}
