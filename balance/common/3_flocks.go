package common

type Flock struct {
	Description string
	ID int
	Expertize struct {
		Cre int 
		Alt int
		Des int
	}
	Experience struct {
		Cre float64 
		Alt float64
		Des float64
	}
	Streams []int
	Heat map[string]float64
	Threshold map[string]float64
}


// CREATE
func DefaultFlock(count int) *Flock {
	var buffer Flock
	buffer.ID = Epoch()
	buffer.Description = "DEFAULT: Contains all the streams and heat. Can not gain any exp."
	buffer.Heat = make(map[string]float64)
	buffer.Threshold = make(map[string]float64)
	for  x:=0; x<count; x++ { buffer.Streams = append(buffer.Streams, x) } 
	return &buffer
}