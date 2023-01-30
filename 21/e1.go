// Реализовать паттерн «адаптер» на любом примере.

package main

import "fmt"

type Connector interface {
	playAudio() //Интерфейс проигрывания аудио
}

type Jack struct {
}

func (j *Jack) playAudio() {
	fmt.Println("Jack: playAudio()") // Реализация существующего интерфейса
}

type Bluetooth struct {
}

func (b *Bluetooth) playAudioWithBluetooth() { // Новый способ воспроизведения
	fmt.Println("Bluetooth: playAudioWithBluetooth()")
}

type BluetoothAdapter struct {
	b *Bluetooth
}

func (ba *BluetoothAdapter) playAudio() { // Адаптация под старый интерфейс
	ba.b.playAudioWithBluetooth()
}

type Client struct {
}

func (c *Client) playAudioWithConnector(con Connector) {
	con.playAudio()
}

func main() {
	c := &Client{}
	j := &Jack{}
	b := &Bluetooth{}
	ba := &BluetoothAdapter{b}

	c.playAudioWithConnector(j)
	c.playAudioWithConnector(ba)
}
