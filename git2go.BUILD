package(default_visibility = ["//visibility:public"])
load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = 'go_default_library',
    srcs = glob(["*.go"]) + ["wrapper.c"],
    importpath = 'github.com/libgit2/git2go/v30',
    cgo = True,
    cdeps = [
      "@libgit2//:libgit2",
    ],
)
