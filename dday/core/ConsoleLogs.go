package core

import "time"

// Logs Logic
type Logs = [][2]string

// Log a message on the console
func (a *Application) Log(s string) {
	a.LogsContent = append(a.LogsContent, [2]string{s, time.Now().Format(time.TimeOnly)})
	a.TeaProgram.Send(LoggedMsg(a.LogsContent))
}

type LoggedMsg [][2]string
