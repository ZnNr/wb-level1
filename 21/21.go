package main

import "fmt"

/*
Реализовать паттерн проектирования «Адаптер» на любом примере.

Описание: паттерн Adapter позволяет сконвертировать интерфейс одного класса в интерфейс другого, который ожидает клиент.

Продемонстрируйте на простом примере в Go: у вас есть существующий интерфейс (или структура) и другой,
несовместимый по интерфейсу потребитель — напишите адаптер,
который реализует нужный интерфейс и делегирует вызовы к встроенному объекту.

Поясните применимость паттерна, его плюсы и минусы, а также приведите реальные примеры использования.
*/

// audioPlayer
type AudioPlayer interface {
	Play(audioType string, filename string)
}
type Mp3Player struct{}

// Play mp3
func (mp *Mp3Player) Play(audioType string, filename string) {
	fmt.Printf("Playing MP3 file: %s\n", filename)
}

type WavPlayer struct{}

func (wp *WavPlayer) PlayWav(filename string) {
	fmt.Printf("Playing WAV file: %s\n", filename)
}

type WavAdapter struct {
	wavPlayer *WavPlayer
}

func (wa *WavAdapter) Play(audioType string, fileName string) {
	if audioType == "wav" {
		wa.wavPlayer.PlayWav(fileName)
	} else {
		fmt.Printf("Error: WavPlayer not initialized\n")
	}
}

type AudioAdapter struct {
	mp3Player  *Mp3Player
	wavAdapter *WavAdapter
}

func (aa *AudioAdapter) Play(audioType string, fileName string) {
	if audioType == "mp3" {
		aa.mp3Player.Play(audioType, fileName)
	} else if audioType == "wav" {
		aa.wavAdapter.Play(audioType, fileName)
	} else {
		fmt.Printf("Unsupported audio format: %s\n", audioType)
	}
}
func main() {
	// Обычный MP3-проигрыватель (уже совместим)
	mp3Player := &Mp3Player{}
	// WAV-проигрыватель (НЕ совместим) - обертка в адаптер
	wavPlayer := &WavPlayer{}
	wavAdapter := &WavAdapter{wavPlayer: wavPlayer}
	audioAdapter := &AudioAdapter{mp3Player: mp3Player, wavAdapter: wavAdapter}

	audioAdapter.Play("mp3", "song.mp3")
	audioAdapter.Play("wav", "track.wav")
	audioAdapter.Play("ogg", "audio.ogg")
}
