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
