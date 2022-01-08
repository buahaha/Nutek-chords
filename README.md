# Nutek
MIDI files generator library written in Go programming language for any kind of musician. Chords and notes for your type of DAW, game or other project with permissive *MIT license* from **Szymon & 🎶**
___

**WARNING!** This library does not create sound itself, the goal of this library is to provide you with programmatical representation of musical tones that can be used extensively in music production using software digital audio workstations like [Waveform Tracktion](https://www.tracktion.com/products/waveform-free) [MuseScore.org](https://musescore.org/)//[MuseScore's github](https://github.com/musescore/MuseScore) or equipment capable of dealing with MIDI file input. For example, you could send files generated by Nutek to your synthesizer giving you free hands to do **more music**
___

## How to use?

> After setting your environment for [go development...](https://go.dev)

1. Using command line terminal of your choice (Windows, Linux, macOS) create main project file (on macOS & Linux -> 

```shell
mkdir Nutek && cd Nutek && \
go mod init midi.example.net && \
touch main.go
```

)

2. While in your project folder install module with go get command

```shell
go get github.com/buahaha/Nutek
```

3. and fill `main.go` file with example code:

```
package main

import (
	"log"

	"github.com/buahaha/Nutek"
)

func main() {
    middleC, err := Nutek.StringToNote("middle c") // full set of available notes below
    if err != nil {
		log.Fatalf("can't convert string to note")
	}

    Nutek.OneNote(middleC, ".") // "." - use current folder, or specify your's
}
```

4. Create your first MIDI file with single note spanning one whole bar

```shell
go run main.go
```

5. This creates MIDI file, no sound, only raw data to be used in Digital Audio Workstation

___
Currently supported chord file generators (you can use any note you want to start with):

- Major 
    * `Nutek.MajorTriad(middleC, ".")`
- Minor
    * `Nutek.MinorTriad(middleC, ".")`
- Diminished
    * `Nutek.Diminished(middleC, ".")`
- Augmented
    * `Nutek.Augmented(middleC, ".")`
- Suspended 2
    * `Nutek.Sus2(middleC, ".")`
- Suspended 4
    * `Nutek.Sus4(middleC, ".")`
- Major 7th
    * `Nutek.Major7th(middleC, ".")`
- Minor 7th
    * `Nutek.Minor7th(middleC, ".")`
- 6th
    * `Nutek.Chord6th(middleC, ".")`
- 7th
    * `Nutek.Chord7th(middleC, ".")`
- 7 Suspended 2
    * `Nutek.Chord7Sus2(middleC, ".")`
- 7 Suspended 4
    * `Nutek.Chord7Sus4(middleC, ".")`
- Minor 6th
    * `Nutek.Minor6th(middleC, ".")`
- Diminished 7
    * `Nutek.Diminished7(middleC, ".")`
- Augmented 7
    * `Nutek.Augmented7(middleC, ".")`
- Diminished 6
    * `Nutek.Diminished7(middleC, ".")`

___
The following table summarizes the MIDI note numbers as defined in the MIDI standard and matched to the Middle C (note number 60) as C4.

The formula connecting the MIDI note number and the base frequency - assuming equal tuning based on A4=a'=440 Hz - is: 
```
f=440*math.Pow(2, (n-69)/12)
```

![MIDI map](/notes.gif)

| MIDI note number  | Key number (Organ)  | Key number (Piano)  | Note names (English)  |  Note names (German) | Frequency (Equal tuning at 440 Hz) |
|---|---|---|---|---|---|
| top of MIDI tuning range  |   |   | G#9/Ab9  | gis’’’’’’/ges’’’’’’  | 13289.75  |
| 127  |   |   | G9  | g’’’’’’ | 12543.85  |
|  126 |   |   | F#9/Gb9  | fis’’’’’’/ges’’’’’’  | 11839.82  |
| 125  |   |   |  F9 | f’’’’’’  | 11175.30  |
|  124 |   |   |  E9 |  e’’’’’’ | 10548.08  |
| 123  |   |   | D#9/Eb9  | dis’’’’’’/es’’’’’’  | 9956.06  |
|  122 |   |   | D9  | d’’’’’’  | 9397.27  |
|  121 |   |   |  C#9/Db9 |  cis’’’’’’/des’’’’’’ | 8869.84  |
120	 |	 |	| C9	| c’’’’’’	 |8372.02 |
| 119	 |	 |	| B8	| h’’’’’	| 7902.13 |
| 118	 |	 |	| A#8/Bb8	| ais’’’’’/b’’’’’	| 7458.62 |
| 117 |	 |	| 	A8 |	a’’’’’ |	7040.00 |
| 116 |	 |	| 	G#8/Ab8	| gis’’’’’/ges’’’’’ |	6644.88 |
| 115 |	 |	| 	G8	| g’’’’’	| 6271.93 |
| 114 |	 |	| 	F#8/Gb8 |	fis’’’’’/ges’’’’’ |	5919.91 |
| 113 |	 |	| 	F8 |	f’’’’’	| 5587.65 |
| 112 |	 |	| 	E8 |	e’’’’’	| 5274.04 |
| 111 |	 |	| 	D#8/Eb8	| dis’’’’’/es’’’’’	| 4978.03 |
| 110 |	 |	| 	D8	| d’’’’’	| 4698.64 |
| 109	 |	 |	| C#8/Db8	| cis’’’’’/des’’’’’	|4434.92 |
| 108 |	 |	88 |	C8 |	c’’’’’	| 4186.01
| 107 |	 |	87 |	B7 |	h’’’’	| 3951.07
| 106 |	 |	86 |	A#7/Bb7 |	ais’’’’/b’’’’	| 3729.31
| 105 |	 |	85 |	A7	| a’’’’	| 3520.00
| 104 | |	 84 |	G#7/Ab7	| gis’’’’/ges’’’’	| 3322.44
| 103 |	 |	83 |	G7	| g’’’’	| 3135.96
| 102 |	 |	82 |	F#7/Gb7	| fis’’’’/ges’’’’	| 2959.96
| 101 |	 |	81 |	F7	| f’’’’	| 2793.83
| 100 |	 |	80 |	E7	| e’’’’	| 2637.02
| 99 |	 |	79 |	D#7/Eb7	| dis’’’’/es’’’’	| 2489.02
| 98 |	 |	78 |	D7	| d’’’’	| 2349.32
| 97 |	 |	77 |	C#7/Db7	| cis’’’’/des’’’’	| 2217.46
| 96	| 61	| 76	| C7	| c’’’’	| 2093.00 |
| 95	| 60	| 75	| B6	| h’’’ |	| 1975.53 |
| 94	| 59	| 74	| A#6/Bb6	| ais’’’/b’’’	| 1864.66 |
| 93	| 58	| 73	| A6	| a’’’	| 1760.00 |
| 92	| 57	| 72	| G#6/Ab6	| gis’’’/as’’’	| 1661.22 |
| 91	| 56	| 71	| G6	| g’’’	| 1567.98 |
| 90	| 55	| 70	| F#6/Gb6	| fis’’’/ges’’’	| 1479.98 |
| 89	| 54	| 69	| F6	| f’’’	| 1396.91 |
| 88	| 53	| 68	| E6	| e’’’	| 1318.51 |
| 87	| 52	| 67	| D#6/Eb6	| dis’’’/es’’’	| 1244.51 |
| 86	| 51	| 66	| D6	| d’’’	| 1174.66 |
| 85	| 50	| 65	| C#6/Db6	| cis’’’/des’’’ |	1108.73 |
| 84	| 49	| 64	| C6	| c’’’	| 1046.50 |
| 83	| 48	| 63	| B5	| h’’	| 987.77 |
| 82	| 47	| 62	| A#5/Bb5	| ais’’/b’’	| 932.33 |
| 81	| 46	| 61	| A5	| a’’	| 880.00 |
| 80	| 45	| 60	| G#5/Ab5	| gis’’/as’’	| 830.61 |
| 79	| 44	| 59	| G5	| g’’	| 783.99 |
| 78	| 43	| 58	| F#5/Gb5	| fis’’/ges’’	| 739.99 |
| 77	| 42	| 57	| F5	| f’’	| 698.46 |
| 76	| 41	| 56	| E5	| e’’	| 659.26 |
| 75	| 40	| 55	| D#5/Eb5	| dis’’/es’’	| 622.25 |
| 74	| 39	| 54	| D5	| d’’	| 587.33 |
| 73	| 38	| 53	| C#5/Db5	| cis’’/des’’	| 554.37 |
| 72	| 37	| 52	| C5	| c’’	| 523.25 |
| 71	| 36	| 51	| B4	| h’	| 493.88 |
| 70	| 35	| 50	| A#4/Bb4	| ais’/b’	| 466.16 |
| 69	| 34	| 49	| A4 concert pitch	| a’ Kammerton	| 440.00 |
| 68	| 33	| 48	| G#4/Ab4	| gis’/as’	| 415.30 |
| 67	| 32	| 47	| G4	| g’	| 392.00 |
| 66	| 31	| 46	| F#4/Gb4	| fis’/ges’	| 369.99 |
| 65	| 30	| 45	| F4	| f’	| 349.23 |
| 64	| 29	| 44	| E4	| e’	| 329.63 |
| 63	| 28	| 43	| D#4/Eb4	| dis’/es’	| 311.13 |
| 62	| 27	| 42	| D4	| d’	| 293.66 |
| 61	| 26	| 41	| C#4/Db4	| cis’/des’	| 277.18 |
| 60	| 25	| 40	| C4 (middle C)	| c’ (Schloss-C) |	261.63 |
| 59	| 24	| 39	| B3	| h	| 246.94 | 
| 58	| 23	| 38	| A#3/Bb3	| ais/b	| 233.08 |
| 57	| 22	| 37	| A3	| a	| 220.00 | 
| 56	| 21	| 36	| G#3/Ab3	| gis/as	| 207.65 |
| 55	| 20	| 35	| G3	| g	| 196.00 |
| 54	| 19	| 34	| F#3/Gb3	| fis/ges	| 185.00 |
| 53	| 18	| 33	| F3	| f	| 174.61 |
| 52	| 17	| 32	| E3	| e	| 164.81 |
| 51	| 16	| 31	| D#3/Eb3	| dis/es	| 155.56 |
| 50	| 15	| 30	| D3	| d	| 146.83 |
| 49	| 14	| 29	| C#3/Db3	| cis/des	| 138.59 |
| 48	| 13	| 28	| C3	| c	| 130.81 |
| 47	| 12	| 27	| B2	| H	| 123.47 |
| 46	| 11	| 26	| A#2/Bb2	| Ais/B	| 116.54 |
| 45	| 10	| 25	| A2	| A	| 110.00 |
| 44	| 9	| 24	| G#2/Ab2	| Gis/As	| 103.83 |
| 43	| 8	| 23	| G2	| G	| 98.00 |
| 42	| 7	| 22	| F#2/Gb2	| Fis/Ges	| 92.50 |
| 41	| 6	| 21	| F2	| F	| 87.31 |
| 40	| 5	| 20	| E2	| E	| 82.41 |
| 39	| 4	| 19	| D#2/Eb2	| Dis/Es	| 77.78 |
| 38	| 3	| 18	| D2	| D	| 73.42 |
| 37	| 2	| 17	| C#2/Db2	| Cis/Des	| 69.30 |
| 36	| 1	| 16	| C2	| C	| 65.41 |
| 35	| 	| 15	| B1	| H1	|61.74 |
| 34 |	 | 14	| A#1/Bb1	| Ais1/b1	| 58.27 |
| 33	| 	| 13	| A1	| A1	| 55.00 |
| 32 |	 | 12 |	G#1/Ab1 | Gis1/As1 | 51.91 |
| 31	| 	| 11	| G1	| G1	| 49.00 |
| 30 |	 	| 10	| F#1/Gb1	| Fis1/Ges1 | 46.25 |
| 29	| 	| 9	| F1	| F1	| 43.65 |
| 28 |	 | 8 | E1 |	E1 | 41.20 |
| 27 | 	| 7	| D#1/Eb1	| Dis1/Es1 | 38.89 |
| 26 |	 | 6 | D1 |	D1 | 36.71 |
| 25 |  | 5	| C#1/Db1	| Cis1/Des1	| 34.65 |
| 24 |	 | 4 |	C1 |	C1 |	| 32.70 |
| 23	| 	| 3	| B0	| H2	| 30.87 |
| 22 |	 | 2 | A#0/Bb0 | Ais2/B2	| 29.14 | 
|21 |	| 1	| A0 |	| A2 | 27.50 |
| 20 |	 |	 |	 |	 | 25.96 |
| 19 | 	| 	 |	| 	| 24.50 |
| 18 | 	| 	| 	 |	| 23.12 |
| 17 |	 |	 |	 |	| 21.83 |
| 16 |	| 	| 	| 	| 20.60 |
| 15 |	 |	 |	 |	| 19.45 |
| 14 |	| 	| 	 |	 | 18.35 |
| 13 | 	 |	 |	 |	| 17.32 |
|12 |	| 	| 	| 	| 16.35 |
| 11 |	 |	 |	 |	| 15.43 |
| 10 |	| 	| 	| 	| 14.57 |
| 9	 |	 |	 |	 |	| 13.75 |
| 8	 |	 |	 |	 |	| 12.98 |
| 7	 |	 |	 |	 |	| 12.25 |
| 6	 |	 |	 |	 |	| 11.56 |
| 5	 |	 |	 |	 |	| 10.91 |
| 4	 |	 |	 |	 |	| 10.30 |
| 3	 |	 |	 |	 |	| 9.72 |
| 2	 |	 |	 |	 |	| 9.18 |
| 1	 |	 |	 |	 |	| 8.66 |
| 0	 |	 |	 |	 |	|8.18 |
					
[table source](https://www.inspiredacoustics.com/en/MIDI_note_numbers_and_center_frequencies)

[explanation or reference](https://studiocode.dev/resources/midi-middle-c/)




