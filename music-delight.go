package musicdelight

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"gitlab.com/gomidi/midi/writer"
)

// MajorTriad create major chord from key startingKey in directory dir
func MajorTriad(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

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
	err := writer.WriteSMF(currentFile, 1, func(wr *writer.SMF) error {

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
	file := workWithDirectories(dir)

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
	err := writer.WriteSMF(currentFile, 1, func(wr *writer.SMF) error {

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

func Diminished(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

	var baseKey uint8 = startingKey

	// chord name
	var chordName string = NoteToName(baseKey) + "dim"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + " "
	chordNameSuffix += NoteToName(baseKey+3) + " "
	chordNameSuffix += NoteToName(baseKey + 6)

	fileName := chordNameSuffix + fileNameSuffix

	// current file to work with
	currentFile := filepath.Join(file, fileName)

	// MIDI chord writer function
	err := writer.WriteSMF(currentFile, 1, func(wr *writer.SMF) error {

		// sets the channel for the next messages
		wr.SetChannel(0)

		// start a Major triad
		writer.NoteOn(wr, baseKey, 100)
		writer.NoteOn(wr, baseKey+3, 100)
		writer.NoteOn(wr, baseKey+6, 100)

		// move for 1 bar
		writer.Forward(wr, 0, 1, 1)

		// release the notes
		writer.NoteOff(wr, baseKey)
		writer.NoteOff(wr, baseKey+3)
		writer.NoteOff(wr, baseKey+6)

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

func Augmented(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

	var baseKey uint8 = startingKey

	// chord name
	var chordName string = NoteToName(baseKey) + "aug"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + " "
	chordNameSuffix += NoteToName(baseKey+4) + " "
	chordNameSuffix += NoteToName(baseKey + 8)

	fileName := chordNameSuffix + fileNameSuffix

	// current file to work with
	currentFile := filepath.Join(file, fileName)

	// MIDI chord writer function
	err := writer.WriteSMF(currentFile, 1, func(wr *writer.SMF) error {

		// sets the channel for the next messages
		wr.SetChannel(0)

		// start a Major triad
		writer.NoteOn(wr, baseKey, 100)
		writer.NoteOn(wr, baseKey+4, 100)
		writer.NoteOn(wr, baseKey+8, 100)

		// move for 1 bar
		writer.Forward(wr, 0, 1, 1)

		// release the notes
		writer.NoteOff(wr, baseKey)
		writer.NoteOff(wr, baseKey+4)
		writer.NoteOff(wr, baseKey+8)

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

func Sus2(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

	var baseKey uint8 = startingKey

	// chord name
	var chordName string = NoteToName(baseKey) + "sus2"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + " "
	chordNameSuffix += NoteToName(baseKey+2) + " "
	chordNameSuffix += NoteToName(baseKey + 7)

	fileName := chordNameSuffix + fileNameSuffix

	// current file to work with
	currentFile := filepath.Join(file, fileName)

	// MIDI chord writer function
	err := writer.WriteSMF(currentFile, 1, func(wr *writer.SMF) error {

		// sets the channel for the next messages
		wr.SetChannel(0)

		// start a Major triad
		writer.NoteOn(wr, baseKey, 100)
		writer.NoteOn(wr, baseKey+2, 100)
		writer.NoteOn(wr, baseKey+7, 100)

		// move for 1 bar
		writer.Forward(wr, 0, 1, 1)

		// release the notes
		writer.NoteOff(wr, baseKey)
		writer.NoteOff(wr, baseKey+2)
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

func Sus4(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

	var baseKey uint8 = startingKey

	// chord name
	var chordName string = NoteToName(baseKey) + "sus4"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + " "
	chordNameSuffix += NoteToName(baseKey+5) + " "
	chordNameSuffix += NoteToName(baseKey + 7)

	fileName := chordNameSuffix + fileNameSuffix

	// current file to work with
	currentFile := filepath.Join(file, fileName)

	// MIDI chord writer function
	err := writer.WriteSMF(currentFile, 1, func(wr *writer.SMF) error {

		// sets the channel for the next messages
		wr.SetChannel(0)

		// start a Major triad
		writer.NoteOn(wr, baseKey, 100)
		writer.NoteOn(wr, baseKey+5, 100)
		writer.NoteOn(wr, baseKey+7, 100)

		// move for 1 bar
		writer.Forward(wr, 0, 1, 1)

		// release the notes
		writer.NoteOff(wr, baseKey)
		writer.NoteOff(wr, baseKey+5)
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

func Major7th(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

	var baseKey uint8 = startingKey

	// chord name
	var chordName string = NoteToName(baseKey) + "maj7"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + " "
	chordNameSuffix += NoteToName(baseKey+4) + " "
	chordNameSuffix += NoteToName(baseKey+7) + " "
	chordNameSuffix += NoteToName(baseKey + 11)

	fileName := chordNameSuffix + fileNameSuffix

	// current file to work with
	currentFile := filepath.Join(file, fileName)

	// MIDI chord writer function
	err := writer.WriteSMF(currentFile, 1, func(wr *writer.SMF) error {

		// sets the channel for the next messages
		wr.SetChannel(0)

		// start a Major triad
		writer.NoteOn(wr, baseKey, 100)
		writer.NoteOn(wr, baseKey+4, 100)
		writer.NoteOn(wr, baseKey+7, 100)
		writer.NoteOn(wr, baseKey+11, 100)

		// move for 1 bar
		writer.Forward(wr, 0, 1, 1)

		// release the notes
		writer.NoteOff(wr, baseKey)
		writer.NoteOff(wr, baseKey+4)
		writer.NoteOff(wr, baseKey+7)
		writer.NoteOff(wr, baseKey+11)

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

func Minor7th(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

	var baseKey uint8 = startingKey

	// chord name
	var chordName string = NoteToName(baseKey) + "min7"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + " "
	chordNameSuffix += NoteToName(baseKey+3) + " "
	chordNameSuffix += NoteToName(baseKey+7) + " "
	chordNameSuffix += NoteToName(baseKey + 10)

	fileName := chordNameSuffix + fileNameSuffix

	// current file to work with
	currentFile := filepath.Join(file, fileName)

	// MIDI chord writer function
	err := writer.WriteSMF(currentFile, 1, func(wr *writer.SMF) error {

		// sets the channel for the next messages
		wr.SetChannel(0)

		// start a Major triad
		writer.NoteOn(wr, baseKey, 100)
		writer.NoteOn(wr, baseKey+3, 100)
		writer.NoteOn(wr, baseKey+7, 100)
		writer.NoteOn(wr, baseKey+10, 100)

		// move for 1 bar
		writer.Forward(wr, 0, 1, 1)

		// release the notes
		writer.NoteOff(wr, baseKey)
		writer.NoteOff(wr, baseKey+3)
		writer.NoteOff(wr, baseKey+7)
		writer.NoteOff(wr, baseKey+10)

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

func Major6th(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

	var baseKey uint8 = startingKey

	// chord name
	var chordName string = NoteToName(baseKey) + "maj6"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + " "
	chordNameSuffix += NoteToName(baseKey+4) + " "
	chordNameSuffix += NoteToName(baseKey+7) + " "
	chordNameSuffix += NoteToName(baseKey + 9)

	fileName := chordNameSuffix + fileNameSuffix

	// current file to work with
	currentFile := filepath.Join(file, fileName)

	// MIDI chord writer function
	err := writer.WriteSMF(currentFile, 1, func(wr *writer.SMF) error {

		// sets the channel for the next messages
		wr.SetChannel(0)

		// start a Major triad
		writer.NoteOn(wr, baseKey, 100)
		writer.NoteOn(wr, baseKey+4, 100)
		writer.NoteOn(wr, baseKey+7, 100)
		writer.NoteOn(wr, baseKey+9, 100)

		// move for 1 bar
		writer.Forward(wr, 0, 1, 1)

		// release the notes
		writer.NoteOff(wr, baseKey)
		writer.NoteOff(wr, baseKey+4)
		writer.NoteOff(wr, baseKey+7)
		writer.NoteOff(wr, baseKey+9)

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

func Chord7th(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

	var baseKey uint8 = startingKey

	// chord name
	var chordName string = NoteToName(baseKey) + "7th"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + " "
	chordNameSuffix += NoteToName(baseKey+4) + " "
	chordNameSuffix += NoteToName(baseKey+7) + " "
	chordNameSuffix += NoteToName(baseKey + 10)

	fileName := chordNameSuffix + fileNameSuffix

	// current file to work with
	currentFile := filepath.Join(file, fileName)

	// MIDI chord writer function
	err := writer.WriteSMF(currentFile, 1, func(wr *writer.SMF) error {

		// sets the channel for the next messages
		wr.SetChannel(0)

		// start a Major triad
		writer.NoteOn(wr, baseKey, 100)
		writer.NoteOn(wr, baseKey+4, 100)
		writer.NoteOn(wr, baseKey+7, 100)
		writer.NoteOn(wr, baseKey+10, 100)

		// move for 1 bar
		writer.Forward(wr, 0, 1, 1)

		// release the notes
		writer.NoteOff(wr, baseKey)
		writer.NoteOff(wr, baseKey+4)
		writer.NoteOff(wr, baseKey+7)
		writer.NoteOff(wr, baseKey+10)

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

func Chord7Sus2(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

	var baseKey uint8 = startingKey

	// chord name
	var chordName string = NoteToName(baseKey) + "7sus2"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + " "
	chordNameSuffix += NoteToName(baseKey+2) + " "
	chordNameSuffix += NoteToName(baseKey+7) + " "
	chordNameSuffix += NoteToName(baseKey + 10)

	fileName := chordNameSuffix + fileNameSuffix

	// current file to work with
	currentFile := filepath.Join(file, fileName)

	// MIDI chord writer function
	err := writer.WriteSMF(currentFile, 1, func(wr *writer.SMF) error {

		// sets the channel for the next messages
		wr.SetChannel(0)

		// start a Major triad
		writer.NoteOn(wr, baseKey, 100)
		writer.NoteOn(wr, baseKey+2, 100)
		writer.NoteOn(wr, baseKey+7, 100)
		writer.NoteOn(wr, baseKey+10, 100)

		// move for 1 bar
		writer.Forward(wr, 0, 1, 1)

		// release the notes
		writer.NoteOff(wr, baseKey)
		writer.NoteOff(wr, baseKey+2)
		writer.NoteOff(wr, baseKey+7)
		writer.NoteOff(wr, baseKey+10)

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

func Chord7Sus4(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

	var baseKey uint8 = startingKey

	// chord name
	var chordName string = NoteToName(baseKey) + "7sus4"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + " "
	chordNameSuffix += NoteToName(baseKey+5) + " "
	chordNameSuffix += NoteToName(baseKey+7) + " "
	chordNameSuffix += NoteToName(baseKey + 10)

	fileName := chordNameSuffix + fileNameSuffix

	// current file to work with
	currentFile := filepath.Join(file, fileName)

	// MIDI chord writer function
	err := writer.WriteSMF(currentFile, 1, func(wr *writer.SMF) error {

		// sets the channel for the next messages
		wr.SetChannel(0)

		// start a Major triad
		writer.NoteOn(wr, baseKey, 100)
		writer.NoteOn(wr, baseKey+5, 100)
		writer.NoteOn(wr, baseKey+7, 100)
		writer.NoteOn(wr, baseKey+10, 100)

		// move for 1 bar
		writer.Forward(wr, 0, 1, 1)

		// release the notes
		writer.NoteOff(wr, baseKey)
		writer.NoteOff(wr, baseKey+5)
		writer.NoteOff(wr, baseKey+7)
		writer.NoteOff(wr, baseKey+10)

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

func Minor6th(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

	var baseKey uint8 = startingKey

	// chord name
	var chordName string = NoteToName(baseKey) + "min6"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + " "
	chordNameSuffix += NoteToName(baseKey+3) + " "
	chordNameSuffix += NoteToName(baseKey+7) + " "
	chordNameSuffix += NoteToName(baseKey + 9)

	fileName := chordNameSuffix + fileNameSuffix

	// current file to work with
	currentFile := filepath.Join(file, fileName)

	// MIDI chord writer function
	err := writer.WriteSMF(currentFile, 1, func(wr *writer.SMF) error {

		// sets the channel for the next messages
		wr.SetChannel(0)

		// start a Major triad
		writer.NoteOn(wr, baseKey, 100)
		writer.NoteOn(wr, baseKey+3, 100)
		writer.NoteOn(wr, baseKey+7, 100)
		writer.NoteOn(wr, baseKey+9, 100)

		// move for 1 bar
		writer.Forward(wr, 0, 1, 1)

		// release the notes
		writer.NoteOff(wr, baseKey)
		writer.NoteOff(wr, baseKey+3)
		writer.NoteOff(wr, baseKey+7)
		writer.NoteOff(wr, baseKey+9)

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

func workWithDirectories(dir string) string {
	// Current working directory
	workingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatalf("working directory error: %v", err.Error())
	}

	// file path to the working directory + new direcotry for midi files
	file := filepath.Join(workingDirectory, dir)

	if _, err := os.Stat(file); os.IsNotExist(err) {
		// path/to/whatever does not exist
		err = os.MkdirAll(file, 0755)
		if err != nil {
			log.Fatalf("making directory error: %v", err.Error())
		}
	}

	return file
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
