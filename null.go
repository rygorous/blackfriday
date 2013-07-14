// Null renderer (useful starting point if you just want the document structure)

package blackfriday

import (
	"bytes"
)

type Null struct {
}

func (n *Null) BlockCode(out *bytes.Buffer, text []byte, lang string) {
}

func (n *Null) BlockQuote(out *bytes.Buffer, text []byte) {
}

func (n *Null) BlockHtml(out *bytes.Buffer, text []byte) {
}

func (n *Null) Header(out *bytes.Buffer, text func() bool, level int) {
	text()
}

func (n *Null) HRule(out *bytes.Buffer) {
}

func (n *Null) List(out *bytes.Buffer, text func() bool, flags int) {
	text()
}

func (n *Null) ListItem(out *bytes.Buffer, text []byte, flags int) {
}

func (n *Null) Paragraph(out *bytes.Buffer, text func() bool) {
	text()
}

func (n *Null) Table(out *bytes.Buffer, header []byte, body []byte, columnData []int) {
}

func (n *Null) TableRow(out *bytes.Buffer, text []byte) {
}

func (n *Null) TableCell(out *bytes.Buffer, text []byte, flags int) {
}

func (n *Null) AutoLink(out *bytes.Buffer, link []byte, kind int) {
}

func (n *Null) CodeSpan(out *bytes.Buffer, text []byte) {
}

func (n *Null) DoubleEmphasis(out *bytes.Buffer, text []byte) {
}

func (n *Null) Emphasis(out *bytes.Buffer, text []byte) {
}

func (n *Null) Image(out *bytes.Buffer, link []byte, title []byte, alt []byte) {
}

func (n *Null) LineBreak(out *bytes.Buffer) {
}

func (n *Null) Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {
}

func (n *Null) RawHtmlTag(out *bytes.Buffer, tag []byte) {
}

func (n *Null) TripleEmphasis(out *bytes.Buffer, text []byte) {
}

func (n *Null) StrikeThrough(out *bytes.Buffer, text []byte) {
}

func (n *Null) DisplayMath(out *bytes.Buffer, text []byte) {
}

func (n *Null) InlineMath(out *bytes.Buffer, text []byte) {
}

func (n *Null) Entity(out *bytes.Buffer, entity []byte) {
}

func (n *Null) NormalText(out *bytes.Buffer, text []byte) {
}

func (n *Null) DocumentHeader(out *bytes.Buffer) {
}

func (n *Null) DocumentFooter(out *bytes.Buffer) {
}
