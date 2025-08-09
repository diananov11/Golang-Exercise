package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool
// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	sum := 0
	for _, v := range cb[file] {
       if v {
          sum++
       }
    }
    return sum
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 && rank > 8 {
        return 0
    }

    occupied := 0
    for _, val := range cb {
        for i, bol := range val {
            if i == rank - 1 {
                if bol {
                    occupied++
                }
            }
        }

    }
    return occupied
    
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    sum := 0
	for range cb {
        sum++
    }
    return sum*sum
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
    occupied := 0
	for _, occ := range cb {
        for _, bool := range occ {
            if bool {
                occupied++
            }
        }
    }
    return occupied
}
