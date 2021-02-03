#!/usr/bin/env bash

set -eu
set -o pipefail

# shellcheck source=SCRIPTDIR/print.sh
source "$(dirname "${BASH_SOURCE[0]}")/print.sh"

function util::tools::path::export() {
  local dir
  dir="${1}"

  if ! echo "${PATH}" | grep -q "${dir}"; then
    PATH="${dir}:$PATH"
    export PATH
  fi
}

function util::tools::pack::install() {
  local dir
  while [[ "${#}" != 0 ]]; do
    case "${1}" in
      --directory)
        dir="${2}"
        shift 2
        ;;

      *)
        util::print::error "unknown argument \"${1}\""
    esac
  done

  mkdir -p "${dir}"
  util::tools::path::export "${dir}"

  local os
  case "$(uname)" in
    "Darwin")
      os="macos"
      ;;

    "Linux")
      os="linux"
      ;;

    *)
      echo "Unknown OS \"$(uname)\""
      exit 1
  esac

  if [[ ! -f "${dir}/pack" ]]; then
    local version
    if [[ ! -f "${SAMPLESDIR}/.github/pack-version" ]]; then
      util::print::error "Pack version file (.github/pack-version) not found"
    fi

    version="v$(cat "${SAMPLESDIR}"/.github/pack-version)"

    util::print::title "Installing pack ${version}"
    curl "https://github.com/buildpacks/pack/releases/download/${version}/pack-${version}-${os}.tgz" \
      --silent \
      --location \
      --output /tmp/pack.tgz
    tar xzf /tmp/pack.tgz -C "${dir}"
    chmod +x "${dir}/pack"
    rm /tmp/pack.tgz
  fi
}

function util::tools::tests::checkfocus() {
  testout="${1}"
  if grep -q 'Focused: [1-9]' "${testout}"; then
    echo "Detected Focused Test(s) - setting exit code to 197"
    rm "${testout}"
    util::print::success "** GO Test Succeeded **" 197
  fi
  rm "${testout}"
}
