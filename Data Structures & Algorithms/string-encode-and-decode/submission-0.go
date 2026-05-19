type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    var sb strings.Builder

    for _, str := range strs {
        sb.WriteString(strconv.Itoa(len(str)))
        sb.WriteByte('#')
        sb.WriteString(str)
    }

    return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
    var result []string
    i := 0

    for i < len(encoded) {
        j := i
        for encoded[j] != '#' {
            j++
        }

        length, _ := strconv.Atoi(encoded[i:j])

        str := encoded[j+1 : j+1+length]
        result = append(result, str)

        i = j + 1 + length
    }

    return result
}
