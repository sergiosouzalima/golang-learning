package position

import (
	"testing"
)

func TestAdvance(t *testing.T) {
	// Cria uma nova posição inicial
	pos := Position{idx: 0, ln: 1, col: 0, fn: "example.txt", ftxt: "Some text"}

	// Testa avanço com caractere normal
	t.Log("Testing advance with a normal character 'a'")
	pos.Advance('a')
	if pos.idx != 1 || pos.col != 1 || pos.ln != 1 {
		t.Errorf("Advance with 'a' failed: expected (1, 1, 1), got (%d, %d, %d)", pos.idx, pos.col, pos.ln)
	}
	t.Logf("After advancing with 'a': idx=%d, col=%d, ln=%d", pos.idx, pos.col, pos.ln)

	// Testa avanço com 'x'
	t.Log("Testing advance with a normal character 'x'")
	pos.Advance('x')
	if pos.idx != 2 || pos.col != 2 || pos.ln != 1 {
		t.Errorf("Advance with 'x' failed: expected (2, 2, 1), got (%d, %d, %d)", pos.idx, pos.col, pos.ln)
	}
	t.Logf("After advancing with 'x': idx=%d, col=%d, ln=%d", pos.idx, pos.col, pos.ln)

	// Testa avanço com caractere de nova linha
	t.Log("Testing advance with a newline character '\\n'")
	pos.Advance('\n')
	if pos.idx != 3 || pos.col != 0 || pos.ln != 2 {
		t.Errorf("Advance with newline failed: expected (2, 0, 2), got (%d, %d, %d)", pos.idx, pos.col, pos.ln)
	}
	t.Logf("After advancing with '\\n': idx=%d, col=%d, ln=%d", pos.idx, pos.col, pos.ln)
}

func TestCopy(t *testing.T) {
	// Cria uma posição original
	original := Position{idx: 5, ln: 10, col: 15, fn: "test.txt", ftxt: "Text content"}

	// Cria uma cópia da posição original
	t.Log("Testing the copy method")
	copy := original.Copy()
	if copy.idx != original.idx || copy.ln != original.ln || copy.col != original.col ||
		copy.fn != original.fn || copy.ftxt != original.ftxt {
		t.Errorf("Copy failed: expected (idx=%d, ln=%d, col=%d, fn=%s, ftxt=%s), got (idx=%d, ln=%d, col=%d, fn=%s, ftxt=%s)",
			original.idx, original.ln, original.col, original.fn, original.ftxt,
			copy.idx, copy.ln, copy.col, copy.fn, copy.ftxt)
	}
	t.Logf("Original Position: idx=%d, ln=%d, col=%d, fn=%s, ftxt=%s", original.idx, original.ln, original.col, original.fn, original.ftxt)
	t.Logf("Copied Position: idx=%d, ln=%d, col=%d, fn=%s, ftxt=%s", copy.idx, copy.ln, copy.col, copy.fn, copy.ftxt)
}
