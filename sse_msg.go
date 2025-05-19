package gcurl

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

type SseMsg struct {
	Id    string
	Event string
	Data  string
	Retry time.Duration
}

func NewSseMsg() *SseMsg {
	ret := &SseMsg{}
	return ret
}

func (m *SseMsg) SetId(id string) *SseMsg {
	m.Id = id
	return m
}

func (m *SseMsg) SetEvent(event string) *SseMsg {
	m.Event = event
	return m
}

func (m *SseMsg) SetData(data string) *SseMsg {
	m.Data = data
	return m
}

func (m *SseMsg) SetRetry(retry time.Duration) *SseMsg {
	m.Retry = retry
	return m
}

func (m *SseMsg) FormatMsg() []byte {
	var data bytes.Buffer
	if len(m.Id) > 0 {
		data.WriteString(fmt.Sprintf("id: %s\n", strings.Replace(m.Id, "\n", "", -1)))
	}
	if len(m.Event) > 0 {
		data.WriteString(fmt.Sprintf("event: %s\n", strings.Replace(m.Event, "\n", "", -1)))
	}
	if len(m.Data) > 0 {
		lines := strings.Split(m.Data, "\n")
		for _, line := range lines {
			data.WriteString(fmt.Sprintf("data: %s\n", line))
		}
	}
	if m.Retry > 0 {
		data.WriteString(fmt.Sprintf("retry: %d\n", m.Retry/time.Millisecond))
	}
	data.WriteString("\n")
	return data.Bytes()
}

type SseRetryMsg struct {
	Retry time.Duration
}

func NewSseRetryMsg(retry time.Duration) *SseRetryMsg {
	ret := &SseRetryMsg{
		Retry: retry,
	}
	return ret
}

func (m *SseRetryMsg) FormatMsg() []byte {
	return []byte(fmt.Sprintf("retry: %d\n\n", m.Retry/time.Millisecond))
}
