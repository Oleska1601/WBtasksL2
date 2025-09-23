package timereceiver

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
)

// TimeReceiver клиент для получения времени с NTP-серверов
type TimeReceiver struct {
	address string
}

// New создает новый экземпляр TimeReceiver с указанным адресом NTP-сервера
func New(address string) *TimeReceiver {
	return &TimeReceiver{
		address: address,
	}
}

// GetTime возвращает текущее время с NTP-сервера
// Возвращает ошибку, если не удалось получить время
func (tr *TimeReceiver) GetTime() (time.Time, error) {
	curTime, err := ntp.Time(tr.address)
	if err != nil {
		return time.Time{}, fmt.Errorf("ntp time error: %w", err)
	}
	return curTime, nil
}
