package Nutek

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
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

var startingNote string = "C3"

func getStartingNoteNum(noteString string) uint8 {
	noteNum, err := StringToNote(startingNote)
	if err != nil {
		log.Fatalf("can't convert string %v to note", startingNote)
	}
	return noteNum
}

var noteNum uint8 = getStartingNoteNum(startingNote)

func TestOneNote(t *testing.T) {

	// to disable logging, pass mid.NoLogger() as option
	rd := reader.New(reader.NoLogger(),
		// set the functions for the messages you are interested in
		reader.NoteOn(p.noteOn),
	)

	OneNote(noteNum, TestDirectory)
	OneNote(noteNum+7, TestDirectory)
	files, err := ioutil.ReadDir(TestDirectory)
	if err != nil {
		t.Errorf("could not read %v directory\n", TestDirectory)
	}
	if len(files) <= 0 {
		t.Errorf("no files created")
	}
	for _, s := range files {
		err := reader.ReadSMFFile(rd, TestDirectory+"/"+s.Name())

		if err != nil {
			t.Errorf("could not read SMF file %v\n", s)
		}
		if strings.HasPrefix(s.Name(), "C") {
			if !strings.Contains(s.Name(), "solo") {
				t.Errorf("file name is wrong, %v should contain %v\n", s.Name(), "solo")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			} else if !strings.Contains(s.Name(), "C3") {
				t.Errorf("wrong note name, should be C3")
			}
		} else if strings.HasPrefix(s.Name(), "G") {
			if !strings.Contains(s.Name(), "solo") {
				t.Errorf("file name is wrong, %v should contain %v\n", s.Name(), "solo")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			}
		} else {
			t.Errorf("wrong filename %v\n", s.Name())
		}
		if Chord[0]-Chord[0] != 0 {
			t.Errorf("failed to create solo from note: %v %v\n",
				NoteToNameAndOctave(Chord[0]), Chord[0])
		}
		Chord = nil
	}

	os.RemoveAll(TestDirectory)
}

// TestMajorTriad
func TestMajorTriad(t *testing.T) {

	// to disable logging, pass mid.NoLogger() as option
	rd := reader.New(reader.NoLogger(),
		// set the functions for the messages you are interested in
		reader.NoteOn(p.noteOn),
	)

	MajorTriad(noteNum, TestDirectory)
	MajorTriad(noteNum+7, TestDirectory)
	files, err := ioutil.ReadDir(TestDirectory)
	if err != nil {
		t.Errorf("could not read %v directory\n", TestDirectory)
	}
	if len(files) <= 0 {
		t.Errorf("no files created")
	}
	for _, s := range files {
		err := reader.ReadSMFFile(rd, TestDirectory+"/"+s.Name())

		if err != nil {
			t.Errorf("could not read SMF file %v\n", s)
		}
		if strings.HasPrefix(s.Name(), "C") {
			if !strings.Contains(s.Name(), "maj") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "maj")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "E") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "E")
			} else if !strings.Contains(s.Name(), "G") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "G")
			} else if !strings.Contains(s.Name(), "C3") {
				t.Errorf("wrong note name, should be C3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "E3") {
				t.Errorf("wrong note name, should be E3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			}
		} else if strings.HasPrefix(s.Name(), "G") {
			if !strings.Contains(s.Name(), "maj") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "maj")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "B") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "B")
			} else if !strings.Contains(s.Name(), "D") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "D")
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "B3") {
				t.Errorf("wrong note name, should be B3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "D4") {
				t.Errorf("wrong note name, should be D4 have %v", s.Name())
			}
		} else {
			t.Errorf("wrong filename %v\n", s.Name())
		}
		if Chord[1]-Chord[0] != 4 || Chord[2]-Chord[0] != 7 {
			t.Errorf("failed to create Major chord from notes: %v %v, %v %v, %v %v\n",
				NoteToNameAndOctave(Chord[0]), Chord[0],
				NoteToNameAndOctave(Chord[1]), Chord[1],
				NoteToNameAndOctave(Chord[2]), Chord[2])
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

	MinorTriad(noteNum, TestDirectory)
	MinorTriad(noteNum+7, TestDirectory)

	files, err := ioutil.ReadDir(TestDirectory)
	if err != nil {
		t.Errorf("could not read %v directory\n", TestDirectory)
	}
	if len(files) <= 0 {
		t.Errorf("no files created")
	}
	for _, s := range files {
		err := reader.ReadSMFFile(rd, TestDirectory+"/"+s.Name())

		if err != nil {
			t.Errorf("could not read SMF file %v\n", s)
		}

		if strings.HasPrefix(s.Name(), "C") {
			if !strings.Contains(s.Name(), "min") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "min")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "D#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "D#")
			} else if !strings.Contains(s.Name(), "G") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "G")
			} else if !strings.Contains(s.Name(), "C3") {
				t.Errorf("wrong note name, should be C3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "D#3") {
				t.Errorf("wrong note name, should be D#3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			}
		} else if strings.HasPrefix(s.Name(), "G") {
			if !strings.Contains(s.Name(), "min") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "min")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "A#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "A#")
			} else if !strings.Contains(s.Name(), "D") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "D")
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "A#3") {
				t.Errorf("wrong note name, should be A#3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "D4") {
				t.Errorf("wrong note name, should be D4 have %v", s.Name())
			}
		} else {
			t.Errorf("wrong filename %v\n", s.Name())
		}
		if Chord[1]-Chord[0] != 3 || Chord[2]-Chord[0] != 7 {
			t.Errorf("failed to create Minor chord from notes: %v %v, %v %v, %v %v",
				NoteToNameAndOctave(Chord[0]), Chord[0],
				NoteToNameAndOctave(Chord[1]), Chord[1],
				NoteToNameAndOctave(Chord[2]), Chord[2])
		}
		Chord = nil
	}

	os.RemoveAll(TestDirectory)
}

