package caches

import (
	"app/base/core"
	"app/base/utils"
	"app/tasks"
)

var (
	enableRefreshPackagesCache  bool
	enableRefreshAdvisoryCaches bool
)

func configure() {
	core.ConfigureApp()
	enableRefreshPackagesCache = utils.GetBoolEnvOrDefault("ENABLE_REFRESH_PACKAGES_CACHE", false)
	enableRefreshAdvisoryCaches = utils.GetBoolEnvOrDefault("ENABLE_REFRESH_ADVISORY_CACHES", false)
}

func RunPackageRefresh() {
	tasks.HandleContextCancel(tasks.WaitAndExit)
	configure()
	if enableRefreshPackagesCache {
		utils.Log().Info("Refreshing package cache")
		RefreshLatestPackagesView()
	}
}

func RunAdvisoryRefresh() {
	tasks.HandleContextCancel(tasks.WaitAndExit)
	configure()
	utils.Log().Info("Refreshing advisory cache")
	RefreshAdvisoryCaches()
}
