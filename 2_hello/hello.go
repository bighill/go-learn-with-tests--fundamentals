package hello

const spanish = "Spanish"
const french = "French"
const enHelloPrefix = "hello "
const esHelloPrefix = "hola "
const frHelloPrefix = "bonjour "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = esHelloPrefix
	case french:
		prefix = frHelloPrefix
	default:
		prefix = enHelloPrefix
	}
	return
}
