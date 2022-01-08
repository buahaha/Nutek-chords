package Nutek

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"gitlab.com/gomidi/midi/writer"
)

// Create one note 'solo' midi file in dir(ectory) of noteNumber in between 0 and 127.
// Use Nutek.StringToNote to convert English note name to MIDI number
func OneNote(noteNumber uint8, dir string) {
	file := workWithDirectories(dir)

	baseKey := noteNumber
	noteName, _ := mapkey(notes, baseKey)

	// chord name
	var chordName string = NoteToName(baseKey) + "solo"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + noteName + " "

	fileName := chordNameSuffix + fileNameSuffix

	// current file to work with
	currentFile := filepath.Join(file, fileName)

	// MIDI chord writer function
	err := writer.WriteSMF(currentFile, 1, func(wr *writer.SMF) error {

		// sets the channel for the next messages
		wr.SetChannel(0)

		// start a Major triad
		writer.NoteOn(wr, baseKey, 100)

		// move for 1 bar
		writer.Forward(wr, 0, 1, 1)

		// release the notes
		writer.NoteOff(wr, baseKey)

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

// Create major chord midi file in dir(ectory) of noteNumber in between 0 and 127.
// Use Nutek.StringToNote to convert English note name to MIDI number
func MajorTriad(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

	baseKey := startingKey
	noteName, _ := mapkey(notes, baseKey)

	// chord name
	var chordName string = NoteToName(baseKey) + "maj"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+4)
	chordNameSuffix += NoteToName(baseKey+4) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+7)
	chordNameSuffix += NoteToName(baseKey+7) + noteName

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

// Create minor chord midi file in dir(ectory) of noteNumber in between 0 and 127.
// Use Nutek.StringToNote to convert English note name to MIDI number
func MinorTriad(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

	baseKey := startingKey
	noteName, _ := mapkey(notes, baseKey)

	// chord name
	var chordName string = NoteToName(baseKey) + "min"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+3)
	chordNameSuffix += NoteToName(baseKey+3) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+7)
	chordNameSuffix += NoteToName(baseKey+7) + noteName

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

// Create diminished chord midi file in dir(ectory) of noteNumber in between 0 and 127.
// Use Nutek.StringToNote to convert English note name to MIDI number
func Diminished(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

	baseKey := startingKey
	noteName, _ := mapkey(notes, baseKey)

	// chord name
	var chordName string = NoteToName(baseKey) + "dim"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+3)
	chordNameSuffix += NoteToName(baseKey+3) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+6)
	chordNameSuffix += NoteToName(baseKey+6) + noteName

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

// Create augmented chord midi file in dir(ectory) of noteNumber in between 0 and 127.
// Use Nutek.StringToNote to convert English note name to MIDI number
func Augmented(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

	baseKey := startingKey
	noteName, _ := mapkey(notes, baseKey)

	// chord name
	var chordName string = NoteToName(baseKey) + "aug"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+4)
	chordNameSuffix += NoteToName(baseKey+4) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+8)
	chordNameSuffix += NoteToName(baseKey+8) + noteName

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

// Create sus2 chord midi file in dir(ectory) of noteNumber in between 0 and 127.
// Use Nutek.StringToNote to convert English note name to MIDI number
func Sus2(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

	baseKey := startingKey
	noteName, _ := mapkey(notes, baseKey)

	// chord name
	var chordName string = NoteToName(baseKey) + "sus2"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+2)
	chordNameSuffix += NoteToName(baseKey+2) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+7)
	chordNameSuffix += NoteToName(baseKey+7) + noteName

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

// Create sus4 chord midi file in dir(ectory) of noteNumber in between 0 and 127.
// Use Nutek.StringToNote to convert English note name to MIDI number
func Sus4(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

	baseKey := startingKey
	noteName, _ := mapkey(notes, baseKey)

	// chord name
	var chordName string = NoteToName(baseKey) + "sus4"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+5)
	chordNameSuffix += NoteToName(baseKey+5) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+7)
	chordNameSuffix += NoteToName(baseKey+7) + noteName

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

