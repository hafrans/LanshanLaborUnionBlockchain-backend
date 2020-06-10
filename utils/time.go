package utils

import "time"

const timeFormat = "2006-01-02 15:04:05"

type Time time.Time

func (t *Time) UnmarshalJSON(b []byte) error{
	tx, err := time.Parse(timeFormat,string(b))
	if err!=nil{
		return err
	}
	*t = Time(tx)
	return nil
}

func (t *Time) MarshalJSON() ([]byte,error){
	b := make([]byte,0,len(timeFormat) + 2)
	b = append(b,'"')
	b = time.Time(*t).AppendFormat(b,timeFormat)
	b = append(b,'"')
	return b, nil
}

func (t Time) FormattedString() string{
	return time.Time(t).Format(timeFormat)
}

func CurrentTimeString() string{
	return time.Now().Format("2006-01-02 15:04:05")
}