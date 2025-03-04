package models

type PlayingPiece struct {
	PieceType PieceType
}

func NewPlayingPiece(pieceType PieceType) *PlayingPiece {
	return &PlayingPiece{
		PieceType: pieceType,
	}
}
