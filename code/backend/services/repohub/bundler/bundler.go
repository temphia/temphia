package bundler

type InstallOptions struct {
}

type BundleInstaller interface {
	New()

	NewTableGroup()
	NewSheetGroup()
	NewPlug()
	NewTarget()
	NewResource()
}

type UpgradeOptions struct {
}

type BundleUpgrader interface {
	Upgrade()

	UpgradeTableGroup()
	UpgradeSheetGroup()
	UpgradePlug()
	UpgradeTarget()
	UpgradeResource()
}