func TestDiminished(t *testing.T) {

	// to disable logging, pass mid.NoLogger() as option
	rd := reader.New(reader.NoLogger(),
		// set the functions for the messages you are interested in
		reader.NoteOn(p.noteOn),
	)

	Diminished(noteNum, TestDirectory)
	Diminished(noteNum+7, TestDirectory)
	files, err := ioutil.ReadDir(TestDirectory)
	if err != nil {
		t.Errorf("could not read %v directory\n", TestDirectory)
	}
	if len(files) <= 0 {
		t.Errorf("no files created")
	}
	for _, s := range files {
		err := reader.ReadSMFFile(rd, TestDirectory+"/"+s.Name())

		if err != nil {
			t.Errorf("could not read SMF file %v\n", s)
		}
		if strings.HasPrefix(s.Name(), "C") {
			if !strings.Contains(s.Name(), "dim") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "dim")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "D#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "D#")
			} else if !strings.Contains(s.Name(), "F#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "F#")
			} else if !strings.Contains(s.Name(), "C3") {
				t.Errorf("wrong note name, should be C3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "D#3") {
				t.Errorf("wrong note name, should be D#3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "F#3") {
				t.Errorf("wrong note name, should be F#3 have %v", s.Name())
			}
		} else if strings.HasPrefix(s.Name(), "G") {
			if !strings.Contains(s.Name(), "dim") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "dim")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "A#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "A#")
			} else if !strings.Contains(s.Name(), "C#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "C#")
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "A#3") {
				t.Errorf("wrong note name, should be A#3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "C#4") {
				t.Errorf("wrong note name, should be C#4 have %v", s.Name())
			}
		} else {
			t.Errorf("wrong filename %v\n", s.Name())
		}
		if Chord[1]-Chord[0] != 3 || Chord[2]-Chord[0] != 6 {
			t.Errorf("failed to create Diminished chord from notes: %v %v, %v %v, %v %v\n",
				NoteToNameAndOctave(Chord[0]), Chord[0],
				NoteToNameAndOctave(Chord[1]), Chord[1],
				NoteToNameAndOctave(Chord[2]), Chord[2])
		}
		Chord = nil
	}

	os.RemoveAll(TestDirectory)
}

func TestAugmented(t *testing.T) {

	// to disable logging, pass mid.NoLogger() as option
	rd := reader.New(reader.NoLogger(),
		// set the functions for the messages you are interested in
		reader.NoteOn(p.noteOn),
	)

	Augmented(noteNum, TestDirectory)
	Augmented(noteNum+7, TestDirectory)
	files, err := ioutil.ReadDir(TestDirectory)
	if err != nil {
		t.Errorf("could not read %v directory\n", TestDirectory)
	}
	if len(files) <= 0 {
		t.Errorf("no files created")
	}
	for _, s := range files {
		err := reader.ReadSMFFile(rd, TestDirectory+"/"+s.Name())

		if err != nil {
			t.Errorf("could not read SMF file %v\n", s)
		}
		if strings.HasPrefix(s.Name(), "C") {
			if !strings.Contains(s.Name(), "aug") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "aug")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "E") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "E")
			} else if !strings.Contains(s.Name(), "G#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "G#")
			} else if !strings.Contains(s.Name(), "C3") {
				t.Errorf("wrong note name, should be C3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "E3") {
				t.Errorf("wrong note name, should be E3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "G#3") {
				t.Errorf("wrong note name, should be G#3 have %v", s.Name())
			}
		} else if strings.HasPrefix(s.Name(), "G") {
			if !strings.Contains(s.Name(), "aug") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "aug")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "B") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "B")
			} else if !strings.Contains(s.Name(), "D#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "D#")
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "B3") {
				t.Errorf("wrong note name, should be B3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "D#4") {
				t.Errorf("wrong note name, should be D#4 have %v", s.Name())
			}
		} else {
			t.Errorf("wrong filename %v\n", s.Name())
		}
		if Chord[1]-Chord[0] != 4 || Chord[2]-Chord[0] != 8 {
			t.Errorf("failed to create Augmented chord from notes: %v %v, %v %v, %v %v\n",
				NoteToNameAndOctave(Chord[0]), Chord[0],
				NoteToNameAndOctave(Chord[1]), Chord[1],
				NoteToNameAndOctave(Chord[2]), Chord[2])
		}
		Chord = nil
	}

	os.RemoveAll(TestDirectory)
}