// Create major 7th chord midi file in dir(ectory) of noteNumber in between 0 and 127.
// Use Nutek.StringToNote to convert English note name to MIDI number
func Major7th(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

	baseKey := startingKey
	noteName, _ := mapkey(notes, baseKey)

	// chord name
	var chordName string = NoteToName(baseKey) + "maj7"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+4)
	chordNameSuffix += NoteToName(baseKey+4) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+7)
	chordNameSuffix += NoteToName(baseKey+7) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+11)
	chordNameSuffix += NoteToName(baseKey+11) + noteName

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

// Create minor 7th chord midi file in dir(ectory) of noteNumber in between 0 and 127.
// Use Nutek.StringToNote to convert English note name to MIDI number
func Minor7th(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

	baseKey := startingKey
	noteName, _ := mapkey(notes, baseKey)

	// chord name
	var chordName string = NoteToName(baseKey) + "min7"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+3)
	chordNameSuffix += NoteToName(baseKey+3) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+7)
	chordNameSuffix += NoteToName(baseKey+7) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+10)
	chordNameSuffix += NoteToName(baseKey+10) + noteName

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

// Create 6th chord midi file in dir(ectory) of noteNumber in between 0 and 127.
// Use Nutek.StringToNote to convert English note name to MIDI number
func Chord6th(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

	baseKey := startingKey
	noteName, _ := mapkey(notes, baseKey)

	// chord name
	var chordName string = NoteToName(baseKey) + "6th"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+4)
	chordNameSuffix += NoteToName(baseKey+4) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+7)
	chordNameSuffix += NoteToName(baseKey+7) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+9)
	chordNameSuffix += NoteToName(baseKey+9) + noteName

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

// Create 7th chord midi file in dir(ectory) of noteNumber in between 0 and 127.
// Use Nutek.StringToNote to convert English note name to MIDI number
func Chord7th(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

	baseKey := startingKey
	noteName, _ := mapkey(notes, baseKey)

	// chord name
	var chordName string = NoteToName(baseKey) + "7th"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+4)
	chordNameSuffix += NoteToName(baseKey+4) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+7)
	chordNameSuffix += NoteToName(baseKey+7) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+10)
	chordNameSuffix += NoteToName(baseKey+10) + noteName

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

// Create 7sus2 chord midi file in dir(ectory) of noteNumber in between 0 and 127.
// Use Nutek.StringToNote to convert English note name to MIDI number
func Chord7Sus2(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

	baseKey := startingKey
	noteName, _ := mapkey(notes, baseKey)

	// chord name
	var chordName string = NoteToName(baseKey) + "7sus2"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+2)
	chordNameSuffix += NoteToName(baseKey+2) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+7)
	chordNameSuffix += NoteToName(baseKey+7) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+10)
	chordNameSuffix += NoteToName(baseKey+10) + noteName

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

// Create 7sus4 chord midi file in dir(ectory) of noteNumber in between 0 and 127.
// Use Nutek.StringToNote to convert English note name to MIDI number
func Chord7Sus4(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

	baseKey := startingKey
	noteName, _ := mapkey(notes, baseKey)

	// chord name
	var chordName string = NoteToName(baseKey) + "7sus4"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+5)
	chordNameSuffix += NoteToName(baseKey+5) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+7)
	chordNameSuffix += NoteToName(baseKey+7) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+10)
	chordNameSuffix += NoteToName(baseKey+10) + noteName

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

// Create minor 6th chord midi file in dir(ectory) of noteNumber in between 0 and 127.
// Use Nutek.StringToNote to convert English note name to MIDI number
func Minor6th(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

	baseKey := startingKey
	noteName, _ := mapkey(notes, baseKey)

	// chord name
	var chordName string = NoteToName(baseKey) + "min6"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+3)
	chordNameSuffix += NoteToName(baseKey+3) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+7)
	chordNameSuffix += NoteToName(baseKey+7) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+9)
	chordNameSuffix += NoteToName(baseKey+9) + noteName

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

