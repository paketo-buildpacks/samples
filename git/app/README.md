For the Paketo Git Buildpack to detect true there must be a `.git`
directory in the application source. To accomplish this a valid `.git`
directory has been added and named `.git.bak`. To make this application
functional rename the `.git.bak` folder to `.git`.
