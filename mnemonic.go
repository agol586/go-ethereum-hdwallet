package hdwallet

import (
	"fmt"
	"strings"

	"github.com/tyler-smith/go-bip39"
)

type Mnemonic struct {
	bitSize    int
	passphrase string
	data       string
}

func NewMenonic(bitSize int, passphrase string) (m *Mnemonic, err error) {
	if bitSize != 128 && bitSize != 256 {
		return nil, fmt.Errorf("invalid bit size: %d", bitSize)
	}
	m = &Mnemonic{
		passphrase: passphrase,
		bitSize:    bitSize,
	}
	entropy, err := bip39.NewEntropy(bitSize)
	if err != nil {
		return nil, err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, err
	}
	m.data = mnemonic
	return
}

func (m *Mnemonic) Words() []string {
	return strings.Split(m.data, " ")
}

func (m *Mnemonic) Data() string {
	return m.data
}

func (m *Mnemonic) Seed() []byte {
	return bip39.NewSeed(m.data, m.passphrase)
}

func (m *Mnemonic) String() string {
	s := strings.Builder{}
	s.WriteString("mnemonic: ")
	s.WriteString(m.data)
	s.WriteRune('\n')
	s.WriteString("passphrase: ")
	s.WriteString(m.passphrase)
	return s.String()
}
