package r

import "vimagination.zapto.org/parser"

type File struct {
	Statements []Expression
	Tokens     Tokens
}

func Parse(t Tokeniser) (*File, error) {
	r, err := newRParser(t)
	if err != nil {
		return nil, err
	}

	f := new(File)
	if err = f.parse(&r); err != nil {
		return nil, err
	}

	return f, nil
}

func (f *File) parse(r *rParser) error {
	for r.AcceptRunWhitespace() != parser.TokenDone {
		var s Expression

		q := r.NewGoal()

		if err := s.parse(&q); err != nil {
			return r.Error("File", err)
		}

		f.Statements = append(f.Statements, s)

		r.Score(q)
	}

	f.Tokens = r.ToTokens()

	return nil
}
