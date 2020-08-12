package musicdelight

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"gitlab.com/gomidi/midi/writer"
)

// MajorTriad create major chord from key startingKey in directory dir
func MajorTriad(startingKey uint8, dir string) {
	// Current working directory
	workingDirectory, err := os.Getwd()

	// file path to the working directory + new direcotry for midi files
	file := filepath.Join(workingDirectory, dir)

	if _, err := os.Stat(file); os.IsNotExist(err) {
		// path/to/whatever does not exist
		err = os.MkdirAll(file, 0755)
	}

	var baseKey uint8 = startingKey

	// chord name
	var chordName string = NoteToName(baseKey) + "maj"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + " "
	chordNameSuffix += NoteToName(baseKey+4) + " "
	chordNameSuffix += NoteToName(baseKey + 7)

	fileName := chordNameSuffix + fileNameSuffix

	// current file to work with
	currentFile := filepath.Join(file, fileName)

	// MIDI chord writer function
	err = writer.WriteSMF(currentFile, 1, func(wr *writer.SMF) error {

		// sets the channel for the next messages
		wr.SetChannel(0)

		// start a Major triad
		writer.NoteOn(wr, baseKey, 100)
		writer.NoteOn(wr, baseKey+4, 100)
		writer.NoteOn(wr, baseKey+7, 100)

		// move for 1 bar
		writer.Forward(wr, 0, 1, 1)

		// release the notes
		writer.NoteOff(wr, baseKey)
		writer.NoteOff(wr, baseKey+4)
		writer.NoteOff(wr, baseKey+7)

		// end of SMF file
		writer.EndOfTrack(wr)
		return nil
	})

	// check if there were errors writing SMF(MIDI) file
	if err != nil {
		fmt.Printf("could not write SMF file %v\n", currentFile)
		return
	}
}

// MinorTriad create minor chord
func MinorTriad(startingKey uint8, dir string) {
	// Current working directory
	workingDirectory, err := os.Getwd()

	// file path to the working directory + new direcotry for midi files
	file := filepath.Join(workingDirectory, dir)

	if _, err := os.Stat(file); os.IsNotExist(err) {
		// path/to/whatever does not exist
		err = os.MkdirAll(file, 0755)
	}

	var baseKey uint8 = startingKey

	// chord name
	var chordName string = NoteToName(baseKey) + "min"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + " "
	chordNameSuffix += NoteToName(baseKey+3) + " "
	chordNameSuffix += NoteToName(baseKey + 7)

	fileName := chordNameSuffix + fileNameSuffix

	// current file to work with
	currentFile := filepath.Join(file, fileName)

	// MIDI chord writer function
	err = writer.WriteSMF(currentFile, 1, func(wr *writer.SMF) error {

		// sets the channel for the next messages
		wr.SetChannel(0)

		// start a Major triad
		writer.NoteOn(wr, baseKey, 100)
		writer.NoteOn(wr, baseKey+3, 100)
		writer.NoteOn(wr, baseKey+7, 100)

		// move for 1 bar
		writer.Forward(wr, 0, 1, 1)

		// release the notes
		writer.NoteOff(wr, baseKey)
		writer.NoteOff(wr, baseKey+3)
		writer.NoteOff(wr, baseKey+7)

		// end of SMF file
		writer.EndOfTrack(wr)
		return nil
	})

	// check if there were errors writing SMF(MIDI) file
	if err != nil {
		fmt.Printf("could not write SMF file %v\n", currentFile)
		return
	}
}

// Notes english names
var Notes = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}

// NoteToName converts a midi note value into its English name
func NoteToName(note uint8) string {
	return Notes[note%12]
}

// NoteToNameAndOctave converts a midi note value into its English name
// and an octave number
func NoteToNameAndOctave(note uint8) string {
	key := Notes[note%12]
	octave := note/12 - 2 // The MIDI scale starts at octave = -2
	return key + strconv.Itoa(int(octave))
}

// C2 in Ableton Live - 48
const C2 uint8 = 48

const fileNameSuffix string = ".mid"
