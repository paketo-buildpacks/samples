#!/usr/bin/env bash

set -eu
set -o pipefail

readonly PROGDIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly SAMPLESDIR="$(cd "${PROGDIR}/.." && pwd)"

# shellcheck source=SCRIPTDIR/.util/tools.sh
source "${PROGDIR}/.util/tools.sh"

# shellcheck source=SCRIPTDIR/.util/print.sh
source "${PROGDIR}/.util/print.sh"

function main() {
  local language

  while [[ "${#}" != 0 ]]; do
    case "${1}" in
      --help|-h)
        shift 1
        usage
        exit 0
        ;;

      --language|-l)
        language="${2}"
        shift 2
        ;;

      "")
        # skip if the argument is empty
        shift 1
        ;;

      *)
        util::print::error "unknown argument \"${1}\""
    esac
  done

  sed -e "s/\${language}/${language}/" "${PROGDIR}/.util/test-workflow-template.yml" > "${SAMPLESDIR}/.github/workflows/test-pull-request-${language}.yml"
}

function usage() {
  cat <<-USAGE
generate-test-workflow.sh [OPTIONS]

Generates a Github Actions workflow that will run tests on PRs that change
files in the <language> subdirectory. The generated workflow will be placed in
.github/workflows/test-pull-request-<language>.yml

OPTIONS
  --help            -h        prints the command usage
  --language <name> -l <name> the directory in which new language family's samples live (e.g. dotnet-core)
USAGE
}

main "${@:-}"
