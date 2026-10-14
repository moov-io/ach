#!/usr/bin/env python3
"""Collapse go test -bench -count=N samples to one median line per benchmark."""

from __future__ import annotations

import re
import sys
from collections import defaultdict

LINE = re.compile(
    r"^(Benchmark\S+)\s+(\d+)\s+(\d+(?:\.\d+)?)\s+ns/op(.*)$"
)
PKG = re.compile(r"^pkg:\s+(\S+)")
BOP = re.compile(r"(\d+(?:\.\d+)?)\s+B/op")
ALLOCS = re.compile(r"(\d+(?:\.\d+)?)\s+allocs/op")
CPU_SUFFIX = re.compile(r"(-\d+)$")


def median(vals: list[float]) -> float:
    s = sorted(vals)
    n = len(s)
    mid = n // 2
    if n % 2:
        return s[mid]
    return (s[mid - 1] + s[mid]) / 2


def fmt_ns(v: float) -> str:
    if abs(v - round(v)) < 1e-9:
        return str(int(round(v)))
    return f"{v:.3f}".rstrip("0").rstrip(".")


def display_name(pkg: str, name: str, pkgs_for_name: set[str]) -> str:
    if len(pkgs_for_name) <= 1 or not pkg:
        return name
    base = pkg.rsplit("/", 1)[-1]
    m = CPU_SUFFIX.search(name)
    if m:
        return name[: m.start()] + "/" + base + m.group(1)
    return name + "/" + base


def main() -> None:
    path = sys.argv[1] if len(sys.argv) > 1 and sys.argv[1] != "-" else None
    src = open(path, encoding="utf-8") if path else sys.stdin
    order: list[tuple[str, str]] = []
    iters: dict[tuple[str, str], list[int]] = defaultdict(list)
    ns: dict[tuple[str, str], list[float]] = defaultdict(list)
    bop: dict[tuple[str, str], list[float]] = defaultdict(list)
    allocs: dict[tuple[str, str], list[float]] = defaultdict(list)
    pkgs_for_name: dict[str, set[str]] = defaultdict(set)
    current_pkg = ""
    with src:
        for raw in src:
            line = raw.rstrip("\n")
            pm = PKG.match(line)
            if pm:
                current_pkg = pm.group(1)
                continue
            m = LINE.match(line)
            if not m:
                continue
            name = m.group(1)
            key = (current_pkg, name)
            if key not in ns:
                order.append(key)
            pkgs_for_name[name].add(current_pkg)
            iters[key].append(int(m.group(2)))
            ns[key].append(float(m.group(3)))
            rest = m.group(4) or ""
            bm = BOP.search(rest)
            am = ALLOCS.search(rest)
            if bm:
                bop[key].append(float(bm.group(1)))
            if am:
                allocs[key].append(float(am.group(1)))
    if not order:
        sys.stderr.write("bench-median: no Benchmark lines found\n")
        sys.exit(1)
    for key in order:
        pkg, name = key
        shown = display_name(pkg, name, pkgs_for_name[name])
        mid = sorted(iters[key])[len(iters[key]) // 2]
        parts = [shown, str(mid), f"{fmt_ns(median(ns[key]))} ns/op"]
        if bop[key]:
            parts.append(f"{int(round(median(bop[key])))} B/op")
        if allocs[key]:
            parts.append(f"{int(round(median(allocs[key])))} allocs/op")
        sys.stdout.write("\t".join(parts) + "\n")


if __name__ == "__main__":
    main()