// Create diminished 7 chord midi file in dir(ectory) of noteNumber in between 0 and 127.
// Use Nutek.StringToNote to convert English note name to MIDI number
func Diminished7(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

	baseKey := startingKey
	noteName, _ := mapkey(notes, baseKey)

	// chord name
	var chordName string = NoteToName(baseKey) + "dim7"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+3)
	chordNameSuffix += NoteToName(baseKey+3) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+6)
	chordNameSuffix += NoteToName(baseKey+6) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+10)
	chordNameSuffix += NoteToName(baseKey+10) + noteName

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
		writer.NoteOn(wr, baseKey+10, 100)

		// move for 1 bar
		writer.Forward(wr, 0, 1, 1)

		// release the notes
		writer.NoteOff(wr, baseKey)
		writer.NoteOff(wr, baseKey+3)
		writer.NoteOff(wr, baseKey+6)
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

// Create augmented 7 chord midi file in dir(ectory) of noteNumber in between 0 and 127.
// Use Nutek.StringToNote to convert English note name to MIDI number
func Augmented7(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

	baseKey := startingKey
	noteName, _ := mapkey(notes, baseKey)

	// chord name
	var chordName string = NoteToName(baseKey) + "aug7"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+4)
	chordNameSuffix += NoteToName(baseKey+4) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+8)
	chordNameSuffix += NoteToName(baseKey+8) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+10)
	chordNameSuffix += NoteToName(baseKey+10) + noteName

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
		writer.NoteOn(wr, baseKey+10, 100)

		// move for 1 bar
		writer.Forward(wr, 0, 1, 1)

		// release the notes
		writer.NoteOff(wr, baseKey)
		writer.NoteOff(wr, baseKey+4)
		writer.NoteOff(wr, baseKey+8)
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

// Create diminished 6 chord midi file in dir(ectory) of noteNumber in between 0 and 127.
// Use Nutek.StringToNote to convert English note name to MIDI number
func Diminished6(startingKey uint8, dir string) {
	file := workWithDirectories(dir)

	baseKey := startingKey
	noteName, _ := mapkey(notes, baseKey)

	// chord name
	var chordName string = NoteToName(baseKey) + "dim6"

	// notes in chord name to chord name suffix
	chordNameSuffix := chordName + " - "
	chordNameSuffix += NoteToName(baseKey) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+3)
	chordNameSuffix += NoteToName(baseKey+3) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+6)
	chordNameSuffix += NoteToName(baseKey+6) + noteName + " "
	noteName, _ = mapkey(notes, baseKey+9)
	chordNameSuffix += NoteToName(baseKey+9) + noteName

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
		writer.NoteOn(wr, baseKey+9, 100)

		// move for 1 bar
		writer.Forward(wr, 0, 1, 1)

		// release the notes
		writer.NoteOff(wr, baseKey)
		writer.NoteOff(wr, baseKey+3)
		writer.NoteOff(wr, baseKey+6)
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

// Setup working directory and folders to work whith files
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
		err := os.MkdirAll(file, 0755)
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

// Convert name of the note to MIDI note number
func StringToNote(input string) (uint8, error) {
	var noteNum uint8
	num, err := strconv.Atoi(input)
	if err != nil {
		if val, ok := notes[input]; ok {
			return val, nil
		} else {
			log.Fatalf("note %v not found in MIDI map", val)
		}
	}
	noteNum = uint8(num)
	if noteNum <= 127 {
		return noteNum, nil
	} else {
		log.Fatalf("note value out of bounds (0 to 127) was %v", noteNum)
	}
	return 128, fmt.Errorf("error finding given note %v", input)
}

// Circular reference to get key (name of the note) from it's MIDI value
func mapkey(m map[string]uint8, value uint8) (key string, ok bool) {
	for k, v := range m {
		if v == value {
			if strings.Contains(k, "middle") {
				continue
			}
			if strings.Contains(k, "b") {
				continue
			}
			if strings.Contains(k, "#") {
				key = k[2:]
				ok = true
				return
			} else {
				key = k[1:]
				ok = true
				return
			}
		}
	}
	return
}