func TestSus2(t *testing.T) {

	// to disable logging, pass mid.NoLogger() as option
	rd := reader.New(reader.NoLogger(),
		// set the functions for the messages you are interested in
		reader.NoteOn(p.noteOn),
	)

	Sus2(noteNum, TestDirectory)
	Sus2(noteNum+7, TestDirectory)
	files, err := ioutil.ReadDir(TestDirectory)
	if err != nil {
		t.Errorf("could not read %v directory\n", TestDirectory)
	}
	if len(files) <= 0 {
		t.Errorf("no files created")
	}
	for _, s := range files {
		err := reader.ReadSMFFile(rd, TestDirectory+"/"+s.Name())

		if err != nil {
			t.Errorf("could not read SMF file %v\n", s)
		}
		if strings.HasPrefix(s.Name(), "C") {
			if !strings.Contains(s.Name(), "sus2") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "sus2")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "D") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "D")
			} else if !strings.Contains(s.Name(), "G") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "G")
			} else if !strings.Contains(s.Name(), "C3") {
				t.Errorf("wrong note name, should be C3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "D3") {
				t.Errorf("wrong note name, should be D3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			}
		} else if strings.HasPrefix(s.Name(), "G") {
			if !strings.Contains(s.Name(), "sus2") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "sus2")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "A") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "A")
			} else if !strings.Contains(s.Name(), "D") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "D")
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "A3") {
				t.Errorf("wrong note name, should be A3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "D4") {
				t.Errorf("wrong note name, should be D4 have %v", s.Name())
			}
		} else {
			t.Errorf("wrong filename %v\n", s.Name())
		}
		if Chord[1]-Chord[0] != 2 || Chord[2]-Chord[0] != 7 {
			t.Errorf("failed to create Sus2 chord from notes: %v %v, %v %v, %v %v\n",
				NoteToNameAndOctave(Chord[0]), Chord[0],
				NoteToNameAndOctave(Chord[1]), Chord[1],
				NoteToNameAndOctave(Chord[2]), Chord[2])
		}
		Chord = nil
	}

	os.RemoveAll(TestDirectory)
}

func TestSus4(t *testing.T) {

	// to disable logging, pass mid.NoLogger() as option
	rd := reader.New(reader.NoLogger(),
		// set the functions for the messages you are interested in
		reader.NoteOn(p.noteOn),
	)

	Sus4(noteNum, TestDirectory)
	Sus4(noteNum+7, TestDirectory)
	files, err := ioutil.ReadDir(TestDirectory)
	if err != nil {
		t.Errorf("could not read %v directory\n", TestDirectory)
	}
	if len(files) <= 0 {
		t.Errorf("no files created")
	}
	for _, s := range files {
		err := reader.ReadSMFFile(rd, TestDirectory+"/"+s.Name())

		if err != nil {
			t.Errorf("could not read SMF file %v\n", s)
		}
		if strings.HasPrefix(s.Name(), "C") {
			if !strings.Contains(s.Name(), "sus4") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "sus4")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "F") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "F")
			} else if !strings.Contains(s.Name(), "G") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "G")
			} else if !strings.Contains(s.Name(), "C3") {
				t.Errorf("wrong note name, should be C3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "F3") {
				t.Errorf("wrong note name, should be F3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			}
		} else if strings.HasPrefix(s.Name(), "G") {
			if !strings.Contains(s.Name(), "sus4") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "sus4")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "C") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "C")
			} else if !strings.Contains(s.Name(), "D") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "D")
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "C4") {
				t.Errorf("wrong note name, should be C4 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "D4") {
				t.Errorf("wrong note name, should be D4 have %v", s.Name())
			}
		} else {
			t.Errorf("wrong filename %v\n", s.Name())
		}
		if Chord[1]-Chord[0] != 5 || Chord[2]-Chord[0] != 7 {
			t.Errorf("failed to create Sus4 chord from notes: %v %v, %v %v, %v %v\n",
				NoteToNameAndOctave(Chord[0]), Chord[0],
				NoteToNameAndOctave(Chord[1]), Chord[1],
				NoteToNameAndOctave(Chord[2]), Chord[2])
		}
		Chord = nil
	}

	os.RemoveAll(TestDirectory)
}

