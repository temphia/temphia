package hook

type HookType struct {
	ClientSide   bool
	Name         string
	TargetPrefix string
}

var HookTypes []HookType

func init() {

	// dtable

	HookTypes = append(HookTypes, HookType{
		ClientSide:   true,
		Name:         "selected_row",
		TargetPrefix: "dtable",
	})

	HookTypes = append(HookTypes, HookType{
		ClientSide:   true,
		Name:         "selected_rows",
		TargetPrefix: "dtable",
	})

	HookTypes = append(HookTypes, HookType{
		ClientSide:   false,
		Name:         "before_modify",
		TargetPrefix: "dtable",
	})

	HookTypes = append(HookTypes, HookType{
		ClientSide:   false,
		Name:         "after_modify",
		TargetPrefix: "dtable",
	})

}
