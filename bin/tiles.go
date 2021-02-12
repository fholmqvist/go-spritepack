package main

type Tiles []Tile

func (tt Tiles) ToSet() (set Tiles) {
	if len(tt) == 0 {
		return set
	}
	set = append(set, tt[0])
	for i := 0; i < len(tt)-1; i++ {
		for ii := i + 1; ii < len(tt); ii++ {
			if !tt[ii].IdenticalInRotations(tt[i]) {
				set = append(set, tt[ii])
			}
		}
	}
	return set
}
