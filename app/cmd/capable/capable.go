package main

import (
	"flag"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"github.com/p4tin/Furiosa/app"
)

var Args app.ArgsStruct

func init() {
	DraftRelease := flag.Bool("draft_release", false, "Create a GitHub Draft release")
	PreRelease := flag.Bool("pre_release", false, "Create a GitHub Pre-release")
	CutRelease := flag.Bool("cut_release", false, "Create a GitHub Release")
	Hotfix := flag.Bool("hotfix", false, "Create a GitHub Release")
	ReleaseBranchName := flag.String("release_branch_name", "X", "Name of release candidate branch. Default is X.")
	Suffix := flag.String("rc", "", "Release Candidate Suffix _number_")
	NoSuffix := flag.Bool("no_rc", false, "No Release Candidate suffix (use with --pre_release)")
	Repository := flag.String("repository", "", "repository (e.g., urbn/CatalogService)")
	SemverBump := flag.String("semver_bump", "minor", "One of major, minor or patch. Specifies which part of Semantic Version to increment.")
	BranchFromName := flag.String("branch_from_name", "", "In any mode creating a release, branch from the named branch to create a new release candidate.")
	BranchFromSha := flag.String("branch_from_sha", "", "In any mode creating a release, branch from sha to create a new release candidate. \n\tYou must supply a full 40-digit SHA1.")

	flag.Parse()

	Args = app.ArgsStruct{
		DraftRelease:      *DraftRelease,
		PreRelease:        *PreRelease,
		CutRelease:        *CutRelease,
		Hotfix:            *Hotfix,
		ReleaseBranchName: *ReleaseBranchName,
		Suffix:            *Suffix,
		NoSuffix:          *NoSuffix,
		Repository:        *Repository,
		SemverBump:        *SemverBump,
		BranchFromName:    *BranchFromName,
		BranchFromSha:     *BranchFromSha,
	}

	log.Debugf("Args: %+v\n", Args)

	dat, err := ioutil.ReadFile("config/capable_config.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	Config := app.ConfigStruct{}
	err = yaml.Unmarshal(dat, &Config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Debugf("Config: %+v\n", Config)
}

func main() {
	endPoint := "http://localhost:50005/health"
	log.Infof("EndPoint is %v\n", endPoint)

	resp, err := http.Get(endPoint)
	if err != nil {
		log.Fatalf("Problem contacting %v (%v)\n", endPoint, err)
	}

	log.Infof("Response is %v\n", resp)
}
