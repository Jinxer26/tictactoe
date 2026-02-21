package game

func (b* Board) Winner() Cell{
	lines := [][][2]int{
		{{0,0},{0,1},{0,2}},
		{{1,0},{1,1},{1,2}},
		{{2,0},{2,1},{2,2}},
		{{0,0},{1,0},{2,0}},
		{{0,1},{1,1},{2,1}},
		{{0,2},{1,2},{2,2}},
		{{0,0},{1,1},{2,2}},
		{{0,2},{1,1},{2,0}},
	}

	for _,line := range lines{
		x, y, z := line[0],line[1],line[2]
		v := b.Cells[x[0]][x[1]]
		if v != Empty && v == b.Cells[y[0]][y[1]] && v == b.Cells[z[0]][z[1]]{
			return v
		}
	}

	return Empty

}

func (b *Board) Draw() bool {
	if b.Winner() != Empty {
		return false
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.Cells[i][j] == Empty {
				return false
			}
		}
	}

	return true
}
