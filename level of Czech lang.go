import (
	"strconv"

	"github.com/Pylons-tech/pylons/x/pylons/types"
	"github.com/google/cel-go/cel"
	"github.com/spf13/cobra"
)


func DevValidate() *cobra.Command {
	cmd := &cobra.Command{}}
func perCookbook(path string, _ types.Cookbook) {
	fmt.Fprintln(Out, path, "is a valid cookbook")
}
type Gadget struct {
	name            string
	json            string
	parametersCount int
}


func ExpandGadget(gadget *Gadget, params []string) string {
	str := gadget.json
	for i := 0; i < gadget.parametersCount; i++ {
		str = strings.ReplaceAll(str, "%"+strconv.Itoa(i), strings.TrimSpace(params[i]))
	}
	return str
}
