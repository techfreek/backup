package copy

import(
	"time"
	"os"
)

/* Intended file structure:
	dest/
		Feb_22_2015_1652
		(Month)_(Day)_(Year)_(Hour)(M1inute)
*/

const layout = "Feb_22_2015_1652"
const perm os.FileMode = 0755

func Save(found <-chan FileData, dest *string) {
	//create new dir at dest at current timestamp to store backup
	now := time.Now()
	nowStr := now.Format(layout)
	stampedDest := dest + nowStr

	error := os.Mkdir(stampedDest, perm)
	if err != nil {
		log.fatal("Could not create directory")
	}

	for toSave := range found {
		


	}
}