func TestMajor7th(t *testing.T) {

	// to disable logging, pass mid.NoLogger() as option
	rd := reader.New(reader.NoLogger(),
		// set the functions for the messages you are interested in
		reader.NoteOn(p.noteOn),
	)

	Major7th(noteNum, TestDirectory)
	Major7th(noteNum+7, TestDirectory)
	files, err := ioutil.ReadDir(TestDirectory)
	if err != nil {
		t.Errorf("could not read %v directory\n", TestDirectory)
	}
	if len(files) <= 0 {
		t.Errorf("no files created")
	}
	for _, s := range files {
		err := reader.ReadSMFFile(rd, TestDirectory+"/"+s.Name())

		if err != nil {
			t.Errorf("could not read SMF file %v\n", s)
		}
		if strings.HasPrefix(s.Name(), "C") {
			if !strings.Contains(s.Name(), "maj7") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "maj7")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "E") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "F")
			} else if !strings.Contains(s.Name(), "G") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "G")
			} else if !strings.Contains(s.Name(), "B") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "B")
			} else if !strings.Contains(s.Name(), "C3") {
				t.Errorf("wrong note name, should be C3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "E3") {
				t.Errorf("wrong note name, should be E3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "B3") {
				t.Errorf("wrong note name, should be B3 have %v", s.Name())
			}
		} else if strings.HasPrefix(s.Name(), "G") {
			if !strings.Contains(s.Name(), "maj7") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "maj7")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "B") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "B")
			} else if !strings.Contains(s.Name(), "D") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "D")
			} else if !strings.Contains(s.Name(), "F#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "F#")
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "B3") {
				t.Errorf("wrong note name, should be B3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "D4") {
				t.Errorf("wrong note name, should be D4 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "F#4") {
				t.Errorf("wrong note name, should be F#4 have %v", s.Name())
			}
		} else {
			t.Errorf("wrong filename %v\n", s.Name())
		}
		if Chord[1]-Chord[0] != 4 || Chord[2]-Chord[0] != 7 || Chord[3]-Chord[0] != 11 {
			t.Errorf("failed to create Major7th chord from notes: %v %v, %v %v, %v %v, %v %v\n",
				NoteToNameAndOctave(Chord[0]), Chord[0],
				NoteToNameAndOctave(Chord[1]), Chord[1],
				NoteToNameAndOctave(Chord[2]), Chord[2],
				NoteToNameAndOctave(Chord[3]), Chord[3])
		}
		Chord = nil
	}

	os.RemoveAll(TestDirectory)
}

func TestMinor7th(t *testing.T) {

	// to disable logging, pass mid.NoLogger() as option
	rd := reader.New(reader.NoLogger(),
		// set the functions for the messages you are interested in
		reader.NoteOn(p.noteOn),
	)

	Minor7th(noteNum, TestDirectory)
	Minor7th(noteNum+7, TestDirectory)
	files, err := ioutil.ReadDir(TestDirectory)
	if err != nil {
		t.Errorf("could not read %v directory\n", TestDirectory)
	}
	if len(files) <= 0 {
		t.Errorf("no files created")
	}
	for _, s := range files {
		err := reader.ReadSMFFile(rd, TestDirectory+"/"+s.Name())

		if err != nil {
			t.Errorf("could not read SMF file %v\n", s)
		}
		if strings.HasPrefix(s.Name(), "C") {
			if !strings.Contains(s.Name(), "min7") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "min7")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "D#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "D#")
			} else if !strings.Contains(s.Name(), "G") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "G")
			} else if !strings.Contains(s.Name(), "A#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "A#")
			} else if !strings.Contains(s.Name(), "C3") {
				t.Errorf("wrong note name, should be C3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "D#3") {
				t.Errorf("wrong note name, should be D#3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "A#3") {
				t.Errorf("wrong note name, should be A#3 have %v", s.Name())
			}
		} else if strings.HasPrefix(s.Name(), "G") {
			if !strings.Contains(s.Name(), "min7") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "min7")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "A#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "A#")
			} else if !strings.Contains(s.Name(), "D") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "D")
			} else if !strings.Contains(s.Name(), "F") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "F")
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "A#3") {
				t.Errorf("wrong note name, should be A#3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "D4") {
				t.Errorf("wrong note name, should be D4 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "F4") {
				t.Errorf("wrong note name, should be F4 have %v", s.Name())
			}
		} else {
			t.Errorf("wrong filename %v\n", s.Name())
		}
		if Chord[1]-Chord[0] != 3 || Chord[2]-Chord[0] != 7 || Chord[3]-Chord[0] != 10 {
			t.Errorf("failed to create Minor7th chord from notes: %v %v, %v %v, %v %v, %v %v\n",
				NoteToNameAndOctave(Chord[0]), Chord[0],
				NoteToNameAndOctave(Chord[1]), Chord[1],
				NoteToNameAndOctave(Chord[2]), Chord[2],
				NoteToNameAndOctave(Chord[3]), Chord[3])
		}
		Chord = nil
	}

	os.RemoveAll(TestDirectory)
}