var notes map[string]uint8 = map[string]uint8{
	"G9":            127,
	"F#9":           126,
	"Gb9":           126,
	"F9":            125,
	"E9":            124,
	"D#9":           123,
	"Eb9":           123,
	"D9":            122,
	"C#9":           121,
	"Db9":           121,
	"C9":            120,
	"B8":            119,
	"A#8":           118,
	"Bb8":           118,
	"A8":            117,
	"G#8":           116,
	"Ab8":           116,
	"G8":            115,
	"F#8":           114,
	"Gb8":           114,
	"F8":            113,
	"E8":            112,
	"D#8":           111,
	"Eb8":           111,
	"D8":            110,
	"C#8":           109,
	"Db8":           109,
	"C8":            108,
	"B7":            107,
	"A#7":           106,
	"Bb7":           106,
	"A7":            105,
	"G#7":           104,
	"Ab7":           104,
	"G7":            103,
	"F#7":           102,
	"Gb7":           102,
	"F7":            101,
	"E7":            100,
	"D#7":           99,
	"Eb7":           99,
	"D7":            98,
	"C#7":           97,
	"Db7":           97,
	"C7":            96,
	"B6":            95,
	"A#6":           94,
	"Bb6":           94,
	"A6":            93,
	"G#6":           92,
	"Ab6":           92,
	"G6":            91,
	"F#6":           90,
	"Gb6":           90,
	"F6":            89,
	"E6":            88,
	"D#6":           87,
	"Eb6":           87,
	"D6":            86,
	"C#6":           85,
	"Db6":           85,
	"C6":            84,
	"B5":            83,
	"A#5":           82,
	"Bb5":           82,
	"A5":            81,
	"G#5":           80,
	"Ab5":           80,
	"G5":            79,
	"F#5":           78,
	"Gb5":           78,
	"F5":            77,
	"E5":            76,
	"D#5":           75,
	"Eb5":           75,
	"D5":            74,
	"C#5":           73,
	"Db5":           73,
	"C5":            72,
	"B4":            71,
	"A#4":           70,
	"Bb4":           70,
	"A4":            69,
	"concert pitch": 69,
	"440":           69,
	"G#4":           68,
	"Ab4":           68,
	"G4":            67,
	"F#4":           66,
	"Gb4":           66,
	"F4":            65,
	"E4":            64,
	"D#4":           63,
	"Eb4":           63,
	"D4":            62,
	"C#4":           61,
	"Db4":           61,
	"C4":            60,
	"middle C":      60,
	"middle c":      60,
	"B3":            59,
	"A#3":           58,
	"Bb3":           58,
	"A3":            57,
	"G#3":           56,
	"Ab3":           56,
	"G3":            55,
	"F#3":           54,
	"Gb3":           54,
	"F3":            53,
	"E3":            52,
	"D#3":           51,
	"Eb3":           51,
	"D3":            50,
	"C#3":           49,
	"Db3":           49,
	"C3":            48,
	"B2":            47,
	"A#2":           46,
	"Bb2":           46,
	"A2":            45,
	"G#2":           44,
	"Ab2":           44,
	"G2":            43,
	"F#2":           42,
	"Gb2":           42,
	"F2":            41,
	"E2":            40,
	"D#2":           39,
	"Eb2":           39,
	"D2":            38,
	"C#2":           37,
	"Db2":           37,
	"C2":            36,
	"B1":            35,
	"A#1":           34,
	"Bb1":           34,
	"A1":            33,
	"G#1":           32,
	"Ab1":           32,
	"G1":            31,
	"F#1":           30,
	"Gb1":           30,
	"F1":            29,
	"E1":            28,
	"D#1":           27,
	"Eb1":           27,
	"D1":            26,
	"C#1":           25,
	"Db1":           25,
	"C1":            24,
	"B0":            23,
	"A#0":           22,
	"Bb0":           22,
	"A0":            21,
	"20":            20,
	"19":            19,
	"18":            18,
	"17":            17,
	"16":            16,
	"15":            15,
	"14":            14,
	"13":            13,
	"12":            12,
	"11":            11,
	"10":            10,
	"9":             9,
	"8":             8,
	"7":             7,
	"6":             6,
	"5":             5,
	"4":             4,
	"3":             3,
	"2":             2,
	"1":             1,
	"0":             0,
}

const fileNameSuffix string = ".mid"
