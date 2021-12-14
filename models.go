package exterr

import "sync"

//extendedErr - base error object
type extendedErr struct {
	msg     string
	altMsg  string
	errCode int
	trace   []traceRow
}

//traceRow - tracing line struct
type traceRow struct {
	Package  string `json:"package"`
	File     string `json:"file"`
	Function string `json:"function"`
	Line     int    `json:"line"`
}

//errorFromFile - error for reading from json
type errorFromFile struct {
	Error    string `json:"error"`
	AltError string `json:"alt-error"`
	Code     int    `json:"code"`
}

type ErrorNames struct {
	RWM   sync.RWMutex
	local string
	list  map[string]*extendedErr
}

func (err *ErrorNames) SetLocal(local string) {
	err.local = local
}

func (err *ErrorNames) GetError(name string) *extendedErr {
	err.RWM.RLock()
	defer err.RWM.RUnlock()
	if _, ok := err.list[name]; !ok {
		return nil
	}
	return err.list[name]
}

func (err *ErrorNames) GetAllKey() []string {
	err.RWM.RLock()
	defer err.RWM.RUnlock()
	listKeys := make([]string, 0)
	for key, _ := range err.list {
		listKeys = append(listKeys, key)
	}
	return listKeys
}

var ErrorByName ErrorNames

// "locals": ["en","ru"],
//     "errors": [

type ReadErrorsModel struct {
	Locale string `json:"locale"`
	Errors []struct {
		Name     string            `json:"name"`
		Error    map[string]string `json:"error"`
		AltError map[string]string `json:"alt-error"`
		Code     int               `json:"code"`
	} `json:"error"`
}