func TestChord6th(t *testing.T) {

	// to disable logging, pass mid.NoLogger() as option
	rd := reader.New(reader.NoLogger(),
		// set the functions for the messages you are interested in
		reader.NoteOn(p.noteOn),
	)

	Chord6th(noteNum, TestDirectory)
	Chord6th(noteNum+7, TestDirectory)
	files, err := ioutil.ReadDir(TestDirectory)
	if err != nil {
		t.Errorf("could not read %v directory\n", TestDirectory)
	}
	if len(files) <= 0 {
		t.Errorf("no files created")
	}
	for _, s := range files {
		err := reader.ReadSMFFile(rd, TestDirectory+"/"+s.Name())

		if err != nil {
			t.Errorf("could not read SMF file %v\n", s)
		}
		if strings.HasPrefix(s.Name(), "C") {
			if !strings.Contains(s.Name(), "6th") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "6th")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "E") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "E")
			} else if !strings.Contains(s.Name(), "G") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "G")
			} else if !strings.Contains(s.Name(), "A") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "A")
			} else if !strings.Contains(s.Name(), "C3") {
				t.Errorf("wrong note name, should be C3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "E3") {
				t.Errorf("wrong note name, should be E3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "A3") {
				t.Errorf("wrong note name, should be A3 have %v", s.Name())
			}
		} else if strings.HasPrefix(s.Name(), "G") {
			if !strings.Contains(s.Name(), "6th") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "6th")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "B") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "B")
			} else if !strings.Contains(s.Name(), "D") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "D")
			} else if !strings.Contains(s.Name(), "E") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "E")
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "B3") {
				t.Errorf("wrong note name, should be B3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "D4") {
				t.Errorf("wrong note name, should be D4 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "E4") {
				t.Errorf("wrong note name, should be E4 have %v", s.Name())
			}
		} else {
			t.Errorf("wrong filename %v\n", s.Name())
		}
		if Chord[1]-Chord[0] != 4 || Chord[2]-Chord[0] != 7 || Chord[3]-Chord[0] != 9 {
			t.Errorf("failed to create 6th chord from notes: %v %v, %v %v, %v %v, %v %v\n",
				NoteToNameAndOctave(Chord[0]), Chord[0],
				NoteToNameAndOctave(Chord[1]), Chord[1],
				NoteToNameAndOctave(Chord[2]), Chord[2],
				NoteToNameAndOctave(Chord[3]), Chord[3])
		}
		Chord = nil
	}

	os.RemoveAll(TestDirectory)
}

func TestChord7th(t *testing.T) {

	// to disable logging, pass mid.NoLogger() as option
	rd := reader.New(reader.NoLogger(),
		// set the functions for the messages you are interested in
		reader.NoteOn(p.noteOn),
	)

	Chord7th(noteNum, TestDirectory)
	Chord7th(noteNum+7, TestDirectory)
	files, err := ioutil.ReadDir(TestDirectory)
	if err != nil {
		t.Errorf("could not read %v directory\n", TestDirectory)
	}
	if len(files) <= 0 {
		t.Errorf("no files created")
	}
	for _, s := range files {
		err := reader.ReadSMFFile(rd, TestDirectory+"/"+s.Name())

		if err != nil {
			t.Errorf("could not read SMF file %v\n", s)
		}
		if strings.HasPrefix(s.Name(), "C") {
			if !strings.Contains(s.Name(), "7th") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "7th")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "E") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "E")
			} else if !strings.Contains(s.Name(), "G") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "G")
			} else if !strings.Contains(s.Name(), "A#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "A#")
			} else if !strings.Contains(s.Name(), "C3") {
				t.Errorf("wrong note name, should be C3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "E3") {
				t.Errorf("wrong note name, should be E3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "A#3") {
				t.Errorf("wrong note name, should be A#3 have %v", s.Name())
			}
		} else if strings.HasPrefix(s.Name(), "G") {
			if !strings.Contains(s.Name(), "7th") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "7th")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "B") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "B")
			} else if !strings.Contains(s.Name(), "D") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "D")
			} else if !strings.Contains(s.Name(), "F") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "F")
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "B3") {
				t.Errorf("wrong note name, should be B3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "D4") {
				t.Errorf("wrong note name, should be D4 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "F4") {
				t.Errorf("wrong note name, should be F4 have %v", s.Name())
			}
		} else {
			t.Errorf("wrong filename %v\n", s.Name())
		}
		if Chord[1]-Chord[0] != 4 || Chord[2]-Chord[0] != 7 || Chord[3]-Chord[0] != 10 {
			t.Errorf("failed to create 7th chord from notes: %v %v, %v %v, %v %v, %v %v\n",
				NoteToNameAndOctave(Chord[0]), Chord[0],
				NoteToNameAndOctave(Chord[1]), Chord[1],
				NoteToNameAndOctave(Chord[2]), Chord[2],
				NoteToNameAndOctave(Chord[3]), Chord[3])
		}
		Chord = nil
	}

	os.RemoveAll(TestDirectory)
}

