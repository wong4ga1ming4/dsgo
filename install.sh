#!/bin/sh
set -eu

version="${DSGO_VERSION:-v0.1.0}"
repo="wong4ga1ming4/dsgo"

os=$(uname -s | tr '[:upper:]' '[:lower:]')
machine=$(uname -m)
case "$machine" in
  x86_64|amd64) arch=amd64 ;;
  arm64|aarch64) arch=arm64 ;;
  *)
    echo "不支持的架构: $machine" >&2
    exit 1
    ;;
esac
case "$os" in
  darwin|linux) ;;
  *)
    echo "不支持的系统: $os" >&2
    exit 1
    ;;
esac

name="dsgo-${os}-${arch}"
url="https://github.com/${repo}/releases/download/${version}/${name}"

if [ -n "${DSGO_BIN:-}" ]; then
  bindir=$DSGO_BIN
elif [ -w /usr/local/bin ]; then
  bindir=/usr/local/bin
else
  bindir="${HOME}/.local/bin"
fi
mkdir -p "$bindir"

tmp=$(mktemp)
trap 'rm -f "$tmp"' EXIT
curl -fsSL "$url" -o "$tmp"
chmod 755 "$tmp"
mv "$tmp" "${bindir}/dsgo"
trap - EXIT

if ! command -v dsgo >/dev/null 2>&1; then
  line="export PATH=\"${bindir}:\$PATH\" # dsgo"
  for rc in "$HOME/.zshrc" "$HOME/.bashrc" "$HOME/.profile"; do
    if [ ! -f "$rc" ]; then
      continue
    fi
    if grep -q '# dsgo' "$rc"; then
      continue
    fi
    printf '\n%s\n' "$line" >> "$rc"
  done
  if ! grep -q '# dsgo' "$HOME/.profile" 2>/dev/null; then
    printf '\n%s\n' "$line" >> "$HOME/.profile"
  fi
fi

echo "已安装 ${bindir}/dsgo (${version})"
if ! command -v dsgo >/dev/null 2>&1; then
  echo "新开一个终端后再运行 dsgo，或执行: export PATH=\"${bindir}:\$PATH\""
fi
