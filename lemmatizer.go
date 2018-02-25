package ling

import (
	"fmt"
	"github.com/liuzl/ling/resources/lemma"
)

type Lemmatizer struct {
}

func (self *Lemmatizer) Process(d *Document) error {
	if d == nil || len(d.Text) == 0 {
		return fmt.Errorf("ducument is empty")
	}
	if len(d.Tokens) == 0 {
		return fmt.Errorf("tokenization required")
	}

	for _, lang := range d.Langs {
		if f, has := lemma.Funcs[lang]; has {
			ret, err := f(d.XTokens("norm"))
			if err != nil {
				return err
			}
			if len(ret) != len(d.Tokens) {
				continue
			}
			for i, str := range ret {
				d.Tokens[i].Annotations["lemma"] = str
			}
		}
	}
	return nil
}