func TestChord7Sus2(t *testing.T) {

	// to disable logging, pass mid.NoLogger() as option
	rd := reader.New(reader.NoLogger(),
		// set the functions for the messages you are interested in
		reader.NoteOn(p.noteOn),
	)

	Chord7Sus2(noteNum, TestDirectory)
	Chord7Sus2(noteNum+7, TestDirectory)
	files, err := ioutil.ReadDir(TestDirectory)
	if err != nil {
		t.Errorf("could not read %v directory\n", TestDirectory)
	}
	if len(files) <= 0 {
		t.Errorf("no files created")
	}
	for _, s := range files {
		err := reader.ReadSMFFile(rd, TestDirectory+"/"+s.Name())

		if err != nil {
			t.Errorf("could not read SMF file %v\n", s)
		}
		if strings.HasPrefix(s.Name(), "C") {
			if !strings.Contains(s.Name(), "7sus2") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "7sus2")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "D") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "D")
			} else if !strings.Contains(s.Name(), "G") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "G")
			} else if !strings.Contains(s.Name(), "A#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "A#")
			} else if !strings.Contains(s.Name(), "C3") {
				t.Errorf("wrong note name, should be C3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "D3") {
				t.Errorf("wrong note name, should be D3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "A#3") {
				t.Errorf("wrong note name, should be A#3 have %v", s.Name())
			}
		} else if strings.HasPrefix(s.Name(), "G") {
			if !strings.Contains(s.Name(), "7sus2") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "7sus2")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "A") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "A")
			} else if !strings.Contains(s.Name(), "D") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "D")
			} else if !strings.Contains(s.Name(), "F") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "F")
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "A3") {
				t.Errorf("wrong note name, should be A3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "D4") {
				t.Errorf("wrong note name, should be D4 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "F4") {
				t.Errorf("wrong note name, should be F4 have %v", s.Name())
			}
		} else {
			t.Errorf("wrong filename %v\n", s.Name())
		}
		if Chord[1]-Chord[0] != 2 || Chord[2]-Chord[0] != 7 || Chord[3]-Chord[0] != 10 {
			t.Errorf("failed to create 7Sus2 chord from notes: %v %v, %v %v, %v %v, %v %v\n",
				NoteToNameAndOctave(Chord[0]), Chord[0],
				NoteToNameAndOctave(Chord[1]), Chord[1],
				NoteToNameAndOctave(Chord[2]), Chord[2],
				NoteToNameAndOctave(Chord[3]), Chord[3])
		}
		Chord = nil
	}

	os.RemoveAll(TestDirectory)
}

func TestChord7Sus4(t *testing.T) {

	// to disable logging, pass mid.NoLogger() as option
	rd := reader.New(reader.NoLogger(),
		// set the functions for the messages you are interested in
		reader.NoteOn(p.noteOn),
	)

	Chord7Sus4(noteNum, TestDirectory)
	Chord7Sus4(noteNum+7, TestDirectory)
	files, err := ioutil.ReadDir(TestDirectory)
	if err != nil {
		t.Errorf("could not read %v directory\n", TestDirectory)
	}
	if len(files) <= 0 {
		t.Errorf("no files created")
	}
	for _, s := range files {
		err := reader.ReadSMFFile(rd, TestDirectory+"/"+s.Name())

		if err != nil {
			t.Errorf("could not read SMF file %v\n", s)
		}
		if strings.HasPrefix(s.Name(), "C") {
			if !strings.Contains(s.Name(), "7sus4") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "7sus4")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "F") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "F")
			} else if !strings.Contains(s.Name(), "G") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "G")
			} else if !strings.Contains(s.Name(), "A#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "A#")
			} else if !strings.Contains(s.Name(), "C3") {
				t.Errorf("wrong note name, should be C3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "F3") {
				t.Errorf("wrong note name, should be F3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "A#3") {
				t.Errorf("wrong note name, should be A#3 have %v", s.Name())
			}
		} else if strings.HasPrefix(s.Name(), "G") {
			if !strings.Contains(s.Name(), "7sus4") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "7sus4")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "C") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "C")
			} else if !strings.Contains(s.Name(), "D") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "D")
			} else if !strings.Contains(s.Name(), "F") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "F")
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "C4") {
				t.Errorf("wrong note name, should be C4 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "D4") {
				t.Errorf("wrong note name, should be D4 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "F4") {
				t.Errorf("wrong note name, should be F4 have %v", s.Name())
			}
		} else {
			t.Errorf("wrong filename %v\n", s.Name())
		}
		if Chord[1]-Chord[0] != 5 || Chord[2]-Chord[0] != 7 || Chord[3]-Chord[0] != 10 {
			t.Errorf("failed to create 7Sus4 chord from notes: %v %v, %v %v, %v %v, %v %v\n",
				NoteToNameAndOctave(Chord[0]), Chord[0],
				NoteToNameAndOctave(Chord[1]), Chord[1],
				NoteToNameAndOctave(Chord[2]), Chord[2],
				NoteToNameAndOctave(Chord[3]), Chord[3])
		}
		Chord = nil
	}

	os.RemoveAll(TestDirectory)
}

