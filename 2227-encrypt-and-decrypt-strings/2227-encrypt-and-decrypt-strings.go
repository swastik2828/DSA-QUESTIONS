
type Encrypter struct {
    mapping map[byte]string
    freq    map[string]int
}

func Constructor(keys []byte, values []string, dictionary []string) Encrypter {
    e := Encrypter{
        mapping: make(map[byte]string),
        freq:    make(map[string]int),
    }

    // Build character-to-encrypted-value mapping
    for i, key := range keys {
        e.mapping[key] = values[i]
    }

    // Pre-encrypt dictionary words and count frequencies
    for _, word := range dictionary {
        encrypted := e.Encrypt(word)
        e.freq[encrypted]++
    }

    return e
}

func (this *Encrypter) Encrypt(word string) string {
    var result strings.Builder

    for i := 0; i < len(word); i++ {
        encrypted, exists := this.mapping[word[i]]

        if !exists {
            return ""
        }

        result.WriteString(encrypted)
    }

    return result.String()
}

func (this *Encrypter) Decrypt(word string) int {
    return this.freq[word]
}
