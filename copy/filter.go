package copy

import(
	"string"
	"os"
)

filters := [...]string[".","..","bin"]

func InitFilter(userfilters []string) {
	filters = append(filters, userfilters)
}

func ShouldCopy(file FileData) {
	filename := file.Info.Name()

	for filter := range filters {
		if filter == filename {
			return false
		}
	}
}