func TestMinor6th(t *testing.T) {

	// to disable logging, pass mid.NoLogger() as option
	rd := reader.New(reader.NoLogger(),
		// set the functions for the messages you are interested in
		reader.NoteOn(p.noteOn),
	)

	Minor6th(noteNum, TestDirectory)
	Minor6th(noteNum+7, TestDirectory)
	files, err := ioutil.ReadDir(TestDirectory)
	if err != nil {
		t.Errorf("could not read %v directory\n", TestDirectory)
	}
	if len(files) <= 0 {
		t.Errorf("no files created")
	}
	for _, s := range files {
		err := reader.ReadSMFFile(rd, TestDirectory+"/"+s.Name())

		if err != nil {
			t.Errorf("could not read SMF file %v\n", s)
		}
		if strings.HasPrefix(s.Name(), "C") {
			if !strings.Contains(s.Name(), "min6") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "min6")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "D#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "D#")
			} else if !strings.Contains(s.Name(), "G") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "G")
			} else if !strings.Contains(s.Name(), "A") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "A")
			} else if !strings.Contains(s.Name(), "C3") {
				t.Errorf("wrong note name, should be C3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "D#3") {
				t.Errorf("wrong note name, should be D#3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "A3") {
				t.Errorf("wrong note name, should be A3 have %v", s.Name())
			}
		} else if strings.HasPrefix(s.Name(), "G") {
			if !strings.Contains(s.Name(), "min6") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "min6")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "A#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "A#")
			} else if !strings.Contains(s.Name(), "D") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "D")
			} else if !strings.Contains(s.Name(), "E") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "E")
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "A#3") {
				t.Errorf("wrong note name, should be A#3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "D4") {
				t.Errorf("wrong note name, should be D4 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "E4") {
				t.Errorf("wrong note name, should be E4 have %v", s.Name())
			}
		} else {
			t.Errorf("wrong filename %v\n", s.Name())
		}
		if Chord[1]-Chord[0] != 3 || Chord[2]-Chord[0] != 7 || Chord[3]-Chord[0] != 9 {
			t.Errorf("failed to create Minor6th chord from notes: %v %v, %v %v, %v %v, %v %v\n",
				NoteToNameAndOctave(Chord[0]), Chord[0],
				NoteToNameAndOctave(Chord[1]), Chord[1],
				NoteToNameAndOctave(Chord[2]), Chord[2],
				NoteToNameAndOctave(Chord[3]), Chord[3])
		}
		Chord = nil
	}

	os.RemoveAll(TestDirectory)
}

func TestDiminished7(t *testing.T) {

	// to disable logging, pass mid.NoLogger() as option
	rd := reader.New(reader.NoLogger(),
		// set the functions for the messages you are interested in
		reader.NoteOn(p.noteOn),
	)

	Diminished7(noteNum, TestDirectory)
	Diminished7(noteNum+7, TestDirectory)
	files, err := ioutil.ReadDir(TestDirectory)
	if err != nil {
		t.Errorf("could not read %v directory\n", TestDirectory)
	}
	if len(files) <= 0 {
		t.Errorf("no files created")
	}
	for _, s := range files {
		err := reader.ReadSMFFile(rd, TestDirectory+"/"+s.Name())

		if err != nil {
			t.Errorf("could not read SMF file %v\n", s)
		}
		if strings.HasPrefix(s.Name(), "C") {
			if !strings.Contains(s.Name(), "dim7") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "dim7")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "D#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "D#")
			} else if !strings.Contains(s.Name(), "F#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "F#")
			} else if !strings.Contains(s.Name(), "A#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "A#")
			} else if !strings.Contains(s.Name(), "C3") {
				t.Errorf("wrong note name, should be C3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "D#3") {
				t.Errorf("wrong note name, should be D#3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "F#3") {
				t.Errorf("wrong note name, should be F#3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "A#3") {
				t.Errorf("wrong note name, should be A#3 have %v", s.Name())
			}
		} else if strings.HasPrefix(s.Name(), "G") {
			if !strings.Contains(s.Name(), "dim7") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "dim7")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "A#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "A#")
			} else if !strings.Contains(s.Name(), "C#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "C#")
			} else if !strings.Contains(s.Name(), "F") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "F")
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "A#3") {
				t.Errorf("wrong note name, should be A#3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "C#4") {
				t.Errorf("wrong note name, should be C#4 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "F4") {
				t.Errorf("wrong note name, should be F4 have %v", s.Name())
			}
		} else {
			t.Errorf("wrong filename %v\n", s.Name())
		}
		if Chord[1]-Chord[0] != 3 || Chord[2]-Chord[0] != 6 || Chord[3]-Chord[0] != 10 {
			t.Errorf("failed to create Diminished7 chord from notes: %v %v, %v %v, %v %v, %v %v\n",
				NoteToNameAndOctave(Chord[0]), Chord[0],
				NoteToNameAndOctave(Chord[1]), Chord[1],
				NoteToNameAndOctave(Chord[2]), Chord[2],
				NoteToNameAndOctave(Chord[3]), Chord[3])
		}
		Chord = nil
	}

	os.RemoveAll(TestDirectory)
}

func TestAugmented7(t *testing.T) {

	// to disable logging, pass mid.NoLogger() as option
	rd := reader.New(reader.NoLogger(),
		// set the functions for the messages you are interested in
		reader.NoteOn(p.noteOn),
	)

	Augmented7(noteNum, TestDirectory)
	Augmented7(noteNum+7, TestDirectory)
	files, err := ioutil.ReadDir(TestDirectory)
	if err != nil {
		t.Errorf("could not read %v directory\n", TestDirectory)
	}
	if len(files) <= 0 {
		t.Errorf("no files created")
	}
	for _, s := range files {
		err := reader.ReadSMFFile(rd, TestDirectory+"/"+s.Name())

		if err != nil {
			t.Errorf("could not read SMF file %v\n", s)
		}
		if strings.HasPrefix(s.Name(), "C") {
			if !strings.Contains(s.Name(), "aug7") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "aug7")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "E") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "E")
			} else if !strings.Contains(s.Name(), "G#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "G#")
			} else if !strings.Contains(s.Name(), "A#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "A#")
			} else if !strings.Contains(s.Name(), "C3") {
				t.Errorf("wrong note name, should be C3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "E3") {
				t.Errorf("wrong note name, should be E3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "G#3") {
				t.Errorf("wrong note name, should be G#3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "A#3") {
				t.Errorf("wrong note name, should be A#3 have %v", s.Name())
			}
		} else if strings.HasPrefix(s.Name(), "G") {
			if !strings.Contains(s.Name(), "aug7") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "aug7")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "B") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "B")
			} else if !strings.Contains(s.Name(), "D#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "D#")
			} else if !strings.Contains(s.Name(), "F") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "F")
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "B3") {
				t.Errorf("wrong note name, should be B3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "D#4") {
				t.Errorf("wrong note name, should be D#4 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "F4") {
				t.Errorf("wrong note name, should be F4 have %v", s.Name())
			}
		} else {
			t.Errorf("wrong filename %v\n", s.Name())
		}
		if Chord[1]-Chord[0] != 4 || Chord[2]-Chord[0] != 8 || Chord[3]-Chord[0] != 10 {
			t.Errorf("failed to create Augmented7 chord from notes: %v %v, %v %v, %v %v, %v %v\n",
				NoteToNameAndOctave(Chord[0]), Chord[0],
				NoteToNameAndOctave(Chord[1]), Chord[1],
				NoteToNameAndOctave(Chord[2]), Chord[2],
				NoteToNameAndOctave(Chord[3]), Chord[3])
		}
		Chord = nil
	}

	os.RemoveAll(TestDirectory)
}

func TestDiminished6(t *testing.T) {

	// to disable logging, pass mid.NoLogger() as option
	rd := reader.New(reader.NoLogger(),
		// set the functions for the messages you are interested in
		reader.NoteOn(p.noteOn),
	)

	Diminished6(noteNum, TestDirectory)
	Diminished6(noteNum+7, TestDirectory)
	files, err := ioutil.ReadDir(TestDirectory)
	if err != nil {
		t.Errorf("could not read %v directory\n", TestDirectory)
	}
	if len(files) <= 0 {
		t.Errorf("no files created")
	}
	for _, s := range files {
		err := reader.ReadSMFFile(rd, TestDirectory+"/"+s.Name())

		if err != nil {
			t.Errorf("could not read SMF file %v\n", s)
		}
		if strings.HasPrefix(s.Name(), "C") {
			if !strings.Contains(s.Name(), "dim6") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "dim6")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "D#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "D#")
			} else if !strings.Contains(s.Name(), "F#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "F#")
			} else if !strings.Contains(s.Name(), "A") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "A")
			} else if !strings.Contains(s.Name(), "C3") {
				t.Errorf("wrong note name, should be C3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "D#3") {
				t.Errorf("wrong note name, should be D#3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "F#3") {
				t.Errorf("wrong note name, should be F#3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "A3") {
				t.Errorf("wrong note name, should be A3 have %v", s.Name())
			}
		} else if strings.HasPrefix(s.Name(), "G") {
			if !strings.Contains(s.Name(), "dim6") {
				t.Errorf("file name is wrong, %v should contain name of the chord %v\n", s.Name(), "dim6")
			} else if !strings.HasSuffix(s.Name(), ".mid") {
				t.Errorf("should end with .mid filename suffix\n")
			}
			if !strings.Contains(s.Name(), "A#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "A#")
			} else if !strings.Contains(s.Name(), "C#") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "C#")
			} else if !strings.Contains(s.Name(), "E") {
				t.Errorf("missing note name in filename: %v, missing: %v", s.Name(), "E")
			} else if !strings.Contains(s.Name(), "G3") {
				t.Errorf("wrong note name, should be G3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "A#3") {
				t.Errorf("wrong note name, should be A#3 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "C#4") {
				t.Errorf("wrong note name, should be C#4 have %v", s.Name())
			} else if !strings.Contains(s.Name(), "E4") {
				t.Errorf("wrong note name, should be E4 have %v", s.Name())
			}
		} else {
			t.Errorf("wrong filename %v\n", s.Name())
		}
		if Chord[1]-Chord[0] != 3 || Chord[2]-Chord[0] != 6 || Chord[3]-Chord[0] != 9 {
			t.Errorf("failed to create Diminished6 chord from notes: %v %v, %v %v, %v %v, %v %v\n",
				NoteToNameAndOctave(Chord[0]), Chord[0],
				NoteToNameAndOctave(Chord[1]), Chord[1],
				NoteToNameAndOctave(Chord[2]), Chord[2],
				NoteToNameAndOctave(Chord[3]), Chord[3])
		}
		Chord = nil
	}

	os.RemoveAll(TestDirectory)
}
